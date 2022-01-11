package Router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang-boilerplate/Controller"
	"net/http"
)

const (
	prefix        = "/api/v1"
	bucketPostfix = "/bucket"
)

func Routes(app *gin.Engine, log *zap.SugaredLogger) {
	router := app.Group(prefix)
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	bucketController := Controller.NewBucketController(log)
	bucket := router.Group(bucketPostfix)
	{
		bucket.POST("/", bucketController.CreateBucket)
	}
}
