package exchanges

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ExchangeClient represents a client for accessing exchange data.
type ExchangeClient struct {
	baseUrl string
	apiKey  string
}

// NewExchangeClient creates a new instance of ExchangeClient.
func NewExchangeClient(baseUrl, apiKey string) ExchangeClient {
	return ExchangeClient{
		baseUrl: baseUrl,
		apiKey:  apiKey,
	}
}

// GetExchanges retrieves a list of exchanges.
func (c *ExchangeClient) GetExchanges() (exchanges []*Exchange, err error) {
	url := fmt.Sprintf("%s/exchanges-list?api_token=%s&fmt=json", c.baseUrl, c.apiKey)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var result []*Exchange

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetSymbols retrieves a list of symbols for a specific exchange.
func (c *ExchangeClient) GetSymbols(exchange string) (symbols []*Symbol, err error) {
	url := fmt.Sprintf("%s/exchange-symbol-list/%s?api_token=%s&fmt=json", c.baseUrl, exchange, c.apiKey)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		if response.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("exchange %s not found", exchange)
		}
		return nil, fmt.Errorf("unexpected status code %d", response.StatusCode)
	}

	var result []*Symbol

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
