package User

type GetUserRequest struct {
	UserName string `json:"username" validate:"required,min=3"`
}
