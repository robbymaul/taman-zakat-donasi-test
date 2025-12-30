package models

import "time"

type WpSejolisaCoupon struct {
	ID             uint       `gorm:"column:ID;primaryKey" json:"id"`
	CreatedAt      time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt      time.Time  `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt      time.Time  `gorm:"column:deleted_at" json:"deletedAt"`
	UserID         int        `gorm:"column:user_id" json:"userId"`
	CouponParentID int        `gorm:"column:coupon_parent_id" json:"couponParentId"`
	Code           string     `gorm:"column:code" json:"code"`
	Rule           string     `gorm:"column:rule" json:"rule"`
	Discount       string     `gorm:"column:discount" json:"discount"`
	Usage          int        `gorm:"column:usage" json:"usage"`
	LimitUse       int        `gorm:"column:limit_use" json:"limitUse"`
	LimitDate      *time.Time `gorm:"column:limit_date" json:"limitDate"`
	Status         string     `gorm:"column:status" json:"status"`
}

func (WpSejolisaCoupon) TableName() string {
	return "wp_sejolisa_coupons"
}
