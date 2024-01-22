package eod

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

// EodClient represents a client for accessing EOD data.
type EodClient struct {
	baseUrl string
	apiKey  string
}

// NewEodClient creates a new instance of EodClient.
func NewEodClient(baseUrl, apiKey string) EodClient {
	return EodClient{
		baseUrl: baseUrl,
		apiKey:  apiKey,
	}
}

func (c *EodClient) GetHistoricalPrice(symbol, exchange, start, end string) (price *Price, err error) {
	if usExchange[exchange] {
		exchange = "US"
	}

	url := fmt.Sprintf("%s/eod/%s.%s?api_token=%s&from=%s&to=%s&fmt=json", c.baseUrl, symbol, exchange, c.apiKey, start, end)
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

	var result *Price

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
