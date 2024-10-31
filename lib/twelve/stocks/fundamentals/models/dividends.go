package models

type Dividends struct {
	Meta struct {
		Symbol           string `json:"symbol"`
		Name             string `json:"name"`
		Currency         string `json:"currency"`
		Exchange         string `json:"exchange"`
		MicCode          string `json:"mic_code"`
		ExchangeTimezone string `json:"exchange_timezone"`
	} `json:"meta"`
	Dividends []struct {
		ExDate string  `json:"ex_date"`
		Amount float64 `json:"amount"`
	} `json:"dividends"`
}
