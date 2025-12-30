package models

import "time"

// WpDjaAffCode represents table `wp_dja_aff_code`.
type WpDjaAffCode struct {
	ID         int        `gorm:"column:id;primaryKey" json:"id"`
	CampaignID *string    `gorm:"column:campaign_id" json:"campaignId"`
	UserID     *int       `gorm:"column:user_id" json:"userId"`
	AffCode    *string    `gorm:"column:aff_code" json:"affCode"`
	CreatedAt  *time.Time `gorm:"column:created_at" json:"createdAt"`
}

func (WpDjaAffCode) TableName() string {
	return "wp_dja_aff_code"
}
