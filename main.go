package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

type ResourceRequest struct {
	Group     string `json:"group"`
	Version   string `json:"version"`
	Kind      string `json:"kind"`
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

func main() {
	// 加载 kubeconfig
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		kubeconfig = os.ExpandEnv("$HOME/.kube/config")
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}

	// 创建 dynamic client
	dynClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.POST("/resource", func(c *gin.Context) {
		var req ResourceRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		gvr := schema.GroupVersionResource{
			Group:    req.Group,
			Version:  req.Version,
			Resource: toResourceName(req.Kind),
		}

		// 查询资源
		resource, err := dynClient.Resource(gvr).Namespace(req.Namespace).Get(context.Background(), req.Name, metav1.GetOptions{})
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		// 返回原始 JSON
		raw, err := json.Marshal(resource.Object)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": string(raw),
		})
	})

	r.Run(":8080")
}

// 将 Kind 转为 resource 名称（简单实现，实际可用 RESTMapping 动态查）
func toResourceName(kind string) string {
	switch kind {
	case "Pod":
		return "pods"
	case "Ingress":
		return "ingresses"
	// 可继续扩展
	default:
		return fmt.Sprintf("%ss", kind)
	}
}
