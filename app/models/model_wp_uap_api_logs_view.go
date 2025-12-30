package models

import "time"

type UapApiLogsView struct {
	AutomatorRecipeID uint64     `gorm:"column:automator_recipe_id" json:"automatorRecipeId"`
	Balance           uint64     `gorm:"column:balance" json:"balance"`
	Completed         uint8      `gorm:"column:completed" json:"completed"`
	Date              *time.Time `gorm:"column:date" json:"date"`
	DisplayName       string     `gorm:"column:display_name" json:"displayName"`
	Endpoint          string     `gorm:"column:endpoint" json:"endpoint"`
	ErrorMessage      string     `gorm:"column:error_message" json:"errorMessage"`
	ID                uint64     `gorm:"column:ID" json:"id"`
	ItemID            uint64     `gorm:"column:item_id" json:"itemId"`
	ItemLogID         uint64     `gorm:"column:item_log_id" json:"itemLogId"`
	Notes             string     `gorm:"column:notes" json:"notes"`
	Params            string     `gorm:"column:params" json:"params"`
	Price             uint64     `gorm:"column:price" json:"price"`
	RecipeCompleted   int8       `gorm:"column:recipe_completed" json:"recipeCompleted"`
	RecipeDateTime    *time.Time `gorm:"column:recipe_date_time" json:"recipeDateTime"`
	RecipeLogID       uint64     `gorm:"column:recipe_log_id" json:"recipeLogId"`
	RecipeRunNumber   uint32     `gorm:"column:recipe_run_number" json:"recipeRunNumber"`
	RecipeTitle       string     `gorm:"column:recipe_title" json:"recipeTitle"`
	Sentence          string     `gorm:"column:sentence" json:"sentence"`
	Status            string     `gorm:"column:status" json:"status"`
	TimeSpent         uint64     `gorm:"column:time_spent" json:"timeSpent"`
	Title             string     `gorm:"column:title" json:"title"`
	Type              string     `gorm:"column:type" json:"type"`
	UserEmail         string     `gorm:"column:user_email" json:"userEmail"`
	UserID            uint64     `gorm:"column:user_id" json:"userId"`
}

func (UapApiLogsView) TableName() string {
	return "wp_uap_api_logs_view"
}
