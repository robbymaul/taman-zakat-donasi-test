package models

import "time"

type WpUapActionLog struct {
	ID                   uint64    `gorm:"column:ID;primaryKey" json:"id"`
	DateTime             time.Time `gorm:"column:date_time" json:"dateTime"`
	UserID               uint64    `gorm:"column:user_id" json:"userId"`
	AutomatorActionID    uint64    `gorm:"column:automator_action_id" json:"automatorActionId"`
	AutomatorRecipeID    uint64    `gorm:"column:automator_recipe_id" json:"automatorRecipeId"`
	AutomatorRecipeLogID *uint64   `gorm:"column:automator_recipe_log_id" json:"automatorRecipeLogId,omitempty"`
	Completed            uint8     `gorm:"column:completed" json:"completed"`
	ErrorMessage         *string   `gorm:"column:error_message" json:"errorMessage,omitempty"`
}

func (WpUapActionLog) TableName() string {
	return "wp_uap_action_log"
}
