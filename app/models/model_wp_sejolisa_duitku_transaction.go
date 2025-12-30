package models

import "time"

type WpSejolisaDuitkuTransaction struct {
	ID        uint      `gorm:"column:ID;primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	LastCheck time.Time `gorm:"column:last_check" json:"lastCheck"`
	OrderID   int       `gorm:"column:order_id" json:"orderId"`
	Status    string    `gorm:"column:status" json:"status"`
	Detail    string    `gorm:"column:detail" json:"detail"`
}

func (WpSejolisaDuitkuTransaction) TableName() string {
	return "wp_sejolisa_duitku_transaction"
}
