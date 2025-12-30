package models

type WpLitespeedURLFile struct {
	ID       uint64 `gorm:"column:id;primaryKey" json:"id"`
	URLID    uint64 `gorm:"column:url_id" json:"urlId"`
	Vary     string `gorm:"column:vary" json:"vary"`         // md5 of final vary
	Filename string `gorm:"column:filename" json:"filename"` // md5 of file content
	Type     uint8  `gorm:"column:type" json:"type"`         // css=1, js=2, ccss=3, ucss=4
	Mobile   uint8  `gorm:"column:mobile" json:"mobile"`     // mobile=1
	Webp     uint8  `gorm:"column:webp" json:"webp"`         // webp=1
	Expired  int    `gorm:"column:expired" json:"expired"`
}

func (WpLitespeedURLFile) TableName() string {
	return "wp_litespeed_url_file"
}
