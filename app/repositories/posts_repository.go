package repositories

import "donasitamanzakattest/app/models"

func (rc *RepositoryContext) PublicListPostsRepository() ([]*models.ListWpPost, error) {
	var posts []*models.ListWpPost
	var err error
	var db = rc.db

	query := "SELECT wpp.ID as id, wpp.post_author as post_author, wpp.post_date as post_date, wpp.post_title as post_title, wpp.post_name as post_name " +
		"FROM wp_posts as wpp " +
		"WHERE wpp.post_status = ? AND wpp.post_author = ? " +
		"ORDER BY wpp.ID DESC "

	query += "LIMIT ?"

	db = db.Raw(query, "publish", 1, 2)

	db = db.Scan(&posts)

	err = db.Error
	if err != nil {
		return nil, err
	}

	return posts, err
}

func (rc *RepositoryContext) PublicGetPostsRepository(postName string) (*models.DetailWpPost, error) {
	var posts *models.DetailWpPost
	var err error
	var db = rc.db

	query := "SELECT wpp.ID as id, wpp.post_author as post_author, wpp.post_date as post_date, wpp.post_title as post_title, wpp.post_name as post_name, wpp.post_content as post_content, wpp.post_status as post_status, wpp.comment_status as comment_status, wpp.ping_status as ping_status " +
		"FROM wp_posts as wpp " +
		"WHERE wpp.post_name = ? "

	db = db.Raw(query, postName)

	db = db.Scan(&posts)

	err = db.Error
	if err != nil {
		return nil, err
	}

	return posts, err
}
