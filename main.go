package main

import (
	"github.com/gin-gonic/gin"
  "net/http"
	"github.com/jinzhu/gorm"
	"github.com/mattn/go-sqlite3"
)

type Todo struct {
	gorm.Model
	Text string
	Status string
}

func dbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けづ！(dbInit)")
	}
	db.AutoMigrate(&Todo{})
	defer db.Close()
}

func dbInsert(text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず(dbInsert)")
	}
	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}


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
