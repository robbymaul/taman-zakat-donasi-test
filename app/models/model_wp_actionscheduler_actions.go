package models

// WpActionschedulerAction represents table `wp_actionscheduler_actions`.
//
// Notes:
// - Datetime columns are mapped as string to avoid issues with MySQL zero dates ('0000-00-00 00:00:00').
type WpActionschedulerAction struct {
	ActionID           uint64  `gorm:"column:action_id;primaryKey" json:"actionId"`
	Hook               string  `gorm:"column:hook" json:"hook"`
	Status             string  `gorm:"column:status" json:"status"`
	ScheduledDateGmt   string  `gorm:"column:scheduled_date_gmt" json:"scheduledDateGmt"`
	ScheduledDateLocal string  `gorm:"column:scheduled_date_local" json:"scheduledDateLocal"`
	Args               *string `gorm:"column:args" json:"args"`
	Schedule           *string `gorm:"column:schedule" json:"schedule"`
	GroupID            uint64  `gorm:"column:group_id" json:"groupId"`
	Attempts           int     `gorm:"column:attempts" json:"attempts"`
	LastAttemptGmt     string  `gorm:"column:last_attempt_gmt" json:"lastAttemptGmt"`
	LastAttemptLocal   string  `gorm:"column:last_attempt_local" json:"lastAttemptLocal"`
	ClaimID            uint64  `gorm:"column:claim_id" json:"claimId"`
	ExtendedArgs       *string `gorm:"column:extended_args" json:"extendedArgs"`
	Priority           uint8   `gorm:"column:priority" json:"priority"`
}

func (WpActionschedulerAction) TableName() string {
	return "wp_actionscheduler_actions"
}
