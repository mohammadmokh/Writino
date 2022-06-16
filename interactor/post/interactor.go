package post

import (
	"context"

	"gitlab.com/gocastsian/writino/contract"
	"gitlab.com/gocastsian/writino/dto"
	"gitlab.com/gocastsian/writino/entity"
	"gitlab.com/gocastsian/writino/errors"
)

type PostInteractor struct {
	store contract.PostStore
}

func New(store contract.PostStore) contract.PostInteractor {
	return PostInteractor{
		store: store,
	}
}

func (i PostInteractor) CreatePost(ctx context.Context, req dto.CreatePostReq) (dto.CreatePostRes, error) {

	post, err := i.store.CreatePost(ctx, entity.Post{
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: req.AuthorID,
		Tags:     req.Tags,
	})

	return dto.CreatePostRes{
		Post: post.Id,
	}, err
}

func (i PostInteractor) UpdatePost(ctx context.Context, req dto.UpdatePostReq) error {

	post, err := i.store.FindPostByID(ctx, req.Id)
	if err != nil {
		return err
	}

	if post.AuthorID != req.AuthorID {
		return errors.ErrNotFound
	}

	if req.Content != nil {
		post.Content = *req.Content
	}
	if req.Tags != nil {
		post.Tags = *req.Tags
	}
	if req.Title != nil {
		post.Title = *req.Title
	}

	err = i.store.UpdatePost(ctx, post)
	return err
}

func (i PostInteractor) FindPostByID(ctx context.Context, req dto.FindPostByIDReq) (dto.FindPostRes, error) {

	post, err := i.store.FindPostByID(ctx, req.ID)
	if err != nil {
		return dto.FindPostRes{}, err
	}

	res := dto.FindPostRes{
		Title:      post.Title,
		Content:    post.Content,
		Tags:       post.Tags,
		CreatedAt:  post.CreatedAt,
		LikesCount: uint(len(post.Likes)),
		IsLiked:    false,
		UpdatedAt:  post.UpdatedAt,
		Author:     post.AuthorID,
	}

	for _, userID := range post.Likes {
		if userID == req.UserID {
			res.IsLiked = true
		}
	}

	return res, nil
}

func (i PostInteractor) DeletePost(ctx context.Context, req dto.DeletePostReq) error {

	post, err := i.store.FindPostByID(ctx, req.ID)
	if err != nil {
		return err
	}

	if post.AuthorID != req.UserID {
		return errors.ErrNotFound
	}

	err = i.store.DeletePost(ctx, req.ID)
	return err
}

func (i PostInteractor) SearchPost(ctx context.Context, req dto.SearchPostReq) (dto.SearchPostRes, error) {

	posts, err := i.store.SearchPost(ctx, contract.SearchPostFilters{
		Query: req.Query,
		Limit: req.Limit,
		Page:  req.Page,
	})

	if err != nil {
		return dto.SearchPostRes{}, err
	}

	res := dto.SearchPostRes{
		TotalCount: len(posts),
	}

	for i := 0; i < len(posts); i++ {

		var description string
		if len(posts[i].Content) < 30 {
			description = posts[i].Content[:len(posts[i].Content)]
		} else {
			description = posts[i].Content[:30]
		}
		summary := dto.SummaryPostRes{
			ID:          posts[i].Id,
			Title:       posts[i].Title,
			Description: description,
			Author:      posts[i].AuthorID,
			CreatedAt:   posts[i].CreatedAt,
		}

		res.Posts = append(res.Posts, summary)
	}

	return res, nil
}

func (i PostInteractor) FindUsersPosts(ctx context.Context, req dto.FindUsersPostsReq) (dto.SearchPostRes, error) {

	posts, err := i.store.FindPostsByUserID(ctx, contract.SearchPostFilters{
		Query: req.UserID,
		Limit: req.Limit,
		Page:  req.Page,
	})

	if err != nil {
		return dto.SearchPostRes{}, err
	}

	res := dto.SearchPostRes{
		TotalCount: len(posts),
	}

	for i := 0; i < len(posts); i++ {

		var description string
		if len(posts[i].Content) < 30 {
			description = posts[i].Content[:len(posts[i].Content)]
		} else {
			description = posts[i].Content[:30]
		}
		summary := dto.SummaryPostRes{
			ID:          posts[i].Id,
			Title:       posts[i].Title,
			Description: description,
			CreatedAt:   posts[i].CreatedAt,
		}

		res.Posts = append(res.Posts, summary)
	}

	return res, nil
}
