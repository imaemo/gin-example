package main

import (
	"github.com/gin-gonic/gin"

	"os"
	"io"
)


func main() {
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	errorFile, _ := os.Create("error.log")

	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(errorFile)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8888") // listen and serve on 0.0.0.0:8080
	
}