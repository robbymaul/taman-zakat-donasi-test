package models

import "time"

type WpDjaCampaignUpdate struct {
	ID          int        `gorm:"column:id;primaryKey" json:"id"`
	CampaignID  *string    `gorm:"column:campaign_id" json:"campaignId"`
	Title       *string    `gorm:"column:title" json:"title"`
	Information *string    `gorm:"column:information" json:"information"`
	UserID      *int64     `gorm:"column:user_id" json:"userId"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"createdAt"`
}

func (WpDjaCampaignUpdate) TableName() string {
	return "wp_dja_campaign_update"
}
