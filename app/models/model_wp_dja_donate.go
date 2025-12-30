package models

import "time"

type WpDjaDonate struct {
	ID                    int        `gorm:"column:id;primaryKey" json:"id"`
	CampaignID            *string    `gorm:"column:campaign_id" json:"campaignId"`
	InvoiceID             *string    `gorm:"column:invoice_id" json:"invoiceId"`
	UserID                *int64     `gorm:"column:user_id" json:"userId"`
	Name                  *string    `gorm:"column:name" json:"name"`
	Whatsapp              *string    `gorm:"column:whatsapp" json:"whatsapp"`
	Email                 *string    `gorm:"column:email" json:"email"`
	Comment               *string    `gorm:"column:comment" json:"comment"`
	Anonim                *int       `gorm:"column:anonim" json:"anonim"`
	PaymentMethod         *string    `gorm:"column:payment_method" json:"paymentMethod"`
	PaymentCode           *string    `gorm:"column:payment_code" json:"paymentCode"`
	PaymentNumber         *string    `gorm:"column:payment_number" json:"paymentNumber"`
	PaymentQrcode         *string    `gorm:"column:payment_qrcode" json:"paymentQrcode"`
	PaymentAccount        *string    `gorm:"column:payment_account" json:"paymentAccount"`
	UniqueNumber          *float64   `gorm:"column:unique_number" json:"uniqueNumber"`
	Nominal               *float64   `gorm:"column:nominal" json:"nominal"`
	Status                *int       `gorm:"column:status" json:"status"`
	CsID                  *int       `gorm:"column:cs_id" json:"csId"`
	F0                    *int       `gorm:"column:f0" json:"f0"`
	F1                    *int       `gorm:"column:f1" json:"f1"`
	F2                    *int       `gorm:"column:f2" json:"f2"`
	F3                    *int       `gorm:"column:f3" json:"f3"`
	F4                    *int       `gorm:"column:f4" json:"f4"`
	F5                    *int       `gorm:"column:f5" json:"f5"`
	PaymentTrxID          *string    `gorm:"column:payment_trx_id" json:"paymentTrxId"`
	PaymentAt             *time.Time `gorm:"column:payment_at" json:"paymentAt"`
	ProcessBy             *string    `gorm:"column:process_by" json:"processBy"`
	CreatedAt             *time.Time `gorm:"column:created_at" json:"createdAt"`
	DeeplinkURL           *string    `gorm:"column:deeplink_url" json:"deeplinkUrl"`
	PaymentGateway        *string    `gorm:"column:payment_gateway" json:"paymentGateway"`
	MainDonate            *float64   `gorm:"column:main_donate" json:"mainDonate"`
	InfoDonate            *string    `gorm:"column:info_donate" json:"infoDonate"`
	ImgConfirmationURL    *string    `gorm:"column:img_confirmation_url" json:"imgConfirmationUrl"`
	ImgConfirmationStatus *int       `gorm:"column:img_confirmation_status" json:"imgConfirmationStatus"`
	ImgConfirmationDate   *time.Time `gorm:"column:img_confirmation_date" json:"imgConfirmationDate"`
	InfoQurban            *string    `gorm:"column:info_qurban" json:"infoQurban"`
	InfoPackage2          *string    `gorm:"column:info_package2" json:"infoPackage2"`
	InfoZfitrah           *string    `gorm:"column:info_zfitrah" json:"infoZfitrah"`
	InfoAddformula        *string    `gorm:"column:info_addformula" json:"infoAddformula"`
	IP                    *string    `gorm:"column:ip" json:"ip"`
	OS                    *string    `gorm:"column:os" json:"os"`
	Browser               *string    `gorm:"column:browser" json:"browser"`
	Mobdesktop            *string    `gorm:"column:mobdesktop" json:"mobdesktop"`
	Sapaan                *string    `gorm:"column:sapaan" json:"sapaan"`
	Currency              *string    `gorm:"column:currency" json:"currency"`
	UtmSource             *string    `gorm:"column:utm_source" json:"utmSource"`
	UtmMedium             *string    `gorm:"column:utm_medium" json:"utmMedium"`
	UtmContent            *string    `gorm:"column:utm_content" json:"utmContent"`
	UtmCampaign           *string    `gorm:"column:utm_campaign" json:"utmCampaign"`
	UtmTerm               *string    `gorm:"column:utm_term" json:"utmTerm"`
	UtmID                 *string    `gorm:"column:utm_id" json:"utmId"`
}

func (WpDjaDonate) TableName() string {
	return "wp_dja_donate"
}
