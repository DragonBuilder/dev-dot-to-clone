package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func main() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", IndexHandler).Methods("GET")
// 	r.HandleFunc("/blog/create", CreateBlogHandler).Methods("GET", "POST")

// 	// http.Handle("/", r)

// 	slog.Info("Starting server on port 8080")

// 	if err := http.ListenAndServe(":8080", r); err != nil {
// 		slog.Error("Error starting server: %v", err)
// 	}
// }

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func CreateBlogHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
