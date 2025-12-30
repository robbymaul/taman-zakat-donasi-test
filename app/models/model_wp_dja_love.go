package models

import "time"

type WpDjaLove struct {
	ID         int        `gorm:"column:id;primaryKey" json:"id"`
	CampaignID *string    `gorm:"column:campaign_id" json:"campaignId"`
	DonateID   *int       `gorm:"column:donate_id" json:"donateId"`
	UserID     *int64     `gorm:"column:user_id" json:"userId"`
	Love       *int       `gorm:"column:love" json:"love"`
	IP         *string    `gorm:"column:ip" json:"ip"`
	OS         *string    `gorm:"column:os" json:"os"`
	Browser    *string    `gorm:"column:browser" json:"browser"`
	MobDesktop *string    `gorm:"column:mobdesktop" json:"mobdesktop"`
	CreatedAt  *time.Time `gorm:"column:created_at" json:"createdAt"`
}

func (WpDjaLove) TableName() string {
	return "wp_dja_love"
}
