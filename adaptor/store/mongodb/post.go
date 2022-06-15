package store

import (
	"context"
	"time"

	"gitlab.com/gocastsian/writino/adaptor/store/mongodb/models"
	"gitlab.com/gocastsian/writino/contract"
	"gitlab.com/gocastsian/writino/entity"
	"gitlab.com/gocastsian/writino/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m MongodbStore) CreatePost(ctx context.Context, post entity.Post) (entity.Post, error) {

	coll := m.db.Collection("posts")

	dbModel := models.MapFromPostEntity(post)

	dbModel.Id = primitive.NewObjectID()
	dbModel.CreatedAt = time.Now()
	_, err := coll.InsertOne(ctx, dbModel)
	return models.MapToPostEntity(dbModel), err
}

func (m MongodbStore) FindPostByID(ctx context.Context, id string) (entity.Post, error) {

	coll := m.db.Collection("posts")

	var post models.Post

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entity.Post{}, err
	}

	filter := bson.D{{"_id", objID}}
	res := coll.FindOne(ctx, filter)

	if res.Err() == mongo.ErrNoDocuments {
		return entity.Post{}, errors.ErrNotFound
	}

	err = res.Decode(&post)
	return models.MapToPostEntity(post), err
}

func (m MongodbStore) FindPostsByUserID(ctx context.Context, filters contract.SearchPostFilters) ([]entity.Post, error) {

	coll := m.db.Collection("posts")

	var dbModels []models.Post
	var posts []entity.Post

	userObjID, err := primitive.ObjectIDFromHex(filters.Query)
	if err != nil {
		return nil, err
	}

	limit := int64(filters.Limit)
	skip := int64(filters.Page*filters.Limit - filters.Limit)
	fOpts := options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	}
	filter := bson.D{{"author_id", userObjID}}
	cur, err := coll.Find(ctx, filter, &fOpts)

	if err != nil {
		return nil, err
	}
	err = cur.All(ctx, &dbModels)

	if len(dbModels) == 0 {
		return nil, errors.ErrNotFound
	}

	for i := 0; i < len(dbModels); i++ {
		posts = append(posts, models.MapToPostEntity(dbModels[i]))
	}
	return posts, err
}

func (m MongodbStore) UpdatePost(ctx context.Context, post entity.Post) error {

	coll := m.db.Collection("posts")

	dbModel := models.MapFromPostEntity(post)

	dbModel.UpdatedAt = time.Now()

	filter := bson.D{{"_id", dbModel.Id}}
	_, err := coll.ReplaceOne(ctx, filter, dbModel)
	return err
}

func (m MongodbStore) DeletePost(ctx context.Context, id string) error {

	coll := m.db.Collection("posts")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", objID}}
	res, err := coll.DeleteOne(ctx, filter)
	if res.DeletedCount != 1 {
		return errors.ErrNotFound
	}
	return err
}

func (m MongodbStore) SearchPost(ctx context.Context, filters contract.SearchPostFilters) ([]entity.Post, error) {

	coll := m.db.Collection("posts")

	var dbModels []models.Post
	var posts []entity.Post

	limit := int64(filters.Limit)
	skip := int64(filters.Page*filters.Limit - filters.Limit)
	fOpts := options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	}
	filter := bson.M{
		"$or": []bson.M{
			{"$text": bson.M{"$search": filters.Query}},
		},
	}

	cur, err := coll.Find(ctx, filter, &fOpts)

	if err != nil {
		return nil, err
	}
	err = cur.All(ctx, &dbModels)

	if len(dbModels) == 0 {
		return nil, errors.ErrNotFound
	}

	for i := 0; i < len(dbModels); i++ {
		posts = append(posts, models.MapToPostEntity(dbModels[i]))
	}
	return posts, err
}
