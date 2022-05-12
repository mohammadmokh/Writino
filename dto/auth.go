package dto

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token    string `json:"access_token"`
	RefToken string `json:"refresh_token"`
}

type RefreshReq struct {
	RefToken string `json:"token"`
}

type RefreshResponse struct {
	Token    string `json:"access_token"`
	RefToken string `json:"refresh_token"`
}
