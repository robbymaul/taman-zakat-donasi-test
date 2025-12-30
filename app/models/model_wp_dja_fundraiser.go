package models

import "time"

type WpDjaFundraiser struct {
	ID        int        `gorm:"column:id" json:"id"`
	CID       *int       `gorm:"column:c_id" json:"cId"`
	CTitle    *string    `gorm:"column:c_title" json:"cTitle"`
	CSlug     *string    `gorm:"column:c_slug" json:"cSlug"`
	CTarget   *int       `gorm:"column:c_target" json:"cTarget"`
	UserID    *int       `gorm:"column:user_id" json:"userId"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"`
}

func (WpDjaFundraiser) TableName() string {
	return "wp_dja_fundraiser"
}
