package handlers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "go-blog/models"
    "go-blog/services"

    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
)

type mockPostService struct {
    posts []models.Post
}

func (m *mockPostService) CreatePost(post *models.Post) (*models.Post, error) {
    post.ID = 1
    m.posts = append(m.posts, *post)
    return post, nil
}

func (m *mockPostService) GetPosts() ([]models.Post, error) {
    return m.posts, nil
}

func TestCreatePost(t *testing.T) {
    mockService := &mockPostService{}
    handler := NewPostHandler(mockService)

    router := mux.NewRouter()
    router.HandleFunc("/posts", handler.CreatePost).Methods("POST")

    newPost := models.Post{Title: "Test Title", Content: "Test Content"}
    jsonPost, _ := json.Marshal(newPost)

    req, err := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonPost))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusCreated, rr.Code, "Expected status 201 Created")

    var createdPost models.Post
    err = json.NewDecoder(rr.Body).Decode(&createdPost)
    if err != nil {
        t.Fatal(err)
    }

    assert.Equal(t, newPost.Title, createdPost.Title, "Expected title to match")
    assert.Equal(t, newPost.Content, createdPost.Content, "Expected content to match")
    assert.Equal(t, 1, createdPost.ID, "Expected ID to be 1")
}

func TestGetPosts(t *testing.T) {
    mockService := &mockPostService{
        posts: []models.Post{
            {ID: 1, Title: "Test Title", Content: "Test Content"},
        },
    }
    handler := NewPostHandler(mockService)

    router := mux.NewRouter()
    router.HandleFunc("/posts", handler.GetPosts).Methods("GET")

    req, err := http.NewRequest("GET", "/posts", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code, "Expected status 200 OK")

    var posts []models.Post
    err = json.NewDecoder(rr.Body).Decode(&posts)
    if err != nil {
        t.Fatal(err)
    }

    assert.Len(t, posts, 1, "Expected length of posts to be 1")
    assert.Equal(t, mockService.posts[0].Title, posts[0].Title, "Expected title to match")
    assert.Equal(t, mockService.posts[0].Content, posts[0].Content, "Expected content to match")
}
