package web

type CampaignFundraiser struct {
	Owner       *string `json:"owner"`
	OwnerImage  *string `json:"ownerImage"`
	OwnerType   *string `json:"ownerType"`
	OwnerStatus string  `json:"ownerStatus"`
}

type CampaignDetailResponse struct {
	ID           int                 `json:"id"`
	CampaignID   *string             `json:"campaignId"`
	Slug         *string             `json:"slug"`
	ImageURL     *string             `json:"imageUrl"`
	Title        *string             `json:"title"`
	Target       *float64            `json:"target"`
	TotalDonasi  *float64            `json:"totalDonasi"`
	TotalDonatur *int                `json:"totalDonatur"`
	Fundraiser   *CampaignFundraiser `json:"fundraiser"`
	Information  *string             `json:"information"`
}
