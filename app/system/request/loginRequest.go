package request

type LoginRequest struct {
	Mobile   string `json:"mobile" validate:"required,min=11,max=11"`
	Password string `json:"password" validate:"required,min=6,max=32"`
}
