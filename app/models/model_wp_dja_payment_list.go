package models

type WpDjaPaymentList struct {
	ID   int     `gorm:"column:id;primaryKey" json:"id"`
	Code *string `gorm:"column:code" json:"code"`
	Name *string `gorm:"column:name" json:"name"`
}

func (WpDjaPaymentList) TableName() string {
	return "wp_dja_payment_list"
}
