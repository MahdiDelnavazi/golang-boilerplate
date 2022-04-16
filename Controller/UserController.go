package Controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang-boilerplate/DTO/Request/User"
	"golang-boilerplate/DTO/Response"
	User2 "golang-boilerplate/DTO/Response/User"
	"golang-boilerplate/Helper"
	"golang-boilerplate/Service"
	"net/http"
)

type UserController struct {
	logger      *zap.SugaredLogger
	userService *Service.UserService
}

func NewUserController(logger *zap.SugaredLogger, userService *Service.UserService) *UserController {
	return &UserController{logger: logger, userService: userService}
}

func (userControler *UserController) CreateUser(context *gin.Context) {
	var userRequest User.CreateUserRequest
	Helper.Decode(context.Request, &userRequest)

	userResponse, responseError := userControler.userService.Create(userRequest)

	if responseError != nil {
		// if username not empty means its userExist error
		if userResponse.UserName != "" {
			response := Response.GeneralResponse{Error: true, Message: "user exist", Data: nil}
			context.JSON(http.StatusBadRequest, gin.H{"response": response})
			return
		}
		// if username is empty means its validation error
		context.JSON(http.StatusBadRequest, gin.H{"response": responseError})
		return
	}

	// all ok
	// create general response
	response := Response.GeneralResponse{Error: false, Message: "user have been created", Data: User2.CreateUserResponse{UserName: userResponse.UserName}}
	context.JSON(http.StatusOK, gin.H{"response": response})
}

func (userControler *UserController) LoginUser(context *gin.Context) {
	var userRequest User.LoginUserRequest
	Helper.Decode(context.Request, &userRequest)

	userResponse, responseError := userControler.userService.LoginUser(userRequest)

	if responseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": responseError})
		return
	}

	// all ok
	// create general response
	response := Response.GeneralResponse{Error: false, Message: "your login is successful", Data: User2.LoginUserResponse{UserName: userResponse.UserName, AccessToken: userResponse.AccessToken}}
	context.JSON(http.StatusOK, gin.H{"response": response})
}

func (userControler *UserController) Logout(context *gin.Context) {
	var userRequest User.LogoutRequest
	Helper.Decode(context.Request, &userRequest)

	logoutResponse, logoutResponseError := userControler.userService.LogoutUser(userRequest)

	if logoutResponseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": logoutResponseError})
		return
	}

	// all ok
	// create general response
	response1 := Response.GeneralResponse{Error: false, Message: logoutResponse}
	context.JSON(http.StatusOK, gin.H{"response": response1})
}
