package services

import (
    "go-blog/models"
    "go-blog/repositories"
)

type CommentService struct {
    CommentRepo *repositories.CommentRepository
}

func NewCommentService(commentRepo *repositories.CommentRepository) *CommentService {
    return &CommentService{CommentRepo: commentRepo}
}

func (s *CommentService) CreateComment(comment *models.Comment) error {
    return s.CommentRepo.Create(comment)
}

func (s *CommentService) GetComments(postID int) ([]models.Comment, error) {
    return s.CommentRepo.GetAllByPostID(postID)
}

func (s *CommentService) GetComment(id int) (*models.Comment, error) {
    return s.CommentRepo.GetByID(id)
}

func (s *CommentService) UpdateComment(comment *models.Comment) error {
    return s.CommentRepo.Update(comment)
}

func (s *CommentService) DeleteComment(id int) error {
    return s.CommentRepo.Delete(id)
}
