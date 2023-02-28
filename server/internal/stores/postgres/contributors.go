package postgres

import (
	"ganhaum.henrybarreto.dev/pkg/models"
)

func (r *Store) CreateContributor(contributor *models.ContributorData, campaignID int) (*models.Contributor, error) {
	var model models.Contributor

	campaign, err := r.GetCampaign(campaignID)
	if err != nil {
		// return nil, fmt.Errorf("failed to check the campaign to contributor: %w", err)
		return nil, NewErrCampaignNotFound(err)
	}

	rows, err := r.db.Query("INSERT INTO contributors (name, email, phone, campaign_id) VALUES ($1, $2, $3, $4) RETURNING *", contributor.Name, contributor.Email, contributor.Phone, campaignID)
	if err != nil {
		// return nil, fmt.Errorf("failed to create the contributor: %w", err)
		return nil, NewErrContributorQuery(err)
	}

	for rows.Next() {
		err = rows.Scan(&model.ID, &model.Name, &model.Email, &model.Phone, &model.Campaign.ID, &model.Confirmed, &model.CreatedAt, &model.UpdatedAt)
		if err != nil {
			// return nil, fmt.Errorf("failed to scan the contributor: %w", err)
			return nil, NewErrContributorScan(err)
		}
	}

	model.Campaign = *campaign

	return &model, nil
}

func (r *Store) GetContributor(id int, campaignID int) (*models.Contributor, error) {
	var model models.Contributor

	rows, err := r.db.Query("SELECT * FROM contributors WHERE id = $1 AND campaign_id = $2", id, campaignID)
	if err != nil {
		// return nil, fmt.Errorf("failed to get the contributor: %w", err)
		return nil, NewErrContributorQuery(err)
	}

	for rows.Next() {
		err := rows.Scan(&model.ID, &model.Name, &model.Email, &model.Phone, &model.Campaign.ID, &model.Confirmed, &model.CreatedAt, &model.UpdatedAt)
		if err != nil {
			// return nil, fmt.Errorf("failed to scan the contributor: %w", err)
			return nil, NewErrContributorScan(err)
		}
	}

	if model.ID == 0 {
		return nil, NewErrContributorNotFound(err)
	}

	return &model, nil
}

func (r *Store) GetContributors(campaignID int) ([]models.Contributor, error) {
	rows, err := r.db.Query("SELECT * FROM contributors WHERE campaign_id = $1", campaignID)
	if err != nil {
		// return nil, err
		return nil, NewErrContributorQuery(err)
	}

	var contributors []models.Contributor
	for rows.Next() {
		var c models.Contributor
		err = rows.Scan(&c.ID, &c.Name, &c.Email, &c.Phone, &c.Campaign.ID, &c.Confirmed, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			// return nil, err
			return nil, NewErrContributorScan(err)
		}

		contributors = append(contributors, c)
	}

	if contributors == nil {
		return nil, NewErrContributorNotFound(err)
	}

	return contributors, nil
}

func (r *Store) ConfirmContributor(id int, campaignID int) (*models.Contributor, error) {
	var model models.Contributor

	var err error
	_, err = r.GetContributor(id, campaignID)
	if err != nil {
		// return nil, fmt.Errorf("failed to check the contributor: %w", err)
		return nil, NewErrContributorNotFound(err)
	}

	campaign, err := r.GetCampaign(campaignID)
	if err != nil {
		// return nil, fmt.Errorf("failed to check the campaign to contributor: %w", err)
		return nil, NewErrCampaignNotFound(err)
	}

	rows, err := r.db.Query("UPDATE contributors SET confirmed = true WHERE id = $1 AND campaign_id = $2 RETURNING *", id, campaignID)
	if err != nil {
		// return nil, fmt.Errorf("failed to confirm the contributor: %w", err)
		return nil, NewErrContributorQuery(err)
	}

	for rows.Next() {
		err := rows.Scan(&model.ID, &model.Name, &model.Email, &model.Phone, &model.Campaign.ID, &model.Confirmed, &model.CreatedAt, &model.UpdatedAt)
		if err != nil {
			// return nil, fmt.Errorf("failed to scan the contributor: %w", err)
			return nil, NewErrContributorScan(err)
		}
	}

	_, err = r.db.Query("UPDATE campaigns SET total_contributors = total_contributors + 1 WHERE id = $1", campaignID)
	if err != nil {
		// return nil, fmt.Errorf("failed to increase the total contributors: %w", err)
		return nil, NewErrCampaignQuery(err)
	}

	model.Campaign = *campaign

	return &model, nil
}
