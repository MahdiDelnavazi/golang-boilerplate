package Repository

import (
	"fmt"
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
	queryResult, queryError := bucketRepository.database.Exec("CALL create_bucket($1)", createBucketRequest.Name)
	fmt.Println("+======________________+++++++++++++++")
	fmt.Println(queryError)
	fmt.Println("+======________________+++++++++++++++")
	if queryError != nil {
		return Entity.Bucket{}, queryError
	}

	fmt.Println("==================================")
	fmt.Println(queryResult.RowsAffected())
	fmt.Println("==================================")

	return Entity.Bucket{}, nil
}
