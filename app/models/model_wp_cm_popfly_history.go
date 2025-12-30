package models

import "time"

// WpCmPopflyHistory represents table `wp_cm_popfly_history`.
type WpCmPopflyHistory struct {
	EventID       int64     `gorm:"column:event_id;primaryKey" json:"eventId"`
	EventType     string    `gorm:"column:event_type" json:"eventType"`
	CampaignID    *int64    `gorm:"column:campaign_id" json:"campaignId"`
	Amount        *int      `gorm:"column:amount" json:"amount"`
	BannerID      *string   `gorm:"column:banner_id" json:"bannerId"`
	RefererURL    string    `gorm:"column:referer_url" json:"refererUrl"`
	RemoteIP      string    `gorm:"column:remote_ip" json:"remoteIp"`
	WebpageURL    string    `gorm:"column:webpage_url" json:"webpageUrl"`
	RemoteCountry string    `gorm:"column:remote_country" json:"remoteCountry"`
	RemoteCity    string    `gorm:"column:remote_city" json:"remoteCity"`
	Regdate       time.Time `gorm:"column:regdate" json:"regdate"`
}

func (WpCmPopflyHistory) TableName() string {
	return "wp_cm_popfly_history"
}
