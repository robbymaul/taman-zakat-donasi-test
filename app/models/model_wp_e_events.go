package models

import "time"

type WpEEvent struct {
	ID        uint64    `gorm:"column:id;primaryKey" json:"id"`
	EventData string    `gorm:"column:event_data" json:"eventData"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
}

func (WpEEvent) TableName() string {
	return "wp_e_events"
}
