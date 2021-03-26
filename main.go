package main

import (
	"log"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"gorm.io/driver/mysql"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"gitlab.com/investio/backend/api/db"
	"gitlab.com/investio/backend/api/v1/controller"
	"gitlab.com/investio/backend/api/v1/model"
	"gitlab.com/investio/backend/api/v1/service"
	"gopkg.in/olahol/melody.v1"
	"gorm.io/gorm"
)

var (
	thaiRegEx, _ = regexp.Compile("([\u0E00-\u0E7F]+)")

	fundService service.FundService = service.NewFundService(thaiRegEx)
	navService  service.NavService  = service.NewNavService()

	fundController controller.FundController = controller.NewFundController(fundService)
	navController  controller.NavController  = controller.NewNavController(navService)

	ws           *melody.Melody              = melody.New()
	wsController controller.SocketController = controller.NewSocketController(ws, fundController)
)

func setupDB() (err error) {
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
		log.Fatalln("Database Init error: ", err)
	} else {
		db.MySQL.AutoMigrate(&model.Fund{})

		db.InfluxClient = influxdb2.NewClient(
			"http://"+os.Getenv("INFLUX_HOST")+":"+os.Getenv("INFLUX_PORT"),
			os.Getenv("INFLUX_TOKEN"),
		)
		db.InfluxQuery = db.InfluxClient.QueryAPI(os.Getenv("INFLUX_ORG"))
	}
	return
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	err = setupDB()
	if err != nil {
		log.Fatal("Fail to connect to DB: ", err.Error())
		return
	}

	// defer db.MySQL.Close()
	defer db.InfluxClient.Close()

	// fmt.Printf("Type: %T", db.MySQL)

	server := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8080", "http://192.168.50.121:3003", "https://investio.dewkul.me", "https://investio.netlify.app"}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for VueJS
	corsConfig.AddAllowMethods("OPTIONS")

	// Register the middleware
	server.Use(cors.New(corsConfig))

	v1 := server.Group("/public/v1")
	{
		f := v1.Group("/funds")
		{
			f.GET("/info/:id", fundController.GetFundByID)
			f.GET("/cats", fundController.ListCat)
			f.GET("/amcs", fundController.ListAmc)
			f.GET("/nav/:id", navController.GetPastNavSeries)
			f.GET("/nav/:id/latest", navController.GetLatestNav)
			f.GET("/top/return", fundController.GetTopReturn)
			f.GET("/search/:fundQuery", fundController.SearchFund)
		}

		ws := v1.Group("/ws")
		{
			ws.GET(":clientID", wsController.HandleSocket)
		}
	}

	server.Run(":" + os.Getenv("API_PORT"))
}
