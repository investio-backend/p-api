package model

type NavDate struct {
	Date string  `json:"date"`
	Nav  float64 `json:"nav"`
}

type NavSeries struct {
	FundCode string    `json:"fund_code"`
	Navs     []NavDate `json:"navs"`
}
