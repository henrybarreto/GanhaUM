package requests

import (
	"ganhaum.henrybarreto.dev/internal/api/http/queries"
	"time"
)

type CampaignCreate struct {
	ProductID            int       `uri:"product_id" binding:"required"`
	Title                string    `json:"title" binding:"required"`
	Value                int       `json:"value" binding:"required"`
	RequiredContributors int       `json:"required_contributors" binding:"required"`
	ClosedAt             time.Time `json:"closed_at" binding:"required"`
}

type CampaignGet struct {
	ID int `uri:"id" binding:"required"`
}

type CampaignsGet struct {
	queries.Pagination
}

type CampaignUpdate struct {
	ID                   int       `uri:"id" binding:"required"`
	ProductID            int       `uri:"product_id" binding:"required"`
	Title                string    `json:"title" binding:"required"`
	Value                int       `json:"value" binding:"required"`
	RequiredContributors int       `json:"required_contributors" binding:"required"`
	ClosedAt             time.Time `json:"closed_at" binding:"required"`
}

type CampaignDelete struct {
	ID int `uri:"id" binding:"required"`
}
