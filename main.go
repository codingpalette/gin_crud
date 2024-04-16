package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/codingpalette/gin_crud/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func setupRouter() *gin.Engine {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBNAME")))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("DB connected", db)

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
