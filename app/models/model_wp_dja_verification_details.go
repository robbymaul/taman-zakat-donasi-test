package models

import "time"

type WpDjaVerificationDetails struct {
	ID               int        `gorm:"column:id;primaryKey" json:"id"`
	UID              *int       `gorm:"column:u_id" json:"uId"`
	UNamaLengkap     *string    `gorm:"column:u_nama_lengkap" json:"uNamaLengkap"`
	UEmail           *string    `gorm:"column:u_email" json:"uEmail"`
	UWhatsapp        *string    `gorm:"column:u_whatsapp" json:"uWhatsapp"`
	UKtp             *string    `gorm:"column:u_ktp" json:"uKtp"`
	UKtpSelfie       *string    `gorm:"column:u_ktp_selfie" json:"uKtpSelfie"`
	UJabatan         *string    `gorm:"column:u_jabatan" json:"uJabatan"`
	UNamaKetua       *string    `gorm:"column:u_nama_ketua" json:"uNamaKetua"`
	UAlamatLengkap   *string    `gorm:"column:u_alamat_lengkap" json:"uAlamatLengkap"`
	UProgramUnggulan *string    `gorm:"column:u_program_unggulan" json:"uProgramUnggulan"`
	UProfile         *string    `gorm:"column:u_profile" json:"uProfile"`
	ULegalitas       *string    `gorm:"column:u_legalitas" json:"uLegalitas"`
	CreatedAt        *time.Time `gorm:"column:created_at" json:"createdAt"`
}

func (WpDjaVerificationDetails) TableName() string {
	return "wp_dja_verification_details"
}
