package models

import "time"

type WpESubmissionActionLog struct {
	ID           uint64    `gorm:"column:id;primaryKey" json:"id"`
	SubmissionID uint64    `gorm:"column:submission_id" json:"submissionId"`
	ActionName   string    `gorm:"column:action_name" json:"actionName"`
	ActionLabel  *string   `gorm:"column:action_label" json:"actionLabel"`
	Status       string    `gorm:"column:status" json:"status"`
	Log          *string   `gorm:"column:log" json:"log"`
	CreatedAtGMT time.Time `gorm:"column:created_at_gmt" json:"createdAtGmt"`
	UpdatedAtGMT time.Time `gorm:"column:updated_at_gmt" json:"updatedAtGmt"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (WpESubmissionActionLog) TableName() string {
	return "wp_e_submissions_actions_log"
}
