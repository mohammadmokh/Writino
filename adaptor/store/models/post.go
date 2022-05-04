package models

import (
	"time"

	"gitlab.com/gocastsian/writino/entity"
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
}

func MapFromPostEntity(post entity.Post) Post {

	var likes []primitive.ObjectID

	objID, _ := primitive.ObjectIDFromHex(post.Id)
	autherObjID, _ := primitive.ObjectIDFromHex(post.AuthorID)

	for i := 0; i < len(post.Likes); i++ {
		likeObjID, _ := primitive.ObjectIDFromHex(post.Likes[i])
		likes = append(likes, likeObjID)
	}

	return Post{
		Id:        objID,
		Title:     post.Title,
		Content:   post.Content,
		AuthorID:  autherObjID,
		Tags:      post.Tags,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		Likes:     likes,
	}
}

func MapToPostEntity(post Post) entity.Post {

	var likes []string

	for i := 0; i < len(post.Likes); i++ {
		likes = append(likes, post.Likes[i].Hex())
	}

	return entity.Post{
		Id:        post.Id.Hex(),
		Title:     post.Title,
		Content:   post.Content,
		AuthorID:  post.AuthorID.Hex(),
		Tags:      post.Tags,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		Likes:     likes,
	}
}
