package models

type WpGlsrAssignedTerm struct {
	RatingID uint64 `gorm:"column:rating_id;primaryKey" json:"ratingId"`
	TermID   uint64 `gorm:"column:term_id;primaryKey" json:"termId"`
}

func (WpGlsrAssignedTerm) TableName() string {
	return "wp_glsr_assigned_terms"
}
