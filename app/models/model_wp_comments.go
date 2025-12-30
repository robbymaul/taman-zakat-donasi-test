package models

// WpComment represents table `wp_comments`.
//
// Notes:
// - Datetime columns are mapped as string to avoid issues with MySQL zero dates ('0000-00-00 00:00:00').
type WpComment struct {
	CommentID          uint64 `gorm:"column:comment_ID;primaryKey" json:"commentId"`
	CommentPostID      uint64 `gorm:"column:comment_post_ID" json:"commentPostId"`
	CommentAuthor      string `gorm:"column:comment_author" json:"commentAuthor"`
	CommentAuthorEmail string `gorm:"column:comment_author_email" json:"commentAuthorEmail"`
	CommentAuthorURL   string `gorm:"column:comment_author_url" json:"commentAuthorUrl"`
	CommentAuthorIP    string `gorm:"column:comment_author_IP" json:"commentAuthorIp"`
	CommentDate        string `gorm:"column:comment_date" json:"commentDate"`
	CommentDateGmt     string `gorm:"column:comment_date_gmt" json:"commentDateGmt"`
	CommentContent     string `gorm:"column:comment_content" json:"commentContent"`
	CommentKarma       int    `gorm:"column:comment_karma" json:"commentKarma"`
	CommentApproved    string `gorm:"column:comment_approved" json:"commentApproved"`
	CommentAgent       string `gorm:"column:comment_agent" json:"commentAgent"`
	CommentType        string `gorm:"column:comment_type" json:"commentType"`
	CommentParent      uint64 `gorm:"column:comment_parent" json:"commentParent"`
	UserID             uint64 `gorm:"column:user_id" json:"userId"`
}

func (WpComment) TableName() string {
	return "wp_comments"
}
