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

type CommentHandler struct {
    CommentService *services.CommentService
}

func NewCommentHandler(commentService *services.CommentService) *CommentHandler {
    return &CommentHandler{CommentService: commentService}
}

func (h *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
    var comment models.Comment
    if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
        utils.RespondWithError(w, utils.NewAppError("Invalid request payload", http.StatusBadRequest))
        return
    }
    vars := mux.Vars(r)
    postID, err := strconv.Atoi(vars["postID"])
    if err != nil {
        utils.RespondWithError(w, utils.NewAppError("Invalid post ID", http.StatusBadRequest))
        return
    }
    comment.PostID = postID
    if err := h.CommentService.CreateComment(&comment); err != nil {
        utils.RespondWithError(w, utils.NewAppError("Failed to create comment", http.StatusInternalServerError))
        return
    }
    utils.RespondWithJSON(w, http.StatusCreated, comment)
}

func (h *CommentHandler) GetComments(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    postID, err := strconv.Atoi(vars["postID"])
    if err != nil {
        utils.RespondWithError(w, utils.NewAppError("Invalid post ID", http.StatusBadRequest))
        return
    }
    comments, err := h.CommentService.GetComments(postID)
    if err != nil {
        utils.RespondWithError(w, utils.NewAppError("Failed to fetch comments", http.StatusInternalServerError))
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, comments)
}

func (h *CommentHandler) GetComment(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        utils.RespondWithError(w, utils.NewAppError("Invalid comment ID", http.StatusBadRequest))
        return
    }
    comment, err := h.CommentService.GetComment(id)
    if err != nil {
        utils.RespondWithError(w, utils.NewAppError("Failed to fetch comment", http.StatusInternalServerError))
        return
    }
    if comment == nil {
        utils.RespondWithError(w, utils.NewAppError("Comment not found", http.StatusNotFound))
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, comment)
}

func (h *CommentHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        utils.RespondWithError(w, utils.NewAppError("Invalid comment ID", http.StatusBadRequest))
        return
    }
    var comment models.Comment
    if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
        utils.RespondWithError(w, utils.NewAppError("Invalid request payload", http.StatusBadRequest))
        return
    }
    comment.ID = id
    if err := h.CommentService.UpdateComment(&comment); err != nil {
        utils.RespondWithError(w, utils.NewAppError("Failed to update comment", http.StatusInternalServerError))
        return
    }
    utils.RespondWithJSON(w, http.StatusOK, comment)
}

func (h *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        utils.RespondWithError(w, utils.NewAppError("Invalid comment ID", http.StatusBadRequest))
        return
    }
    if err := h.CommentService.DeleteComment(id); err != nil {
        utils.RespondWithError(w, utils.NewAppError("Failed to delete comment", http.StatusInternalServerError))
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
