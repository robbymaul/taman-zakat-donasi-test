package models

import "time"

type WpSejolisaLicense struct {
	ID        uint      `gorm:"column:ID;primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deletedAt"`
	OrderID   int       `gorm:"column:order_id" json:"orderId"`
	UserID    int       `gorm:"column:user_id" json:"userId"`
	ProductID int       `gorm:"column:product_id" json:"productId"`
	Code      string    `gorm:"column:code" json:"code"`
	String    string    `gorm:"column:string" json:"string"`
	Status    string    `gorm:"column:status" json:"status"`
	MetaData  string    `gorm:"column:meta_data" json:"metaData"`
}

func (WpSejolisaLicense) TableName() string {
	return "wp_sejolisa_licenses"
}
