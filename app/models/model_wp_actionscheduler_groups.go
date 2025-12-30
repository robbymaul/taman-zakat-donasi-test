package models

// WpActionschedulerGroup represents table `wp_actionscheduler_groups`.
type WpActionschedulerGroup struct {
	GroupID uint64 `gorm:"column:group_id;primaryKey" json:"groupId"`
	Slug    string `gorm:"column:slug" json:"slug"`
}

func (WpActionschedulerGroup) TableName() string {
	return "wp_actionscheduler_groups"
}
