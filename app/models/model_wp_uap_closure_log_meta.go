package models

type WpUapClosureLogMeta struct {
	ID                    uint64 `gorm:"column:ID;primaryKey" json:"id"`
	UserID                uint64 `gorm:"column:user_id" json:"userId"`
	AutomatorClosureID    uint64 `gorm:"column:automator_closure_id" json:"automatorClosureId"`
	AutomatorClosureLogID uint64 `gorm:"column:automator_closure_log_id" json:"automatorClosureLogId"`
	MetaKey               string `gorm:"column:meta_key" json:"metaKey"`
	MetaValue             string `gorm:"column:meta_value" json:"metaValue"`
}

// TableName overrides the table name
func (WpUapClosureLogMeta) TableName() string {
	return "wp_uap_closure_log_meta"
}
