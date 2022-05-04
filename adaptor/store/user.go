package store

import (
	"context"

	"gitlab.com/gocastsian/writino/adaptor/store/models"
	"gitlab.com/gocastsian/writino/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m MongodbStore) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {

	coll := m.db.Collection("users")

	dbModel := models.MapFromUserEntity(user)
	dbModel.Id = primitive.NewObjectID()

	_, err := coll.InsertOne(ctx, dbModel)
	if err != nil {
		return entity.User{}, err
	}
	return models.MapToUserEntity(dbModel), nil
}

func (m MongodbStore) FindUser(ctx context.Context, id string) (entity.User, error) {

	coll := m.db.Collection("users")

	var user models.User

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entity.User{}, err
	}
	filter := bson.D{{"_id", objID}}
	res := coll.FindOne(ctx, filter, nil)
	err = res.Decode(&user)
	return models.MapToUserEntity(user), err
}

func (m MongodbStore) UpdateUser(ctx context.Context, user entity.User) error {

	coll := m.db.Collection("users")
	dbModel := models.MapFromUserEntity(user)

	filter := bson.D{{"_id", dbModel.Id}}
	_, err := coll.ReplaceOne(ctx, filter, dbModel, nil)
	return err
}

func (m MongodbStore) DeleteUser(ctx context.Context, id string) error {

	coll := m.db.Collection("users")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", objID}}
	_, err = coll.DeleteOne(ctx, filter, nil)
	return err
}
