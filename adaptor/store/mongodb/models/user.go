package models

import (
	"time"

	"gitlab.com/gocastsian/writino/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID `bson:"_id"`
	Password    string             `bson:"password"`
	Username    string             `bson:"username"`
	DisplayName string             `bson:"display_name"`
	ProfilePic  string             `bson:"profile_pic,omitempty"`
	Bio         string             `bson:"bio,omitempty"`
	Email       string             `bson:"email"`
	IsVerified  bool               `bson:"is_verified"`
	CreatedAt   time.Time          `bson:"created_at"`
}

func MapFromUserEntity(user entity.User) User {

	ObjID, _ := primitive.ObjectIDFromHex(user.Id)

	return User{
		Id:          ObjID,
		Password:    user.Password,
		Username:    user.Username,
		DisplayName: user.DisplayName,
		ProfilePic:  user.ProfilePic,
		Bio:         user.Bio,
		Email:       user.Email,
		CreatedAt:   user.CreatedAt,
		IsVerified:  user.IsVerified,
	}
}

func MapToUserEntity(user User) entity.User {

	strID := user.Id.Hex()

	return entity.User{
		Id:          strID,
		Password:    user.Password,
		Username:    user.Username,
		DisplayName: user.DisplayName,
		ProfilePic:  user.ProfilePic,
		Bio:         user.Bio,
		Email:       user.Email,
		CreatedAt:   user.CreatedAt,
		IsVerified:  user.IsVerified,
	}
}
