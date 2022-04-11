package User

type LoginUserRequest struct {
	UserName string `json:"userName" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=8"`
}
