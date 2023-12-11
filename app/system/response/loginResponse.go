package response

type LoginResponse struct {
	Token      string `json:"token"`
	ExpireTime int64  `json:"expire_time"`
}
