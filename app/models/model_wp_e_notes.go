package models

import "time"

type WpENote struct {
	ID                uint64     `gorm:"column:id;primaryKey" json:"id"`
	RouteURL          string     `gorm:"column:route_url" json:"routeUrl"`
	RouteTitle        string     `gorm:"column:route_title" json:"routeTitle"`
	RoutePostID       *uint64    `gorm:"column:route_post_id" json:"routePostId,omitempty"`
	PostID            *uint64    `gorm:"column:post_id" json:"postId,omitempty"`
	ElementID         string     `gorm:"column:element_id" json:"elementId"`
	ParentID          uint64     `gorm:"column:parent_id" json:"parentId"`
	AuthorID          *uint64    `gorm:"column:author_id" json:"authorId,omitempty"`
	AuthorDisplayName string     `gorm:"column:author_display_name" json:"authorDisplayName"`
	Status            string     `gorm:"column:status" json:"status"`
	Position          string     `gorm:"column:position" json:"position"` // JSON string
	Content           string     `gorm:"column:content" json:"content"`
	IsResolved        bool       `gorm:"column:is_resolved" json:"isResolved"`
	IsPublic          bool       `gorm:"column:is_public" json:"isPublic"`
	LastActivityAt    *time.Time `gorm:"column:last_activity_at" json:"lastActivityAt,omitempty"`
	CreatedAt         time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt         time.Time  `gorm:"column:updated_at" json:"updatedAt"`
}

func (WpENote) TableName() string {
	return "wp_e_notes"
}
