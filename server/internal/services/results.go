package services

import (
	"errors"
	"ganhaum.henrybarreto.dev/internal/stores"
	"ganhaum.henrybarreto.dev/pkg/models"
	"time"
)

var (
	ErrResultInternal = errors.New("could not complete the operation on result due a internal error")
	ErrResultNotFound = errors.New("could not find the result")
	ErrResultEmpty    = errors.New("could not find any result")
)

var NewErrResultInternal = func(err error) error {
	return errors.Join(ErrResultInternal, err)
}

var NewErrResultNotFound = func(err error) error {
	return errors.Join(ErrResultNotFound, err)
}

var NewErrResultEmpty = func(err error) error {
	return errors.Join(ErrResultEmpty, err)
}

func (s *Service) CreateResult(campaignID int, receiverID int) (*models.Result, error) {
	campaign, err := s.store.GetCampaign(campaignID)
	if err != nil {
		return nil, NewErrCampaignNotFound(err)
	}

	if campaign.ClosedAt.Unix() < time.Now().Unix() {
		return nil, NewErrCampaignClosedAt(campaign.ClosedAt)
	}

	model, err := s.store.CreateResult(campaign.ID, receiverID)
	if err != nil {
		return nil, NewErrResultInternal(err)
	}

	return model, nil
}

func (s *Service) GetResult(id int) (*models.Result, error) {
	model, err := s.store.GetResult(id)
	if err != nil {
		switch {
		case errors.Is(err, stores.ErrQuery) && errors.Is(err, stores.ErrScan):
			return nil, NewErrResultInternal(err)
		case errors.Is(err, stores.ErrNotFound):
			return nil, NewErrResultNotFound(err)
		default:
			return nil, NewErrResultNotFound(err)
		}
	}

	return model, nil
}

func (s *Service) GetResults() ([]models.Result, error) {
	model, err := s.store.GetResults()
	if err != nil {
		switch {
		case errors.Is(err, stores.ErrQuery) && errors.Is(err, stores.ErrScan):
			return nil, NewErrResultInternal(err)
		case errors.Is(err, stores.ErrEmpty):
			return nil, NewErrResultEmpty(err)
		default:
			return nil, NewErrResultInternal(err)
		}
	}

	return model, nil
}
