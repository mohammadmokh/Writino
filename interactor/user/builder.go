package user

import "gitlab.com/gocastsian/writino/contract"

type UserBuilder struct {
	user UserIntractor
}

func NewBuilder() *UserBuilder {
	return &UserBuilder{
		user: UserIntractor{},
	}
}

func (u *UserBuilder) Build() UserIntractor {
	return u.user
}

func (u *UserBuilder) SetUserStore(store contract.UserStore) *UserBuilder {
	u.user.store = store
	return u
}

func (u *UserBuilder) SetMailService(mail contract.EmailService) *UserBuilder {
	u.user.mail = mail
	return u
}

func (u *UserBuilder) SetProfilePicStore(store contract.ProfilePicStore) *UserBuilder {
	u.user.profilePic = store
	return u
}

func (u *UserBuilder) SetVerCodeService(verification contract.VerificationCodeInteractor) *UserBuilder {
	u.user.verificationCode = verification
	return u
}
func (u *UserBuilder) SetPostSerivce(post PostInteractor) *UserBuilder {
	u.user.postInteractor = post
	return u
}

func (u *UserBuilder) SetCommentService(comment CommentInteractor) *UserBuilder {
	u.user.commentInteractor = comment
	return u
}
