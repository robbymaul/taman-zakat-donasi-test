package models

type WpLitespeedAvatar struct {
	ID       uint64 `gorm:"column:id;primaryKey" json:"id"`
	URL      string `gorm:"column:url" json:"url"`
	MD5      string `gorm:"column:md5" json:"md5"`
	Dateline int    `gorm:"column:dateline" json:"dateline"`
}

func (WpLitespeedAvatar) TableName() string {
	return "wp_litespeed_avatar"
}
