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
		post.Tags = req.Tags
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

	searchRes, err := i.store.SearchPost(ctx, contract.SearchPostFilters{
		Query: req.Query,
		Limit: req.Limit,
		Page:  req.Page,
	})

	if err != nil {
		return dto.SearchPostRes{}, err
	}

	res := dto.SearchPostRes{
		TotalCount: searchRes.Count,
	}

	for i := 0; i < len(searchRes.Posts); i++ {

		var description string
		if len(searchRes.Posts[i].Content) < 30 {
			description = searchRes.Posts[i].Content[:len(searchRes.Posts[i].Content)]
		} else {
			description = searchRes.Posts[i].Content[:30]
		}
		summary := dto.SummaryPostRes{
			ID:          searchRes.Posts[i].Id,
			Title:       searchRes.Posts[i].Title,
			Description: description,
			Author:      searchRes.Posts[i].AuthorID,
			CreatedAt:   searchRes.Posts[i].CreatedAt,
		}

		res.Posts = append(res.Posts, summary)
	}

	return res, nil
}

func (i PostInteractor) FindUsersPosts(ctx context.Context, req dto.FindUsersPostsReq) (dto.SearchPostRes, error) {

	findRes, err := i.store.FindPostsByUserID(ctx, contract.SearchPostFilters{
		Query: req.UserID,
		Limit: req.Limit,
		Page:  req.Page,
	})

	if err != nil {
		return dto.SearchPostRes{}, err
	}

	res := dto.SearchPostRes{
		TotalCount: findRes.Count,
	}

	for i := 0; i < len(findRes.Posts); i++ {

		var description string
		if len(findRes.Posts[i].Content) < 30 {
			description = findRes.Posts[i].Content[:len(findRes.Posts[i].Content)]
		} else {
			description = findRes.Posts[i].Content[:30]
		}
		summary := dto.SummaryPostRes{
			ID:          findRes.Posts[i].Id,
			Title:       findRes.Posts[i].Title,
			Description: description,
			Author:      findRes.Posts[i].AuthorID,
			CreatedAt:   findRes.Posts[i].CreatedAt,
		}

		res.Posts = append(res.Posts, summary)
	}

	return res, nil
}

func (i PostInteractor) FindAll(ctx context.Context, req dto.SearchPostReq) (dto.SearchPostRes, error) {

	findRes, err := i.store.FindAll(ctx, contract.SearchPostFilters{
		Query: req.Query,
		Limit: req.Limit,
		Page:  req.Page,
	})

	if err != nil {
		return dto.SearchPostRes{}, err
	}

	res := dto.SearchPostRes{
		TotalCount: findRes.Count,
	}

	for i := 0; i < len(findRes.Posts); i++ {

		var description string
		if len(findRes.Posts[i].Content) < 30 {
			description = findRes.Posts[i].Content[:len(findRes.Posts[i].Content)]
		} else {
			description = findRes.Posts[i].Content[:30]
		}
		summary := dto.SummaryPostRes{
			ID:          findRes.Posts[i].Id,
			Title:       findRes.Posts[i].Title,
			Description: description,
			Author:      findRes.Posts[i].AuthorID,
			CreatedAt:   findRes.Posts[i].CreatedAt,
		}

		res.Posts = append(res.Posts, summary)
	}

	return res, nil

}
