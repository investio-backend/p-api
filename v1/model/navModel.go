package model

type NavDate struct {
	Date string  `json:"date"`
	Nav  float64 `json:"nav"`
}

type NavSeries struct {
	FundID int32     `json:"fund_id"`
	Navs   []NavDate `json:"navs"`
}
