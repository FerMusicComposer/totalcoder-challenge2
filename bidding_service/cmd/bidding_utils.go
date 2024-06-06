package main

import (
	"fmt"
	"math/rand"
)

func generateBidPrice() string {
	rawPrice := rand.Float64() * 1000.0
	return fmt.Sprintf("%.2f$", rawPrice)
}

func generateAdId() string {
	numBody := rand.Intn(100000)

	return fmt.Sprintf("ad-%d", numBody)
}
