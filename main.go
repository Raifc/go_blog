package main

import (
    "net/http"

    "github.com/gorilla/mux"
    "github.com/sirupsen/logrus"
    "go-blog/config"
    "go-blog/db"
    "go-blog/handlers"
    "go-blog/middlewares"
)

func main() {
    config.Init()
    db.Init()
    defer db.Close()

    logrus.SetFormatter(&logrus.JSONFormatter{})
    logrus.SetLevel(logrus.InfoLevel)

    r := mux.NewRouter()
    r.Use(middlewares.Logger)
    r.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
    r.HandleFunc("/posts", handlers.CreatePost).Methods("POST")
    r.HandleFunc("/posts/{id:[0-9]+}", handlers.GetPost).Methods("GET")
    r.HandleFunc("/posts/{id:[0-9]+}", handlers.UpdatePost).Methods("PUT")
    r.HandleFunc("/posts/{id:[0-9]+}", handlers.DeletePost).Methods("DELETE")

    logrus.Info("Server started at :8080")
    logrus.Fatal(http.ListenAndServe(":8080", r))
}
