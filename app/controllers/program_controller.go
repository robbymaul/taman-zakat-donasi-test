package controllers

import (
	"donasitamanzakattest/app/models"
	"donasitamanzakattest/app/services"
	"donasitamanzakattest/app/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) CreateProgram(ctx *gin.Context) {
	req := new(web.ProgramUpsertRequest)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.ResponseWeb{Message: err.Error(), Success: false})
		return
	}

	programService := services.NewProgramService(ctx, c.repo, c.cfg)
	cat, err := programService.Create(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ResponseWeb{Message: err.Error(), Success: false})
		return
	}

	ctx.JSON(http.StatusCreated, web.ResponseWeb{Message: "OK", Success: true, Data: cat})
}

func (c *Controller) ListPrograms(ctx *gin.Context) {
	programService := services.NewProgramService(ctx, c.repo, c.cfg)
	cats, err := programService.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.ResponseWeb{Message: err.Error(), Success: false})
		return
	}

	ctx.JSON(http.StatusOK, web.ResponseWeb{Message: "OK", Success: true, Data: cats})
}

func (c *Controller) GetProgram(ctx *gin.Context) {
	id := ctx.Param("id")
	programService := services.NewProgramService(ctx, c.repo, c.cfg)

	cat, err := programService.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.ResponseWeb{Message: err.Error(), Success: false})
		return
	}

	ctx.JSON(http.StatusOK, web.ResponseWeb{Message: "OK", Success: true, Data: cat})
}

func (c *Controller) UpdateProgram(ctx *gin.Context) {
	id := ctx.Param("id")

	req := new(models.Program)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.ResponseWeb{Message: err.Error(), Success: false})
		return
	}

	programService := services.NewProgramService(ctx, c.repo, c.cfg)
	cat, err := programService.Update(id, req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.ResponseWeb{Message: err.Error(), Success: false})
		return
	}

	ctx.JSON(http.StatusOK, web.ResponseWeb{Message: "OK", Success: true, Data: cat})
}

func (c *Controller) DeleteProgram(ctx *gin.Context) {
	id := ctx.Param("id")
	programService := services.NewProgramService(ctx, c.repo, c.cfg)

	cat, err := programService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.ResponseWeb{Message: err.Error(), Success: false})
		return
	}

	ctx.JSON(http.StatusOK, web.ResponseWeb{Message: "OK", Success: true, Data: cat})
}
