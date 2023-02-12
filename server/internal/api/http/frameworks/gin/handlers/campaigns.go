package handlers

import (
	"ganhaum.henrybarreto.dev/internal/api/http/requests"
	"ganhaum.henrybarreto.dev/internal/api/http/responses"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCampaign(c *gin.Context) {
	var req requests.CampaignCreate

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, responses.NewError(err))

		return
	}

	res, err := h.Service.CreateCampaign(req.Title, req.Value, req.RequiredContributors, req.ClosedAt, req.ProductID)
	if err != nil {
		c.JSON(500, responses.Error{
			Error: err.Error(),
		})

		return
	}

	c.JSON(201, res)
}

func (h *Handler) GetCampaign(c *gin.Context) {
	var req requests.CampaignGet

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, responses.NewError(err))

		return
	}

	res, err := h.Service.GetCampaign(req.ID)
	if err != nil {
		c.JSON(500, responses.NewError(err))

		return
	}

	c.JSON(200, res)
}

func (h *Handler) GetCampaigns(c *gin.Context) {
	var req requests.CampaignsGet

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, responses.NewError(err))

		return
	}

	res, err := h.Service.GetCampaigns(req.Page, req.Limit)
	if err != nil {
		c.JSON(500, responses.NewError(err))

		return
	}

	c.JSON(200, res)
}

func (h *Handler) UpdateCampaign(c *gin.Context) {
	var req requests.CampaignUpdate

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, responses.NewError(err))

		return
	}

	res, err := h.Service.UpdateCampaign(req.ID, req.Title, req.Value, req.RequiredContributors, req.ClosedAt)
	if err != nil {
		c.JSON(500, responses.NewError(err))

		return
	}

	c.JSON(200, res)
}

func (h *Handler) DeleteCampaign(c *gin.Context) {
	var req requests.CampaignDelete

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, responses.NewError(err))

		return
	}

	err := h.Service.DeleteCampaign(req.ID)
	if err != nil {
		c.JSON(500, responses.NewError(err))

		return
	}

	c.JSON(204, nil)
}
