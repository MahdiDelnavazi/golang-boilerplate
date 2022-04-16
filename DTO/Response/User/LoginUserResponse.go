package User

type LoginUserResponse struct {
	UserName    string `json:"username" validate:"required,min=3"`
	AccessToken string `json:"accessToken" binding:"required"`
}
