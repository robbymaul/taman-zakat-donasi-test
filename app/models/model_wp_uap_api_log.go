package models

import "time"

type WpUapApiLog struct {
	ID          uint64    `gorm:"column:ID;primaryKey" json:"id"`
	DateTime    time.Time `gorm:"column:date_time" json:"dateTime"`
	Type        string    `gorm:"column:type" json:"type"`
	RecipeLogID uint64    `gorm:"column:recipe_log_id" json:"recipeLogId"`
	ItemLogID   uint64    `gorm:"column:item_log_id" json:"itemLogId"`
	Endpoint    *string   `gorm:"column:endpoint" json:"endpoint,omitempty"`
	Params      *string   `gorm:"column:params" json:"params,omitempty"`
	Request     *string   `gorm:"column:request" json:"request,omitempty"`
	Response    *string   `gorm:"column:response" json:"response,omitempty"`
	Status      *string   `gorm:"column:status" json:"status,omitempty"`
	Price       *uint64   `gorm:"column:price" json:"price,omitempty"`
	Balance     *uint64   `gorm:"column:balance" json:"balance,omitempty"`
	TimeSpent   *uint64   `gorm:"column:time_spent" json:"timeSpent,omitempty"`
	Notes       *string   `gorm:"column:notes" json:"notes,omitempty"`
}

func (WpUapApiLog) TableName() string {
	return "wp_uap_api_log"
}
