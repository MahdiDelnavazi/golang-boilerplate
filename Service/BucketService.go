package Service

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"golang-boilerplate/DTO/Request"
	"golang-boilerplate/Repository"
)

type BucketService struct {
	bucketRepository *Repository.BucketRepository
}

func NewBucketService(logger *zap.SugaredLogger, database sqlx.DB) *BucketService {
	return &BucketService{
		bucketRepository: Repository.NewBucketRepository(logger, database),
	}
}

func (bucketService BucketService) Create(createBucketRequest Request.CreateBucketRequest) error {
	return nil
}
