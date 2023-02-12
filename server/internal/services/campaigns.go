package services

import (
	"errors"
	"fmt"
	"ganhaum.henrybarreto.dev/internal/stores"
	"ganhaum.henrybarreto.dev/pkg/models"
	"time"
)

var (
	ErrCampaignInternal = errors.New("could not complete the operation on campaign due a internal error")
	ErrCampaignNotFound = errors.New("could not find the campaign")
	ErrCampaignEmpty    = errors.New("could not find any campaign")
	ErrCampaignClosedAt = errors.New("close date must be in the future")
)

var NewErrCampaignInternal = func(err error) error {
	return errors.Join(ErrCampaignInternal, err)
}

var NewErrCampaignNotFound = func(err error) error {
	return errors.Join(ErrCampaignNotFound, err)
}

var NewErrCampaignEmpty = func(err error) error {
	return errors.Join(ErrCampaignEmpty, err)
}

var NewErrCampaignClosedAt = func(closedAt time.Time) error {
	return errors.Join(ErrCampaignClosedAt, fmt.Errorf("close date must be in the future: %s", closedAt))
}

func (s *Service) CreateCampaign(title string, value int, requiredContributors int, closedAt time.Time, productID int) (*models.Campaign, error) {
	if closedAt.Before(time.Now()) {
		return nil, fmt.Errorf("close date must be in the future")
	}

	model, err := s.store.CreateCampaign(&models.CampaignData{
		Title:                title,
		Value:                value,
		RequiredContributors: requiredContributors,
		TotalContributors:    0,
		ClosedAt:             closedAt,
	}, productID)
	if err != nil {
		return nil, NewErrCampaignInternal(err)
	}

	return model, nil
}

func (s *Service) GetCampaign(id int) (*models.Campaign, error) {
	model, err := s.store.GetCampaign(id)
	if err != nil {
		switch {
		case errors.Is(err, stores.ErrQuery) && errors.Is(err, stores.ErrScan):
			return nil, NewErrCampaignInternal(err)
		case errors.Is(err, stores.ErrNotFound):
			return nil, NewErrCampaignNotFound(err)
		default:
			return nil, NewErrCampaignInternal(err)
		}
	}

	return model, nil
}

func (s *Service) GetCampaigns(page int, limit int) ([]models.Campaign, error) {
	models, err := s.store.GetCampaigns(page, limit)
	if err != nil {
		switch {
		case errors.Is(err, stores.ErrQuery) && errors.Is(err, stores.ErrScan):
			return nil, NewErrCampaignInternal(err)
		case errors.Is(err, stores.ErrNotFound):
			return nil, NewErrCampaignEmpty(err)
		}
	}

	return models, nil
}

func (s *Service) UpdateCampaign(id int, title string, value int, requiredContributors int, closedAt time.Time) (*models.Campaign, error) {
	if closedAt.Before(time.Now()) {
		// return nil, fmt.Errorf("close date must be in the future")
		return nil, NewErrCampaignClosedAt(closedAt)
	}

	model, err := s.store.UpdateCampaign(id, &models.CampaignData{
		Title:                title,
		Value:                value,
		RequiredContributors: requiredContributors,
		TotalContributors:    0,
		ClosedAt:             closedAt,
	})
	if err != nil {
		return nil, NewErrCampaignNotFound(err)
	}

	return model, nil
}

func (s *Service) DeleteCampaign(id int) error {
	if err := s.store.DeleteCampaign(id); err != nil {
		return NewErrCampaignInternal(err)
	}

	return nil
}
