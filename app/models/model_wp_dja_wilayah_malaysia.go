package models

type WpDjaWilayahMalaysia struct {
	ID     int     `gorm:"column:id;primaryKey" json:"id"`
	Negeri *string `gorm:"column:negeri" json:"negeri"`
	Bandar *string `gorm:"column:bandar" json:"bandar"`
	Daerah *string `gorm:"column:daerah" json:"daerah"`
}

func (WpDjaWilayahMalaysia) TableName() string {
	return "wp_dja_wilayah_malaysia"
}
