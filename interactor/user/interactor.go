package user

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"gitlab.com/gocastsian/writino/contract"
	"gitlab.com/gocastsian/writino/dto"
	"gitlab.com/gocastsian/writino/entity"
	"gitlab.com/gocastsian/writino/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserIntractor struct {
	store            contract.UserStore
	mail             contract.EmailService
	image            contract.ImageStore
	verificationCode contract.VerificationCodeInteractor
}

func New(store contract.UserStore, mail contract.EmailService, image contract.ImageStore,
	verificationCode contract.VerificationCodeInteractor) contract.UserInteractor {
	return UserIntractor{
		store:            store,
		mail:             mail,
		image:            image,
		verificationCode: verificationCode,
	}
}

func (i UserIntractor) Register(ctx context.Context, req dto.RegisterReq) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	uuid := strings.Replace(uuid.New().String(), "-", "", -1)

	user := entity.User{

		Email:       req.Email,
		Password:    string(hashedPassword),
		Username:    uuid,
		DisplayName: req.Email,
		IsVerified:  false,
	}

	body, err := i.verificationCode.Create(ctx, user.Email)
	if err != nil {
		return err
	}

	err = i.store.CreateUser(ctx, user)
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

	user, err := i.store.FindUserByEmail(ctx, req.Email)
	if err != nil {
		if err == errors.ErrNotFound {
			return dto.CheckEmailRes{IsUnique: true}, nil
		}
		return dto.CheckEmailRes{}, err
	}

	// if user registered but not verified we delete the user and free email Address
	if !user.IsVerified && (time.Since(user.CreatedAt).Minutes() > 5) {
		err := i.store.DeleteUser(ctx, user.Id)
		if err != nil {
			return dto.CheckEmailRes{}, err
		}
		return dto.CheckEmailRes{IsUnique: true}, nil
	}

	return dto.CheckEmailRes{IsUnique: false}, nil
}

func (i UserIntractor) Update(ctx context.Context, req dto.UpdateUserReq) error {

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

func (i UserIntractor) UpdatePassword(ctx context.Context, req dto.UpdatePasswordReq) error {

	user, err := i.store.FindUser(ctx, req.ID)
	if err != nil {
		return err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Old)) != nil {
		return errors.ErrInvalidCredentials
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.New), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	err = i.store.UpdateUser(ctx, user)
	return err
}

func (i UserIntractor) VerifyUser(ctx context.Context, req dto.VerifyUserReq) error {

	code, err := i.verificationCode.Find(ctx, req.Email)
	if err != nil {
		return err
	}
	if code != req.VerificationCode {
		return errors.ErrInvalidCredentials
	}

	user, err := i.store.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return err
	}
	user.IsVerified = true

	err = i.store.UpdateUser(ctx, user)
	return err
}

func (i UserIntractor) UpdateProfilePic(ctx context.Context, req dto.UpdateProfilePicReq) (
	dto.UpdateProfilePicRes, error) {

	user, err := i.store.FindUser(ctx, req.ID)
	if err != nil {
		return dto.UpdateProfilePicRes{}, err
	}

	filename := user.Id + "." + req.Format
	err = i.image.SaveImage(req.Image, "avatars/"+filename)
	if err != nil {
		return dto.UpdateProfilePicRes{}, err
	}
	user.ProfilePic = filename

	err = i.store.UpdateUser(ctx, user)
	return dto.UpdateProfilePicRes{
		Link: filename,
	}, err
}
