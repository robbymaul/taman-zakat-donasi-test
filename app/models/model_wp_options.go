package models

type WpOption struct {
	OptionID    uint64 `gorm:"column:option_id;primaryKey" json:"optionID"`
	OptionName  string `gorm:"column:option_name" json:"optionName"`
	OptionValue string `gorm:"column:option_value" json:"optionValue"`
	Autoload    string `gorm:"column:autoload" json:"autoload"`
}

func (WpOption) TableName() string {
	return "wp_options"
}
