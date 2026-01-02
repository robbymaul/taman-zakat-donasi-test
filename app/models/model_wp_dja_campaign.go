package models

import "time"

type WpDjaCampaign struct {
	ID                          int        `gorm:"column:id;primaryKey" json:"id"`
	CampaignID                  *string    `gorm:"column:campaign_id" json:"campaignId"`
	Title                       *string    `gorm:"column:title" json:"title"`
	Slug                        *string    `gorm:"column:slug" json:"slug"`
	Target                      *float64   `gorm:"column:target" json:"target"`
	ImageURL                    *string    `gorm:"column:image_url" json:"imageUrl"`
	Information                 *string    `gorm:"column:information" json:"information"`
	LocationName                *string    `gorm:"column:location_name" json:"locationName"`
	LocationGmaps               *string    `gorm:"column:location_gmaps" json:"locationGmaps"`
	PublishStatus               *int       `gorm:"column:publish_status" json:"publishStatus"`
	ReachedStatus               *int       `gorm:"column:reached_status" json:"reachedStatus"`
	EndDate                     *time.Time `gorm:"column:end_date" json:"endDate"`
	FormType                    *string    `gorm:"column:form_type" json:"formType"`
	Packaged                    *int       `gorm:"column:packaged" json:"packaged"`
	PackagedTitle               *string    `gorm:"column:packaged_title" json:"packagedTitle"`
	Currency                    *string    `gorm:"column:currency" json:"currency"`
	CategoryID                  *int       `gorm:"column:category_id" json:"categoryId"`
	UserID                      *int64     `gorm:"column:user_id" json:"userId"`
	FundraiserID                *int64     `gorm:"column:fundraiser_id" json:"fundraiserId"`
	PaymentStatus               *int       `gorm:"column:payment_status" json:"paymentStatus"`
	MethodStatus                *string    `gorm:"column:method_status" json:"methodStatus"`
	BankAccount                 *string    `gorm:"column:bank_account" json:"bankAccount"`
	FormStatus                  *int       `gorm:"column:form_status" json:"formStatus"`
	FormText                    *string    `gorm:"column:form_text" json:"formText"`
	UniqueNumberSetting         *int       `gorm:"column:unique_number_setting" json:"uniqueNumberSetting"`
	UniqueNumberValue           *string    `gorm:"column:unique_number_value" json:"uniqueNumberValue"`
	NotificationStatus          *int       `gorm:"column:notification_status" json:"notificationStatus"`
	WanotifMessage              *string    `gorm:"column:wanotif_message" json:"wanotifMessage"`
	PixelStatus                 *int       `gorm:"column:pixel_status" json:"pixelStatus"`
	FbPixel                     *string    `gorm:"column:fb_pixel" json:"fbPixel"`
	FbEvent                     *string    `gorm:"column:fb_event" json:"fbEvent"`
	CreatedAt                   *time.Time `gorm:"column:created_at" json:"createdAt"`
	FormBase                    *int       `gorm:"column:form_base" json:"formBase"`
	PengeluaranSetting          *int       `gorm:"column:pengeluaran_setting" json:"pengeluaranSetting"`
	PengeluaranTitle            *string    `gorm:"column:pengeluaran_title" json:"pengeluaranTitle"`
	GtmStatus                   *int       `gorm:"column:gtm_status" json:"gtmStatus"`
	GtmID                       *string    `gorm:"column:gtm_id" json:"gtmId"`
	TiktokStatus                *int       `gorm:"column:tiktok_status" json:"tiktokStatus"`
	TiktokPixel                 *string    `gorm:"column:tiktok_pixel" json:"tiktokPixel"`
	Socialproof                 *int       `gorm:"column:socialproof" json:"socialproof"`
	SocialproofText             *string    `gorm:"column:socialproof_text" json:"socialproofText"`
	SocialproofPosition         *string    `gorm:"column:socialproof_position" json:"socialproofPosition"`
	ZakatSetting                *int       `gorm:"column:zakat_setting" json:"zakatSetting"`
	ZakatPercent                *float64   `gorm:"column:zakat_percent" json:"zakatPercent"`
	FundraiserSetting           *int       `gorm:"column:fundraiser_setting" json:"fundraiserSetting"`
	FundraiserOn                *int       `gorm:"column:fundraiser_on" json:"fundraiserOn"`
	FundraiserText              *string    `gorm:"column:fundraiser_text" json:"fundraiserText"`
	FundraiserButton            *string    `gorm:"column:fundraiser_button" json:"fundraiserButton"`
	FundraiserCommissionOn      *int       `gorm:"column:fundraiser_commission_on" json:"fundraiserCommissionOn"`
	FundraiserCommissionType    *int       `gorm:"column:fundraiser_commission_type" json:"fundraiserCommissionType"`
	FundraiserCommissionPercent *float64   `gorm:"column:fundraiser_commission_percent" json:"fundraiserCommissionPercent"`
	FundraiserCommissionFixed   *int       `gorm:"column:fundraiser_commission_fixed" json:"fundraiserCommissionFixed"`
	AdditionalInfo              *string    `gorm:"column:additional_info" json:"additionalInfo"`
	AdditionalFormula           *string    `gorm:"column:additional_formula" json:"additionalFormula"`
	AdditionalField             *string    `gorm:"column:additional_field" json:"additionalField"`
	CustomFieldSetting          *string    `gorm:"column:custom_field_setting" json:"customFieldSetting"`
	GeneralStatus               *int       `gorm:"column:general_status" json:"generalStatus"`
	AllocationTitle             *string    `gorm:"column:allocation_title" json:"allocationTitle"`
	AllocationOthersTitle       *string    `gorm:"column:allocation_others_title" json:"allocationOthersTitle"`
	DonaturName                 *int       `gorm:"column:donatur_name" json:"donaturName"`
	DonaturOthersName           *string    `gorm:"column:donatur_others_name" json:"donaturOthersName"`
	HomeIconURL                 *string    `gorm:"column:home_icon_url" json:"homeIconUrl"`
	BackIconURL                 *string    `gorm:"column:back_icon_url" json:"backIconUrl"`
	MinimumDonate               *int       `gorm:"column:minimum_donate" json:"minimumDonate"`
	OptNominal                  *string    `gorm:"column:opt_nominal" json:"optNominal"`
	CsRotator                   *string    `gorm:"column:cs_rotator" json:"csRotator"`
	WanotifDevice               *int       `gorm:"column:wanotif_device" json:"wanotifDevice"`
	OptQurban                   *string    `gorm:"column:opt_qurban" json:"optQurban"`
	OptPackage2                 *string    `gorm:"column:opt_package2" json:"optPackage2"`
	OptZfitrah                  *string    `gorm:"column:opt_zfitrah" json:"optZfitrah"`
	PackagedUnit                *int       `gorm:"column:packaged_unit" json:"packagedUnit"`
	PackagedCustom              *string    `gorm:"column:packaged_custom" json:"packagedCustom"`
	MetapixelOnly               *int       `gorm:"column:metapixel_only" json:"metapixelOnly"`
	MetapixelConvertion         *int       `gorm:"column:metapixel_convertion" json:"metapixelConvertion"`
	MetapixelConvertionData     *string    `gorm:"column:metapixel_convertion_data" json:"metapixelConvertionData"`
	IconSetting                 *string    `gorm:"column:icon_setting" json:"iconSetting"`
	WanotifMessage2             *string    `gorm:"column:wanotif_message2" json:"wanotifMessage2"`
	ZakatPenghasilanType        *string    `gorm:"column:zakat_penghasilan_type" json:"zakatPenghasilanType"`
	ZakatPenghasilanCustomTitle *int       `gorm:"column:zakat_penghasilan_custom_title" json:"zakatPenghasilanCustomTitle"`
	Pendapatan1Title            *string    `gorm:"column:pendapatan1_title" json:"pendapatan1Title"`
	Pendapatan2Title            *string    `gorm:"column:pendapatan2_title" json:"pendapatan2Title"`
	ZakatHasilPertanian         *string    `gorm:"column:zakat_hasil_pertanian" json:"zakatHasilPertanian"`
	ZakatHargaPerKg             *int       `gorm:"column:zakat_harga_per_kg" json:"zakatHargaPerKg"`
	ZakatPengairan              *string    `gorm:"column:zakat_pengairan" json:"zakatPengairan"`
	ZakatNisabKg                *int       `gorm:"column:zakat_nisab_kg" json:"zakatNisabKg"`
	ZakatHargaEmas              *int       `gorm:"column:zakat_harga_emas" json:"zakatHargaEmas"`
	FollowupStatus              *int       `gorm:"column:followup_status" json:"followupStatus"`
	FollowupText                *string    `gorm:"column:followup_text" json:"followupText"`
}

func (WpDjaCampaign) TableName() string {
	return "wp_dja_campaign"
}

type ListPriorityWpDjaCampaign struct {
	ID         int     `gorm:"column:id;primaryKey" json:"id"`
	UserID     *int64  `gorm:"column:user_id" json:"-"`
	CampaignID *string `gorm:"column:campaign_id" json:"campaignId"`
	Title      *string `gorm:"column:title" json:"title"`
	Slug       *string `gorm:"column:slug" json:"slug"`
	ImageURL   *string `gorm:"column:image_url" json:"imageUrl"`
	Owner      *string `gorm:"column:owner" json:"owner"`
}

type DetailWpDjaCampaign struct {
	ID                int      `gorm:"column:id;primaryKey" json:"id"`
	UserID            *int64   `gorm:"column:user_id" json:"-"`
	CampaignID        *string  `gorm:"column:campaign_id" json:"campaignId"`
	Slug              *string  `gorm:"column:slug" json:"slug"`
	ImageURL          *string  `gorm:"column:image_url" json:"imageUrl"`
	Title             *string  `gorm:"column:title" json:"title"`
	Target            *float64 `gorm:"column:target" json:"target"`
	Owner             *string  `gorm:"column:owner" json:"owner"`
	TotalDonasi       *float64 `gorm:"total_donasi" json:"totalDonasi"`
	TotalDonatur      *int     `gorm:"total_donatur" json:"totalDonatur"`
	OwnerImage        *string  `gorm:"column:owner_image" json:"ownerImage"`
	OwnerType         *string  `gorm:"column:owner_type" json:"ownerType"`
	OwnerVerification int      `gorm:"column:owner_verification" json:"ownerVerification"`
	Information       *string  `gorm:"column:information" json:"information"`
}
