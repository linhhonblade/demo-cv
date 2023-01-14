package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-hello/helper"
	"net/http"
)

func JWTAuthMiddleware(c context.Context) gin.HandlerFunc {
	return func(context *gin.Context) {
		err := helper.ValidateJWT(context, c)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}
		context.Next()
	}
}
