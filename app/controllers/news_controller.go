package controllers

import (
	"donasitamanzakattest/app/models"
	"donasitamanzakattest/app/services"
	"donasitamanzakattest/app/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) CreateNews(ctx *gin.Context) {
	req := new(web.NewsUpsertRequest)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.ResponseWeb{Message: err.Error(), Success: false})
		return
	}

	newsService := services.NewNewsService(ctx, c.repo, c.cfg)
	cat, err := newsService.Create(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ResponseWeb{Message: err.Error(), Success: false})
		return
	}

	ctx.JSON(http.StatusCreated, web.ResponseWeb{Message: "OK", Success: true, Data: cat})
}

func (c *Controller) ListNews(ctx *gin.Context) {
	newsService := services.NewNewsService(ctx, c.repo, c.cfg)
	cats, err := newsService.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.ResponseWeb{Message: err.Error(), Success: false})
		return
	}

	ctx.JSON(http.StatusOK, web.ResponseWeb{Message: "OK", Success: true, Data: cats})
}

func (c *Controller) GetNews(ctx *gin.Context) {
	id := ctx.Param("id")
	newsService := services.NewNewsService(ctx, c.repo, c.cfg)

	cat, err := newsService.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.ResponseWeb{Message: err.Error(), Success: false})
		return
	}

	ctx.JSON(http.StatusOK, web.ResponseWeb{Message: "OK", Success: true, Data: cat})
}

func (c *Controller) UpdateNews(ctx *gin.Context) {
	id := ctx.Param("id")

	req := new(models.News)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.ResponseWeb{Message: err.Error(), Success: false})
		return
	}

	newsService := services.NewNewsService(ctx, c.repo, c.cfg)
	cat, err := newsService.Update(id, req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.ResponseWeb{Message: err.Error(), Success: false})
		return
	}

	ctx.JSON(http.StatusOK, web.ResponseWeb{Message: "OK", Success: true, Data: cat})
}

func (c *Controller) DeleteNews(ctx *gin.Context) {
	id := ctx.Param("id")
	newsService := services.NewNewsService(ctx, c.repo, c.cfg)

	cat, err := newsService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.ResponseWeb{Message: err.Error(), Success: false})
		return
	}

	ctx.JSON(http.StatusOK, web.ResponseWeb{Message: "OK", Success: true, Data: cat})
}
