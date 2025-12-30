package models

import "time"

type WpSejolisaOrder struct {
	ID             uint       `gorm:"column:ID;primaryKey" json:"id"`
	CreatedAt      time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" json:"updatedAt"` // default 0000-00-00 → nullable
	DeletedAt      *time.Time `gorm:"column:deleted_at" json:"deletedAt"` // default 0000-00-00 → nullable
	OrderParentID  int        `gorm:"column:order_parent_id" json:"orderParentId"`
	ProductID      int        `gorm:"column:product_id" json:"productId"`
	UserID         int        `gorm:"column:user_id" json:"userId"`
	AffiliateID    int        `gorm:"column:affiliate_id" json:"affiliateId"`
	CouponID       int        `gorm:"column:coupon_id" json:"couponId"`
	PaymentGateway string     `gorm:"column:payment_gateway" json:"paymentGateway"`
	GrandTotal     float64    `gorm:"column:grand_total" json:"grandTotal"`
	Quantity       int        `gorm:"column:quantity" json:"quantity"`
	Type           string     `gorm:"column:type" json:"type"`
	Status         string     `gorm:"column:status" json:"status"`
	MetaData       string     `gorm:"column:meta_data" json:"metaData"`
}

func (WpSejolisaOrder) TableName() string {
	return "wp_sejolisa_orders"
}
