package models

import "time"

type CustomFollowupScheduler struct {
	ID            int        `gorm:"column:id;primaryKey" json:"id"`
	Status        *string    `gorm:"column:status" json:"status"`
	IsFirstRun    int        `gorm:"column:is_first_run" json:"isFirstRun"`
	CodeID        *string    `gorm:"column:code_id" json:"codeId"`
	Name          *string    `gorm:"column:name" json:"name"`
	Schedule      *string    `gorm:"column:schedule" json:"schedule"`
	Format        *int       `gorm:"column:format" json:"format"`
	DateTime      *time.Time `gorm:"column:date_time" json:"dateTime"`
	CustomFormat  *string    `gorm:"column:custom_format" json:"customFormat"`
	Target        *string    `gorm:"column:target" json:"target"`
	GroupID       *string    `gorm:"column:group_id" json:"groupId"`
	CampaignID    *string    `gorm:"column:campaign_id" json:"campaignId"`
	DonaturStatus *string    `gorm:"column:donatur_status" json:"donaturStatus"`
	BroadcastMode *string    `gorm:"column:broadcast_mode" json:"broadcastMode"`
	BatchSize     *int       `gorm:"column:batch_size" json:"batchSize"`
	SendDelay     *float64   `gorm:"column:send_delay" json:"sendDelay"`
	CreatedAt     *time.Time `gorm:"column:created_at" json:"createdAt"`
}

// TableName overrides the default table name
func (CustomFollowupScheduler) TableName() string {
	return "wp_wp_dja_custom_followup_scheduler"
}
