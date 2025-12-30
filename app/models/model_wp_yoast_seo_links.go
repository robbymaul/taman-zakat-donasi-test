package models

type YoastSeoLink struct {
	ID                uint64  `gorm:"column:id;primaryKey" json:"id"`
	URL               *string `gorm:"column:url" json:"url"`
	PostID            *uint64 `gorm:"column:post_id" json:"postId"`
	TargetPostID      *uint64 `gorm:"column:target_post_id" json:"targetPostId"`
	Type              *string `gorm:"column:type" json:"type"`
	IndexableID       *uint   `gorm:"column:indexable_id" json:"indexableId"`
	TargetIndexableID *uint   `gorm:"column:target_indexable_id" json:"targetIndexableId"`
	Height            *uint   `gorm:"column:height" json:"height"`
	Width             *uint   `gorm:"column:width" json:"width"`
	Size              *uint   `gorm:"column:size" json:"size"`
	Language          *string `gorm:"column:language" json:"language"`
	Region            *string `gorm:"column:region" json:"region"`
}

func (YoastSeoLink) TableName() string {
	return "wp_yoast_seo_links"
}
