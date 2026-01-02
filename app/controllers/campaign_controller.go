package controllers

import (
	"donasitamanzakattest/app/helpers"
	"donasitamanzakattest/app/services"
	"donasitamanzakattest/app/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) PublicListCampaignController(ctx *gin.Context) {

	campaignService := services.NewCampaignService(ctx, c.repo, c.cfg)

	response, err := campaignService.PublicListCampaignService()
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

func (c *Controller) PublicGetCampaignController(ctx *gin.Context) {
	slug := ctx.Param("slug")

	campaignService := services.NewCampaignService(ctx, c.repo, c.cfg)

	response, err := campaignService.PublicGetCampaignService(slug)
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

func (c *Controller) PublicListGetCampaignDonaturController(ctx *gin.Context) {
	sort := ctx.DefaultQuery("sort", "terbaru")
	campaignId := ctx.Param("campaign-id")

	campaignService := services.NewCampaignService(ctx, c.repo, c.cfg)

	response, err := campaignService.PublicListGetCampaignDonaturService(campaignId, sort)
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

func (c *Controller) PublicListGetCampaignCommentController(ctx *gin.Context) {
	campaignId := ctx.Param("campaign-id")

	campaignService := services.NewCampaignService(ctx, c.repo, c.cfg)

	response, err := campaignService.PublicListGetCampaignCommentService(campaignId)
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
