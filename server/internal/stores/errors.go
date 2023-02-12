package stores

import "fmt"

var (
	ErrConnectDatabase = fmt.Errorf("failed to connect to the database")
	ErrPingDatabase    = fmt.Errorf("failed to ping the database")
)

var (
	ErrQuery    = fmt.Errorf("the query to the database failed")
	ErrScan     = fmt.Errorf("it could not to populate the models with the data from the database")
	ErrScans    = fmt.Errorf("it could not to populate the model with the data from the database")
	ErrNotFound = fmt.Errorf("this resource does not exist")
	ErrEmpty    = fmt.Errorf("this resource is empty")
)
