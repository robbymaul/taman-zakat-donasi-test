package models

import "time"

type WpDjaUsers struct {
	ID               int     `gorm:"column:id;primaryKey" json:"id"`
	UserID           *int64  `gorm:"column:user_id" json:"userId"`
	UserRandid       *string `gorm:"column:user_randid" json:"userRandid"`
	UserType         *string `gorm:"column:user_type" json:"userType"`
	UserVerification *int    `gorm:"column:user_verification" json:"userVerification"`
	UserBio          *string `gorm:"column:user_bio" json:"userBio"`
	UserWa           *string `gorm:"column:user_wa" json:"userWa"`
	UserProvinsi     *string `gorm:"column:user_provinsi" json:"userProvinsi"`
	UserKabkota      *string `gorm:"column:user_kabkota" json:"userKabkota"`
	UserKecamatan    *string `gorm:"column:user_kecamatan" json:"userKecamatan"`

	UserProvinsiID  *int `gorm:"column:user_provinsi_id" json:"userProvinsiId"`
	UserKabkotaID   *int `gorm:"column:user_kabkota_id" json:"userKabkotaId"`
	UserKecamatanID *int `gorm:"column:user_kecamatan_id" json:"userKecamatanId"`

	UserAlamat   *string `gorm:"column:user_alamat" json:"userAlamat"`
	UserBankName *string `gorm:"column:user_bank_name" json:"userBankName"`
	UserBankNo   *string `gorm:"column:user_bank_no" json:"userBankNo"`
	UserBankAn   *string `gorm:"column:user_bank_an" json:"userBankAn"`

	UserPpImg    *string `gorm:"column:user_pp_img" json:"userPpImg"`
	UserCoverImg *string `gorm:"column:user_cover_img" json:"userCoverImg"`

	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"`

	UserAnonimF     *int    `gorm:"column:user_anonim_f" json:"userAnonimF"`
	UserCommissionF *int    `gorm:"column:user_commission_f" json:"userCommissionF"`
	UserSapaan      *string `gorm:"column:user_sapaan" json:"userSapaan"`
}

func (WpDjaUsers) TableName() string {
	return "wp_dja_users"
}
