package controllers

import (
	"donasitamanzakattest/app/helpers"
	"donasitamanzakattest/app/services"
	"donasitamanzakattest/app/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) PublicListPostsController(ctx *gin.Context) {

	postsService := services.NewPostsService(ctx, c.repo, c.cfg)

	response, err := postsService.PublicListPostsService()
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

func (c *Controller) PublicGetPostsController(ctx *gin.Context) {
	postName := ctx.Param("post-name")

	postsService := services.NewPostsService(ctx, c.repo, c.cfg)

	response, err := postsService.PublicGetPostsService(postName)
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
