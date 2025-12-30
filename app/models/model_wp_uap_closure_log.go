package models

import "time"

type WpUapClosureLog struct {
	ID                   uint64    `gorm:"column:ID;primaryKey" json:"id"`
	DateTime             time.Time `gorm:"column:date_time" json:"dateTime"`
	UserID               uint64    `gorm:"column:user_id" json:"userId"`
	AutomatorClosureID   uint64    `gorm:"column:automator_closure_id" json:"automatorClosureId"`
	AutomatorRecipeID    uint64    `gorm:"column:automator_recipe_id" json:"automatorRecipeId"`
	AutomatorRecipeLogID uint64    `gorm:"column:automator_recipe_log_id" json:"automatorRecipeLogId"`
	Completed            uint8     `gorm:"column:completed" json:"completed"`
}

func (WpUapClosureLog) TableName() string {
	return "wp_uap_closure_log"
}
