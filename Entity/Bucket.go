package Entity

import (
	"github.com/google/uuid"
	"time"
)

type Bucket struct {
	Id        uuid.UUID
	Name      string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
