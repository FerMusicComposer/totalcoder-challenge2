package main

import (
	"github.com/labstack/echo/v4"
)

const port = ":4200"

func main() {
	srv := echo.New()

	srv.POST("/bid", handleBidRequest)

	srv.Logger.Fatal(srv.Start(port))
}
