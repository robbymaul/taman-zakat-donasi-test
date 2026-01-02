package services

import (
	"context"
	"donasitamanzakattest/app/enums"
	"donasitamanzakattest/app/helpers"
	"donasitamanzakattest/app/models"
	"donasitamanzakattest/app/repositories"
	"donasitamanzakattest/app/web"
	"donasitamanzakattest/config"
	"net/http"
)

type ICampaignService interface {
	PublicListCampaignService() ([]*models.ListPriorityWpDjaCampaign, error)
	PublicGetCampaignService(slug string) (*web.CampaignDetailResponse, error)
	PublicListGetCampaignDonaturService(campaignId, sort string) ([]*models.ListCampaignWpDjaDonate, error)
	PublicListGetCampaignCommentService(campaignId string) ([]*models.ListCampaignWpDjaComment, error)
}

type CampaignService struct {
	*Service
}

func NewCampaignService(ctx context.Context, repo *repositories.RepositoryContext, cfg *config.Config) ICampaignService {
	return &CampaignService{
		Service: NewService(ctx, repo, cfg),
	}
}

func (s *CampaignService) PublicListCampaignService() ([]*models.ListPriorityWpDjaCampaign, error) {
	var campaigns []*models.ListPriorityWpDjaCampaign
	var err error

	campaigns, err = s.repository.PublicListCampaignRepository()
	if err != nil {
		return nil, helpers.NewErrorTrace(err, "list campaign").WithStatusCode(http.StatusInternalServerError)
	}

	return campaigns, err
}

func (s *CampaignService) PublicGetCampaignService(slug string) (*web.CampaignDetailResponse, error) {
	var campaign *models.DetailWpDjaCampaign
	var err error

	campaign, err = s.repository.PublicGetCampaignRepository(slug)
	if err != nil {
		return nil, helpers.NewErrorTrace(err, "campaign").WithStatusCode(http.StatusInternalServerError)
	}

	var response *web.CampaignDetailResponse

	response = s.mapCampaignDetailResponse(campaign)

	return response, err
}

func (s *CampaignService) PublicListGetCampaignDonaturService(campaignId, sort string) ([]*models.ListCampaignWpDjaDonate, error) {
	var donatur []*models.ListCampaignWpDjaDonate
	var err error

	donatur, err = s.repository.PublicListGetCampaignDonaturRepository(campaignId, sort)
	if err != nil {
		return nil, helpers.NewErrorTrace(err, "list donatur").WithStatusCode(http.StatusInternalServerError)
	}

	return donatur, err
}

func (s *CampaignService) PublicListGetCampaignCommentService(campaignId string) ([]*models.ListCampaignWpDjaComment, error) {
	var comments []*models.ListCampaignWpDjaComment
	var err error

	comments, err = s.repository.PublicListGetCampaignCommentRepository(campaignId)
	if err != nil {
		return nil, helpers.NewErrorTrace(err, "list comment").WithStatusCode(http.StatusInternalServerError)
	}

	return comments, err
}

func (s *CampaignService) mapCampaignDetailResponse(campaign *models.DetailWpDjaCampaign) *web.CampaignDetailResponse {
	return &web.CampaignDetailResponse{
		ID:           campaign.ID,
		CampaignID:   campaign.CampaignID,
		Slug:         campaign.Slug,
		ImageURL:     campaign.ImageURL,
		Title:        campaign.Title,
		Target:       campaign.Target,
		TotalDonasi:  campaign.TotalDonasi,
		TotalDonatur: campaign.TotalDonatur,
		Fundraiser: &web.CampaignFundraiser{
			Owner:       campaign.Owner,
			OwnerImage:  campaign.OwnerImage,
			OwnerType:   campaign.OwnerType,
			OwnerStatus: enums.GetUserVerification(campaign.OwnerVerification),
		},
		Information: campaign.Information,
	}
}
