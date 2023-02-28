package postgres

import (
	"ganhaum.henrybarreto.dev/pkg/models"
)

func (r *Store) CreateResult(campaignID int, receiverID int) (*models.Result, error) {
	var model models.Result

	campaign, err := r.GetCampaign(campaignID)
	if err != nil {
		// return nil, fmt.Errorf("failed to check the campaign to result: %w", err)
		return nil, NewErrCampaignNotFound(err)
	}

	receiver, err := r.GetContributor(receiverID, campaignID)
	if err != nil {
		// return nil, fmt.Errorf("failed to check the receiver to result: %w", err)
		return nil, NewErrContributorNotFound(err)
	}

	amount := campaign.Value / campaign.TotalContributors

	rows, err := r.db.Query("INSERT INTO results (campaign_id, receiver_id, amount) VALUES ($1, $2, $3) RETURNING *", campaignID, receiverID, amount)
	if err != nil {
		// return nil, fmt.Errorf("failed to create the result: %w", err)
		return nil, NewErrResultQuery(err)
	}

	for rows.Next() {
		err = rows.Scan(&model.ID, &model.Campaign.ID, &model.Amount, &model.Receiver.ID, &model.ClosedAt)
		if err != nil {
			// return nil, fmt.Errorf("failed to scan the result: %w", err)
			return nil, NewErrResultScan(err)
		}
	}

	model.Campaign = *campaign
	model.Receiver = *receiver

	return &model, nil
}

func (r *Store) GetResult(id int) (*models.Result, error) {
	var model models.Result

	rows, err := r.db.Query("SELECT * FROM results WHERE id = $1", id)
	if err != nil {
		// return nil, fmt.Errorf("failed to get the result: %w", err)
		return nil, NewErrResultQuery(err)
	}

	for rows.Next() {
		err = rows.Scan(&model.ID, &model.Campaign.ID, &model.Amount, &model.Receiver.ID, &model.ClosedAt)
		if err != nil {
			// return nil, fmt.Errorf("failed to scan the result: %w", err)
			return nil, NewErrResultScan(err)
		}
	}

	if model.ID == 0 {
		return nil, NewErrResultNotFound(err)
	}

	return &model, nil
}

func (r *Store) GetResults() ([]models.Result, error) {
	rows, err := r.db.Query("SELECT * FROM results")
	if err != nil {
		// return nil, err
		return nil, NewErrResultQuery(err)
	}

	var results []models.Result
	for rows.Next() {
		var r models.Result
		err = rows.Scan(&r.ID, &r.Campaign.ID, &r.Amount, &r.Receiver.ID, &r.ClosedAt)
		if err != nil {
			// return nil, err
			return nil, NewErrResultScan(err)
		}

		results = append(results, r)
	}

	if results == nil {
		return nil, NewErrResultNotFound(err)
	}

	return results, nil
}
