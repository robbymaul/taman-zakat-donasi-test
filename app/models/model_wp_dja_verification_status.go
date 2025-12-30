package models

type WpDjaVerificationStatus struct {
	ID     int     `gorm:"column:id;primaryKey" json:"id"`
	Status *int    `gorm:"column:status" json:"status"`
	Name   *string `gorm:"column:name" json:"name"`
}

func (WpDjaVerificationStatus) TableName() string {
	return "wp_dja_verification_status"
}
