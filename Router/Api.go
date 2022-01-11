package Router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang-boilerplate/Controller"
)

const (
	prefix        = "/api/v1"
	bucketPostfix = "/bucket"
)

func Routes(app *gin.Engine, log *zap.SugaredLogger) {
	router := app.Group(prefix)

	bucketController := Controller.NewBucketController(log)
	bucket := router.Group(bucketPostfix)
	{
		bucket.POST("/", bucketController.CreateBucket)
		bucket.GET("/", bucketController.ListBucket)
		bucket.GET("/:id", bucketController.GetOneBucket)
		bucket.PUT("/:id", bucketController.UpdateBucket)
		bucket.PATCH("/:id", bucketController.ToggleActiveBucket)
		bucket.DELETE("/:id", bucketController.DeleteBucket)
	}
}
