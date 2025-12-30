package models

import "time"

type WpUapTriggerLogsView struct {
	AutomatorRecipeID  uint64    `gorm:"column:automator_recipe_id" json:"automatorRecipeId"`
	AutomatorTriggerID uint64    `gorm:"column:automator_trigger_id" json:"automatorTriggerId"`
	DisplayName        string    `gorm:"column:display_name" json:"displayName"`
	ID                 uint64    `gorm:"column:ID" json:"id"`
	RecipeCompleted    bool      `gorm:"column:recipe_completed" json:"recipeCompleted"`
	RecipeDateTime     time.Time `gorm:"column:recipe_date_time" json:"recipeDateTime"`
	RecipeLogID        uint64    `gorm:"column:recipe_log_id" json:"recipeLogId"`
	RecipeRunNumber    uint32    `gorm:"column:recipe_run_number" json:"recipeRunNumber"`
	RecipeTitle        string    `gorm:"column:recipe_title" json:"recipeTitle"`
	TriggerCompleted   uint8     `gorm:"column:trigger_completed" json:"triggerCompleted"`
	TriggerDate        time.Time `gorm:"column:trigger_date" json:"triggerDate"`
	TriggerRunNumber   uint32    `gorm:"column:trigger_run_number" json:"triggerRunNumber"`
	TriggerRunTime     time.Time `gorm:"column:trigger_run_time" json:"triggerRunTime"`
	TriggerSentence    string    `gorm:"column:trigger_sentence" json:"triggerSentence"`
	TriggerTitle       string    `gorm:"column:trigger_title" json:"triggerTitle"`
	TriggerTotalTimes  string    `gorm:"column:trigger_total_times" json:"triggerTotalTimes"`
	UserEmail          string    `gorm:"column:user_email" json:"userEmail"`
	UserID             uint64    `gorm:"column:user_id" json:"userId"`
}

func (WpUapTriggerLogsView) TableName() string {
	return "wp_uap_trigger_logs_view"
}
