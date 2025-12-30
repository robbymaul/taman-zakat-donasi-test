package models

import "time"

type WpPost struct {
	ID                  uint64    `gorm:"column:ID;primaryKey" json:"id"`
	PostAuthor          uint64    `gorm:"column:post_author" json:"postAuthor"`
	PostDate            time.Time `gorm:"column:post_date" json:"postDate"`
	PostDateGMT         time.Time `gorm:"column:post_date_gmt" json:"postDateGMT"`
	PostContent         string    `gorm:"column:post_content" json:"postContent"`
	PostTitle           string    `gorm:"column:post_title" json:"postTitle"`
	PostExcerpt         string    `gorm:"column:post_excerpt" json:"postExcerpt"`
	PostStatus          string    `gorm:"column:post_status" json:"postStatus"`
	CommentStatus       string    `gorm:"column:comment_status" json:"commentStatus"`
	PingStatus          string    `gorm:"column:ping_status" json:"pingStatus"`
	PostPassword        string    `gorm:"column:post_password" json:"postPassword"`
	PostName            string    `gorm:"column:post_name" json:"postName"`
	ToPing              string    `gorm:"column:to_ping" json:"toPing"`
	Pinged              string    `gorm:"column:pinged" json:"pinged"`
	PostModified        time.Time `gorm:"column:post_modified" json:"postModified"`
	PostModifiedGMT     time.Time `gorm:"column:post_modified_gmt" json:"postModifiedGMT"`
	PostContentFiltered string    `gorm:"column:post_content_filtered" json:"postContentFiltered"`
	PostParent          uint64    `gorm:"column:post_parent" json:"postParent"`
	Guid                string    `gorm:"column:guid" json:"guid"`
	MenuOrder           int       `gorm:"column:menu_order" json:"menuOrder"`
	PostType            string    `gorm:"column:post_type" json:"postType"`
	PostMimeType        string    `gorm:"column:post_mime_type" json:"postMimeType"`
	CommentCount        int64     `gorm:"column:comment_count" json:"commentCount"`
}

func (WpPost) TableName() string {
	return "wp_posts"
}
