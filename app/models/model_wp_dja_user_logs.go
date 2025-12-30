package models

import "time"

type WpDjaUserLogs struct {
	ID               int     `gorm:"column:id;primaryKey" json:"id"`
	DelBy            *int    `gorm:"column:del_by" json:"delBy"`
	UserLogin        *string `gorm:"column:user_login" json:"userLogin"`
	UserEmail        *string `gorm:"column:user_email" json:"userEmail"`
	UserID           *int64  `gorm:"column:user_id" json:"userId"`
	UserFullname     *string `gorm:"column:user_fullname" json:"userFullname"`
	UserRandid       *string `gorm:"column:user_randid" json:"userRandid"`
	UserType         *string `gorm:"column:user_type" json:"userType"`
	UserVerification *int    `gorm:"column:user_verification" json:"userVerification"`

	UserBio       *string `gorm:"column:user_bio" json:"userBio"`
	UserWa        *string `gorm:"column:user_wa" json:"userWa"`
	UserProvinsi  *string `gorm:"column:user_provinsi" json:"userProvinsi"`
	UserKabkota   *string `gorm:"column:user_kabkota" json:"userKabkota"`
	UserKecamatan *string `gorm:"column:user_kecamatan" json:"userKecamatan"`

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
}

func (WpDjaUserLogs) TableName() string {
	return "wp_dja_user_logs"
}
