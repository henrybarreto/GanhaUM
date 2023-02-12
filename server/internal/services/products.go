package services

import (
	"errors"
	"ganhaum.henrybarreto.dev/internal/stores"
	"ganhaum.henrybarreto.dev/pkg/models"
)

var (
	ErrProductInternal = errors.New("could not complete the operation on product due a internal error")
	ErrProductNotFound = errors.New("could not find the product")
	ErrProductsEmpty   = errors.New("there are not any products yet")
)

var NewErrProductInternal = func(err error) error {
	return errors.Join(ErrProductInternal, err)
}

var NewErrProductNotFound = func(err error) error {
	return errors.Join(ErrProductNotFound, err)
}

var NewErrProductsEmpty = func(err error) error {
	return errors.Join(ErrProductsEmpty, err)
}

func (s *Service) CreateProduct(thumbnail string, name string, description string, kind string, value int) (*models.Product, error) {
	model, err := s.store.CreateProduct(&models.ProductData{
		Thumbnail:   thumbnail,
		Name:        name,
		Description: description,
		Type:        kind,
		Value:       value,
	})
	if err != nil {
		return nil, NewErrProductInternal(err)
	}

	return model, nil
}

func (s *Service) GetProduct(id int) (*models.Product, error) {
	model, err := s.store.GetProduct(id)
	if err != nil {
		switch {
		case errors.Is(err, stores.ErrQuery) && errors.Is(err, stores.ErrScan):
			return nil, NewErrProductInternal(err)
		case errors.Is(err, stores.ErrNotFound):
			return nil, NewErrProductNotFound(err)
		default:
			return nil, NewErrProductInternal(err)
		}
	}

	return model, nil
}

func (s *Service) GetProducts(page int, limit int) ([]models.Product, error) {
	model, err := s.store.GetProducts(page, limit)
	if err != nil {
		switch {
		case errors.Is(err, stores.ErrQuery) && errors.Is(err, stores.ErrScan):
			return nil, NewErrProductInternal(err)
		case errors.Is(err, stores.ErrEmpty):
			return nil, NewErrProductsEmpty(err)
		default:
			return nil, NewErrProductInternal(err)
		}
	}

	return model, nil
}

func (s *Service) UpdateProduct(id int, thumbnail string, name string, description string, kind string, value int) (*models.Product, error) {
	model, err := s.store.UpdateProduct(id, &models.ProductData{
		Thumbnail:   thumbnail,
		Name:        name,
		Description: description,
		Type:        kind,
		Value:       value,
	})
	if err != nil {
		return nil, NewErrProductInternal(err)
	}

	return model, nil
}

func (s *Service) DeleteProduct(id int) error {
	err := s.store.DeleteProduct(id)
	if err != nil {
		return NewErrProductInternal(err)
	}

	return nil
}
