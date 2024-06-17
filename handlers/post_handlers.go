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

type PostHandler struct {
    PostService *services.PostService
}

func NewPostHandler(postService *services.PostService) *PostHandler {
    return &PostHandler{PostService: postService}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
    var post models.BlogPost
    if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
        utils.RespondWithError(w, utils.NewAppError("Invalid request payload", http.StatusBadRequest))
        return
    }
    utils.RespondWithJSON(w, http.StatusCreated, post)
}

func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
    posts, err := h.PostService.GetPosts()
    if err != nil {
        utils.RespondWithError(w, utils.NewAppError("Failed to fetch posts", http.StatusInternalServerError))
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, posts)
}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        utils.RespondWithError(w, utils.NewAppError("Invalid post ID", http.StatusBadRequest))
        return
    }
    post, err := h.PostService.GetPost(id)
    if err != nil {
        utils.RespondWithError(w, utils.NewAppError("Failed to fetch post", http.StatusInternalServerError))
        return
    }
    if post == nil {
        utils.RespondWithError(w, utils.NewAppError("Post not found", http.StatusNotFound))
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, post)
}

func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
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
    if err := h.PostService.UpdatePost(&post); err != nil {
        utils.RespondWithError(w, utils.NewAppError("Failed to update post", http.StatusInternalServerError))
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, post)
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        utils.RespondWithError(w, utils.NewAppError("Invalid post ID", http.StatusBadRequest))
        return
    }
    if err := h.PostService.DeletePost(id); err != nil {
        utils.RespondWithError(w, utils.NewAppError("Failed to delete post", http.StatusInternalServerError))
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
