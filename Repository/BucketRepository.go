package Repository

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type BucketRepository struct {
	logger   *zap.SugaredLogger
	database sqlx.DB
}

func NewBucketRepository(logger *zap.SugaredLogger, database sqlx.DB) *BucketRepository {
	return &BucketRepository{
		logger:   logger,
		database: database,
	}
}

func (bucketRepository *BucketRepository) CreateBucket() {}

func (bucketRepository *BucketRepository) ListBucket() {}
