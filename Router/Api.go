package Router

import (
	"github.com/gin-gonic/gin"
	"golang-boilerplate/Controller"
)

const (
	prefix        = "/api/v1"
	bucketPostfix = "/bucket"
)

func Routes(app *gin.Engine) {
	router := app.Group(prefix)

	bucket := router.Group(bucketPostfix)
	{
		bucket.POST("/", Controller.CreateBucket)
		bucket.GET("/", Controller.ListBucket)
		bucket.GET("/:id", Controller.GetOneBucket)
		bucket.PUT("/:id", Controller.UpdateBucket)
		bucket.PATCH("/:id", Controller.ToggleActiveBucket)
		bucket.DELETE("/:id", Controller.DeleteBucket)
	}
}
