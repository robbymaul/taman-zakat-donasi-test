package models

import "time"

type WpDjaPaymentLog struct {
	ID        int        `gorm:"column:id;primaryKey" json:"id"`
	IdDonate  *int       `gorm:"column:id_donate" json:"idDonate"`
	Hit       *string    `gorm:"column:hit" json:"hit"`
	Status    *string    `gorm:"column:status" json:"status"`
	Message   *string    `gorm:"column:message" json:"message"`
	Log       *string    `gorm:"column:log" json:"log"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"`
}

func (WpDjaPaymentLog) TableName() string {
	return "wp_dja_payment_log"
}
