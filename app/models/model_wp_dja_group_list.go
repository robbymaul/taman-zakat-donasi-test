package models

import "time"

type WpDjaGroupList struct {
	ID        int        `gorm:"column:id;primaryKey" json:"id"`
	GroupName *string    `gorm:"column:group_name" json:"group_name"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
}

func (WpDjaGroupList) TableName() string {
	return "wp_dja_group_list"
}
