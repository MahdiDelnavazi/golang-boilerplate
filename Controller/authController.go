package Controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"golang-boilerplate/DTO/Request/User"
	"golang-boilerplate/DTO/Response"
	"golang-boilerplate/Helper"
	"golang-boilerplate/Middleware/token"
	"net/http"
	"time"
)

type AuthController struct {
	logger *zap.SugaredLogger
	token  token.Maker
	redis  *redis.Client
	//userService *Service.UserService
}

func NewAuthController(logger *zap.SugaredLogger, token token.Maker, redis *redis.Client) *AuthController {
	return &AuthController{logger: logger, token: token, redis: redis}
}

func (authController *AuthController) AccessToken(context *gin.Context) {
	var accessTokenReq User.AccessTokenRequest
	Helper.Decode(context.Request, &accessTokenReq)

	//logoutResponse, logoutResponseError := userControler.userService.LogoutUser(userRequest)

	//if logoutResponseError != nil {
	//	context.JSON(http.StatusBadRequest, gin.H{"response": logoutResponseError})
	//	return
	//}

	payload, err := authController.token.VerifyToken(accessTokenReq.AccessToken)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": err})
		return
	}
	newToken, err := authController.token.CreateToken(payload.Username, time.Hour)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": err})
		return
	}

	// all ok
	// create general response
	response1 := Response.GeneralResponse{Error: false, Message: "successful", Data: Response.AccessTokenResponse{AccessToken: newToken}}
	context.JSON(http.StatusOK, gin.H{"response": response1})
}
