package models

import "time"

type WpDjaShortcode struct {
	ID          int        `gorm:"column:id;primaryKey" json:"id"`
	SId         *string    `gorm:"column:s_id" json:"sId"`
	SName       *string    `gorm:"column:s_name" json:"sName"`
	SCategory   *string    `gorm:"column:s_category" json:"sCategory"`
	SStyle      *string    `gorm:"column:s_style" json:"sStyle"`
	SShow       *int       `gorm:"column:s_show" json:"sShow"`
	SLoadmore   *int       `gorm:"column:s_loadmore" json:"sLoadmore"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updatedAt"`
	SButtonOn   *int       `gorm:"column:s_button_on" json:"sButtonOn"`
	SButtonText *string    `gorm:"column:s_button_text" json:"sButtonText"`
	SDataLoad   *int       `gorm:"column:s_data_load" json:"sDataLoad"`
	SCampaign   *string    `gorm:"column:s_campaign" json:"sCampaign"`
	SGridOn     *int       `gorm:"column:s_grid_on" json:"sGridOn"`
	SGridMobile *int       `gorm:"column:s_grid_mobile" json:"sGridMobile"`
}

func (WpDjaShortcode) TableName() string {
	return "wp_dja_shortcode"
}
