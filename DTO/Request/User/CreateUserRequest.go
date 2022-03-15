package User

type CreateUserRequest struct {
	UserName string `json:"userName" validate:"required,min=3"`
}
