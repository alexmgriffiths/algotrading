package stocks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/alexmgriffiths/algotrading/lib/twelve/stocks/models"
)

type Stocks struct {
	client  *http.Client
	baseURL string
	apiKey  string
}

func NewStocks() *Stocks {
	return &Stocks{
		client:  http.DefaultClient,
		baseURL: "https://api.twelvedata.com",
		apiKey:  os.Getenv("TWELVE_API_KEY"),
	}
}

func (t *Stocks) MakeRequest(endpoint string) ([]byte, error) {
	req, reqErr := http.NewRequest("GET", fmt.Sprintf("%s/%s", t.baseURL, endpoint), nil)
	if reqErr != nil {
		return nil, fmt.Errorf("failed to create request: %w", reqErr)
	}

	req.Header.Set("Authorization", fmt.Sprintf("apikey %s", t.apiKey))
	res, resErr := t.client.Do(req)
	if resErr != nil {
		return nil, fmt.Errorf("request failed: %w", resErr)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code %d", res.StatusCode)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, fmt.Errorf("failed to read response body: %w", readErr)
	}
	return body, nil
}

func (t *Stocks) UnmarshalData(data []byte, v interface{}) error {
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("failed to parse data: %w", err)
	}
	return nil
}

func (t *Stocks) GetStocks() ([]models.Symbol, error) {
	resBody, err := t.MakeRequest("stocks")
	if err != nil {
		return nil, err
	}

	var stocks []models.Symbol
	if err := t.UnmarshalData(resBody, &stocks); err != nil {
		return nil, err
	}
	return stocks, nil
}

func (t *Stocks) GetStockPrice(symbol string) (models.Price, error) {
	resBody, err := t.MakeRequest(fmt.Sprintf("price?symbol=%s", symbol))
	if err != nil {
		return models.Price{}, err
	}

	var price models.Price
	if err := t.UnmarshalData(resBody, &price); err != nil {
		return models.Price{}, err
	}
	return price, nil
}

// Only available on pro and enterprise -AG
func (t *Stocks) GetMarketMovers() (models.MarketMovers, error) {
	resBody, err := t.MakeRequest("market_movers/stocks")
	if err != nil {
		return models.MarketMovers{}, err
	}

	var marketMovers models.MarketMovers
	if err := t.UnmarshalData(resBody, &marketMovers); err != nil {
		return models.MarketMovers{}, err
	}
	return marketMovers, nil
}

func (t *Stocks) GetEOD(symbol string) (models.EOD, error) {
	resBody, err := t.MakeRequest(fmt.Sprintf("eod?symbol=%s", symbol))
	if err != nil {
		return models.EOD{}, err
	}

	var eod models.EOD
	if err := t.UnmarshalData(resBody, &eod); err != nil {
		return models.EOD{}, err
	}
	return eod, nil
}

func (t *Stocks) GetTimeSeries(symbol string, interval string) (models.TimeSeries, error) {
	resBody, err := t.MakeRequest(fmt.Sprintf("time_series?symbol=%s:NASDAQ&interval=%s", symbol, interval))
	if err != nil {
		return models.TimeSeries{}, err
	}

	var timeSeries models.TimeSeries
	if err := t.UnmarshalData(resBody, &timeSeries); err != nil {
		return models.TimeSeries{}, err
	}
	return timeSeries, nil
}
