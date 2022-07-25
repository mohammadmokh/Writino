package dto

type RegisterReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type VerifyUserReq struct {
	Email            string `json:"email"`
	VerificationCode string `json:"verfication_code"`
}

type CheckUsernameReq struct {
	Username string `json:"username"`
}

type CheckUsernameRes struct {
	IsUnique bool `json:"is_unique"`
}

type CheckEmailReq struct {
	Email string `json:"email"`
}

type CheckEmailRes struct {
	IsUnique bool `json:"is_unique"`
}

type UpdateUserReq struct {
	ID          string  `json:"-"`
	ProfilePic  *string `json:"profile_pic"`
	Username    *string `json:"username"`
	DisplayName *string `json:"display_name"`
	Bio         *string `json:"bio"`
}

type DeleteUserReq struct {
	Id string
}

type FindUserReq struct {
	Id string
}

type FindUserRes struct {
	ProfilePic  string `json:"profile_pic,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Bio         string `json:"bio,omitempty"`
	Email       string `json:"email,omitempty"`
}

type UpdatePasswordReq struct {
	ID  string `json:"-"`
	Old string `json:"old"`
	New string `json:"new"`
}

type UpdateProfilePicReq struct {
	ID     string
	Image  []byte
	Format string
}

type UpdateProfilePicRes struct {
	Link string
}
