package controllers

import (
	"donasitamanzakattest/app/services"
	"donasitamanzakattest/app/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) HealthController(ctx *gin.Context) {
	healthService := services.NewHealthService(ctx, c.startTime, c.repo, c.cfg)

	response := healthService.GetHealthService()

	ctx.JSON(http.StatusOK, web.ResponseWeb{
		Message: "OK",
		Success: true,
		Data:    response,
	})
}
