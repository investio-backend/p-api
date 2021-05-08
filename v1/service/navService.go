package service

import (
	"context"

	"gitlab.com/investio/backend/api/db"
	"gitlab.com/investio/backend/api/v1/model"
)

type NavService interface {
	FindPastNavWithAsset(navList *[]model.NavDate, fundID, dataRange string) (err error)
	FindPastNav(nav *[]float64, date *[]string, fundID, dataRange string) error
	FindLatestNavByFundID(navList *model.NavDate, fundID string) error
	FindNAVByDate(navList *model.NavDate, dataDate model.Date, fundID string) error
}

type navService struct {
	bucket string
}

func NewNavService() NavService {
	return &navService{
		bucket: "DailyFund",
	}
}

func (s *navService) FindPastNavWithAsset(navList *[]model.NavDate, fundID, dataRange string) (err error) {
	navs := make(map[string]*model.NavDate)

	result, err := db.InfluxQuery.Query(
		context.Background(),
		`from(bucket: "`+s.bucket+`")
		|> range(start: -`+dataRange+`)
		|> filter(fn: (r) => r._measurement == "PastNAV")
		|> filter(fn: (r) => r._field == "nav" or r._field == "asset_amount")
		|> filter(fn: (r) => r.fund_id == "`+fundID+`")
	`)

	if err != nil {
		return err
	}

	var (
		isNav    bool
		dateList []string
	)
	// Iterate over query response
	for result.Next() {
		// Notice when group key has changed
		if result.TableChanged() {
			if result.Record().Field() == "nav" {
				isNav = true
			}
		}

		// Access data
		if isNav {
			nav := result.Record().Value().(float64)
			date := result.Record().Time().Format("2006-01-02")
			navs[date].Nav = nav
		} else {
			asset := result.Record().Value().(int64)
			date := result.Record().Time().Format("2006-01-02")

			navDate := &model.NavDate{
				Date: date, Asset: asset,
			}
			navs[date] = navDate
			dateList = append(dateList, date)
			// *navList = append(*navList, navDate)
		}
	}
	// check for an error
	if result.Err() != nil {
		// fmt.Printf("query parsing error: %s\n", result.Err().Error())
		err = result.Err()
		return
	}

	// Slice nav value from a map
	for _, k := range dateList {
		*navList = append(*navList, *navs[k])
	}
	return
}

func (s *navService) FindPastNav(navList *[]float64, dateList *[]string, fundID, dataRange string) error {
	result, err := db.InfluxQuery.Query(
		context.Background(),
		`from(bucket:"`+s.bucket+`")
		|> range(start: -`+dataRange+`)
		|> filter(fn: (r) => r._measurement == "PastNAV")
		|> filter(fn: (r) => r._field == "nav"
			and r.fund_id == "`+fundID+`")`)

	if err != nil {
		return err
	}

	// Iterate over query response
	for result.Next() {
		nav := result.Record().Value().(float64)
		date := result.Record().Time().Format("2006-01-02")

		*navList = append(*navList, nav)
		*dateList = append(*dateList, date)
	}
	// check for an error
	if result.Err() != nil {
		// fmt.Printf("query parsing error: %s\n", result.Err().Error())
		return result.Err()
	}
	return nil
}

func (s *navService) FindLatestNavByFundID(navList *model.NavDate, fundID string) error {
	var (
		nav   float64
		date  string
		asset int64
	)
	result, err := db.InfluxQuery.Query(
		context.Background(),
		`from(bucket:"`+s.bucket+`")
		|> range(start: -2mo)
		|> filter(fn: (r) => r._field == "nav" or r._field == "asset_amount")
		|> filter(fn: (r) => r.fund_id == "`+fundID+`")
		|> last(column: "_time")`)

	if err != nil {
		return err
	}

	for result.Next() {
		if result.Record().Field() == "nav" {
			nav = result.Record().Value().(float64)
			date = result.Record().Time().Format("2006-01-02")
		} else {
			asset = result.Record().Value().(int64)
		}
	}

	*navList = model.NavDate{
		Date: date, Nav: nav, Asset: asset,
	}

	return nil
}

func (s *navService) FindNAVByDate(navList *model.NavDate, dataDate model.Date, fundID string) error {
	var (
		nav   float64
		date  string
		asset int64
	)
	result, err := db.InfluxQuery.Query(
		context.Background(),
		`from(bucket:"`+s.bucket+`")
		|> range(start: `+dataDate.String()+`T00:00:00Z, stop: `+dataDate.String()+`T23:59:59Z)
		|> filter(fn: (r) => r._field == "nav" or r._field == "asset_amount")
		|> filter(fn: (r) => r.fund_id == "`+fundID+`")
	`)

	if err != nil {
		return err
	}

	for result.Next() {
		if result.Record().Field() == "nav" {
			nav = result.Record().Value().(float64)
			date = result.Record().Time().Format("2006-01-02")
		} else {
			asset = result.Record().Value().(int64)
		}
	}

	*navList = model.NavDate{
		Date: date, Nav: nav, Asset: asset,
	}

	return nil
}
