package Entity

import (
	"github.com/google/uuid"
	"time"
	"unsafe"
)

type Bucket struct {
	Id        uuid.UUID
	Name      string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func toBucket(bucket Bucket) Bucket {
	bucketCreated := (*Bucket)(unsafe.Pointer(&bucket))
	return *bucketCreated
}
