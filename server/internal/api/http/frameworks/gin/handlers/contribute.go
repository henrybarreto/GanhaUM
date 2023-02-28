package handlers

import (
	"fmt"
	"ganhaum.henrybarreto.dev/internal/api/http/requests"
	"ganhaum.henrybarreto.dev/internal/api/http/responses"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
	"net/http"
	"time"
)

func (h *Handler) Contribute(c *gin.Context) {
	var req requests.Contribute
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, responses.NewError(err))

		return
	}

	campaign, err := h.Service.GetCampaign(req.CampaignID)
	if campaign == nil || err != nil {
		c.JSON(404, responses.NewError(err))

		return
	}

	if campaign.ClosedAt.Before(time.Now()) {
		c.JSON(400, responses.NewError(fmt.Errorf("campaign already closed")))

		return
	}

	domain := "http://localhost:3000"
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(""),
				Quantity: stripe.Int64(1),
			},
		},
		PhoneNumberCollection: &stripe.CheckoutSessionPhoneNumberCollectionParams{
			Enabled: stripe.Bool(true),
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "/?contribution=success"),
		CancelURL:  stripe.String(domain + "/?contribution=canceled"),
	}

	params.AddMetadata("campaign_id", string(req.CampaignID))

	checkout, err := session.New(params)
	if err != nil {
		c.JSON(500, responses.NewError(err))

		return
	}

	c.Redirect(http.StatusTemporaryRedirect, checkout.URL)

	return
}
