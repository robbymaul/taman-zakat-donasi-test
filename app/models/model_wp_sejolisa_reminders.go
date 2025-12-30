package models

import "time"

type SejolisaReminder struct {
	ID           uint      `gorm:"column:ID;primaryKey" json:"id"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"createdAt"`
	SentAt       time.Time `gorm:"column:sent_at" json:"sentAt"`
	SendDay      int       `gorm:"column:send_day" json:"sendDay"`
	OrderID      int       `gorm:"column:order_id" json:"orderId"`
	Title        string    `gorm:"column:title" json:"title"`
	Content      string    `gorm:"column:content" json:"content"`
	MediaType    string    `gorm:"column:media_type" json:"mediaType"`
	ReminderType string    `gorm:"column:reminder_type" json:"reminderType"`
	Status       string    `gorm:"column:status" json:"status"`
}

func (SejolisaReminder) TableName() string {
	return "wp_sejolisa_reminders"
}
