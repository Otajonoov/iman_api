package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/iman_api/pkg/logger"
)

type handler struct {
	log logger.Logger
}

type HandlerOptions struct {
	Log logger.Logger
}

func New(options *HandlerOptions) *handler {
	return &handler{
		log: options.Log,
	}
}

// @Description 2025 yil 1 yanvargacha qolgan kunlar sonini hisoblab qaytaradi.
// @Tags 		Day
// @Security    BearerAuth
// @Router /v1/days [get]
func (h *handler) Days(ctx *gin.Context) {

	targetDate := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	currentTime := time.Now()
	daysLeft := int(targetDate.Sub(currentTime).Hours() / 24)

	ctx.JSON(http.StatusOK, gin.H{
		"is_ok": daysLeft,
	})
}
