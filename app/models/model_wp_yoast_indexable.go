package models

import "time"

type YoastIndexable struct {
	ID                          uint       `gorm:"column:id;primaryKey" json:"id"`
	Permalink                   *string    `gorm:"column:permalink" json:"permalink"`
	PermalinkHash               *string    `gorm:"column:permalink_hash" json:"permalinkHash"`
	ObjectID                    *int64     `gorm:"column:object_id" json:"objectId"`
	ObjectType                  string     `gorm:"column:object_type" json:"objectType"`
	ObjectSubType               *string    `gorm:"column:object_sub_type" json:"objectSubType"`
	AuthorID                    *int64     `gorm:"column:author_id" json:"authorId"`
	PostParent                  *int64     `gorm:"column:post_parent" json:"postParent"`
	Title                       *string    `gorm:"column:title" json:"title"`
	Description                 *string    `gorm:"column:description" json:"description"`
	BreadcrumbTitle             *string    `gorm:"column:breadcrumb_title" json:"breadcrumbTitle"`
	PostStatus                  *string    `gorm:"column:post_status" json:"postStatus"`
	IsPublic                    *bool      `gorm:"column:is_public" json:"isPublic"`
	IsProtected                 *bool      `gorm:"column:is_protected" json:"isProtected"`
	HasPublicPosts              *bool      `gorm:"column:has_public_posts" json:"hasPublicPosts"`
	NumberOfPages               *uint      `gorm:"column:number_of_pages" json:"numberOfPages"`
	Canonical                   *string    `gorm:"column:canonical" json:"canonical"`
	PrimaryFocusKeyword         *string    `gorm:"column:primary_focus_keyword" json:"primaryFocusKeyword"`
	PrimaryFocusKeywordScore    *int       `gorm:"column:primary_focus_keyword_score" json:"primaryFocusKeywordScore"`
	ReadabilityScore            *int       `gorm:"column:readability_score" json:"readabilityScore"`
	IsCornerstone               *bool      `gorm:"column:is_cornerstone" json:"isCornerstone"`
	IsRobotsNoindex             *bool      `gorm:"column:is_robots_noindex" json:"isRobotsNoindex"`
	IsRobotsNofollow            *bool      `gorm:"column:is_robots_nofollow" json:"isRobotsNofollow"`
	IsRobotsNoarchive           *bool      `gorm:"column:is_robots_noarchive" json:"isRobotsNoarchive"`
	IsRobotsNoimageindex        *bool      `gorm:"column:is_robots_noimageindex" json:"isRobotsNoimageindex"`
	IsRobotsNosnippet           *bool      `gorm:"column:is_robots_nosnippet" json:"isRobotsNosnippet"`
	TwitterTitle                *string    `gorm:"column:twitter_title" json:"twitterTitle"`
	TwitterImage                *string    `gorm:"column:twitter_image" json:"twitterImage"`
	TwitterDescription          *string    `gorm:"column:twitter_description" json:"twitterDescription"`
	TwitterImageID              *string    `gorm:"column:twitter_image_id" json:"twitterImageId"`
	TwitterImageSource          *string    `gorm:"column:twitter_image_source" json:"twitterImageSource"`
	OpenGraphTitle              *string    `gorm:"column:open_graph_title" json:"openGraphTitle"`
	OpenGraphDescription        *string    `gorm:"column:open_graph_description" json:"openGraphDescription"`
	OpenGraphImage              *string    `gorm:"column:open_graph_image" json:"openGraphImage"`
	OpenGraphImageID            *string    `gorm:"column:open_graph_image_id" json:"openGraphImageId"`
	OpenGraphImageSource        *string    `gorm:"column:open_graph_image_source" json:"openGraphImageSource"`
	OpenGraphImageMeta          *string    `gorm:"column:open_graph_image_meta" json:"openGraphImageMeta"`
	LinkCount                   *int       `gorm:"column:link_count" json:"linkCount"`
	IncomingLinkCount           *int       `gorm:"column:incoming_link_count" json:"incomingLinkCount"`
	ProminentWordsVersion       *uint      `gorm:"column:prominent_words_version" json:"prominentWordsVersion"`
	CreatedAt                   *time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt                   time.Time  `gorm:"column:updated_at" json:"updatedAt"`
	BlogID                      uint64     `gorm:"column:blog_id" json:"blogId"`
	Language                    *string    `gorm:"column:language" json:"language"`
	Region                      *string    `gorm:"column:region" json:"region"`
	SchemaPageType              *string    `gorm:"column:schema_page_type" json:"schemaPageType"`
	SchemaArticleType           *string    `gorm:"column:schema_article_type" json:"schemaArticleType"`
	HasAncestors                *bool      `gorm:"column:has_ancestors" json:"hasAncestors"`
	EstimatedReadingTimeMinutes *int       `gorm:"column:estimated_reading_time_minutes" json:"estimatedReadingTimeMinutes"`
	Version                     *int       `gorm:"column:version" json:"version"`
	ObjectLastModified          *time.Time `gorm:"column:object_last_modified" json:"objectLastModified"`
	ObjectPublishedAt           *time.Time `gorm:"column:object_published_at" json:"objectPublishedAt"`
	InclusiveLanguageScore      *int       `gorm:"column:inclusive_language_score" json:"inclusiveLanguageScore"`
}

func (YoastIndexable) TableName() string {
	return "wp_yoast_indexable"
}
