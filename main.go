package main

import (
	"log"
	"os"
	"regexp"
	_ "time/tzdata"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"gitlab.com/investio/backend/api/db"
	"gitlab.com/investio/backend/api/v1/controller"
	"gitlab.com/investio/backend/api/v1/service"
)

var (
	thaiRegEx, _ = regexp.Compile("([\u0E00-\u0E7F]+)")

	fundService service.FundService = service.NewFundService(thaiRegEx)
	navService  service.NavService  = service.NewNavService()
	statService service.StatService = service.NewStatService()

	fundController controller.FundController = controller.NewFundController(fundService)
	navController  controller.NavController  = controller.NewNavController(navService)
	statController controller.StatController = controller.NewStatController(statService)

	// ws           *melody.Melody              = melody.New()
	// wsController controller.SocketController = controller.NewSocketController(ws, fundController)
)

func main() {
	var err error
	if os.Getenv("GIN_MODE") != "release" {
		err = godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
			return
		}
	}

	err = db.SetupDB()
	if err != nil {
		log.Fatal("Fail to connect to DB: ", err.Error())
		return
	}

	defer db.InfluxClient.Close()

	server := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:2564", "http://192.168.50.121:3003", "https://investio.dewkul.me", "https://investio.netlify.app"}
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
			f.GET("/nav/series", navController.GetPastNav)
			f.GET("/nav", navController.GetNavByDate)
			f.GET("/search/:fundQuery", fundController.SearchFund)
			f.GET("/top/return", statController.GetTopReturn)
			f.GET("/stat/:fundID", statController.GetStatInfo)
		}

		// ws := v1.Group("/ws")
		// {
		// 	ws.GET(":clientID", wsController.HandleSocket)
		// }
	}

	internalV1 := server.Group("/s2gsoeq93f/v1")
	{
		internalV1.GET("/index/set", navController.GetPastSetIndex)
		internalV1.GET("/nav/series/ast", navController.GetPastNavWithAsset)
	}

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "5005"
	}
	server.Run(":" + port)
}
