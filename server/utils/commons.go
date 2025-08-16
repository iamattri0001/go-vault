package utils

import (
	"errors"
	"go-vault/constants"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserIdFromContext(ctx *gin.Context) (uuid.UUID, error) {
	userId, exists := ctx.Get(constants.UserID)
	if !exists {
		return uuid.Nil, errors.New("user_id not found in context")
	}

	id, ok := userId.(uuid.UUID)
	if !ok || id == uuid.Nil {
		return uuid.Nil, errors.New("user_id not found in context")
	}

	return id, nil
}
