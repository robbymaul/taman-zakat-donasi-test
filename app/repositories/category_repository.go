package repositories

import "donasitamanzakattest/app/models"

func (rc *RepositoryContext) PublicListCategoriesRepositories() ([]*models.WpDjaCategory, error) {
	var categories []*models.WpDjaCategory
	var err error

	query := "SELECT wp_dja_category.id as id, " +
		"wp_dja_category.category as category, " +
		"wp_dja_category.private_category as private_category FROM wp_dja_category"

	err = rc.db.Raw(query).Scan(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, err
}
