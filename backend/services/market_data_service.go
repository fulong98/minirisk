package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/minirisk/models"
)

// MarketDataUpdater handles fetching and updating market data
type MarketDataUpdater struct {
	DB             *sql.DB
	APIKey         string
	APIURL         string
	UpdateInterval time.Duration
}

// NewMarketDataUpdater creates a new MarketDataUpdater instance
func NewMarketDataUpdater(db *sql.DB) *MarketDataUpdater {
	return &MarketDataUpdater{
		DB:             db,
		APIKey:         os.Getenv("MARKET_DATA_API_KEY"),
		APIURL:         os.Getenv("MARKET_DATA_API_URL"),
		UpdateInterval: time.Duration(getEnvInt("MARKET_DATA_UPDATE_INTERVAL", 60)) * time.Second,
	}
}

// Start begins the market data update process
func (mdu *MarketDataUpdater) Start() {
	ticker := time.NewTicker(mdu.UpdateInterval)
	go func() {
		for range ticker.C {
			if err := mdu.UpdateMarketData(); err != nil {
				fmt.Printf("Error updating market data: %v\n", err)
			}
		}
	}()
}

// UpdateMarketData fetches and updates market data for all tracked symbols
func (mdu *MarketDataUpdater) UpdateMarketData() error {
	// Get all unique symbols from positions
	symbols, err := mdu.getTrackedSymbols()
	if err != nil {
		return fmt.Errorf("failed to get tracked symbols: %v", err)
	}

	// Fetch current prices for all symbols
	prices, err := mdu.fetchMarketPrices(symbols)
	if err != nil {
		return fmt.Errorf("failed to fetch market prices: %v", err)
	}

	// Update market data in database
	marketDataService := &models.MarketDataService{DB: mdu.DB}
	for symbol, price := range prices {
		marketData := &models.MarketData{
			Symbol:       symbol,
			CurrentPrice: price,
		}
		if err := marketDataService.UpdateMarketData(marketData); err != nil {
			fmt.Printf("Failed to update market data for %s: %v\n", symbol, err)
		}
	}

	return nil
}

// getTrackedSymbols retrieves all unique symbols from positions
func (mdu *MarketDataUpdater) getTrackedSymbols() ([]string, error) {
	query := "SELECT DISTINCT symbol FROM positions"
	rows, err := mdu.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var symbols []string
	for rows.Next() {
		var symbol string
		if err := rows.Scan(&symbol); err != nil {
			return nil, err
		}
		symbols = append(symbols, symbol)
	}

	return symbols, nil
}

// fetchMarketPrices retrieves current prices from the market data API
func (mdu *MarketDataUpdater) fetchMarketPrices(symbols []string) (map[string]float64, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	prices := make(map[string]float64)

	for _, symbol := range symbols {
		url := fmt.Sprintf("%s/quote/%s?apikey=%s", mdu.APIURL, symbol, mdu.APIKey)
		resp, err := client.Get(url)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch price for %s: %v", symbol, err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response for %s: %v", symbol, err)
		}

		var result struct {
			Price float64 `json:"price"`
		}
		if err := json.Unmarshal(body, &result); err != nil {
			return nil, fmt.Errorf("failed to parse response for %s: %v", symbol, err)
		}

		prices[symbol] = result.Price
	}

	return prices, nil
}

// getEnvInt gets an environment variable as an integer with a default value
func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	var result int
	fmt.Sscanf(value, "%d", &result)
	return result
}
