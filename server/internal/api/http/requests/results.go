package requests

type ResultCreate struct {
	CampaignID int `json:"campaign_id" binding:"required"`
	ReceiverID int `json:"receiver_id" binding:"required"`
}

type ResultGet struct {
	ID         int `uri:"id" binding:"required"`
	CampaignID int `uri:"campaign_id" binding:"required"`
}

type ResultsGet struct{}
