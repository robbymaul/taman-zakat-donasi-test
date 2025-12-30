package models

// WpTermmeta represents table `wp_termmeta`.
type WpTermmeta struct {
	MetaID    uint64  `gorm:"column:meta_id;primaryKey" json:"metaId"`
	TermID    uint64  `gorm:"column:term_id" json:"termId"`
	MetaKey   *string `gorm:"column:meta_key" json:"metaKey"`
	MetaValue *string `gorm:"column:meta_value" json:"metaValue"`
}

func (WpTermmeta) TableName() string {
	return "wp_termmeta"
}
