package models

import "time"

// WpSejolisaSubscription represents table `wp_sejolisa_subscriptions`.
//
// Notes:
// - Some datetime columns use MySQL zero-date default ('0000-00-00 00:00:00'), mapped as string to avoid parsing issues.
type WpSejolisaSubscription struct {
	ID            uint      `gorm:"column:ID;primaryKey" json:"id"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt     string    `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt     string    `gorm:"column:deleted_at" json:"deletedAt"`
	OrderID       int       `gorm:"column:order_id" json:"orderId"`
	OrderParentID int       `gorm:"column:order_parent_id" json:"orderParentId"`
	UserID        int       `gorm:"column:user_id" json:"userId"`
	ProductID     int       `gorm:"column:product_id" json:"productId"`
	Type          string    `gorm:"column:type" json:"type"`
	EndDate       string    `gorm:"column:end_date" json:"endDate"`
	Status        string    `gorm:"column:status" json:"status"`
}

func (WpSejolisaSubscription) TableName() string {
	return "wp_sejolisa_subscriptions"
}
