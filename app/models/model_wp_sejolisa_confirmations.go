package models

import "time"

type WpSejolisaConfirmation struct {
	ID        uint      `gorm:"column:ID;primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	OrderID   int       `gorm:"column:order_id" json:"orderId"`
	ProductID int       `gorm:"column:product_id" json:"productId"`
	UserID    int       `gorm:"column:user_id" json:"userId"`
	Detail    string    `gorm:"column:detail" json:"detail"`
}

func (WpSejolisaConfirmation) TableName() string {
	return "wp_sejolisa_confirmations"
}
