package services

import (
    "database/sql"
    "go-blog/db"
    "go-blog/models"
)

func CreatePost(post *models.BlogPost) error {
    stmt, err := db.DB.Prepare("INSERT INTO blogpost (title, content) VALUES ($1, $2) RETURNING id")
    if err != nil {
        return err
    }
    err = stmt.QueryRow(post.Title, post.Content).Scan(&post.ID)
    if err != nil {
        return err
    }
    return nil
}

func GetPosts() ([]models.BlogPost, error) {
    rows, err := db.DB.Query("SELECT id, title, content FROM blogpost")
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

func GetPost(id int) (*models.BlogPost, error) {
    var post models.BlogPost
    err := db.DB.QueryRow("SELECT id, title, content FROM blogpost WHERE id = $1", id).Scan(&post.ID, &post.Title, &post.Content)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return &post, nil
}

func UpdatePost(post *models.BlogPost) error {
    stmt, err := db.DB.Prepare("UPDATE blogpost SET title = $1, content = $2 WHERE id = $3")
    if err != nil {
        return err
    }
    _, err = stmt.Exec(post.Title, post.Content, post.ID)
    if err != nil {
        return err
    }
    return nil
}

func DeletePost(id int) error {
    stmt, err := db.DB.Prepare("DELETE FROM blogpost WHERE id = $1")
    if err != nil {
        return err
    }
    _, err = stmt.Exec(id)
    if err != nil {
        return err
    }
    return nil
}
