package services

import (
    "database/sql"
    "go-blog/db"
    "go-blog/models"
)

func CreatePost(post *models.BlogPost) error {
    stmt, err := db.DB.Prepare("INSERT INTO blogpost(title, content) VALUES(?, ?)")
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
    err := db.DB.QueryRow("SELECT id, title, content FROM blogpost WHERE id = ?", id).Scan(&post.ID, &post.Title, &post.Content)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return &post, nil
}

func UpdatePost(post *models.BlogPost) error {
    stmt, err := db.DB.Prepare("UPDATE blogpost SET title = ?, content = ? WHERE id = ?")
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
    stmt, err := db.DB.Prepare("DELETE FROM blogpost WHERE id = ?")
    if err != nil {
        return err
    }
    _, err = stmt.Exec(id)
    if err != nil {
        return err
    }
    return nil
}
