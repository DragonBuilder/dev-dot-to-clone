package main

import (
	"html/template"
	"log"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	// r.LoadHTMLGlob("templates/*")
	html := template.Must(template.ParseFiles(
		"templates/index.html",
		"templates/create-blog.html",
	))
	r.SetHTMLTemplate(html)

	r.GET("/", IndexHandler)
	r.GET("/ping", PingHandler)
	r.GET("/blog/create", CreateBlogFormHandler)
	r.POST("/blog/create", CreateBlogHandler)

	slog.Info("Starting server on port 8080")
	r.Run() // listen and serve on 0.0.0.0:8080
}

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func CreateBlogFormHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "create-blog.html", gin.H{})
}

func CreateBlogHandler(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	log.Printf("Title: %s, \nContent:\n%s", title, content)

	// log.Info(fmt.Sprintf("Title: %s, \nContent:\n%s", title, content))
	c.HTML(http.StatusOK, "create-blog.html", gin.H{})
}
