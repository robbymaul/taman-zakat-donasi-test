package controllers

import (
	"donasitamanzakattest/app/helpers"
	"donasitamanzakattest/app/services"
	"donasitamanzakattest/app/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

type categoryUpsertRequest struct {
	Name string `json:"name"`
}

func (c *Controller) CreateCategory(ctx *gin.Context) {
	var req categoryUpsertRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.ResponseWeb{Message: err.Error(), Success: false})
		return
	}

	categoryService := services.NewCategoryService(ctx, c.repo, c.cfg)
	cat, err := categoryService.Create(req.Name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ResponseWeb{Message: err.Error(), Success: false})
		return
	}

	ctx.JSON(http.StatusCreated, web.ResponseWeb{Message: "OK", Success: true, Data: cat})
}

func (c *Controller) ListCategories(ctx *gin.Context) {
	categoryService := services.NewCategoryService(ctx, c.repo, c.cfg)
	cats, err := categoryService.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.ResponseWeb{Message: err.Error(), Success: false})
		return
	}

	ctx.JSON(http.StatusOK, web.ResponseWeb{Message: "OK", Success: true, Data: cats})
}

func (c *Controller) GetCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	categoryService := services.NewCategoryService(ctx, c.repo, c.cfg)

	cat, err := categoryService.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.ResponseWeb{Message: err.Error(), Success: false})
		return
	}

	ctx.JSON(http.StatusOK, web.ResponseWeb{Message: "OK", Success: true, Data: cat})
}

func (c *Controller) UpdateCategory(ctx *gin.Context) {
	id := ctx.Param("id")

	var req categoryUpsertRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.ResponseWeb{Message: err.Error(), Success: false})
		return
	}

	categoryService := services.NewCategoryService(ctx, c.repo, c.cfg)
	cat, err := categoryService.Update(id, req.Name)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.ResponseWeb{Message: err.Error(), Success: false})
		return
	}

	ctx.JSON(http.StatusOK, web.ResponseWeb{Message: "OK", Success: true, Data: cat})
}

func (c *Controller) DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	categoryService := services.NewCategoryService(ctx, c.repo, c.cfg)

	cat, err := categoryService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.ResponseWeb{Message: err.Error(), Success: false})
		return
	}

	ctx.JSON(http.StatusOK, web.ResponseWeb{Message: "OK", Success: true, Data: cat})
}

func (c *Controller) PublicListCategoriesController(ctx *gin.Context) {
	categoryService := services.NewCategoryService(ctx, c.repo, c.cfg)

	response, err := categoryService.PublicListCategoriesService()
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
