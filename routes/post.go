package routes

import (
	"github.com/codingpalette/gin_crud/models"
	"github.com/gin-gonic/gin"
)

func PostRoutes(router *gin.Engine) {
	v1 := router.Group("/v1/post")
	{
		v1.GET("/", func(c *gin.Context) {
			result := models.GetPost()

			c.JSON(200, gin.H{
				"message": result,
			})
		})
	}
}
