package Controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang-boilerplate/DTO/Request"
	"golang-boilerplate/Helper"
	"golang-boilerplate/Service"
	"net/http"
)

type BucketController struct {
	logger        *zap.SugaredLogger
	bucketService *Service.BucketService
}

func NewBucketController(logger *zap.SugaredLogger, bucketService *Service.BucketService) *BucketController {
	return &BucketController{logger: logger, bucketService: bucketService}
}

func (bucketController *BucketController) CreateBucket(context *gin.Context) {
	var bucketRequest Request.CreateBucketRequest
	Helper.Decode(context.Request, &bucketRequest)
	//if requestDecoderError :=  requestDecoderError == nil {
	//	context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": requestDecoderError})
	//	return
	//}
	bucketResponse, responseError := bucketController.bucketService.Create(bucketRequest)
	if responseError != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": responseError})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": bucketResponse})
}
