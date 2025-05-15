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

// UserJson 用户信息 Json 格式
type UserJson struct {
	// Balance 存款数
	Balance int32 `json:"balance"`

	// Name 用户名
	Name string `json:"name"`
}

// 实现 ServerInterface
func (r *Registry) GetUser(c *gin.Context) {
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
		var user User
		user.Name = req.User.Name
		user.Balance = 100
		userjson := UserJson{
			Balance: user.Balance,
			Name:    user.Name,
		}
		userjsonBytes, _ := json.Marshal(userjson)
		userJsonStr := string(userjsonBytes)
		user.UserJson = &userJsonStr

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
