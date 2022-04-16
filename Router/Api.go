package Router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"golang-boilerplate/Controller"
	"golang-boilerplate/Middleware"
	"golang-boilerplate/Middleware/token"
	"golang-boilerplate/Repository"
	"golang-boilerplate/Service"
	"net/http"
)

const (
	prefix        = "/api/v1"
	bucketPostfix = "/bucket"
)

func Routes(app *gin.Engine, log *zap.SugaredLogger, db *sqlx.DB, token token.Maker, redis *redis.Client) {
	router := app.Group(prefix)
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	newUserRepository := Repository.NewUserRepository(log, db)
	newUserService := Service.NewUserService(log, newUserRepository, token, redis)
	newUserController := Controller.NewUserController(log, newUserService)

	//newBucketRepository := Repository.NewBucketRepository(log, db)
	//newBucketService := Service.NewBucketService(log, newBucketRepository)
	//bucketController := Controller.NewBucketController(log, newBucketService)

	newTicketRepository := Repository.NewTicketRepository(log, db)
	newTicketService := Service.NewTicketService(log, newUserService, newTicketRepository)
	newTicketController := Controller.NewTicketController(log, newTicketService)

	authRoutes := router.Group("/").Use(Middleware.AuthMiddleware(token, redis))
	router.POST("/createUser", newUserController.CreateUser)
	router.POST("/loginUser", newUserController.LoginUser)
	router.POST("/logout", newUserController.Logout)
	authRoutes.POST("/addTicket", newTicketController.CreateTicket)
	//bucket := router.Group(bucketPostfix)
	//{
	//	bucket.POST("/", bucketController.CreateBucket)
	//	bucket.POST("/createUser", newUserController.CreateUser)
	//	bucket.POST("/loginUser", newUserController.LoginUser)
	//	bucket.POST("/addTicket", newTicketController.CreateTicket)
	//}
}
