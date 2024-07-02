package repositories

import (
	"database/sql"
	"go-blog/models"

	"github.com/sirupsen/logrus"
)

type PostRepository struct {
	DB *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (r *PostRepository) Create(post *models.Post) error {
    stmt, err := r.DB.Prepare("INSERT INTO posts(title, content) VALUES($1, $2) RETURNING id")
    if err != nil {
        logrus.Errorf("Error preparing statement: %v", err)
        return err
    }
    defer stmt.Close()

    err = stmt.QueryRow(post.Title, post.Content).Scan(&post.ID)
    if err != nil {
        logrus.Errorf("Error executing statement: %v", err)
        return err
    }

    return nil
}

func (r *PostRepository) GetAll() ([]models.Post, error) {
	rows, err := r.DB.Query("SELECT id, title, content FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *PostRepository) GetByID(id int) (*models.Post, error) {
	var post models.Post
	err := r.DB.QueryRow("SELECT id, title, content FROM posts WHERE id = $1", id).Scan(&post.ID, &post.Title, &post.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) Update(post *models.Post) error {
	stmt, err := r.DB.Prepare("UPDATE posts SET title = $1, content = $2 WHERE id = $3")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(post.Title, post.Content, post.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) Delete(id int) error {
	stmt, err := r.DB.Prepare("DELETE FROM posts WHERE id = $1")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
