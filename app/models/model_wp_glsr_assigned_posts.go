package models

type WpGlsrAssignedPost struct {
	RatingID    uint64 `gorm:"column:rating_id;primaryKey" json:"ratingId"`
	PostID      uint64 `gorm:"column:post_id;primaryKey" json:"postId"`
	IsPublished bool   `gorm:"column:is_published" json:"isPublished"`
}

func (WpGlsrAssignedPost) TableName() string {
	return "wp_glsr_assigned_posts"
}
