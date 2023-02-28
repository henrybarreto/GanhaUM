package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stripe/stripe-go/v74"

	"ganhaum.henrybarreto.dev/internal/api/http"
	"ganhaum.henrybarreto.dev/internal/services"
	"ganhaum.henrybarreto.dev/internal/stores/postgres"
)

func main() {
	fmt.Println("checking environment variables...")
	if key, ok := os.LookupEnv("STRIPE_SECRET_KEY"); !ok {
		panic("STRIPE_SECRET_KEY not set")
	} else {
		stripe.Key = key
	}

	if _, ok := os.LookupEnv("STRIPE_PRICE_ID"); !ok {
		panic("STRIPE_PRICE_ID not set")
	}

	fmt.Println("starting postgres store...")
	store, err := postgres.NewStore("postgres", "password", "postgres")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("loading services...")
	service := services.NewService(store)

	product, _ := service.CreateProduct(
		"https://i.imgur.com/3ZQZ9Z0.png",
		"Produto teste",
		"Descrição do produto teste",
		"giftcard",
		100,
	)

	service.CreateCampaign(
		"Contribuicao teste",
		100,
		10,
		time.Now().Add(time.Hour*24*7),
		product.ID,
	)

	fmt.Println("watching for the HTTP server...")
	if err := http.NewHTTPServer(service, http.ServerOptions{
		Address:                ":8080",
		AllowedExternalAddress: nil,
	}); err != nil {
		return
	}
}
