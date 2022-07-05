package dto

type CreateCommentReq struct {
	Comment string `json:"comment"`
	UserID  string `josn:"-"`
	PostID  string `json:"-"`
}

type FindCommentReq struct {
	PostID string
	Limit  int
	Page   int
}

type Comment struct {
	Comment string `json:"comment"`
	User    string `json:"user"`
}

type FindCommentRes struct {
	TotalCount int       `json:"total"`
	Comments   []Comment `json:"comments"`
}
