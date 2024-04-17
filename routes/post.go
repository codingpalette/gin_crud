package routes

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/codingpalette/gin_crud/models"
	"github.com/gin-gonic/gin"
)

func PostRoutes(router *gin.Engine, db *sql.DB) {
	v1 := router.Group("/v1/post")
	{
		v1.GET("/list", func(c *gin.Context) {
			fmt.Println("list")
			result, err := models.GetPostList(db)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(200, gin.H{
				"message": "조회 성공",
				"data":    result,
			})
		})

		// JSON 바인딩
		type PostType struct {
			Title   string `form:"title" json:"title" xml:"title" binding:"required"`
			Content string `form:"content" json:"title" xml:"title" binding:"required"`
		}

		v1.POST("/create", func(c *gin.Context) {
			// title := c.PostForm("title")
			// content := c.PostForm("content")
			var form PostType
			if err := c.ShouldBind(&form); err != nil {
				fmt.Println("err", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			result, err := models.PostCreate(db, form.Title, form.Content)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			fmt.Println("post_info", result)

			c.JSON(http.StatusOK, gin.H{
				"message": "생성 성공",
				"data":    result,
			})
		})

		v1.GET(":id", func(c *gin.Context) {
			id := c.Param("id")

			result, err := models.GetPost(db, id)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			if result.Id == 0 {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "게시글이 존재하지 않습니다.",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "조회 성공",
				"data":    result,
			})

		})

		v1.PUT(":id", func(c *gin.Context) {
			var form PostType
			if err := c.ShouldBind(&form); err != nil {
				fmt.Println("err", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			idStr := c.Param("id")
			id, err := strconv.Atoi(idStr)

			result, err := models.PostUpdate(db, id, form.Title, form.Content)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "수정 성공",
				"data":    result,
			})
		})

		v1.DELETE(":id", func(c *gin.Context) {
			idStr := c.Param("id")
			id, err := strconv.Atoi(idStr)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			err = models.PostDelete(db, id)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "삭제 성공",
			})
		})

	}
}
