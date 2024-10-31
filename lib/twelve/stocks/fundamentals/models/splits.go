package models

type Splits struct {
	Meta struct {
		Symbol           string `json:"symbol"`
		Name             string `json:"name"`
		Currency         string `json:"currency"`
		Exchange         string `json:"exchange"`
		MicCode          string `json:"mic_code"`
		ExchangeTimezone string `json:"exchange_timezone"`
	} `json:"meta"`
	Splits []struct {
		Date        string  `json:"date"`
		Description string  `json:"description"`
		Ratio       float64 `json:"ratio"`
		FromFactor  int     `json:"from_factor"`
		ToFactor    int     `json:"to_factor"`
	} `json:"splits"`
}
