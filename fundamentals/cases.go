// Package fundamentals provides a client for accessing fundamentals data.

package fundamentals

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	usExchange = map[string]bool{
		"NYSE":   true,
		"NASDAQ": true,
		"AMEX":   true,
	}
)

// FundamentalsClient represents a client for accessing fundamentals data.
type FundamentalsClient struct {
	baseUrl string
	apiKey  string
}

// NewFundamentalsClient creates a new instance of FundamentalsClient.
func NewFundamentalsClient(baseUrl, apiKey string) FundamentalsClient {
	return FundamentalsClient{
		baseUrl: baseUrl,
		apiKey:  apiKey,
	}
}

// GetEtf retrieves ETF fundamentals data for the specified symbol and exchange.
func (c *FundamentalsClient) GetEtf(symbol, exchange string) (etf *Etf, err error) {
	if usExchange[exchange] {
		exchange = "US"
	}

	url := fmt.Sprintf("%s/fundamentals/%s.%s?api_token=%s&fmt=json", c.baseUrl, symbol, exchange, c.apiKey)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		if response.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("symbol %s in exchange %s not found", symbol, exchange)
		}
		return nil, fmt.Errorf("unexpected status code %d", response.StatusCode)
	}

	var result *Etf

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	if result.General.Type != "ETF" {
		return nil, fmt.Errorf("symbol %s in exchange %s is not an ETF", symbol, exchange)
	}

	return result, nil
}
