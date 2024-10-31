package models

type Profile struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Exchange    string `json:"exchange"`
	MicCode     string `json:"mic_code"`
	Sector      string `json:"sector"`
	Industry    string `json:"industry"`
	Employees   int    `json:"employees"`
	Website     string `json:"website"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Ceo         string `json:"CEO"`
	Address     string `json:"address"`
	Address2    string `json:"address2"`
	City        string `json:"city"`
	Zip         string `json:"zip"`
	State       string `json:"state"`
	Country     string `json:"country"`
	Phone       string `json:"phone"`
}
