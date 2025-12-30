package models

import "time"

type WpUapRecipeLogsView struct {
	AutomatorRecipeID uint64    `gorm:"column:automator_recipe_id" json:"automatorRecipeId"`
	Completed         bool      `gorm:"column:completed" json:"completed"`
	DisplayName       string    `gorm:"column:display_name" json:"displayName"`
	RecipeCompleted   bool      `gorm:"column:recipe_completed" json:"recipeCompleted"`
	RecipeDateTime    time.Time `gorm:"column:recipe_date_time" json:"recipeDateTime"`
	RecipeLogID       uint64    `gorm:"column:recipe_log_id" json:"recipeLogId"`
	RecipeTitle       string    `gorm:"column:recipe_title" json:"recipeTitle"`
	RunNumber         uint32    `gorm:"column:run_number" json:"runNumber"`
	UserEmail         string    `gorm:"column:user_email" json:"userEmail"`
	UserID            uint64    `gorm:"column:user_id" json:"userId"`
}

func (WpUapRecipeLogsView) TableName() string {
	return "wp_uap_recipe_logs_view"
}
