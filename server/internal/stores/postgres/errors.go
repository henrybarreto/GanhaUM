package postgres

import (
	"errors"
	"ganhaum.henrybarreto.dev/internal/stores"
)

var NewErrProductQuery = func(err error) error {
	return errors.Join(stores.ErrQuery, err)
}

var NewErrProductScan = func(err error) error {
	return errors.Join(stores.ErrScan, err)
}

var NewErrProductsScan = func(err error) error {
	return errors.Join(stores.ErrScans, err)
}

var NewErrProductNotFound = func(err error) error {
	return errors.Join(stores.ErrNotFound, err)
}

var NewErrProductsEmpty = func(err error) error {
	return errors.Join(stores.ErrEmpty, err)
}

var NewErrContributorNotFound = func(err error) error {
	return errors.Join(stores.ErrNotFound, err)
}

var NewErrContributorQuery = func(err error) error {
	return errors.Join(stores.ErrQuery, err)
}

var NewErrContributorScan = func(err error) error {
	return errors.Join(stores.ErrScan, err)
}

var NewErrCampaignNotFound = func(err error) error {
	return errors.Join(stores.ErrNotFound, err)
}

var NewErrCampaignQuery = func(err error) error {
	return errors.Join(stores.ErrQuery, err)
}

var NewErrCampaignScan = func(err error) error {
	return errors.Join(stores.ErrScan, err)
}

var NewErrResultNotFound = func(err error) error {
	return errors.Join(stores.ErrNotFound, err)
}

var NewErrResultQuery = func(err error) error {
	return errors.Join(stores.ErrQuery, err)
}

var NewErrResultScan = func(err error) error {
	return errors.Join(stores.ErrScan, err)
}

var NewErrCampaignEmpty = func(err error) error {
	return errors.Join(stores.ErrEmpty, err)
}
