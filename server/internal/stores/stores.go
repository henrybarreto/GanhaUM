package stores

import (
	"ganhaum.henrybarreto.dev/pkg/models"
)

type Product interface {
	CreateProduct(product *models.ProductData) (*models.Product, error)
	GetProduct(id int) (*models.Product, error)
	GetProducts(page int, limit int) ([]models.Product, error)
	UpdateProduct(id int, product *models.ProductData) (*models.Product, error)
	DeleteProduct(id int) error
}
type Campaign interface {
	CreateCampaign(campaign *models.CampaignData, productID int) (*models.Campaign, error)
	GetCampaign(id int) (*models.Campaign, error)
	GetCampaigns(page int, limit int) ([]models.Campaign, error)
	UpdateCampaign(id int, campaign *models.CampaignData) (*models.Campaign, error)
	DeleteCampaign(id int) error
}
type Contributor interface {
	CreateContributor(contributor *models.ContributorData, campaignID int) (*models.Contributor, error)
	GetContributor(id int, campaignID int) (*models.Contributor, error)
	GetContributors(campaignID int) ([]models.Contributor, error)
	ConfirmContributor(id int, campaignID int) (*models.Contributor, error)
}

type Result interface {
	CreateResult(campaignID int, receiverID int) (*models.Result, error)
	GetResult(id int) (*models.Result, error)
	GetResults() ([]models.Result, error)
}
type Stores interface {
	Product
	Campaign
	Contributor
	Result
}
