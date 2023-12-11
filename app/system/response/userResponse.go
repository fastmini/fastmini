package response

type UserResponse struct {
	ID              int64  `json:"id"`
	Mobile          string `json:"mobile"`
	PasswordVersion int8   `json:"password_version"`
	NickName        string `json:"nick_name"`
}
