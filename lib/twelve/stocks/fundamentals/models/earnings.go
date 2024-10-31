package models

type Earnings struct {
	Meta struct {
		Symbol           string `json:"symbol"`
		Name             string `json:"name"`
		Currency         string `json:"currency"`
		Exchange         string `json:"exchange"`
		MicCode          string `json:"mic_code"`
		ExchangeTimezone string `json:"exchange_timezone"`
	} `json:"meta"`
	Earnings []struct {
		Date        string  `json:"date"`
		Time        string  `json:"time"`
		EpsEstimate float64 `json:"eps_estimate"`
		EpsActual   float64 `json:"eps_actual"`
		Difference  float64 `json:"difference"`
		SurprisePrc float64 `json:"surprise_prc"`
	} `json:"earnings"`
	Status string `json:"status"`
}
