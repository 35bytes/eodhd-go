package eodhd

import (
	"github.com/karlbehrensg/eodhd-go/eod"
	"github.com/karlbehrensg/eodhd-go/exchanges"
	"github.com/karlbehrensg/eodhd-go/fundamentals"
)

// baseUrl is the base URL for the EODHD API.
const (
	baseUrl = "https://eodhd.com/api"
)

// Client represents the EODHD client.
type Client struct {
	apiKey       string
	Exchanges    exchanges.ExchangeClient
	Fundamentals fundamentals.FundamentalsClient
	Eod          eod.EodClient
}

// NewEodClient creates a new instance of the EODHD client.
func NewEodClient(apiKey string) *Client {
	return &Client{
		apiKey:       apiKey,
		Exchanges:    exchanges.NewExchangeClient(baseUrl, apiKey),
		Fundamentals: fundamentals.NewFundamentalsClient(baseUrl, apiKey),
		Eod:          eod.NewEodClient(baseUrl, apiKey),
	}
}
