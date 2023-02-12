package handlers

import (
	"ganhaum.henrybarreto.dev/internal/api/http/requests"
	"ganhaum.henrybarreto.dev/internal/api/http/responses"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateContributor(c *gin.Context) {
	var req requests.ContributorCreate

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, responses.Error{
			Error: err.Error(),
		})

		return
	}

	res, err := h.Service.CreateContributor(req.Name, req.Email, req.Phone, req.CampaignID)
	if err != nil {
		c.JSON(500, responses.Error{
			Error: err.Error(),
		})

		return
	}

	c.JSON(201, res)
}

func (h *Handler) GetContributor(c *gin.Context) {
	var req requests.ContributorGet

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, responses.Error{
			Error: err.Error(),
		})

		return
	}

	res, err := h.Service.GetContributor(req.ID, req.CampaignID)
	if err != nil {
		c.JSON(500, responses.Error{
			Error: err.Error(),
		})

		return
	}

	c.JSON(200, res)
}

func (h *Handler) GetContributors(c *gin.Context) {
	var req requests.ContributorsGet

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, responses.Error{
			Error: err.Error(),
		})

		return
	}

	res, err := h.Service.GetContributors(req.CampaignID)
	if err != nil {
		c.JSON(500, responses.Error{
			Error: err.Error(),
		})

		return
	}

	c.JSON(200, res)
}

func (h *Handler) ConfirmContributor(c *gin.Context) {
	var req requests.ContributorGet

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, responses.Error{
			Error: err.Error(),
		})

		return
	}

	res, err := h.Service.ConfirmContributor(req.ID, req.CampaignID)
	if err != nil {
		c.JSON(500, responses.Error{
			Error: err.Error(),
		})

		return
	}

	c.JSON(200, res)
}
