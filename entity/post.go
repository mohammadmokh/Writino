package entity

import (
	"time"
)

type Post struct {
	Id        string
	Title     string
	Content   string
	AuthorID  string
	Tags      []string
	CreatedAt time.Time
	UpdatedAt time.Time
	Likes     []string
}
