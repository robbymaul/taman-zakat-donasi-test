package models

type WpUapActionLogMeta struct {
	ID                   uint64  `gorm:"column:ID;primaryKey" json:"id"`
	UserID               uint64  `gorm:"column:user_id" json:"userId"`
	AutomatorActionLogID *uint64 `gorm:"column:automator_action_log_id" json:"automatorActionLogId,omitempty"`
	AutomatorActionID    uint64  `gorm:"column:automator_action_id" json:"automatorActionId"`
	MetaKey              string  `gorm:"column:meta_key" json:"metaKey"`
	MetaValue            *string `gorm:"column:meta_value" json:"metaValue,omitempty"`
}

func (WpUapActionLogMeta) TableName() string {
	return "wp_uap_action_log_meta"
}
