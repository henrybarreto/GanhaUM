package http

import (
	"errors"
	"fmt"
	"ganhaum.henrybarreto.dev/internal/api/http/frameworks/gin"
	"ganhaum.henrybarreto.dev/internal/services"
)

type ServerOptions struct {
	Address                string
	AllowedExternalAddress []string
}

var (
	ErrCreateHTTPServer = fmt.Errorf("failed to create HTTP server")
)

var NewErrCreateHTTPServer = func(err error) error {
	return errors.Join(ErrCreateHTTPServer, err)
}

func NewHTTPServer(service services.Services, options ServerOptions) error { // TODO: add options structure.
	if err := gin.NewGinServer(service, gin.Options{
		Address:                options.Address,
		AllowedExternalAddress: options.AllowedExternalAddress,
	}); err != nil {
		return NewErrCreateHTTPServer(err)
	}

	return nil
}
