package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	gorm.Model
	Text   string
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

func dbUpdate(id int, text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず(dbupdate)")
	}
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

func dbDelete(id int) {
	fmt.Println("hello")
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず dbdelete")
	}
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

func dbGetAll() []Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず dbgetall")
	}
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

func dbGetOne(id int) Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず db get one")
	}
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("web/template/*")
	dbInit()

	engine.GET("/", func(c *gin.Context) {
		todos := dbGetAll()
		c.HTML(http.StatusOK, "index.html", gin.H{
			"todos": todos,
		})
	})

	engine.POST("/new", func(c *gin.Context) {
		text := c.PostForm("text")
		status := c.PostForm("status")
		dbInsert(text, status)
		c.Redirect(302, "/")
	})

	engine.GET("/detail/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := dbGetOne(id)
		c.HTML(http.StatusOK, "detail.html", gin.H{
			"todo": todo,
		})
	})

	engine.POST("/update/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		text := c.PostForm("text")
		status := c.PostForm("status")
		dbUpdate(id, text, status)
		c.Redirect(302, "/")
	})

	engine.GET("/delete_check/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		todo := dbGetOne(id)
		c.HTML(200, "delete.html", gin.H{
			"todo": todo,
		})
	})

	engine.POST("/delete/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERRPR")
		}
		dbDelete(id)
		c.Redirect(302, "/")
	})

	engine.Run(":3000")
}
