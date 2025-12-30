package models

type WpSejolisaAcquisitionOrder struct {
	OrderID int    `gorm:"column:order_id;primaryKey" json:"orderID"`
	Source  string `gorm:"column:source" json:"source"`
	Media   string `gorm:"column:media" json:"media"`
}

func (WpSejolisaAcquisitionOrder) TableName() string {
	return "wp_sejolisa_acquisition_order"
}
