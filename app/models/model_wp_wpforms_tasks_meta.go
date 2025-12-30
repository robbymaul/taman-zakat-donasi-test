package models

import "time"

type WpWpformsTasksMeta struct {
	ID     uint64    `gorm:"column:id;primaryKey" json:"id"`
	Action string    `gorm:"column:action" json:"action"`
	Data   string    `gorm:"column:data" json:"data"`
	Date   time.Time `gorm:"column:date" json:"date"`
}

func (WpWpformsTasksMeta) TableName() string {
	return "wp_wpforms_tasks_meta"
}
