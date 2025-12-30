package models

// WpCommentmeta represents table `wp_commentmeta`.
type WpCommentmeta struct {
	MetaID    uint64  `gorm:"column:meta_id;primaryKey" json:"metaId"`
	CommentID uint64  `gorm:"column:comment_id" json:"commentId"`
	MetaKey   *string `gorm:"column:meta_key" json:"metaKey"`
	MetaValue *string `gorm:"column:meta_value" json:"metaValue"`
}

func (WpCommentmeta) TableName() string {
	return "wp_commentmeta"
}
