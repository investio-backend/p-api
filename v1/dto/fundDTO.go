package dto

type QueryStrStat struct {
	Cat   string `form:"cat" json:"cat"`
	Range string `form:"range" json:"range"`
}
