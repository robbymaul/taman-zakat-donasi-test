package models

import "time"

type WpDjaPaymentCallback struct {
	ID        int        `gorm:"column:id;primaryKey" json:"id"`
	IDDonate  *int       `gorm:"column:id_donate" json:"idDonate"`
	PG        *string    `gorm:"column:pg" json:"pg"`
	Callback  *string    `gorm:"column:callback" json:"callback"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"`
}

func (WpDjaPaymentCallback) TableName() string {
	return "wp_dja_payment_callback"
}
