package models

type WpESubmissionValue struct {
	ID           uint64  `gorm:"column:id;primaryKey" json:"id"`
	SubmissionID uint64  `gorm:"column:submission_id" json:"submissionId"`
	Key          *string `gorm:"column:key" json:"key"`
	Value        *string `gorm:"column:value" json:"value"`
}

func (WpESubmissionValue) TableName() string {
	return "wp_e_submissions_values"
}
