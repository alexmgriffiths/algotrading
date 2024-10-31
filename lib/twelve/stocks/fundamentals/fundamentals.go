package fundamentals

import (
	"fmt"

	"github.com/alexmgriffiths/algotrading/lib/twelve/stocks"
	"github.com/alexmgriffiths/algotrading/lib/twelve/stocks/fundamentals/models"
)

type Fundamentals struct {
	Stocks *stocks.Stocks
}

func NewFundamentals(stocks *stocks.Stocks) *Fundamentals {
	return &Fundamentals{
		Stocks: stocks,
	}
}

func (f *Fundamentals) GetProfile(symbol string) (models.Profile, error) {
	resBody, err := f.Stocks.MakeRequest(fmt.Sprintf("profile?symbol=%s", symbol))
	if err != nil {
		return models.Profile{}, err
	}

	var profile models.Profile
	if err := f.Stocks.UnmarshalData(resBody, &profile); err != nil {
		return models.Profile{}, err
	}
	return profile, nil
}

func (f *Fundamentals) GetDividends(symbol string) (models.Dividends, error) {
	resBody, err := f.Stocks.MakeRequest(fmt.Sprintf("dividends?symbol=%s", symbol))
	if err != nil {
		return models.Dividends{}, err
	}

	var dividends models.Dividends
	if err := f.Stocks.UnmarshalData(resBody, &dividends); err != nil {
		return models.Dividends{}, err
	}
	return dividends, nil
}

func (f *Fundamentals) GetSplits(symbol string) (models.Splits, error) {
	resBody, err := f.Stocks.MakeRequest(fmt.Sprintf("splits?symbol=%s", symbol))
	if err != nil {
		return models.Splits{}, err
	}

	var splits models.Splits
	if err := f.Stocks.UnmarshalData(resBody, &splits); err != nil {
		return models.Splits{}, err
	}
	return splits, nil
}

func (f *Fundamentals) GetEarnings(symbol string) (models.Earnings, error) {
	resBody, err := f.Stocks.MakeRequest(fmt.Sprintf("earnings?symbol=%s", symbol))
	if err != nil {
		return models.Earnings{}, err
	}

	var earnings models.Earnings
	if err := f.Stocks.UnmarshalData(resBody, &earnings); err != nil {
		return models.Earnings{}, err
	}
	return earnings, nil
}
