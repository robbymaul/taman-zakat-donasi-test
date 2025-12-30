package models

import "time"

type WpUapTriggerLogMeta struct {
	ID                    uint64    `gorm:"column:ID;primaryKey" json:"id"`
	UserID                uint64    `gorm:"column:user_id" json:"userId"`
	AutomatorTriggerLogID *uint64   `gorm:"column:automator_trigger_log_id" json:"automatorTriggerLogId"`
	AutomatorTriggerID    uint64    `gorm:"column:automator_trigger_id" json:"automatorTriggerId"`
	MetaKey               string    `gorm:"column:meta_key" json:"metaKey"`
	MetaValue             string    `gorm:"column:meta_value" json:"metaValue"`
	RunNumber             uint32    `gorm:"column:run_number" json:"runNumber"`
	RunTime               time.Time `gorm:"column:run_time" json:"runTime"`
}

func (WpUapTriggerLogMeta) TableName() string {
	return "wp_uap_trigger_log_meta"
}
