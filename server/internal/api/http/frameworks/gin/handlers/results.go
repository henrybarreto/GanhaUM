package handlers

import (
	"ganhaum.henrybarreto.dev/internal/api/http/requests"
	"ganhaum.henrybarreto.dev/internal/api/http/responses"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateResult(c *gin.Context) {
	var req requests.ResultCreate

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, responses.Error{
			Error: err.Error(),
		})

		return
	}

	res, err := h.Service.CreateResult(req.CampaignID, req.ReceiverID)
	if err != nil {
		c.JSON(500, responses.Error{
			Error: err.Error(),
		})

		return
	}

	c.JSON(201, res)
}

func (h *Handler) GetResult(c *gin.Context) {
	var req requests.ResultGet

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, responses.Error{
			Error: err.Error(),
		})

		return
	}

	res, err := h.Service.GetResult(req.ID)
	if err != nil {
		c.JSON(500, responses.Error{
			Error: err.Error(),
		})

		return
	}

	c.JSON(200, res)
}

func (h *Handler) GetResults(c *gin.Context) {
	var req requests.ResultsGet

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, responses.Error{
			Error: err.Error(),
		})

		return
	}

	res, err := h.Service.GetResults()
	if err != nil {
		c.JSON(500, responses.Error{
			Error: err.Error(),
		})

		return
	}

	c.JSON(200, res)
}
