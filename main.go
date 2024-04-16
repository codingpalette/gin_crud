package main

import (
	"github.com/codingpalette/gin_crud/routes"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	routes.PostRoutes(r)
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
