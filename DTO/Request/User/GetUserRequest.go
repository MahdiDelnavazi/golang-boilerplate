package User

type GetUserRequest struct {
	UserName string `json:"userName" validate:"required,min=3"`
}
