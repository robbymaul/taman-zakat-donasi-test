package models

import "time"

type WpDjaDonateTrash struct {
	ID                    int        `gorm:"column:id" json:"id"`
	WhoDelID              *int       `gorm:"column:who_del_id" json:"whoDelId"`
	WhoDelName            string     `gorm:"column:who_del_name" json:"whoDelName"`
	DonateID              *int       `gorm:"column:donate_id" json:"donateId"`
	CampaignID            string     `gorm:"column:campaign_id" json:"campaignId"`
	InvoiceID             string     `gorm:"column:invoice_id" json:"invoiceId"`
	UserID                *int64     `gorm:"column:user_id" json:"userId"`
	Name                  string     `gorm:"column:name" json:"name"`
	Whatsapp              string     `gorm:"column:whatsapp" json:"whatsapp"`
	Email                 string     `gorm:"column:email" json:"email"`
	Comment               string     `gorm:"column:comment" json:"comment"`
	Anonim                *int       `gorm:"column:anonim" json:"anonim"`
	PaymentMethod         string     `gorm:"column:payment_method" json:"paymentMethod"`
	PaymentCode           string     `gorm:"column:payment_code" json:"paymentCode"`
	PaymentNumber         string     `gorm:"column:payment_number" json:"paymentNumber"`
	PaymentQrcode         string     `gorm:"column:payment_qrcode" json:"paymentQrcode"`
	PaymentAccount        string     `gorm:"column:payment_account" json:"paymentAccount"`
	UniqueNumber          string     `gorm:"column:unique_number" json:"uniqueNumber"`
	Nominal               string     `gorm:"column:nominal" json:"nominal"`
	MainDonate            string     `gorm:"column:main_donate" json:"mainDonate"`
	InfoDonate            string     `gorm:"column:info_donate" json:"infoDonate"`
	InfoQurban            string     `gorm:"column:info_qurban" json:"infoQurban"`
	InfoPackage2          string     `gorm:"column:info_package2" json:"infoPackage2"`
	InfoZfitrah           string     `gorm:"column:info_zfitrah" json:"infoZfitrah"`
	InfoAddformula        string     `gorm:"column:info_addformula" json:"infoAddformula"`
	Status                *int       `gorm:"column:status" json:"status"`
	CsID                  *int       `gorm:"column:cs_id" json:"csId"`
	F0                    *int       `gorm:"column:f0" json:"f0"`
	F1                    *int       `gorm:"column:f1" json:"f1"`
	F2                    *int       `gorm:"column:f2" json:"f2"`
	F3                    *int       `gorm:"column:f3" json:"f3"`
	F4                    *int       `gorm:"column:f4" json:"f4"`
	F5                    *int       `gorm:"column:f5" json:"f5"`
	PaymentTrxID          string     `gorm:"column:payment_trx_id" json:"paymentTrxId"`
	PaymentAt             *time.Time `gorm:"column:payment_at" json:"paymentAt"`
	ProcessBy             string     `gorm:"column:process_by" json:"processBy"`
	DeeplinkURL           string     `gorm:"column:deeplink_url" json:"deeplinkUrl"`
	PaymentGateway        string     `gorm:"column:payment_gateway" json:"paymentGateway"`
	ImgConfirmationURL    string     `gorm:"column:img_confirmation_url" json:"imgConfirmationUrl"`
	ImgConfirmationStatus *int       `gorm:"column:img_confirmation_status" json:"imgConfirmationStatus"`
	ImgConfirmationDate   *time.Time `gorm:"column:img_confirmation_date" json:"imgConfirmationDate"`
	IP                    string     `gorm:"column:ip" json:"ip"`
	OS                    string     `gorm:"column:os" json:"os"`
	Browser               string     `gorm:"column:browser" json:"browser"`
	Mobdesktop            string     `gorm:"column:mobdesktop" json:"mobdesktop"`
	Sapaan                string     `gorm:"column:sapaan" json:"sapaan"`
	Currency              string     `gorm:"column:currency" json:"currency"`
	CreatedAt             *time.Time `gorm:"column:created_at" json:"createdAt"`
	DeletedAt             *time.Time `gorm:"column:deleted_at" json:"deletedAt"`
	ApprovedAt            *time.Time `gorm:"column:approved_at" json:"approvedAt"`
}

func (WpDjaDonateTrash) TableName() string {
	return "wp_dja_donate_trash"
}
