package models

import "time"

// WpDjaAffClick represents table `wp_dja_aff_click`.
type WpDjaAffClick struct {
	ID         int        `gorm:"column:id;primaryKey" json:"id"`
	CampaignID *string    `gorm:"column:campaign_id" json:"campaignId"`
	AffcodeID  *string    `gorm:"column:affcode_id" json:"affcodeId"`
	IP         *string    `gorm:"column:ip" json:"ip"`
	OS         *string    `gorm:"column:os" json:"os"`
	Browser    *string    `gorm:"column:browser" json:"browser"`
	CreatedAt  *time.Time `gorm:"column:created_at" json:"createdAt"`
}

func (WpDjaAffClick) TableName() string {
	return "wp_dja_aff_click"
}
