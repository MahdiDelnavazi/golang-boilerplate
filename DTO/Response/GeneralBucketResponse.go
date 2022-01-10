package Response

import (
	"github.com/google/uuid"
	"time"
)

type GeneralBucketResponse struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
