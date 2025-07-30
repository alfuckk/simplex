package hdl

import (
	"simplex/pkg/jwt"
	"simplex/pkg/logx"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	logger *logx.Logger
}

func NewHandler(
	logger *logx.Logger,
) *Handler {
	return &Handler{
		logger: logger,
	}
}

func GetUserIdFromCtx(ctx *gin.Context) string {
	v, exists := ctx.Get("claims")
	if !exists {
		return ""
	}
	return v.(*jwt.MyCustomClaims).UserId
}
