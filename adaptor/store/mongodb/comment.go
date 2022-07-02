package store

import (
	"context"

	"gitlab.com/gocastsian/writino/adaptor/store/mongodb/models"
	"gitlab.com/gocastsian/writino/contract"
	"gitlab.com/gocastsian/writino/entity"
	"gitlab.com/gocastsian/writino/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m MongodbStore) CreateComment(ctx context.Context, comment entity.Comment, postID string) error {

	coll := m.db.Collection("comments")

	dbModel := models.MapFromCommentEntity(comment)

	dbModel.Id = primitive.NewObjectID()
	_, err := coll.InsertOne(ctx, dbModel)
	if err != nil {
		return err
	}

	coll = m.db.Collection("posts")
	postObjID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", postObjID}}
	update := bson.D{{"$push", bson.D{{"comments", dbModel.Id}}}}
	_, err = coll.UpdateOne(ctx, filter, update)
	return err
}

func (m MongodbStore) FindCommentsByPostID(ctx context.Context, filters contract.FindCommentfilters) (
	contract.FindCommentRes, error) {

	coll := m.db.Collection("posts")

	type commentID struct {
		Comments []primitive.ObjectID `bson:"comments"`
	}

	var commentIDs commentID
	var dbModels []models.Comment
	var comments []entity.Comment

	postObjID, err := primitive.ObjectIDFromHex(filters.PostID)
	if err != nil {
		return contract.FindCommentRes{}, err
	}
	filter := bson.D{{"_id", postObjID}}
	opts := options.FindOne().SetProjection(bson.D{{"comments", 1}, {"_id", 0}})
	res := coll.FindOne(ctx, filter, opts)

	err = res.Decode(&commentIDs)
	if err != nil {
		return contract.FindCommentRes{}, err
	}
	coll = m.db.Collection("comments")
	limit := int64(filters.Limit)
	skip := int64(filters.Page*filters.Limit - filters.Limit)
	fOpts := options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	}
	filter = bson.D{{"_id", bson.D{{"$in", commentIDs.Comments}}}}
	cur, err := coll.Find(ctx, filter, &fOpts)
	if err != nil {
		return contract.FindCommentRes{}, err
	}
	count, err := coll.CountDocuments(ctx, bson.D{{}})
	if err != nil {
		return contract.FindCommentRes{}, err
	}
	if count == 0 {
		return contract.FindCommentRes{}, errors.ErrNotFound
	}

	err = cur.All(ctx, &dbModels)

	for i := 0; i < len(dbModels); i++ {
		comments = append(comments, models.MapToCommentEntity(dbModels[i]))
	}

	return contract.FindCommentRes{
		Comments:   comments,
		TotalCount: int(count),
	}, err
}
