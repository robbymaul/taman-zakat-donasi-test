package models

type WpDjaCategory struct {
	ID              int     `gorm:"column:id;primaryKey" json:"id"`
	Category        *string `gorm:"column:category" json:"category"`
	PrivateCategory *int    `gorm:"column:private_category" json:"privateCategory"`
}

func (WpDjaCategory) TableName() string {
	return "wp_dja_category"
}
