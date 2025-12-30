package models

import "time"

type WpDjaPasswordReset struct {
	ID          int        `gorm:"column:id;primaryKey" json:"id"`
	ResetEmail  *string    `gorm:"column:reset_email" json:"resetEmail"`
	ResetCode   *string    `gorm:"column:reset_code" json:"resetCode"`
	ResetStatus *int       `gorm:"column:reset_status" json:"resetStatus"`
	IP          *string    `gorm:"column:ip" json:"ip"`
	OS          *string    `gorm:"column:os" json:"os"`
	Browser     *string    `gorm:"column:browser" json:"browser"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"createdAt"`
}

func (WpDjaPasswordReset) TableName() string {
	return "wp_dja_password_reset"
}
