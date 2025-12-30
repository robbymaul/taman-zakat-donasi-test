package models

type YoastMigration struct {
	ID      uint   `gorm:"column:id;primaryKey" json:"id"`
	Version string `gorm:"column:version" json:"version"`
}

func (YoastMigration) TableName() string {
	return "wp_yoast_migrations"
}
