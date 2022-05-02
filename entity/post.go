package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id        primitive.ObjectID   `bson:"_id"`
	Title     string               `bson:"title"`
	Content   string               `bson:"content"`
	AuthorID  primitive.ObjectID   `bson:"author_id"`
	Tags      []string             `bson:"tags"`
	CreatedAt time.Time            `bson:"created_at"`
	UpdatedAt time.Time            `bson:"updated_at,omitempty"`
	Likes     []primitive.ObjectID `bson:"likes"`
	Comments  []primitive.ObjectID `bson:"comments"`
}
