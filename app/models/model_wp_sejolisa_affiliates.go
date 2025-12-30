package models

import "time"

type WpSejolisaAffiliate struct {
	ID          uint      `gorm:"column:ID;primaryKey" json:"id"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt   time.Time `gorm:"column:deleted_at" json:"deletedAt"`
	OrderID     int       `gorm:"column:order_id" json:"orderID"`
	AffiliateID int       `gorm:"column:affiliate_id" json:"affiliateID"`
	ProductID   int       `gorm:"column:product_id" json:"productID"`
	Tier        int       `gorm:"column:tier" json:"tier"`
	Commission  float64   `gorm:"column:commission" json:"commission"`
	Status      string    `gorm:"column:status" json:"status"`
	PaidStatus  int       `gorm:"column:paid_status" json:"paidStatus"`
}

func (WpSejolisaAffiliate) TableName() string {
	return "wp_sejolisa_affiliates"
}
