package requests

type ContributorCreate struct {
	CampaignID int    `json:"campaign_id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
}

type ContributorGet struct {
	ID         int `uri:"id" binding:"required"`
	CampaignID int `uri:"campaign_id" binding:"required"`
}

type ContributorsGet struct {
	CampaignID int `uri:"campaign_id" binding:"required"`
}
