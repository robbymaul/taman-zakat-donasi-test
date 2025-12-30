package models

type WpGlsrAssignedUser struct {
	RatingID uint64 `gorm:"column:rating_id;primaryKey" json:"ratingId"`
	UserID   uint64 `gorm:"column:user_id;primaryKey" json:"userId"`
}

func (WpGlsrAssignedUser) TableName() string {
	return "wp_glsr_assigned_users"
}
