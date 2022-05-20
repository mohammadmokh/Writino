package user

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"gitlab.com/gocastsian/writino/contract"
	"gitlab.com/gocastsian/writino/dto"
	"gitlab.com/gocastsian/writino/entity"
	"gitlab.com/gocastsian/writino/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserIntractor struct {
	store      contract.UserStore
	mail       contract.EmailService
	verCodeGen contract.VerificationCodeGen
	parseTempl contract.ParseVerificationTempl
}

func New(store contract.UserStore, mail contract.EmailService, verCodeGen contract.VerificationCodeGen,
	parseTempl contract.ParseVerificationTempl) contract.UserInteractor {
	return UserIntractor{
		store:      store,
		mail:       mail,
		verCodeGen: verCodeGen,
		parseTempl: parseTempl,
	}
}

func (i UserIntractor) Register(ctx context.Context, req dto.RegisterReq, validator contract.ValidateRegisterUser) error {

	err := validator(req)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	uuid := strings.Replace(uuid.New().String(), "-", "", -1)
	verCode, err := i.verCodeGen()
	if err != nil {
		return err
	}

	user := entity.User{

		Email:            req.Email,
		Password:         string(hashedPassword),
		Username:         uuid,
		DisplayName:      req.Email,
		VerificationCode: verCode,
		IsVerified:       false,
	}

	err = i.store.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	body, err := i.parseTempl(user)
	if err != nil {
		return err
	}
	err = i.mail.SendEmail(user.Email, "Verification Code", body)
	return err
}

func (i UserIntractor) CheckUsername(ctx context.Context, req dto.CheckUsernameReq) (dto.CheckUsernameRes, error) {

	_, err := i.store.FindUserByUsername(ctx, req.Username)
	if err != nil {
		if err == errors.ErrNotFound {
			return dto.CheckUsernameRes{IsUnique: true}, nil
		}
		return dto.CheckUsernameRes{}, err
	}
	return dto.CheckUsernameRes{IsUnique: false}, nil
}

func (i UserIntractor) CheckEmail(ctx context.Context, req dto.CheckEmailReq) (dto.CheckEmailRes, error) {

	_, err := i.store.FindUserByEmail(ctx, req.Email)
	if err != nil {
		if err == errors.ErrNotFound {
			return dto.CheckEmailRes{IsUnique: true}, nil
		}
		return dto.CheckEmailRes{}, err
	}
	return dto.CheckEmailRes{IsUnique: false}, nil
}

func (i UserIntractor) Update(ctx context.Context, req dto.UpdateUserReq, validator contract.ValidateUpdateUser) error {

	err := validator(req)
	if err != nil {
		return err
	}

	user, err := i.store.FindUser(ctx, req.ID)
	if err != nil {
		return err
	}

	if req.Bio != nil {
		user.Bio = *req.Bio
	}
	if req.DisplayName != nil {
		user.DisplayName = *req.DisplayName
	}
	if req.Username != nil {
		user.Username = *req.Username
	}
	if req.ProfilePic != nil {
		user.ProfilePic = *req.ProfilePic
	}

	err = i.store.UpdateUser(ctx, user)
	return err
}

func (i UserIntractor) DeleteAccount(ctx context.Context, req dto.DeleteUserReq) error {

	err := i.store.DeleteUser(ctx, req.Id)
	return err
}

func (i UserIntractor) Find(ctx context.Context, req dto.FindUserReq) (dto.FindUserRes, error) {

	user, err := i.store.FindUserByUsername(ctx, req.Username)
	if err != nil {
		return dto.FindUserRes{}, err
	}
	return dto.FindUserRes{
		ProfilePic:  user.ProfilePic,
		DisplayName: user.DisplayName,
		Bio:         user.Bio,
		Email:       user.Email,
	}, nil
}

func (i UserIntractor) UpdatePassword(ctx context.Context, req dto.UpdatePasswordReq,
	validator contract.ValidateUpdatePassword) error {

	err := validator(req)
	if err != nil {
		return err
	}

	if req.New != req.Old {
		return errors.ErrInvalidCredentials
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.New), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user, err := i.store.FindUser(ctx, req.ID)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	err = i.store.UpdateUser(ctx, user)
	return err
}

func (i UserIntractor) VerifyUser(ctx context.Context, req dto.VerifyUserReq) error {

	user, err := i.store.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return err
	}
	if user.VerificationCode != req.VerificationCode {
		return errors.ErrInvalidCredentials
	}

	user.IsVerified = true
	verCode, err := i.verCodeGen()
	if err != nil {
		return err
	}
	user.VerificationCode = verCode

	err = i.store.UpdateUser(ctx, user)
	return err
}
