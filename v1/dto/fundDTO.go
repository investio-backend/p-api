package dto

type QueryStrStat struct {
	Cat   string `form:"cat" json:"cat"`
	Amc   string `form:"amc" json:"amc"`
	Range string `form:"range" json:"range"`
}
