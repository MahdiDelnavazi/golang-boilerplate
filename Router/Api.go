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
	auth          = "/api/v1/authentication"
	user          = "/api/v1/user"
	bucketPostfix = "/bucket"
)

func Routes(app *gin.Engine, log *zap.SugaredLogger, db *sqlx.DB, token token.Maker, redis *redis.Client) {
	router := app.Group(prefix)
	routerAuth := app.Group(auth)
	routerUser := app.Group(user)
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	newUserRepository := Repository.NewUserRepository(log, db)
	newUserService := Service.NewUserService(log, newUserRepository, token, redis)
	newUserController := Controller.NewUserController(log, newUserService)
	newAuthController := Controller.NewAuthController(log, token, redis)

	//newBucketRepository := Repository.NewBucketRepository(log, db)
	//newBucketService := Service.NewBucketService(log, newBucketRepository)
	//bucketController := Controller.NewBucketController(log, newBucketService)

	newTicketRepository := Repository.NewTicketRepository(log, db)
	newTicketService := Service.NewTicketService(log, newUserService, newTicketRepository)
	newTicketController := Controller.NewTicketController(log, newTicketService)

	authRoutes := router.Group("/").Use(Middleware.AuthMiddleware(token, redis))
	routerUser.POST("/create", newUserController.CreateUser)
	routerUser.POST("/login", newUserController.LoginUser)
	routerAuth.POST("/logout", newUserController.Logout)
	routerAuth.POST("/accessToken", newAuthController.AccessToken)
	authRoutes.POST("/addTicket", newTicketController.CreateTicket)

}
