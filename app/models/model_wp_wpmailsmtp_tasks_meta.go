package models

import "time"

type WpWpmailsmtpTasksMeta struct {
	ID     uint64    `gorm:"column:id;primaryKey" json:"id"`
	Action string    `gorm:"column:action" json:"action"`
	Data   string    `gorm:"column:data" json:"data"`
	Date   time.Time `gorm:"column:date" json:"date"`
}

func (WpWpmailsmtpTasksMeta) TableName() string {
	return "wp_wpmailsmtp_tasks_meta"
}
