package models

type BlogPost struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Comments []Comment `json:"comments,omitempty"`
}

type Comment struct {
	ID      int    `json:"id"`
	PostID  int    `json:"post_id"`
	Content string `json:"content"`
}
