package Middleware

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"golang-boilerplate/Middleware/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

// AuthMiddleware creates a gin middleware for authorization
func AuthMiddleware(tokenMaker token.Maker, redis *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		//val2 , _ :=redis.Get("token").Result()
		//fmt.Println("redis resuuuuult : " , val2)
		//if err != nil {
		//	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err})
		//	return
		//}

		fmt.Println("this is token ------------>", authorizationHeader)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": err})
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": err})
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": err})
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": err})
			return
		}

		val, _ := redis.Get(payload.Username).Result()
		if val == accessToken {
			err := fmt.Errorf("token is expired")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": err})
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
