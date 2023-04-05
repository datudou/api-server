package response

type User struct {
	UserID   string `json:"user_id"`
	NickName string `json:"nickname,omitempty"`
	Comment  string `json:"comment,omitempty"`
}
