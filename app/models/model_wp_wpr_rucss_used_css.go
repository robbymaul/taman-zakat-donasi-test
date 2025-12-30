package models

import (
	"database/sql"
)

type WpWprRucssUsedCss struct {
	ID             uint64       `gorm:"column:id;primaryKey" json:"id"`
	URL            string       `gorm:"column:url" json:"url"`
	CSS            *string      `gorm:"column:css" json:"css"`
	Hash           *string      `gorm:"column:hash" json:"hash"`
	ErrorCode      *string      `gorm:"column:error_code" json:"errorCode"`
	ErrorMessage   *string      `gorm:"column:error_message" json:"errorMessage"`
	UnprocessedCSS *string      `gorm:"column:unprocessedcss" json:"unprocessedCss"`
	Retries        bool         `gorm:"column:retries" json:"retries"`
	IsMobile       bool         `gorm:"column:is_mobile" json:"isMobile"`
	JobID          string       `gorm:"column:job_id" json:"jobId"`
	QueueName      string       `gorm:"column:queue_name" json:"queueName"`
	Status         string       `gorm:"column:status" json:"status"`
	Modified       sql.NullTime `gorm:"column:modified" json:"modified"`
	LastAccessed   sql.NullTime `gorm:"column:last_accessed" json:"lastAccessed"`
}

func (WpWprRucssUsedCss) TableName() string {
	return "wp_wpr_rucss_used_css"
}
