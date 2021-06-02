package tests

import (
	"encoding/json"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/joho/godotenv"
	"gitlab.com/investio/backend/api/db"
	"gitlab.com/investio/backend/api/v1/controller"
	"gitlab.com/investio/backend/api/v1/model"
	"gitlab.com/investio/backend/api/v1/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// func init() (err error) {

// }

func setupDB(t *testing.T) (err error) {
	db.MySQL, err = gorm.Open(
		mysql.Open(
			db.MySqlURL(db.BuildDbConfig(
				os.Getenv("MYSQL_HOST"),
				os.Getenv("MYSQL_PORT"),
				os.Getenv("MYSQL_USER"),
				os.Getenv("MYSQL_PWD"),
				os.Getenv("MYSQL_DB"),
			)),
		),
		&gorm.Config{},
	)

	if err != nil {
		t.Fatal("Database Init error: ", err)
	} else {
		db.MySQL.AutoMigrate(&model.Fund{})

		db.InfluxClient = influxdb2.NewClient(
			os.Getenv("INFLUX_HOST"),
			os.Getenv("INFLUX_TOKEN"),
		)
		db.InfluxQuery = db.InfluxClient.QueryAPI(os.Getenv("INFLUX_ORG"))
	}
	return
}

var (
	thaiRegEx, _ = regexp.Compile("([\u0E00-\u0E7F]+)")

	fundService    service.FundService       = service.NewFundService(thaiRegEx)
	fundController controller.FundController = controller.NewFundController(fundService)
)

func TestGetCatList(t *testing.T) {
	var err = godotenv.Load("../.env")
	if err != nil {
		t.Fatal("Error loading .env file ", err)
		return
	}

	err = setupDB(t)
	if err != nil {
		t.Fatal("Fail to connect to DB: ", err.Error())
		return
	}

	defer db.InfluxClient.Close()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	fundController.ListCat(c)
	assert.Equal(t, 200, w.Code)

	// var recieved gin.H
	var recieved []interface{}
	err = json.Unmarshal(w.Body.Bytes(), &recieved)
	if err != nil {
		t.Fatal(err)
	}
	// log.Print(recieved)
	// assert
}

func TestGetAmc(t *testing.T) {
	var err = godotenv.Load("../.env")
	if err != nil {
		t.Fatal("Error loading .env file ", err)
	}

	err = setupDB(t)
	if err != nil {
		t.Fatal("Fail to connect to DB: ", err.Error())
	}

	defer db.InfluxClient.Close()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	fundController.ListAmc(c)
	assert.Equal(t, 200, w.Code)

	// var recieved gin.H
	var recieved []interface{}
	err = json.Unmarshal(w.Body.Bytes(), &recieved)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetFund(t *testing.T) {

	if err := godotenv.Load("../.env"); err != nil {
		t.Fatal("Error loading .env file ", err)
	}

	if err := setupDB(t); err != nil {
		t.Fatal("Fail to connect to DB: ", err.Error())
	}

	defer db.InfluxClient.Close()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	// id := "970ed651"
	fundController.GetFundByID(c)
}
