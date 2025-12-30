package models

import "time"

type WpSejolisaAcquisition struct {
	ID          uint      `gorm:"column:ID;primaryKey" json:"id"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt"`
	AffiliateID int       `gorm:"column:affiliate_id" json:"affiliateID"`
	ProductID   int       `gorm:"column:product_id" json:"productID"`
	Source      string    `gorm:"column:source" json:"source"`
	Media       string    `gorm:"column:media" json:"media"`
	View        int       `gorm:"column:view" json:"view"`
	Lead        int       `gorm:"column:lead" json:"lead"`
	Sales       int       `gorm:"column:sales" json:"sales"`
}

func (WpSejolisaAcquisition) TableName() string {
	return "wp_sejolisa_acquisition"
}
