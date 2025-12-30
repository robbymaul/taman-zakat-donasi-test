package models

import "time"

type WpDjaGroupData struct {
	ID        int        `gorm:"column:id;primaryKey" json:"id"`
	GroupID   string     `gorm:"column:group_id" json:"group_id"`
	Name      *string    `gorm:"column:name" json:"name"`
	Whatsapp  *string    `gorm:"column:whatsapp" json:"whatsapp"`
	Sapaan    *string    `gorm:"column:sapaan" json:"sapaan"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
}

func (WpDjaGroupData) TableName() string {
	return "wp_dja_group_data"
}
