package verificationCode

import (
	"context"

	"github.com/mohammadmokh/writino/contract"
	"github.com/mohammadmokh/writino/entity"
)

type verificationCodeInteractor struct {
	store  contract.VerficationCodeStore
	random contract.Random
	parser contract.ParseVerificationEmailTmpl
}

func New(store contract.VerficationCodeStore, random contract.Random,
	parser contract.ParseVerificationEmailTmpl) contract.VerificationCodeInteractor {

	return verificationCodeInteractor{
		store:  store,
		random: random,
		parser: parser,
	}
}

func (v verificationCodeInteractor) Find(ctx context.Context, email string) (string, error) {
	code, err := v.store.FindVerCode(ctx, email)
	if err != nil {
		return "", err
	}
	return code, nil
}

func (v verificationCodeInteractor) Create(ctx context.Context, email string) (string, error) {

	//create a six digit code
	code, err := v.random()
	if err != nil {
		return "", err
	}
	err = v.store.CreateVerCode(ctx, entity.VerificationCode{
		Email: email,
		Code:  code,
	})
	if err != nil {
		return "", err
	}

	//parse a html template to send a email
	Templ, err := v.parser(entity.VerificationCode{
		Email: email,
		Code:  code,
	})
	return Templ, err
}
