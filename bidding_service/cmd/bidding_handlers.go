package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdRequest struct {
	AdPlacementId string `json:"adPlacementId"`
}

type AdObject struct {
	AdId     string `json:"adId"`
	BidPrice string `json:"bidPrice"`
}

func handleBidRequest(c echo.Context) error {
	req := new(AdRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Input")
	}

	if rand.Float64() < 0.5 {
		fmt.Println("not bidding")
		return c.NoContent(http.StatusNoContent)
	}

	bidPrice := generateBidPrice()
	adObject := AdObject{
		AdId:     generateAdId(),
		BidPrice: bidPrice,
	}

	return c.JSON(http.StatusOK, adObject)

}
