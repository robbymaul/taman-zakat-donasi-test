package models

type WpGlsrRating struct {
	ID         uint64 `gorm:"column:ID;primaryKey" json:"id"`
	ReviewID   uint64 `gorm:"column:review_id" json:"reviewId"`
	Rating     int    `gorm:"column:rating" json:"rating"`
	Type       string `gorm:"column:type" json:"type"`
	IsApproved bool   `gorm:"column:is_approved" json:"isApproved"`
	IsPinned   bool   `gorm:"column:is_pinned" json:"isPinned"`
	Name       string `gorm:"column:name" json:"name"`
	Email      string `gorm:"column:email" json:"email"`
	Avatar     string `gorm:"column:avatar" json:"avatar"`
	IPAddress  string `gorm:"column:ip_address" json:"ipAddress"`
	URL        string `gorm:"column:url" json:"url"`
	Terms      bool   `gorm:"column:terms" json:"terms"`
}

func (WpGlsrRating) TableName() string {
	return "wp_glsr_ratings"
}
