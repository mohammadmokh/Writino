package models

import (
	"gitlab.com/gocastsian/writino/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	Id     primitive.ObjectID `bson:"_id"`
	UserID primitive.ObjectID `bson:"user_id"`
	Text   string             `bson:"text"`
}

func MapFromCommentEntity(comment entity.Comment) Comment {

	objID, _ := primitive.ObjectIDFromHex(comment.Id)
	userObjID, _ := primitive.ObjectIDFromHex(comment.UserID)

	return Comment{
		Id:     objID,
		UserID: userObjID,
		Text:   comment.Text,
	}
}

func MapToCommentEntity(comment Comment) entity.Comment {

	return entity.Comment{
		Id:     comment.Id.Hex(),
		UserID: comment.UserID.Hex(),
		Text:   comment.Text,
	}
}
