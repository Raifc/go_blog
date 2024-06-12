package repositories

import (
    "database/sql"
    "go-blog/models"
)

type PostRepository struct {
    DB *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
    return &PostRepository{DB: db}
}

func (r *PostRepository) Create(post *models.BlogPost) error {
    stmt, err := r.DB.Prepare("INSERT INTO blogpost(title, content) VALUES($1, $2)")
    if err != nil {
        return err
    }
    res, err := stmt.Exec(post.Title, post.Content)
    if err != nil {
        return err
    }
    id, err := res.LastInsertId()
    if err != nil {
        return err
    }
    post.ID = int(id)
    return nil
}

func (r *PostRepository) GetAll() ([]models.BlogPost, error) {
    rows, err := r.DB.Query("SELECT id, title, content FROM blogpost")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var posts []models.BlogPost
    for rows.Next() {
        var post models.BlogPost
        err := rows.Scan(&post.ID, &post.Title, &post.Content)
        if err != nil {
            return nil, err
        }
        posts = append(posts, post)
    }
    return posts, nil
}

func (r *PostRepository) GetByID(id int) (*models.BlogPost, error) {
    var post models.BlogPost
    err := r.DB.QueryRow("SELECT id, title, content FROM blogpost WHERE id = $1", id).Scan(&post.ID, &post.Title, &post.Content)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return &post, nil
}

func (r *PostRepository) Update(post *models.BlogPost) error {
    stmt, err := r.DB.Prepare("UPDATE blogpost SET title = $1, content = $2 WHERE id = $3")
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
    stmt, err := r.DB.Prepare("DELETE FROM blogpost WHERE id = $1")
    if err != nil {
        return err
    }
    _, err = stmt.Exec(id)
    if err != nil {
        return err
    }
    return nil
}
