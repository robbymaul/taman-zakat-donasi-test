package models

import "time"

// WpDjaAffPayout represents table `wp_dja_aff_payout`.
//
// Notes:
// - nominal_payout (decimal(21,0)) is mapped as string to avoid precision loss without extra decimal dependencies.
type WpDjaAffPayout struct {
	ID            int        `gorm:"column:id;primaryKey" json:"id"`
	PayoutNumber  *string    `gorm:"column:payout_number" json:"payoutNumber"`
	UserID        *int64     `gorm:"column:user_id" json:"userId"`
	AffSubmitID   *string    `gorm:"column:aff_submit_id" json:"affSubmitId"`
	NominalPayout *string    `gorm:"column:nominal_payout" json:"nominalPayout"`
	Status        *int       `gorm:"column:status" json:"status"`
	Image         *string    `gorm:"column:image" json:"image"`
	BankName      *string    `gorm:"column:bank_name" json:"bankName"`
	BankNo        *string    `gorm:"column:bank_no" json:"bankNo"`
	BankAn        *string    `gorm:"column:bank_an" json:"bankAn"`
	CreatedAt     *time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt     *time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (WpDjaAffPayout) TableName() string {
	return "wp_dja_aff_payout"
}
