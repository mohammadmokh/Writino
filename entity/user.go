package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID `bson:"_id"`
	Password    string             `bson:"password"`
	Username    string             `bson:"username"`
	DisplayName string             `bson:"display_name"`
	ProfilePic  string             `bson:"profile_pic,omitempty"`
	Bio         string             `bson:"bio,omitempty"`
	Email       string             `bson:"email,omitempty"`
}
