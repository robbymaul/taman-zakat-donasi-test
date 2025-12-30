package models

import "time"

type WpDjaAffSubmit struct {
	ID                int        `gorm:"column:id;primaryKey" json:"id"`
	DonateID          *int       `gorm:"column:donate_id" json:"donateId"`
	CampaignID        *string    `gorm:"column:campaign_id" json:"campaignId"`
	AffcodeID         *int       `gorm:"column:affcode_id" json:"affcodeId"`
	PayoutStatus      *int       `gorm:"column:payout_status" json:"payoutStatus"`
	NominalCommission *float64   `gorm:"column:nominal_commission" json:"nominalCommission"`
	CreatedAt         *time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt         *time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (WpDjaAffSubmit) TableName() string {
	return "wp_dja_aff_submit"
}
