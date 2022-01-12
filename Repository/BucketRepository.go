package Repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"golang-boilerplate/DTO/Request"
	"golang-boilerplate/Entity"
)

type BucketRepository struct {
	logger   *zap.SugaredLogger
	database *sqlx.DB
}

func NewBucketRepository(logger *zap.SugaredLogger, database *sqlx.DB) *BucketRepository {
	return &BucketRepository{
		logger:   logger,
		database: database,
	}
}

func (bucketRepository *BucketRepository) CreateBucket(createBucketRequest Request.CreateBucketRequest) (Entity.Bucket, error) {
	bucketRepository.database.Exec("CALL create_bucket($1)", createBucketRequest.Name)
	return Entity.Bucket{}, nil
}
