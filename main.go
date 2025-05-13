package main

import (
	"log"
	"play/api"
	"play/pkg"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := pkg.NewStore()
	registry := api.NewRegistry(store)
	api.RegisterHandlers(r, registry)

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	log.Println("Server started at :8080")
	r.Run(":8080")
}
