package models

type WpLitespeedURL struct {
	ID        uint64 `gorm:"column:id;primaryKey" json:"id"`
	URL       string `gorm:"column:url" json:"url"`
	CacheTags string `gorm:"column:cache_tags" json:"cacheTags"`
}

func (WpLitespeedURL) TableName() string {
	return "wp_litespeed_url"
}
