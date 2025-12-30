package models

import "time"

type WpSejolisaBcaTransaction struct {
	ID         uint      `gorm:"column:ID;primaryKey" json:"id"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"UpdatedAt"`
	OrderID    int       `gorm:"column:order_id" json:"orderID"`
	UserID     *int      `gorm:"column:user_id" json:"userID"` // nullable
	Account    string    `gorm:"column:account" json:"account"`
	Total      float64   `gorm:"column:total" json:"total"`
	UniqueCode int       `gorm:"column:unique_code" json:"uniqueCode"`
	MetaData   string    `gorm:"column:meta_data" json:"metaData"`
}

func (WpSejolisaBcaTransaction) TableName() string {
	return "wp_sejolisa_bca_transaction"
}
