package User

import "github.com/google/uuid"

type GetUserResponse struct {
	UserId   uuid.UUID `json:"subject" validate:"required"`
	UserName string    `json:"userName" validate:"required,min=3"`
}
