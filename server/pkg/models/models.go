package models

import (
	"time"
)

type ProductData struct {
	Thumbnail   string `json:"thumbnail" db:"thumbnail"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Type        string `json:"type" db:"type"`
	Value       int    `json:"value" db:"value"`
}

type Product struct {
	ID int `json:"id" db:"id"`
	ProductData
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CampaignData struct {
	Title                string    `json:"title" db:"title"`
	Value                int       `json:"value" db:"value"`
	RequiredContributors int       `json:"required_participants" db:"required_participants"`
	TotalContributors    int       `json:"total_participants" db:"total_participants"`
	ClosedAt             time.Time `json:"closed_at" db:"closed_at"`
}
type Campaign struct {
	ID int `json:"id" db:"id"`
	CampaignData
	Product   Product   `json:"products" db:"products"`
	StartedAt time.Time `json:"started_at" db:"started_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type ContributorData struct {
	Name      string `json:"name" db:"name"`
	Email     string `json:"email" db:"email"`
	Phone     string `json:"phone" db:"phone"`
	Confirmed bool   `json:"confirmed" db:"confirmed"`
}

type Contributor struct {
	ID int `json:"id" db:"id"`
	ContributorData
	Campaign  Campaign  `json:"campaign" db:"campaign"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Result struct {
	ID       int         `json:"id" db:"id"`
	Campaign Campaign    `json:"campaign" db:"campaign"`
	Receiver Contributor `json:"receiver" db:"receiver"`
	Amount   int         `json:"amount" db:"amount"`
	ClosedAt time.Time   `json:"closed_at" db:"closed_at"`
}
