package services

import (
    "go-blog/models"
    "go-blog/repositories"
)

type PostService struct {
    PostRepo *repositories.PostRepository
}

func NewPostService(postRepo *repositories.PostRepository) *PostService {
    return &PostService{PostRepo: postRepo}
}

func (s *PostService) CreatePost(post *models.BlogPost) error {
    return s.PostRepo.Create(post)
}

func (s *PostService) GetPosts() ([]models.BlogPost, error) {
    return s.PostRepo.GetAll()
}

func (s *PostService) GetPost(id int) (*models.BlogPost, error) {
    return s.PostRepo.GetByID(id)
}

func (s *PostService) UpdatePost(post *models.BlogPost) error {
    return s.PostRepo.Update(post)
}

func (s *PostService) DeletePost(id int) error {
    return s.PostRepo.Delete(id)
}
