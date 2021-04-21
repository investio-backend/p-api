package model

type NavDate struct {
	Date  string  `json:"date"`
	Nav   float64 `json:"nav"`
	Asset int64   `json:"asset"`
}

type NavSeries struct {
	FundID   string    `json:"fund_id,omitempty"`
	FundCode string    `json:"fund_code,omitempty"`
	Navs     []NavDate `json:"navs"`
}
