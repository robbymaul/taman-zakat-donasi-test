package models

import "time"

type WpDjaSettings struct {
	ID        int        `gorm:"column:id;primaryKey" json:"id"`
	Type      *string    `gorm:"column:type" json:"type"`
	Data      *string    `gorm:"column:data" json:"data"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"`
}

func (WpDjaSettings) TableName() string {
	return "wp_dja_settings"
}
