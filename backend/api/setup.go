package api

import (
	"database/sql"
	"log" // Import the log package
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/minirisk/models"
)

// SetupRoutes configures all API routes
func SetupRoutes(router *gin.Engine) {
	// Market data endpoints
	marketDataGroup := router.Group("/api/market-data")
	{
		marketDataGroup.GET("/:symbol", GetMarketData)
		marketDataGroup.POST("/", UpdateMarketData)
	}

	// Position endpoints
	positionGroup := router.Group("/api/positions")
	{
		positionGroup.GET("/:clientId", GetPositions)
		positionGroup.POST("/", CreatePosition)
		positionGroup.PUT("/:id", UpdatePosition)
		positionGroup.DELETE("/:id", DeletePosition)
	}

	// Margin endpoints
	marginGroup := router.Group("/api/margin")
	{
		marginGroup.GET("/status/:clientId", GetMarginStatus)
		marginGroup.POST("/", UpdateMargin)
	}
}

// GetMarketData retrieves current market data for a symbol
func GetMarketData(c *gin.Context) {
	symbol := c.Param("symbol")
	db := c.MustGet("db").(*sql.DB)
	marketDataService := &models.MarketDataService{DB: db}

	marketData, err := marketDataService.GetCurrentPrice(symbol)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve market data"})
		return
	}
	if marketData == nil {
		c.JSON(404, gin.H{"error": "Market data not found"})
		return
	}

	c.JSON(200, marketData)
}

// UpdateMarketData updates market data for symbols
func UpdateMarketData(c *gin.Context) {
	var marketData models.MarketData
	if err := c.ShouldBindJSON(&marketData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	db := c.MustGet("db").(*sql.DB)
	marketDataService := &models.MarketDataService{DB: db}
	if err := marketDataService.UpdateMarketData(&marketData); err != nil {
		c.JSON(500, gin.H{"error": "Failed to update market data"})
		return
	}

	c.JSON(200, gin.H{"message": "Market data updated successfully"})
}

// GetPositions retrieves all positions for a client
func GetPositions(c *gin.Context) {
	clientIDStr := c.Param("clientId")
	clientID, err := strconv.ParseInt(clientIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid client ID"})
		return
	}

	db := c.MustGet("db").(*sql.DB)
	positionService := &models.PositionService{DB: db}

	positions, err := positionService.GetPositionsByClientID(clientID)
	if err != nil {
		log.Printf("Error retrieving positions for client %d: %v", clientID, err) // Log the specific error
		c.JSON(500, gin.H{"error": "Failed to retrieve positions"})
		return
	}

	c.JSON(200, positions)
}

// CreatePosition creates a new position
func CreatePosition(c *gin.Context) {
	var position models.Position
	if err := c.ShouldBindJSON(&position); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	db := c.MustGet("db").(*sql.DB)
	positionService := &models.PositionService{DB: db}
	if err := positionService.CreatePosition(&position); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create position"})
		return
	}

	c.JSON(201, position)
}

// UpdatePosition updates an existing position
func UpdatePosition(c *gin.Context) {
	var position models.Position
	if err := c.ShouldBindJSON(&position); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	db := c.MustGet("db").(*sql.DB)
	positionService := &models.PositionService{DB: db}
	if err := positionService.UpdatePosition(&position); err != nil {
		c.JSON(500, gin.H{"error": "Failed to update position"})
		return
	}

	c.JSON(200, position)
}

// DeletePosition deletes a position
func DeletePosition(c *gin.Context) {
	positionIDStr := c.Param("id")
	clientIDStr := c.Query("clientId")

	positionID, err := strconv.ParseInt(positionIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid position ID"})
		return
	}

	clientID, err := strconv.ParseInt(clientIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid client ID"})
		return
	}

	db := c.MustGet("db").(*sql.DB)
	positionService := &models.PositionService{DB: db}
	if err := positionService.DeletePosition(positionID, clientID); err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete position"})
		return
	}

	c.JSON(200, gin.H{"message": "Position deleted successfully"})
}

// GetMarginStatus retrieves the current margin status for a client
func GetMarginStatus(c *gin.Context) {
	clientIDStr := c.Param("clientId")
	clientID, err := strconv.ParseInt(clientIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid client ID"})
		return
	}

	db := c.MustGet("db").(*sql.DB)

	// Get positions
	positionService := &models.PositionService{DB: db}
	positions, err := positionService.GetPositionsByClientID(clientID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve positions"})
		return
	}

	// Get market prices for positions
	var symbols []string
	for _, position := range positions {
		symbols = append(symbols, position.Symbol)
	}

	marketDataService := &models.MarketDataService{DB: db}
	marketPrices, err := marketDataService.GetMarketDataForSymbols(symbols)
	if err != nil {
		log.Printf("Error retrieving market data for symbols %v: %v", symbols, err) // Log the specific error
		c.JSON(500, gin.H{"error": "Failed to retrieve market data"})
		return
	}

	// Calculate margin status
	marginService := &models.MarginService{DB: db}
	marginStatus, err := marginService.CalculateMarginStatus(clientID, positions, marketPrices)
	if err != nil {
		log.Printf("Error calculating margin status for client %d: %v", clientID, err) // Log the specific error
		c.JSON(500, gin.H{"error": "Failed to calculate margin status"})
		return
	}

	c.JSON(200, marginStatus)
}

// UpdateMargin updates margin data for a client
func UpdateMargin(c *gin.Context) {
	var margin models.Margin
	if err := c.ShouldBindJSON(&margin); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	db := c.MustGet("db").(*sql.DB)
	marginService := &models.MarginService{DB: db}
	if err := marginService.UpdateMargin(&margin); err != nil {
		c.JSON(500, gin.H{"error": "Failed to update margin data"})
		return
	}

	c.JSON(200, gin.H{"message": "Margin data updated successfully"})
}
