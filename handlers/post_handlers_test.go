package handlers

import (
    "bytes"
    "database/sql"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "os"
    "strconv"
    "testing"

    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
    _ "github.com/mattn/go-sqlite3"
    "go-blog/config"
    "go-blog/db"
    "go-blog/models"
    "go-blog/services"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
    os.Setenv("DB_DRIVER", "sqlite3")
    os.Setenv("DB_NAME", "./test_blog.db")

    config.DBDriver = os.Getenv("DB_DRIVER")
    config.DBName = os.Getenv("DB_NAME")

    db.Init()
    testDB = db.DB

    code := m.Run()

    db.Close()

    os.Remove("./test_blog.db")

    os.Exit(code)
}

func clearTable() {
    testDB.Exec("DELETE FROM blogpost")
    testDB.Exec("DELETE FROM sqlite_sequence WHERE name='blogpost'")
}


func TestCreatePost(t *testing.T) {
    clearTable()

    post := models.BlogPost{Title: "Test Title", Content: "Test Content"}
    payload, _ := json.Marshal(post)

    req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(payload))
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(CreatePost)

    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestGetPosts(t *testing.T) {
    clearTable()
    services.CreatePost(&models.BlogPost{Title: "Title 1", Content: "Content 1"})
    services.CreatePost(&models.BlogPost{Title: "Title 2", Content: "Content 2"})

    req, _ := http.NewRequest("GET", "/posts", nil)
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(GetPosts)

    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
    var posts []models.BlogPost
    err := json.Unmarshal(rr.Body.Bytes(), &posts)
    assert.NoError(t, err)
    assert.Len(t, posts, 2)
}

func TestGetPost(t *testing.T) {
    clearTable()
    post := models.BlogPost{Title: "Test Title", Content: "Test Content"}
    services.CreatePost(&post)

    req, _ := http.NewRequest("GET", "/posts/"+strconv.Itoa(post.ID), nil)
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(GetPost)

    vars := map[string]string{
        "id": strconv.Itoa(post.ID),
    }
    req = mux.SetURLVars(req, vars)

    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
    var fetchedPost models.BlogPost
    err := json.Unmarshal(rr.Body.Bytes(), &fetchedPost)
    assert.NoError(t, err)
    assert.Equal(t, post.Title, fetchedPost.Title)
    assert.Equal(t, post.Content, fetchedPost.Content)
}

func TestUpdatePost(t *testing.T) {
    clearTable()
    post := models.BlogPost{Title: "Initial Title", Content: "Initial Content"}
    services.CreatePost(&post)

    updatedPost := models.BlogPost{Title: "Updated Title", Content: "Updated Content"}
    payload, _ := json.Marshal(updatedPost)

    req, _ := http.NewRequest("PUT", "/posts/"+strconv.Itoa(post.ID), bytes.NewBuffer(payload))
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(UpdatePost)

    vars := map[string]string{
        "id": strconv.Itoa(post.ID),
    }
    req = mux.SetURLVars(req, vars)

    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
    var fetchedPost models.BlogPost
    err := json.Unmarshal(rr.Body.Bytes(), &fetchedPost)
    assert.NoError(t, err)
    assert.Equal(t, updatedPost.Title, fetchedPost.Title)
    assert.Equal(t, updatedPost.Content, fetchedPost.Content)
}

func TestDeletePost(t *testing.T) {
    clearTable()
    post := models.BlogPost{Title: "Test Title", Content: "Test Content"}
    services.CreatePost(&post)

    req, _ := http.NewRequest("DELETE", "/posts/"+strconv.Itoa(post.ID), nil)
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(DeletePost)

    vars := map[string]string{
        "id": strconv.Itoa(post.ID),
    }
    req = mux.SetURLVars(req, vars)

    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusNoContent, rr.Code)
}
