package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SystemHandler struct{}

// HealthCheck godoc
// @Summary 死活監視用
// @Tags healthcheck
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /v1/health [get]
func (h *SystemHandler) Health(ctx *gin.Context) {
	// gin.Hはmap[string]]interface{}のショートカット
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"Status": "ok",
	// })
	ctx.JSON(http.StatusOK, Response{Status: "ok"})
}

func NewSystemHandler() *SystemHandler {
	return &SystemHandler{}
}
