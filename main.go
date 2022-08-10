package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	engine:= gin.Default()
	// engine.GET("/someGet", getting)
	engine.LoadHTMLGlob("web/template/*")
	engine.GET("/fff", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "hello world",
		})
	})
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "wwwwww",
		})
	})
	engine.Run(":3000")
}

git remote add origin https://github.com/kenboo0426/gin_study.git