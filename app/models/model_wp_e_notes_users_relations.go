package models

import "time"

type WpENotesUsersRelation struct {
	ID        uint64    `gorm:"column:id;primaryKey" json:"id"`
	Type      string    `gorm:"column:type" json:"type"`
	NoteID    uint64    `gorm:"column:note_id" json:"noteId"`
	UserID    uint64    `gorm:"column:user_id" json:"userId"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (WpENotesUsersRelation) TableName() string {
	return "wp_e_notes_users_relations"
}
