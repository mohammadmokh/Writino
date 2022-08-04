package auth

import (
	"context"

	"github.com/mohammadmokh/writino/contract"
	"github.com/mohammadmokh/writino/dto"
	"github.com/mohammadmokh/writino/errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthInteractor struct {
	store          contract.AuthStore
	secret         []byte
	tokenGen       contract.GenerateTokenPair
	refTokenParser contract.ParseRefToken
}

func New(store contract.AuthStore, secret []byte,
	tokenGen contract.GenerateTokenPair, refTokenParser contract.ParseRefToken) contract.AuthInteractor {
	return AuthInteractor{
		store:          store,
		secret:         secret,
		tokenGen:       tokenGen,
		refTokenParser: refTokenParser,
	}
}

func (i AuthInteractor) Login(ctx context.Context, req dto.LoginReq) (dto.LoginResponse, error) {

	user, err := i.store.FindUserByEmail(ctx, req.Email)
	if err != nil {
		if err == errors.ErrNotFound {
			return dto.LoginResponse{}, errors.ErrInvalidCredentials
		}
		return dto.LoginResponse{}, err
	}

	if !user.IsVerified {
		return dto.LoginResponse{}, errors.ErrInvalidCredentials
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return dto.LoginResponse{}, errors.ErrInvalidCredentials
	}

	tokenPair, err := i.tokenGen(i.secret, user)
	return dto.LoginResponse{
		Token:    tokenPair["access_token"],
		RefToken: tokenPair["refresh_token"],
	}, err
}

func (i AuthInteractor) RefreshToken(ctx context.Context, req dto.RefreshReq) (dto.RefreshResponse, error) {

	id, err := i.refTokenParser(i.secret, req.RefToken)
	if err != nil {
		return dto.RefreshResponse{}, err
	}
	user, err := i.store.FindUser(ctx, id)
	if err != nil {
		if err == errors.ErrNotFound {
			return dto.RefreshResponse{}, errors.ErrInvalidToken
		}
		return dto.RefreshResponse{}, err
	}

	tokenPair, err := i.tokenGen(i.secret, user)
	return dto.RefreshResponse{
		Token:    tokenPair["access_token"],
		RefToken: tokenPair["refresh_token"],
	}, err
}
