package services

import (
	"ganhaum.henrybarreto.dev/internal/stores"
	"ganhaum.henrybarreto.dev/pkg/models"
	"time"
)

type Product interface {
	CreateProduct(thumbnail string, name string, description string, kind string, value int) (*models.Product, error)
	GetProduct(id int) (*models.Product, error)
	GetProducts(page int, limit int) ([]models.Product, error)
	UpdateProduct(id int, thumbnail string, name string, description string, kind string, value int) (*models.Product, error)
	DeleteProduct(id int) error
}

type Campaign interface {
	CreateCampaign(title string, value int, requiredContributors int, closedAt time.Time, productID int) (*models.Campaign, error)
	GetCampaign(id int) (*models.Campaign, error)
	GetCampaigns(page int, limit int) ([]models.Campaign, error)
	UpdateCampaign(id int, title string, value int, requiredContributors int, closedAt time.Time) (*models.Campaign, error) // TODO: add param for ProductID
	DeleteCampaign(id int) error
}

type Contributor interface {
	CreateContributor(name string, email string, phone string, campaignID int) (*models.Contributor, error)
	GetContributor(id int, campaignID int) (*models.Contributor, error)
	GetContributors(campaignID int) ([]models.Contributor, error)
	ConfirmContributor(id int, campaignID int) (*models.Contributor, error)
}

type Results interface {
	CreateResult(campaignID int, receiverID int) (*models.Result, error)
	GetResult(id int) (*models.Result, error)
	GetResults() ([]models.Result, error)
}

type Services interface {
	Product
	Campaign
	Contributor
	Results
}

type Service struct {
	store stores.Stores
}

func NewService(store stores.Stores) Services {
	return &Service{store: store}
}
