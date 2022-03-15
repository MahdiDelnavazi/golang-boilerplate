package User

type CreateUserResponse struct {
	UserName string `json:"userName" binding:"required"`
}
