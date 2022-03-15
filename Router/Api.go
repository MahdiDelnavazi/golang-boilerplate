package Router

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"golang-boilerplate/Controller"
	"golang-boilerplate/Repository"
	"golang-boilerplate/Service"
	"net/http"
)

const (
	prefix        = "/api/v1"
	bucketPostfix = "/bucket"
)

func Routes(app *gin.Engine, log *zap.SugaredLogger, db *sqlx.DB) {
	router := app.Group(prefix)
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	newUserRepository := Repository.NewUserRepository(log, db)
	newUserService := Service.NewUserService(log, newUserRepository)
	newUserController := Controller.NewUserController(log, newUserService)

	newBucketRepository := Repository.NewBucketRepository(log, db)
	newBucketService := Service.NewBucketService(log, newBucketRepository)
	bucketController := Controller.NewBucketController(log, newBucketService)

	newTicketRepository := Repository.NewTicketRepository(log, db)
	newTicketService := Service.NewTicketService(log, newUserService, newTicketRepository)
	newTicketController := Controller.NewTicketController(log, newTicketService)

	router.POST("/adduser", newUserController.CreateUser)
	router.POST("/addTicket", newTicketController.CreateTicket)
	bucket := router.Group(bucketPostfix)
	{
		bucket.POST("/", bucketController.CreateBucket)
		bucket.POST("/addUser", newUserController.CreateUser)
		bucket.POST("/addTicket", newTicketController.CreateTicket)
	}
}
