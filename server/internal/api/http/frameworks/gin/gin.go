package gin

import (
	"fmt"
	"ganhaum.henrybarreto.dev/internal/api/http/frameworks/gin/handlers"
	middleware2 "ganhaum.henrybarreto.dev/internal/api/http/frameworks/gin/middleware"
	paths2 "ganhaum.henrybarreto.dev/internal/api/http/paths"
	"ganhaum.henrybarreto.dev/internal/services"

	"github.com/gin-gonic/gin"
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

	api := server.Group(paths2.API)

	authenticated := api.Group(paths2.Authenticated)
	authenticated.Use(middleware2.AuthenticationJWT)

	authenticated.POST(paths2.CreateProduct, handler.CreateProduct)
	authenticated.POST(paths2.CreateCampaign, handler.CreateCampaign)
	authenticated.POST(paths2.CreateResult, handler.CreateResult)

	authenticated.PUT(paths2.UpdateProduct, handler.UpdateProduct)
	authenticated.PUT(paths2.UpdateCampaign, handler.UpdateCampaign)
	authenticated.DELETE(paths2.DeleteProduct, handler.DeleteProduct)
	authenticated.DELETE(paths2.DeleteCampaign, handler.DeleteCampaign)

	public := api.Group(paths2.Public)
	public.POST(paths2.CreateContributor, handler.CreateContributor)

	public.GET(paths2.GetProduct, handler.GetProduct)
	public.GET(paths2.GetProducts, handler.GetProducts)
	public.GET(paths2.GetCampaign, handler.GetCampaign)
	public.GET(paths2.GetCampaigns, handler.GetCampaigns)
	public.GET(paths2.GetContributor, handler.GetContributor)
	public.GET(paths2.GetContributors, handler.GetContributors)
	public.GET(paths2.GetResult, handler.GetResult)
	public.GET(paths2.GetResults, handler.GetResults)

	external := api.Group(paths2.External)
	external.Use(middleware2.PermissionAllowOnly(options.AllowedExternalAddress...))

	external.POST(paths2.ConfirmContributor, handler.ConfirmContributor)
	external.POST(paths2.CreateResult, handler.CreateResult)

	if err := server.Run(options.Address); err != nil {
		return fmt.Errorf("failed to create the Gin HTTP server: %w", err)
	}

	return nil
}
