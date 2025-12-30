package models

import "time"

type WpDjaRegister struct {
	ID           int        `gorm:"column:id;primaryKey" json:"id"`
	RNamaLengkap *string    `gorm:"column:r_nama_lengkap" json:"rNamaLengkap"`
	REmail       *string    `gorm:"column:r_email" json:"rEmail"`
	RWhatsapp    *string    `gorm:"column:r_whatsapp" json:"rWhatsapp"`
	RPassword    *string    `gorm:"column:r_password" json:"rPassword"`
	RCode        *string    `gorm:"column:r_code" json:"rCode"`
	RStatus      *int       `gorm:"column:r_status" json:"rStatus"`
	CreatedAt    *time.Time `gorm:"column:created_at" json:"createdAt"`
}

func (WpDjaRegister) TableName() string {
	return "wp_dja_register"
}
