package models

import "time"

type WpWpmailsmtpDebugEvents struct {
	ID        uint64    `gorm:"column:id;primaryKey" json:"id"`
	Content   *string   `gorm:"column:content" json:"content"`
	Initiator *string   `gorm:"column:initiator" json:"initiator"`
	EventType uint8     `gorm:"column:event_type" json:"eventType"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
}

func (WpWpmailsmtpDebugEvents) TableName() string {
	return "wp_wpmailsmtp_debug_events"
}
