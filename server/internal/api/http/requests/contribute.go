package requests

type Contribute struct {
	CampaignID int `form:"campaign_id" binding:"required"`
}
