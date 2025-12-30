package models

import "time"

type WpUapRecipeLog struct {
	ID                uint64    `gorm:"column:ID;primaryKey" json:"id"`
	DateTime          time.Time `gorm:"column:date_time" json:"dateTime"`
	UserID            uint64    `gorm:"column:user_id" json:"userId"`
	AutomatorRecipeID uint64    `gorm:"column:automator_recipe_id" json:"automatorRecipeId"`
	Completed         bool      `gorm:"column:completed" json:"completed"`
	RunNumber         uint32    `gorm:"column:run_number" json:"runNumber"`
}

func (WpUapRecipeLog) TableName() string {
	return "wp_uap_recipe_log"
}
