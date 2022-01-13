package main

import (
	"github.com/gin-gonic/gin"
	"github.com/josesalasdev/beer-api/cmd/api/clients"
	"github.com/josesalasdev/beer-api/cmd/api/config"
	"github.com/josesalasdev/beer-api/cmd/api/controllers"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT
func main() {
	r := gin.Default()

	DB := clients.NewConnectDataBase()

	currencyService := clients.NewCurrencyClient()
	beerController := controllers.NewBeerController(currencyService, DB)

	// routes
	r.GET("/v1/ping", controllers.Ping)
	apiBeer := r.Group("/v1/beers")
	{
		apiBeer.POST("/", beerController.CreateBeer)
		apiBeer.GET("/", beerController.ListBeer)
		apiBeer.GET("/:id", beerController.RetrieveBeer)
		apiBeer.GET("/:id/boxprice", beerController.CalculateBeerBox)
	}

	r.Run(config.Port) // nolint
}
