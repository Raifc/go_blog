package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"

    "go-blog/models"
    "go-blog/services"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
    var post models.BlogPost
    if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := services.CreatePost(&post); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    respondWithJSON(w, http.StatusCreated, post)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
    posts, err := services.GetPosts()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    respondWithJSON(w, http.StatusOK, posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/posts/"))
    if err != nil {
        http.NotFound(w, r)
        return
    }
    post, err := services.GetPost(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if post == nil {
        http.NotFound(w, r)
        return
    }
    respondWithJSON(w, http.StatusOK, post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/posts/"))
    if err != nil {
        http.NotFound(w, r)
        return
    }
    var post models.BlogPost
    if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    post.ID = id
    if err := services.UpdatePost(&post); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    respondWithJSON(w, http.StatusOK, post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/posts/"))
    if err != nil {
        http.NotFound(w, r)
        return
    }
    if err := services.DeletePost(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
    response, err := json.Marshal(payload)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write(response)
}
