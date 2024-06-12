package repositories

import (
    "database/sql"
    "go-blog/models"
)

type CommentRepository struct {
    DB *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
    return &CommentRepository{DB: db}
}

func (r *CommentRepository) Create(comment *models.Comment) error {
    stmt, err := r.DB.Prepare("INSERT INTO comments(post_id, content) VALUES($1, $2)")
    if err != nil {
        return err
    }
    res, err := stmt.Exec(comment.PostID, comment.Content)
    if err != nil {
        return err
    }
    id, err := res.LastInsertId()
    if err != nil {
        return err
    }
    comment.ID = int(id)
    return nil
}

func (r *CommentRepository) GetAllByPostID(postID int) ([]models.Comment, error) {
    rows, err := r.DB.Query("SELECT id, post_id, content FROM comments WHERE post_id = $1", postID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var comments []models.Comment
    for rows.Next() {
        var comment models.Comment
        err := rows.Scan(&comment.ID, &comment.PostID, &comment.Content)
        if err != nil {
            return nil, err
        }
        comments = append(comments, comment)
    }
    return comments, nil
}

func (r *CommentRepository) GetByID(id int) (*models.Comment, error) {
    var comment models.Comment
    err := r.DB.QueryRow("SELECT id, post_id, content FROM comments WHERE id = $1", id).Scan(&comment.ID, &comment.PostID, &comment.Content)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return &comment, nil
}

func (r *CommentRepository) Update(comment *models.Comment) error {
    stmt, err := r.DB.Prepare("UPDATE comments SET content = $1 WHERE id = $2")
    if err != nil {
        return err
    }
    _, err = stmt.Exec(comment.Content, comment.ID)
    if err != nil {
        return err
    }
    return nil
}

func (r *CommentRepository) Delete(id int) error {
    stmt, err := r.DB.Prepare("DELETE FROM comments WHERE id = $1")
    if err != nil {
        return err
    }
    _, err = stmt.Exec(id)
    if err != nil {
        return err
    }
    return nil
}
