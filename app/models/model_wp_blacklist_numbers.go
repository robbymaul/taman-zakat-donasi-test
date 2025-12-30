package models

import "time"

// WpBlacklistNumber represents table `wp_blacklist_numbers`.
type WpBlacklistNumber struct {
	ID          int        `gorm:"column:id;primaryKey" json:"id"`
	PhoneNumber string     `gorm:"column:phone_number" json:"phoneNumber"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"createdAt"`
}

func (WpBlacklistNumber) TableName() string {
	return "wp_blacklist_numbers"
}
