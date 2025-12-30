package models

import (
	"database/sql"
)

type WpWprRocketCache struct {
	ID           uint64       `gorm:"column:id;primaryKey" json:"id"`
	URL          string       `gorm:"column:url" json:"url"`
	Status       string       `gorm:"column:status" json:"status"`
	Modified     sql.NullTime `gorm:"column:modified" json:"modified"`
	LastAccessed sql.NullTime `gorm:"column:last_accessed" json:"lastAccessed"`
	IsLocked     bool         `gorm:"column:is_locked" json:"isLocked"`
}

func (WpWprRocketCache) TableName() string {
	return "wp_wpr_rocket_cache"
}
