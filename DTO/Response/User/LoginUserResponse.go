package User

type LoginUserResponse struct {
	UserName    string `json:"userName" validate:"required,min=3"`
	AccessToken string `json:"AccessToken" binding:"required"`
}
