package middleware

import (
	"go-vault/constants"
	"go-vault/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Implement your authentication logic here
		token, err := ctx.Cookie("token")
		if err != nil || token == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		// Validate the token using the service
		claims, err := jwt.ParseToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}

		userID := uuid.UUID(claims.ID)
		if userID == uuid.Nil {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid user ID"})
			return
		}

		// Set the user information in the context
		ctx.Set(constants.Username, claims.Username)
		ctx.Set(constants.UserID, userID)
		ctx.Next()
	}
}
