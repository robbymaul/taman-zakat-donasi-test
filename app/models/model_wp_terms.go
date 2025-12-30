package models

// WpTerm represents table `wp_terms`.
type WpTerm struct {
	TermID    uint64 `gorm:"column:term_id;primaryKey" json:"termId"`
	Name      string `gorm:"column:name" json:"name"`
	Slug      string `gorm:"column:slug" json:"slug"`
	TermGroup int64  `gorm:"column:term_group" json:"termGroup"`
}

func (WpTerm) TableName() string {
	return "wp_terms"
}
