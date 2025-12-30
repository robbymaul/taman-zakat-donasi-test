package models

// WpTermRelationship represents table `wp_term_relationships`.
//
// Notes:
// - Composite primary key (object_id, term_taxonomy_id)
type WpTermRelationship struct {
	ObjectID       uint64 `gorm:"column:object_id;primaryKey" json:"objectId"`
	TermTaxonomyID uint64 `gorm:"column:term_taxonomy_id;primaryKey" json:"termTaxonomyId"`
	TermOrder      int    `gorm:"column:term_order" json:"termOrder"`
}

func (WpTermRelationship) TableName() string {
	return "wp_term_relationships"
}
