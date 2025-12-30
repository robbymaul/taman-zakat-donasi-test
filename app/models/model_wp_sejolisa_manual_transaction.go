package models

import "time"

type WpSejolisaManualTransaction struct {
	ID         uint      `gorm:"column:ID;primaryKey" json:"id"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updatedAt"`
	OrderID    int       `gorm:"column:order_id" json:"orderId"`
	UserID     *int      `gorm:"column:user_id" json:"userId"` // nullable
	Bank       string    `gorm:"column:bank" json:"bank"`
	Total      float64   `gorm:"column:total" json:"total"`
	UniqueCode int       `gorm:"column:unique_code" json:"uniqueCode"`
	MetaData   string    `gorm:"column:meta_data" json:"metaData"`
}

func (WpSejolisaManualTransaction) TableName() string {
	return "wp_sejolisa_manual_transaction"
}
