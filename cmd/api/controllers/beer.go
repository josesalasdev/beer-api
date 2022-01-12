package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josesalasdev/beer-api/cmd/api/clients"
	"github.com/josesalasdev/beer-api/cmd/api/models"
)

// List beers controller.
// @Summary List beers.
// @Description You can list beers.
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.BeerItem
// @Router /v1/beers [get]
func ListBeer(c *gin.Context) {
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
func CreateBeer(c *gin.Context) {
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
func RetrieveBeer(c *gin.Context) {
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
// @Router /v1/beers/{beerID} [get]
func CalculateBeerBox(c *gin.Context) {
	var beer models.BeerItem
	if err := clients.DB.Where("id = ?", c.Param("id")).First(&beer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Beer not found!"})
		return
	}

	c.JSON(http.StatusOK, beer)
}
