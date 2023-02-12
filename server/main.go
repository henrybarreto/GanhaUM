package main

import (
	"fmt"

	"ganhaum.henrybarreto.dev/internal/api/http"
	"ganhaum.henrybarreto.dev/internal/services"
	"ganhaum.henrybarreto.dev/internal/stores/postgres"
)

func main() {
	fmt.Println("starting postgres store...")
	store, err := postgres.NewStore("postgres", "password", "postgres")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("loading services...")
	service := services.NewService(store)

	fmt.Println("watching for the HTTP server...")
	if err := http.NewHTTPServer(service, http.ServerOptions{
		Address:                ":8080",
		AllowedExternalAddress: nil,
	}); err != nil {
		return
	}
}
