package dto

import (
	"time"
)

type CreatePostReq struct {
	AuthorID string   `json:"-"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Tags     []string `json:"tags"`
}

type CreatePostRes struct {
	Post string `json:"post"`
}

type UpdatePostReq struct {
	AuthorID string    `json:"-"`
	Id       string    `json:"id"`
	Title    *string   `json:"title"`
	Content  *string   `json:"content"`
	Tags     *[]string `json:"tags"`
}

type FindPostByIDReq struct {
	UserID string
	ID     string
}

type FindPostRes struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Author     string    `json:"author"`
	Tags       []string  `json:"tags"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	LikesCount uint      `json:"likes_count"`
	IsLiked    bool      `json:"is_liked"`
}

type DeletePostReq struct {
	ID     string
	UserID string
}

type SearchPostReq struct {
	Query string
	Limit int
	Page  int
}

type SummaryPostRes struct {
	ID          string    `json:"-"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Author      string    `json:"author,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	Link        string    `json:"link"`
}

type SearchPostRes struct {
	TotalCount int              `json:"total"`
	Posts      []SummaryPostRes `json:"posts"`
}

type FindUsersPostsReq struct {
	UserID string
	Limit  int
	Page   int
}
