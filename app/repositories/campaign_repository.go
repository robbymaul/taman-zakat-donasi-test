package repositories

import "donasitamanzakattest/app/models"

func (rc *RepositoryContext) PublicListCampaignRepository() ([]*models.ListPriorityWpDjaCampaign, error) {
	var campaigns []*models.ListPriorityWpDjaCampaign
	var err error
	var db = rc.db

	query := "SELECT wp_dja_campaign.id as id, wp_dja_campaign.user_id as user_id, wp_dja_campaign.campaign_id as campaign_id, wp_dja_campaign.title as title, wp_dja_campaign.slug as slug, wp_dja_campaign.image_url as image_url,wp_users.display_name as owner " +
		"FROM wp_dja_campaign " +
		"JOIN wp_users ON wp_dja_campaign.user_id = wp_users.id " +
		"WHERE wp_dja_campaign.publish_status = ? " +
		"LIMIT ?"

	db = db.Raw(query, 1, 10)

	db = db.Scan(&campaigns)

	err = db.Error
	if err != nil {
		return nil, err
	}

	return campaigns, err
}

func (rc *RepositoryContext) PublicGetCampaignRepository(slug string) (*models.DetailWpDjaCampaign, error) {
	var campaign *models.DetailWpDjaCampaign
	var err error
	var db = rc.db

	query := "SELECT wdc.id as id, wdc.user_id as user_id, wdc.campaign_id as campaign_id, wdc.slug as slug, wdc.image_url as image_url, wdc.title as title, wdc.target as target, COALESCE(SUM(wdd.nominal), 0) AS total_donasi, COUNT(wdd.id) AS total_donatur, wu.display_name as owner, wu.user_url as owner_image, wdu.user_type as owner_type, wdu.user_verification as owner_verification, wdc.information as information " +
		"FROM wp_dja_campaign AS wdc " +
		"JOIN wp_users AS wu ON wdc.user_id = wu.id " +
		"JOIN wp_dja_users AS wdu ON wdu.user_id = wu.id " +
		"LEFT JOIN wp_dja_donate AS wdd ON wdd.campaign_id = wdc.campaign_id AND wdd.payment_at IS NOT NULL " +
		"WHERE wdc.slug = ? " +
		"GROUP BY wdc.id, wdc.user_id, wdc.campaign_id, wdc.slug, wdc.image_url, wdc.title, wdc.target, wu.display_name, wu.user_url, wdu.user_type, wdu.user_verification, wdc.information"

	db = db.Raw(query, slug)

	db = db.Scan(&campaign)

	err = db.Error
	if err != nil {
		return nil, err
	}

	return campaign, err
}

func (rc *RepositoryContext) PublicListGetCampaignDonaturRepository(campaignId string, sort string) ([]*models.ListCampaignWpDjaDonate, error) {
	var donatur []*models.ListCampaignWpDjaDonate
	var err error
	var db = rc.db

	query := "SELECT wdd.id as id, wdd.campaign_id as campaign_id, wdd.name as name, wdd.nominal as nominal, wdd.payment_at as payment_at, wdd.created_at as created_at " +
		"FROM wp_dja_donate as wdd " +
		"WHERE wdd.campaign_id = ? AND wdd.payment_at IS NOT NULL "

	if sort == "terbaru" {
		query += "ORDER BY wdd.payment_at DESC "
	} else if sort == "terbesar" {
		query += "ORDER BY wdd.nominal DESC "
	} else {
		query += "ORDER BY wdd.payment_at DESC "
	}

	query += "LIMIT ? "

	db = db.Raw(query, campaignId, 10)

	db = db.Scan(&donatur)

	err = db.Error
	if err != nil {
		return nil, err
	}

	return donatur, err
}

func (rc *RepositoryContext) PublicListGetCampaignCommentRepository(campaignId string) ([]*models.ListCampaignWpDjaComment, error) {
	var comments []*models.ListCampaignWpDjaComment
	var err error
	var db = rc.db

	query := "SELECT wdd.id as id, wdd.campaign_id as campaign_id, wdd.name as name, wdd.comment as comment, wdd.created_at as created_at " +
		"FROM wp_dja_donate as wdd " +
		"WHERE wdd.campaign_id = ? AND wdd.payment_at IS NOT NULL AND wdd.comment != ? " +
		"ORDER BY wdd.created_at DESC "

	query += "LIMIT ? "

	db = db.Raw(query, campaignId, "", 10)

	db = db.Scan(&comments)

	err = db.Error
	if err != nil {
		return nil, err
	}

	return comments, err
}
