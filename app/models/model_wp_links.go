package models

type WpLink struct {
	LinkID          uint64 `gorm:"column:link_id;primaryKey" json:"linkId"`
	LinkURL         string `gorm:"column:link_url" json:"linkUrl"`
	LinkName        string `gorm:"column:link_name" json:"linkName"`
	LinkImage       string `gorm:"column:link_image" json:"linkImage"`
	LinkTarget      string `gorm:"column:link_target" json:"linkTarget"`
	LinkDescription string `gorm:"column:link_description" json:"linkDescription"`
	LinkVisible     string `gorm:"column:link_visible" json:"linkVisible"`
	LinkOwner       uint64 `gorm:"column:link_owner" json:"linkOwner"`
	LinkRating      int    `gorm:"column:link_rating" json:"linkRating"`
	LinkUpdated     string `gorm:"column:link_updated" json:"linkUpdated"` // bisa juga time.Time jika ingin konversi datetime
	LinkRel         string `gorm:"column:link_rel" json:"linkRel"`
	LinkNotes       string `gorm:"column:link_notes" json:"linkNotes"`
	LinkRSS         string `gorm:"column:link_rss" json:"linkRss"`
}

func (WpLink) TableName() string {
	return "wp_links"
}
