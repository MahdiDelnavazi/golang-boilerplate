package Controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang-boilerplate/Service"
)

type BucketController struct {
	logger        *zap.SugaredLogger
	bucketService Service.BucketService
}

func NewBucketController(logger *zap.SugaredLogger) *BucketController {
	return &BucketController{logger: logger}
}

func (bucketController *BucketController) CreateBucket(context *gin.Context) {
	//bucketController.bucketService.CreateBucket(context)
}
