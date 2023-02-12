package services

import (
	"errors"
	"ganhaum.henrybarreto.dev/internal/stores"
	"ganhaum.henrybarreto.dev/pkg/models"
)

var (
	ErrContributorInternal = errors.New("could not complete the operation on contributor due a internal error")
	ErrContributorNotFound = errors.New("could not find the contributor")
	ErrContributorEmpty    = errors.New("there are not any contributors yet")
)

var NewErrContributorInternal = func(err error) error {
	return errors.Join(ErrContributorInternal, err)
}

var NewErrContributorNotFound = func(err error) error {
	return errors.Join(ErrContributorNotFound, err)
}

var NewErrContributorEmpty = func(err error) error {
	return errors.Join(ErrContributorEmpty, err)
}

func (s *Service) CreateContributor(name string, email string, phone string, campaignID int) (*models.Contributor, error) {
	model, err := s.store.CreateContributor(&models.ContributorData{
		Name:  name,
		Email: email,
		Phone: phone,
	}, campaignID)
	if err != nil {
		return nil, NewErrContributorInternal(err)
	}

	return model, nil
}

func (s *Service) GetContributor(id int, campaignID int) (*models.Contributor, error) {
	model, err := s.store.GetContributor(id, campaignID)
	if err != nil {
		switch {
		case errors.Is(err, stores.ErrQuery) && errors.Is(err, stores.ErrScan):
			return nil, NewErrContributorInternal(err)
		case errors.Is(err, stores.ErrNotFound):
			return nil, NewErrContributorNotFound(err)
		default:
			return nil, NewErrContributorInternal(err)
		}
	}

	return model, nil
}

func (s *Service) GetContributors(campaignID int) ([]models.Contributor, error) {
	model, err := s.store.GetContributors(campaignID)
	if err != nil {
		switch {
		case errors.Is(err, stores.ErrQuery) && errors.Is(err, stores.ErrScan):
			return nil, NewErrContributorInternal(err)
		case errors.Is(err, stores.ErrNotFound):
			return nil, NewErrContributorEmpty(err)
		default:
			return nil, NewErrContributorInternal(err)
		}
	}

	return model, nil
}

func (s *Service) ConfirmContributor(id int, campaignID int) (*models.Contributor, error) {
	model, err := s.store.ConfirmContributor(id, campaignID)
	if err != nil {
		switch {
		case errors.Is(err, stores.ErrQuery) && errors.Is(err, stores.ErrScan):
			return nil, NewErrContributorInternal(err)
		case errors.Is(err, stores.ErrNotFound):
			return nil, NewErrContributorNotFound(err)
		default:
			return nil, NewErrContributorInternal(err)
		}
	}

	return model, nil
}
