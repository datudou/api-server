package request

type CreateUserParam struct {
	UserID   string `json:"user_id" binding:"required,min=6,max=20"`
	Password string `json:"password" binding:"required,min=8,max=20,ascii"`
	NickName string `json:"nickname,omitempty"`
	Comment  string `json:"comment,omitempty"`
}

type UpdateUserParam struct {
	NickName string `json:"nickname" binding:"required,max=30"`
	Comment  string `json:"comment" binding:"required,max=100,ascii"`
}
