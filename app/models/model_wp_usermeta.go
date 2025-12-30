package models

type WpUserMeta struct {
	UmetaID   uint64  `gorm:"column:umeta_id;primaryKey" json:"umetaId"`
	UserID    uint64  `gorm:"column:user_id" json:"userId"`
	MetaKey   *string `gorm:"column:meta_key" json:"metaKey"`
	MetaValue *string `gorm:"column:meta_value" json:"metaValue"`
}

func (WpUserMeta) TableName() string {
	return "wp_usermeta"
}
