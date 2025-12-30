package models

import "time"

type YoastPrimaryTerm struct {
	ID        uint       `gorm:"column:id;primaryKey" json:"id"`
	PostID    *uint64    `gorm:"column:post_id" json:"postId"`
	TermID    *uint64    `gorm:"column:term_id" json:"termId"`
	Taxonomy  string     `gorm:"column:taxonomy" json:"taxonomy"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updatedAt"`
	BlogID    uint64     `gorm:"column:blog_id" json:"blogId"`
}

func (YoastPrimaryTerm) TableName() string {
	return "wp_yoast_primary_term"
}
