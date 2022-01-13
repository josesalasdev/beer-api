package main

import (
	"github.com/gin-gonic/gin"
	"github.com/josesalasdev/beer-api/cmd/api/clients"
	"github.com/josesalasdev/beer-api/cmd/api/config"
	"github.com/josesalasdev/beer-api/cmd/api/controllers"
	"github.com/josesalasdev/beer-api/cmd/api/docs"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT
func main() {
	r := gin.Default()

	DB := clients.NewConnectDataBase()

	currencyService := clients.NewCurrencyClient()
	beerController := controllers.NewBeerController(currencyService, DB)

	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Version = "v1"
	docs.SwaggerInfo.Host = "http://localhost"
	docs.SwaggerInfo.Title = "Beer API"

	// routes
	r.GET("/v1/ping", controllers.Ping)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiBeer := r.Group("/v1/beers")
	{
		apiBeer.POST("/", beerController.CreateBeer)
		apiBeer.GET("/", beerController.ListBeer)
		apiBeer.GET("/:id", beerController.RetrieveBeer)
		apiBeer.GET("/:id/boxprice", beerController.CalculateBeerBox)
	}

	r.Run(config.Port) // nolint
}
