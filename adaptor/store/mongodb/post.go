package store

import (
	"context"
	"time"

	"github.com/mohammadmokh/writino/adaptor/store/mongodb/models"
	"github.com/mohammadmokh/writino/contract"
	"github.com/mohammadmokh/writino/entity"
	"github.com/mohammadmokh/writino/errors"
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

func (m MongodbStore) FindPostsByUserID(ctx context.Context, filters contract.SearchPostFilters) (
	contract.SearchPostRes, error) {

	coll := m.db.Collection("posts")

	var dbModels []models.Post
	var posts []entity.Post

	userObjID, err := primitive.ObjectIDFromHex(filters.Query)
	if err != nil {
		return contract.SearchPostRes{}, err
	}

	limit := int64(filters.Limit)
	skip := int64(filters.Page*filters.Limit - filters.Limit)
	fOpts := options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	}
	filter := bson.D{{"author_id", userObjID}}

	count, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		return contract.SearchPostRes{}, err
	}
	if count == 0 {
		return contract.SearchPostRes{}, errors.ErrNotFound
	}

	cur, err := coll.Find(ctx, filter, &fOpts)

	if err != nil {
		return contract.SearchPostRes{}, err
	}
	err = cur.All(ctx, &dbModels)

	for i := 0; i < len(dbModels); i++ {
		posts = append(posts, models.MapToPostEntity(dbModels[i]))
	}
	return contract.SearchPostRes{
		Posts: posts,
		Count: int(count),
	}, err
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

func (m MongodbStore) SearchPost(ctx context.Context, filters contract.SearchPostFilters) (contract.SearchPostRes, error) {

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

	count, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		return contract.SearchPostRes{}, err
	}
	if count == 0 {
		return contract.SearchPostRes{}, errors.ErrNotFound
	}

	cur, err := coll.Find(ctx, filter, &fOpts)

	if err != nil {
		return contract.SearchPostRes{}, err
	}
	err = cur.All(ctx, &dbModels)

	for i := 0; i < len(dbModels); i++ {
		posts = append(posts, models.MapToPostEntity(dbModels[i]))
	}

	return contract.SearchPostRes{
		Posts: posts,
		Count: int(count),
	}, err
}

func (m MongodbStore) FindAll(ctx context.Context, filters contract.SearchPostFilters) (contract.SearchPostRes, error) {

	coll := m.db.Collection("posts")

	var dbModels []models.Post
	var posts []entity.Post

	limit := int64(filters.Limit)
	skip := int64(filters.Page*filters.Limit - filters.Limit)
	fOpts := options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	}

	switch filters.Query {

	case "newest":
		fOpts.SetSort(bson.D{{"created_at", -1}})

	case "oldest":
		fOpts.SetSort(bson.D{{"created_at", 1}})
	}

	count, err := coll.CountDocuments(ctx, bson.D{{}})
	if err != nil {
		return contract.SearchPostRes{}, err
	}
	if count == 0 {
		return contract.SearchPostRes{}, errors.ErrNotFound
	}

	cur, err := coll.Find(ctx, bson.D{{}}, &fOpts)

	if err != nil {
		return contract.SearchPostRes{}, err
	}
	err = cur.All(ctx, &dbModels)

	if len(dbModels) == 0 {
		return contract.SearchPostRes{}, errors.ErrNotFound
	}

	for i := 0; i < len(dbModels); i++ {
		posts = append(posts, models.MapToPostEntity(dbModels[i]))
	}
	return contract.SearchPostRes{
		Posts: posts,
		Count: int(count),
	}, err

}

func (m MongodbStore) LikePost(ctx context.Context, postID string, userID string) error {

	coll := m.db.Collection("posts")

	postObjID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return err
	}
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": postObjID}
	res, err := coll.UpdateOne(ctx, filter, bson.M{"$push": bson.M{"likes": userObjID}})
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return errors.ErrNotFound
	}

	return nil

}

func (m MongodbStore) DeletePostsByUserID(ctx context.Context, userID string) error {

	coll := m.db.Collection("posts")

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	filter := bson.M{"author_id": userObjID}
	_, err = coll.DeleteMany(ctx, filter)
	return err
}
