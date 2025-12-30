package models

import "time"

// WpCriticalCssForWpURL represents table `wp_critical_css_for_wp_urls`.
//
// Notes:
// - created_at is datetime (no zero-date default in DDL), mapped as time.Time.
type WpCriticalCssForWpURL struct {
	ID          uint64    `gorm:"column:id;primaryKey" json:"id"`
	URLID       uint64    `gorm:"column:url_id" json:"urlId"`
	Type        *string   `gorm:"column:type" json:"type"`
	TypeName    *string   `gorm:"column:type_name" json:"typeName"`
	URL         string    `gorm:"column:url" json:"url"`
	Status      string    `gorm:"column:status" json:"status"`
	CachedName  *string   `gorm:"column:cached_name" json:"cachedName"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updatedAt"`
	FailedError string    `gorm:"column:failed_error" json:"failedError"`
}

func (WpCriticalCssForWpURL) TableName() string {
	return "wp_critical_css_for_wp_urls"
}
