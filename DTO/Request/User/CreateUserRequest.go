package User

type CreateUserRequest struct {
	UserName string `json:"username" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=8"`
}
