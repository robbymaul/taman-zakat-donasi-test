package controllers

import (
	"donasitamanzakattest/app/helpers"
	"donasitamanzakattest/app/services"
	"donasitamanzakattest/app/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) RegistrationController(ctx *gin.Context) {
	req := new(web.RegistrationRequest)

	err := ctx.ShouldBindJSON(req)
	if err != nil {
		helpers.ErrorResponse(ctx, helpers.NewErrorTrace(err, "registration").WithStatusCode(http.StatusBadRequest))
		return
	}

	authService := services.NewAuthService(ctx, c.repo, c.cfg)

	response, err := authService.RegistrationService(req)
	if err != nil {
		helpers.ErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, web.ResponseWeb{
		Message: "OK",
		Success: true,
		Data:    response,
	})

}

func (c *Controller) LoginController(ctx *gin.Context) {
	req := new(web.LoginRequest)

	err := ctx.ShouldBindJSON(req)
	if err != nil {
		helpers.ErrorResponse(ctx, helpers.NewErrorTrace(err, "login").WithStatusCode(http.StatusBadRequest))
		return
	}

	authService := services.NewAuthService(ctx, c.repo, c.cfg)

	response, err := authService.LoginService(req)
	if err != nil {
		helpers.ErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, web.ResponseWeb{
		Message: "OK",
		Success: true,
		Data:    response,
	})
}

func (c *Controller) PrivateGetProfileController(ctx *gin.Context) {
	authService := services.NewAuthService(ctx, c.repo, c.cfg)

	response, err := authService.PrivateGetProfileService()
	if err != nil {
		helpers.ErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, web.ResponseWeb{
		Message: "OK",
		Success: true,
		Data:    response,
	})
}
