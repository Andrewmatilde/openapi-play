package api

import (
	"encoding/json"
	"net/http"
	"play/pkg"

	"github.com/gin-gonic/gin"
)

type Registry struct {
	Store *pkg.Store
}

func NewRegistry(store *pkg.Store) *Registry {
	return &Registry{Store: store}
}

// 实现 ServerInterface
func (r *Registry) UserActions(c *gin.Context) {
	var raw map[string]interface{}
	if err := c.ShouldBindJSON(&raw); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求体解析失败"})
		return
	}
	action, _ := raw["action"].(string)
	userRaw, _ := raw["user"]
	userBytes, _ := json.Marshal(userRaw)

	switch action {
	case "create":
		var req UserNameActionRequest
		if err := json.Unmarshal(userBytes, &req.User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user 字段解析失败"})
			return
		}
		r.Store.Set(req.User.Name, string(userBytes))
		c.JSON(http.StatusOK, req.User)
	case "get":
		var req UserNameActionRequest
		if err := json.Unmarshal(userBytes, &req.User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user 字段解析失败"})
			return
		}
		val, ok := r.Store.Get(req.User.Name)
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
			return
		}
		var user User
		_ = json.Unmarshal([]byte(val), &user)
		c.JSON(http.StatusOK, user)
	case "put":
		var req UserNameActionRequest
		if err := json.Unmarshal(userBytes, &req.User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user 字段解析失败"})
			return
		}
		r.Store.Set(req.User.Name, string(userBytes))
		c.JSON(http.StatusOK, req.User)
	case "delete":
		var req UserNameActionRequest
		if err := json.Unmarshal(userBytes, &req.User); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user 字段解析失败"})
			return
		}
		r.Store.Delete(req.User.Name)
		c.JSON(http.StatusOK, gin.H{"msg": "删除成功"})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "未知 action"})
	}
}
