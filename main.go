package main

import (
	"fmt"
	"log"
	"net/http"

	"go-blog/db"
	"go-blog/handlers"
)

func main() {
	db.Init()
	defer db.Close()

	http.HandleFunc("/posts", postsHandler)
	http.HandleFunc("/posts/", postHandler)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handlers.GetPosts(w, r)
	case "POST":
		handlers.CreatePost(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handlers.GetPost(w, r)
	case "PUT":
		handlers.UpdatePost(w, r)
	case "DELETE":
		handlers.DeletePost(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
