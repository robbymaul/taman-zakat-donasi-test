package models

type YoastIndexableHierarchy struct {
	IndexableID uint   `gorm:"column:indexable_id;primaryKey" json:"indexableId"`
	AncestorID  uint   `gorm:"column:ancestor_id;primaryKey" json:"ancestorId"`
	Depth       *uint  `gorm:"column:depth" json:"depth"`
	BlogID      uint64 `gorm:"column:blog_id" json:"blogId"`
}

func (YoastIndexableHierarchy) TableName() string {
	return "wp_yoast_indexable_hierarchy"
}
