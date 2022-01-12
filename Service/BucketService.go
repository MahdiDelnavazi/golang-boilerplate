package Service

import (
	"fmt"
	"go.uber.org/zap"
	"golang-boilerplate/DTO/Request"
	"golang-boilerplate/DTO/Response"
	"golang-boilerplate/Repository"
)

type BucketService struct {
	bucketRepository *Repository.BucketRepository
	logger           *zap.SugaredLogger
}

func NewBucketService(logger *zap.SugaredLogger, bucketRepository *Repository.BucketRepository) *BucketService {
	return &BucketService{logger: logger, bucketRepository: bucketRepository}
}

func (bucketService BucketService) Create(createBucketRequest Request.CreateBucketRequest) (Response.GeneralBucketResponse, error) {
	validationError := ValidationCheck(createBucketRequest)
	if validationError != nil {
		return Response.GeneralBucketResponse{}, validationError
	}
	bucketResponse, bucketRepositoryError := bucketService.bucketRepository.CreateBucket(createBucketRequest)
	if bucketRepositoryError != nil {
		return Response.GeneralBucketResponse{}, validationError
	}
	fmt.Println("bucketResponse: ", bucketResponse)
	// we need a transformer
	return Response.GeneralBucketResponse{}, nil
}
