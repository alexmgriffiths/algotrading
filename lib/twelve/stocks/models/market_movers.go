package models

type MarketMovers struct {
	Values []struct {
		Symbol        string  `json:"symbol"`
		Name          string  `json:"name"`
		Exchange      string  `json:"exchange"`
		MicCode       string  `json:"mic_code"`
		Datetime      string  `json:"datetime"`
		Last          float64 `json:"last"`
		High          float64 `json:"high"`
		Low           float64 `json:"low"`
		Volume        int     `json:"volume"`
		Change        float64 `json:"change"`
		PercentChange float64 `json:"percent_change"`
	} `json:"values"`
	Status string `json:"status"`
}
