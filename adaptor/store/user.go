package store

import (
	"context"

	"gitlab.com/gocastsian/writino/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m MongodbStore) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {

	coll := m.db.Collection("users")

	user.Id = primitive.NewObjectID()
	res, err := coll.InsertOne(ctx, user)
	if err != nil {
		return entity.User{}, err
	}
	user.Id = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (m MongodbStore) FindUser(ctx context.Context, id string) (entity.User, error) {

	var user entity.User
	coll := m.db.Collection("users")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}
	filter := bson.D{{"_id", objID}}
	res := coll.FindOne(ctx, filter, nil)
	err = res.Decode(&user)
	return user, err
}

func (m MongodbStore) UpdateUser(ctx context.Context, user entity.User) error {

	coll := m.db.Collection("users")

	filter := bson.D{{"_id", user.Id}}
	_, err := coll.ReplaceOne(ctx, filter, user, nil)
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
