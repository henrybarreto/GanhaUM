package handlers

import (
	"errors"
	"ganhaum.henrybarreto.dev/internal/api/http/requests"
	"ganhaum.henrybarreto.dev/internal/api/http/responses"
	"ganhaum.henrybarreto.dev/internal/services"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateProduct(c *gin.Context) {
	var body requests.ProductCreate

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(400, responses.NewError(err))

		return
	}

	res, err := h.Service.CreateProduct(body.Thumbnail, body.Name, body.Description, body.Kind, body.Value)
	if err != nil {
		c.JSON(500, responses.NewError(err))

		return
	}

	c.JSON(201, res)
}

func (h *Handler) GetProduct(c *gin.Context) {
	var req requests.ProductGet

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, responses.NewError(err))

		return
	}

	res, err := h.Service.GetProduct(req.ID)
	if err != nil {
		var code int
		switch {
		case errors.Is(err, services.ErrProductInternal):
			code = 500
		case errors.Is(err, services.ErrProductNotFound):
			code = 404
		}

		c.JSON(code, responses.NewError(err))

		return
	}

	c.JSON(200, res)
}

func (h *Handler) GetProducts(c *gin.Context) {
	var req requests.ProductsGet

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, responses.NewError(err))

		return
	}

	res, err := h.Service.GetProducts(req.Page, req.Limit)
	if err != nil {
		c.JSON(204, responses.NewError(err))

		return
	}

	c.JSON(200, res)
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	var req requests.ProductUpdate

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, responses.NewError(err))

		return
	}

	res, err := h.Service.UpdateProduct(req.ID, req.Thumbnail, req.Name, req.Description, req.Kind, req.Value)
	if err != nil {
		c.JSON(500, responses.NewError(err))

		return
	}

	c.JSON(200, res)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	var req requests.ProductDelete

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, responses.NewError(err))

		return
	}

	err := h.Service.DeleteProduct(req.ID)
	if err != nil {
		c.JSON(500, responses.NewError(err))

		return
	}

	c.JSON(204, nil)
}
