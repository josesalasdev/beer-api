package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/josesalasdev/beer-api/cmd/api/clients"
	"github.com/josesalasdev/beer-api/cmd/api/models"
)

type BeerController interface {
	ListBeer(c *gin.Context)
	CreateBeer(c *gin.Context)
	RetrieveBeer(c *gin.Context)
	CalculateBeerBox(c *gin.Context)
}

type beerController struct {
	currencyService clients.Currency
}

func NewBeerController(currencyService clients.Currency) BeerController {
	return &beerController{
		currencyService: currencyService,
	}
}

// List beers controller.
// @Summary List beers.
// @Description You can list beers.
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.BeerItem
// @Router /v1/beers [get]
func (cntl *beerController) ListBeer(c *gin.Context) {
	var beers []models.BeerItem
	clients.DB.Find(&beers)

	c.JSON(http.StatusOK, beers)
}

// Create beers controller
// @Summary Create beers.
// @Description You can create beer.
// @Accept  json
// @Produce  json
// @Param body body models.BeerItem true "Create beers"
// @Success 201 {object} models.BeerItem
// @Router /v1/beers [post]
func (cntl *beerController) CreateBeer(c *gin.Context) {
	// Validate input
	var beerItem models.BeerItem
	if err := c.ShouldBindJSON(&beerItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	clients.DB.Create(&beerItem)

	c.JSON(http.StatusCreated, beerItem)
}

// Retrieve beer controller.
// @Summary Retrieve a beer.
// @Description You can retrieve a beer.
// @Accept  json
// @Produce  json
// @Param beerID path int true "Beer ID"
// @Success 200 {object} models.BeerItem
// @Router /v1/beers/{beerID} [get]
func (cntl *beerController) RetrieveBeer(c *gin.Context) {
	var beer models.BeerItem
	if err := clients.DB.Where("id = ?", c.Param("id")).First(&beer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Beer not found!"})
		return
	}

	c.JSON(http.StatusOK, beer)
}

// Calculate Beer Box controller.
// @Summary Calculate Beer Box.
// @Description You can Calculate Beer Boxes.
// @Accept  json
// @Produce  json
// @Success 200 {object} models.BeerBox
// @Router /v1/beers/{beerID}/boxprice [get]
func (cntl *beerController) CalculateBeerBox(c *gin.Context) {
	var beer models.BeerItem
	if err := clients.DB.Where("id = ?", c.Param("id")).First(&beer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Beer not found!"})
		return
	}
	quantity, err := strconv.ParseFloat(c.Query("quantity"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "quantity param is required!"})
		return
	}
	currencyTo := c.Query("currency")

	total, err := cntl.currencyService.Calculate(beer.Price, quantity, beer.Currency, currencyTo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to calculate currencies!"})
		return
	}
	c.JSON(http.StatusOK, models.BeerBox{PriceTotal: total})
}
