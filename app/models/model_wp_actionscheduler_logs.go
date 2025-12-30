package models

// WpActionschedulerLog represents table `wp_actionscheduler_logs`.
//
// Notes:
// - Datetime columns are mapped as string to avoid issues with MySQL zero dates ('0000-00-00 00:00:00').
type WpActionschedulerLog struct {
	LogID        uint64 `gorm:"column:log_id;primaryKey" json:"logId"`
	ActionID     uint64 `gorm:"column:action_id" json:"actionId"`
	Message      string `gorm:"column:message" json:"message"`
	LogDateGmt   string `gorm:"column:log_date_gmt" json:"logDateGmt"`
	LogDateLocal string `gorm:"column:log_date_local" json:"logDateLocal"`
}

func (WpActionschedulerLog) TableName() string {
	return "wp_actionscheduler_logs"
}
