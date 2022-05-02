package store

import (
	"context"
	"time"

	"gitlab.com/gocastsian/writino/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m MongodbStore) CreatePost(ctx context.Context, post entity.Post) (entity.Post, error) {

	coll := m.db.Collection("posts")
	post.Id = primitive.NewObjectID()
	post.CreatedAt = time.Now()
	res, err := coll.InsertOne(ctx, post)
	post.Id = res.InsertedID.(primitive.ObjectID)
	return post, err
}

func (m MongodbStore) FindPostByID(ctx context.Context, id string) (entity.Post, error) {

	var post entity.Post
	coll := m.db.Collection("posts")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entity.Post{}, err
	}

	filter := bson.D{{"_id", objID}}
	res := coll.FindOne(ctx, filter)
	err = res.Decode(&post)
	return post, err
}

func (m MongodbStore) FindPostsByUserID(ctx context.Context, userID string) ([]entity.Post, error) {

	var posts []entity.Post

	coll := m.db.Collection("posts")

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{"author_id", userObjID}}
	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	err = cur.All(ctx, &posts)
	return posts, err
}

func (m MongodbStore) UpdatePost(ctx context.Context, post entity.Post) error {

	coll := m.db.Collection("posts")
	post.UpdatedAt = time.Now()

	filter := bson.D{{"_id", post.Id}}
	_, err := coll.ReplaceOne(ctx, filter, post)
	return err
}

func (m MongodbStore) DeletePost(ctx context.Context, id string) error {

	coll := m.db.Collection("posts")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", objID}}
	_, err = coll.DeleteOne(ctx, filter)
	return err
}
