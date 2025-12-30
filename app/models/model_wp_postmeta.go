package models

type WpPostMeta struct {
	MetaID    uint64 `gorm:"column:meta_id;primaryKey" json:"metaID"`
	PostID    uint64 `gorm:"column:post_id" json:"postID"`
	MetaKey   string `gorm:"column:meta_key" json:"metaKey"`
	MetaValue string `gorm:"column:meta_value" json:"metaValue"`
}

func (WpPostMeta) TableName() string {
	return "wp_postmeta"
}
