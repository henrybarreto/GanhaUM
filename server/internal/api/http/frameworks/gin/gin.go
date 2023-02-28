package gin

import (
	"fmt"

	"ganhaum.henrybarreto.dev/internal/api/http/frameworks/gin/handlers"
	"ganhaum.henrybarreto.dev/internal/api/http/frameworks/gin/middleware"
	"ganhaum.henrybarreto.dev/internal/api/http/paths"
	"ganhaum.henrybarreto.dev/internal/services"
	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

type Options struct {
	Address                string
	AllowedExternalAddress []string
}

func NewGinServer(service services.Services, options Options) error {
	handler := handlers.Handler{
		Service: service,
	}

	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	api := server.Group(paths.API)
    api.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
    }))

	public := api.Group(paths.Public)
	public.POST(paths.CreateContributor, handler.CreateContributor)

	public.GET(paths.GetProduct, handler.GetProduct)
	public.GET(paths.GetProducts, handler.GetProducts)
	public.GET(paths.GetCampaign, handler.GetCampaign)
	public.GET(paths.GetCampaigns, handler.GetCampaigns)
	public.GET(paths.GetContributor, handler.GetContributor)
	public.GET(paths.GetContributors, handler.GetContributors)
	public.GET(paths.GetResult, handler.GetResult)
	public.GET(paths.GetResults, handler.GetResults)

	public.GET(paths.Contribute, handler.Contribute)

	authenticated := api.Group(paths.Authenticated)
	authenticated.Use(middleware.AuthenticationJWT)

	authenticated.POST(paths.CreateProduct, handler.CreateProduct)
	authenticated.POST(paths.CreateCampaign, handler.CreateCampaign)
	authenticated.POST(paths.CreateResult, handler.CreateResult)

	authenticated.PUT(paths.UpdateProduct, handler.UpdateProduct)
	authenticated.PUT(paths.UpdateCampaign, handler.UpdateCampaign)
	authenticated.DELETE(paths.DeleteProduct, handler.DeleteProduct)
	authenticated.DELETE(paths.DeleteCampaign, handler.DeleteCampaign)

	webhook := api.Group(paths.Webhook)
	webhook.Use(middleware.PermissionAllowOnly(options.AllowedExternalAddress...))

	webhook.POST(paths.ConfirmContributor, handler.ConfirmContributor)
	webhook.POST(paths.CreateResult, handler.CreateResult)

	if err := server.Run(options.Address); err != nil {
		return fmt.Errorf("failed to create the Gin HTTP server: %w", err)
	}

	return nil
}
