package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
    "go-blog/models"
    "go-blog/services"
    "go-blog/utils"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
    var post models.BlogPost
    if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
        utils.RespondWithError(w, utils.NewAppError("Invalid request payload", http.StatusBadRequest))
        return
    }
    if err := services.CreatePost(&post); err != nil {
        utils.RespondWithError(w, utils.NewAppError("Failed to create post", http.StatusInternalServerError))
        return
    }
    respondWithJSON(w, http.StatusCreated, post)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
    posts, err := services.GetPosts()
    if err != nil {
        utils.RespondWithError(w, utils.NewAppError("Failed to fetch posts", http.StatusInternalServerError))
        return
    }
    respondWithJSON(w, http.StatusOK, posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        utils.RespondWithError(w, utils.NewAppError("Invalid post ID", http.StatusBadRequest))
        return
    }
    post, err := services.GetPost(id)
    if err != nil {
        utils.RespondWithError(w, utils.NewAppError("Failed to fetch post", http.StatusInternalServerError))
        return
    }
    if post == nil {
        utils.RespondWithError(w, utils.NewAppError("Post not found", http.StatusNotFound))
        return
    }
    respondWithJSON(w, http.StatusOK, post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        utils.RespondWithError(w, utils.NewAppError("Invalid post ID", http.StatusBadRequest))
        return
    }
    var post models.BlogPost
    if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
        utils.RespondWithError(w, utils.NewAppError("Invalid request payload", http.StatusBadRequest))
        return
    }
    post.ID = id
    if err := services.UpdatePost(&post); err != nil {
        utils.RespondWithError(w, utils.NewAppError("Failed to update post", http.StatusInternalServerError))
        return
    }
    respondWithJSON(w, http.StatusOK, post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        utils.RespondWithError(w, utils.NewAppError("Invalid post ID", http.StatusBadRequest))
        return
    }
    if err := services.DeletePost(id); err != nil {
        utils.RespondWithError(w, utils.NewAppError("Failed to delete post", http.StatusInternalServerError))
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
    response, err := json.Marshal(payload)
    if err != nil {
        utils.RespondWithError(w, utils.NewAppError("Failed to encode response", http.StatusInternalServerError))
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write(response)
}
