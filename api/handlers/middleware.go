package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	token "gitlab.com/iman_api/pkg/token"
)

func (h *handler) AuthMiddleWare(ctx *gin.Context) {
	accessToken := ctx.GetHeader("Authorization")

	if len(accessToken) == 0 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "authorization header is not provided",
		})
		return
	}

	err := token.VerifyToken(accessToken)
	if err != nil {
		h.log.Error(err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Token is invalid",
		})
		return
	}

	ctx.Next()
}
