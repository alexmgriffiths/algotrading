package models

type Symbol struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Currency string `json:"currency"`
	Exchange string `json:"exchange"`
	MicCode  string `json:"mic_code"`
	Country  string `json:"country"`
	Type     string `json:"type"`
	FigiCode string `json:"figi_code"`
}
