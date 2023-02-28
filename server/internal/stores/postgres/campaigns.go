package postgres

import (
	"ganhaum.henrybarreto.dev/pkg/models"
)

func (r *Store) CreateCampaign(campaign *models.CampaignData, productID int) (*models.Campaign, error) {
	var model models.Campaign

	product, err := r.GetProduct(productID)
	if err != nil {
		// return nil, fmt.Errorf("failed to load the product to compaing: %w", err)
		return nil, NewErrCampaignNotFound(err)
	}

	rows, err := r.db.Query("INSERT INTO campaigns (title, value, required_contributors, total_contributors, closed_at, product_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *", campaign.Title, campaign.Value, campaign.RequiredContributors, campaign.TotalContributors, campaign.ClosedAt, productID)
	if err != nil {
		// return nil, fmt.Errorf("failed to create the campaign: %w", err)
		return nil, NewErrCampaignQuery(err)
	}

	for rows.Next() {
		err = rows.Scan(&model.ID, &model.Title, &model.Value, &model.Product.ID, &model.RequiredContributors, &model.TotalContributors, &model.StartedAt, &model.UpdatedAt, &model.ClosedAt)
		if err != nil {
			// return nil, fmt.Errorf("failed to scan the campaign: %w", err)
			return nil, NewErrCampaignScan(err)
		}
	}

	model.Product = *product

	return &model, nil
}

func (r *Store) GetCampaign(id int) (*models.Campaign, error) {
	var model models.Campaign

	rows, err := r.db.Query("SELECT * FROM campaigns WHERE id = $1", id)
	if err != nil {
		// return nil, fmt.Errorf("failed to get the campaign: %w", err)
		return nil, NewErrCampaignQuery(err)
	}

	for rows.Next() {
		err = rows.Scan(&model.ID, &model.Title, &model.Value, &model.Product.ID, &model.RequiredContributors, &model.TotalContributors, &model.StartedAt, &model.UpdatedAt, &model.ClosedAt)
		if err != nil {
			// return nil, fmt.Errorf("failed to scan the campaign: %w", err)
			return nil, NewErrCampaignScan(err)
		}
	}

	if model.ID == 0 {
		return nil, NewErrCampaignNotFound(err)
	}

	return &model, nil
}

func (r *Store) GetCampaigns(page int, limit int) ([]models.Campaign, error) {
	rows, err := r.db.Query("SELECT * FROM campaigns LIMIT $1 OFFSET $2", limit, (page-1)*limit)
	if err != nil {
		// return nil, fmt.Errorf("failed to get the campaigns: %w", err)
		return nil, NewErrCampaignQuery(err)
	}

	var campaigns []models.Campaign
	for rows.Next() {
		var c models.Campaign
		err = rows.Scan(&c.ID, &c.Title, &c.Value, &c.Product.ID, &c.RequiredContributors, &c.TotalContributors, &c.StartedAt, &c.UpdatedAt, &c.ClosedAt)
		if err != nil {
			// return nil, fmt.Errorf("failed to scan the campaign: %w", err)
			return nil, NewErrCampaignScan(err)
		}

		campaigns = append(campaigns, c)
	}

	if campaigns == nil {
		return nil, NewErrCampaignNotFound(err)
	}

	return campaigns, nil
}

func (r *Store) UpdateCampaign(id int, campaign *models.CampaignData) (*models.Campaign, error) {
	var model models.Campaign

	rows, err := r.db.Query("UPDATE campaigns SET title = $1, value = $2, required_contributors = $4, total_contributors = $5, closed_at = $6 WHERE id = $7 RETURNING *", campaign.Title, campaign.Value, campaign.RequiredContributors, campaign.TotalContributors, campaign.ClosedAt, id)
	if err != nil {
		// return nil, fmt.Errorf("failed to update the campaign: %w", err)
		return nil, NewErrCampaignQuery(err)
	}

	for rows.Next() {
		err = rows.Scan(&model.ID, &model.Title, &model.Value, &model.Product, &model.RequiredContributors, &model.TotalContributors, &model.StartedAt, &model.UpdatedAt)
		if err != nil {
			// return nil, fmt.Errorf("failed to scan the campaign: %w", err)
			return nil, NewErrCampaignScan(err)
		}
	}

	return &model, nil
}

func (r *Store) DeleteCampaign(id int) error {
	_, err := r.db.Query("DELETE FROM campaigns WHERE id = $1", id)
	if err != nil {
		// return fmt.Errorf("failed to delete the campaign: %w", err)
		return NewErrCampaignQuery(err)
	}

	return nil
}
