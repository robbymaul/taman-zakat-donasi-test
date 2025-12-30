package models

// WpActionschedulerClaim represents table `wp_actionscheduler_claims`.
//
// Notes:
// - Datetime columns are mapped as string to avoid issues with MySQL zero dates ('0000-00-00 00:00:00').
type WpActionschedulerClaim struct {
	ClaimID        uint64 `gorm:"column:claim_id;primaryKey" json:"claimId"`
	DateCreatedGmt string `gorm:"column:date_created_gmt" json:"dateCreatedGmt"`
}

func (WpActionschedulerClaim) TableName() string {
	return "wp_actionscheduler_claims"
}
