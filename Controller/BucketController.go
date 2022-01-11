package Controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BucketController struct {
	logger *zap.SugaredLogger
}

func NewBucketController(logger *zap.SugaredLogger) *BucketController {
	return &BucketController{logger: logger}
}

func (bucketController *BucketController) CreateBucket(context *gin.Context)   {}
func (bucketController *BucketController) ListBucket(context *gin.Context)     {}
func (bucketController *BucketController) GetOneBucket(ctx *gin.Context)       {}
func (bucketController *BucketController) UpdateBucket(ctx *gin.Context)       {}
func (bucketController *BucketController) ToggleActiveBucket(ctx *gin.Context) {}
func (bucketController *BucketController) DeleteBucket(ctx *gin.Context)       {}
