package main

import (
    "net/http"

    "go-blog/config"
    "go-blog/db"
    "go-blog/handlers"
    "go-blog/middlewares"
    "go-blog/repositories"
    "go-blog/services"

    "github.com/gorilla/mux"
    "github.com/sirupsen/logrus"
    _ "github.com/lib/pq"
)

func main() {
    config.Init()
    db.Init()
    defer db.Close()

    logrus.SetFormatter(&logrus.JSONFormatter{})
    logrus.SetLevel(logrus.InfoLevel)

    postRepo := repositories.NewPostRepository(db.DB)
    postService := services.NewPostService(postRepo)
    postHandler := handlers.NewPostHandler(postService)

    commentRepo := repositories.NewCommentRepository(db.DB)
    commentService := services.NewCommentService(commentRepo)
    commentHandler := handlers.NewCommentHandler(commentService)

    r := mux.NewRouter()
    r.Use(middlewares.Logger)

    r.HandleFunc("/posts", postHandler.CreatePost).Methods("POST")
    r.HandleFunc("/posts", postHandler.GetPosts).Methods("GET")
    r.HandleFunc("/posts/{id:[0-9]+}", postHandler.GetPost).Methods("GET")
    r.HandleFunc("/posts/{id:[0-9]+}", postHandler.UpdatePost).Methods("PUT")
    r.HandleFunc("/posts/{id:[0-9]+}", postHandler.DeletePost).Methods("DELETE")

    r.HandleFunc("/posts/{postID:[0-9]+}/comments", commentHandler.CreateComment).Methods("POST")
    r.HandleFunc("/posts/{postID:[0-9]+}/comments", commentHandler.GetComments).Methods("GET")
    r.HandleFunc("/comments/{id:[0-9]+}", commentHandler.GetComment).Methods("GET")
    r.HandleFunc("/comments/{id:[0-9]+}", commentHandler.UpdateComment).Methods("PUT")
    r.HandleFunc("/comments/{id:[0-9]+}", commentHandler.DeleteComment).Methods("DELETE")

    logrus.Info("Server started at :8080")
    logrus.Fatal(http.ListenAndServe(":8080", r))
}
