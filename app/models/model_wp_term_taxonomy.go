package models

type WpTermTaxonomy struct {
	TermTaxonomyID uint64 `gorm:"column:term_taxonomy_id;primaryKey" json:"termTaxonomyId"`
	TermID         uint64 `gorm:"column:term_id" json:"termId"`
	Taxonomy       string `gorm:"column:taxonomy" json:"taxonomy"`
	Description    string `gorm:"column:description" json:"description"`
	Parent         uint64 `gorm:"column:parent" json:"parent"`
	Count          int64  `gorm:"column:count" json:"count"`
}

func (WpTermTaxonomy) TableName() string {
	return "wp_term_taxonomy"
}
