package models

type WpDjaUserType struct {
	ID   int     `gorm:"column:id;primaryKey" json:"id"`
	Type *string `gorm:"column:type" json:"type"`
	Name *string `gorm:"column:name" json:"name"`
}

func (WpDjaUserType) TableName() string {
	return "wp_dja_user_type"
}
