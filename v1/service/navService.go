package service

import (
	// influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"context"
	"fmt"

	"gitlab.com/investio/backend/api/db"
	"gitlab.com/investio/backend/api/v1/model"
)

type NavService interface {
	GetPastNavByFundID(navList *[]model.NavDate, fundID, dataRange string) error
	QueryLatestNavByFundID(navList *model.NavDate, fundID string) error
}

type navService struct {
	//
}

func NewNavService() NavService {
	return &navService{}
}

func (s *navService) GetPastNavByFundCode(navList *[]model.NavDate, fundCode, dataRange string) error {
	result, err := db.InfluxQuery.Query(
		context.Background(),
		`from(bucket:"fund-3Y")
		|> range(start: -`+dataRange+`)
		|> filter(fn: (r) => r._field == "nav"
			and r.fund_code == "`+fundCode+`")`)

	if err != nil {
		return err
	}

	// Iterate over query response
	for result.Next() {
		// // Notice when group key has changed
		// if result.TableChanged() {
		// 	fmt.Printf("table: %s\n", result.TableMetadata().String())
		// }
		// Access data
		nav := result.Record().Value().(float64)
		date := result.Record().Time().Format("2006-01-02")

		navDate := model.NavDate{
			Date: date, Nav: nav,
		}
		*navList = append(*navList, navDate)
	}
	// check for an error
	if result.Err() != nil {
		fmt.Printf("query parsing error: %s\n", result.Err().Error())
		return result.Err()
	}
	return nil
}

func (s *navService) GetPastNavByFundID(navList *[]model.NavDate, fundID, dataRange string) error {
	result, err := db.InfluxQuery.Query(
		context.Background(),
		`from(bucket:"fund-3Y")
		|> range(start: -`+dataRange+`)
		|> filter(fn: (r) => r._field == "nav"
			and r.fund_id == "`+fundID+`")`)

	if err != nil {
		return err
	}

	// Iterate over query response
	for result.Next() {
		// Access data
		nav := result.Record().Value().(float64)
		date := result.Record().Time().Format("2006-01-02")

		navDate := model.NavDate{
			Date: date, Nav: nav,
		}
		*navList = append(*navList, navDate)
	}
	// check for an error
	if result.Err() != nil {
		fmt.Printf("query parsing error: %s\n", result.Err().Error())
		return result.Err()
	}
	return nil
}

func (s *navService) QueryLatestNavByFundID(navList *model.NavDate, fundID string) error {
	result, err := db.InfluxQuery.Query(
		context.Background(),
		`from(bucket:"fund-3Y")
		|> range(start: -2w)
		|> filter(fn: (r) => r._field == "nav"
			and r.fund_id == "`+fundID+`")
		|> last(column: "_time")`)

	if err != nil {
		return err
	}

	for result.Next() {
		nav := result.Record().Value().(float64)
		date := result.Record().Time().Format("2006-01-02")
		*navList = model.NavDate{
			Date: date, Nav: nav,
		}
	}

	return nil
}
