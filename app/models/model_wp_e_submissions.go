package models

import "time"

type WpESubmission struct {
	ID                    uint64    `gorm:"column:id;primaryKey" json:"id"`
	Type                  *string   `gorm:"column:type" json:"type"`
	HashID                string    `gorm:"column:hash_id" json:"hashId"`
	MainMetaID            uint64    `gorm:"column:main_meta_id" json:"mainMetaId"`
	PostID                uint64    `gorm:"column:post_id" json:"postId"`
	Referer               string    `gorm:"column:referer" json:"referer"`
	RefererTitle          *string   `gorm:"column:referer_title" json:"refererTitle"`
	ElementID             string    `gorm:"column:element_id" json:"elementId"`
	FormName              string    `gorm:"column:form_name" json:"formName"`
	CampaignID            uint64    `gorm:"column:campaign_id" json:"campaignId"`
	UserID                *uint64   `gorm:"column:user_id" json:"userId"`
	UserIP                string    `gorm:"column:user_ip" json:"userIp"`
	UserAgent             string    `gorm:"column:user_agent" json:"userAgent"`
	ActionsCount          *int      `gorm:"column:actions_count" json:"actionsCount"`
	ActionsSucceededCount *int      `gorm:"column:actions_succeeded_count" json:"actionsSucceededCount"`
	Status                string    `gorm:"column:status" json:"status"`
	IsRead                int8      `gorm:"column:is_read" json:"isRead"`
	Meta                  *string   `gorm:"column:meta" json:"meta"`
	CreatedAtGMT          time.Time `gorm:"column:created_at_gmt" json:"createdAtGmt"`
	UpdatedAtGMT          time.Time `gorm:"column:updated_at_gmt" json:"updatedAtGmt"`
	CreatedAt             time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt             time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (WpESubmission) TableName() string {
	return "wp_e_submissions"
}
