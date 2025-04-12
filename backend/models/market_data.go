package models

import (
	"database/sql"
	"time"
)

// MarketData represents real-time market data for a symbol
type MarketData struct {
	ID           int64     `json:"id"`
	Symbol       string    `json:"symbol"`
	CurrentPrice float64   `json:"current_price"`
	Timestamp    time.Time `json:"timestamp"`
}

// MarketDataService handles database operations for market data
type MarketDataService struct {
	DB *sql.DB
}

// GetCurrentPrice retrieves the current price for a symbol
func (mds *MarketDataService) GetCurrentPrice(symbol string) (*MarketData, error) {
	query := `
		SELECT id, symbol, current_price, timestamp
		FROM market_data
		WHERE symbol = ?
		ORDER BY timestamp DESC
		LIMIT 1
	`

	var md MarketData
	err := mds.DB.QueryRow(query, symbol).Scan(
		&md.ID,
		&md.Symbol,
		&md.CurrentPrice,
		&md.Timestamp,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &md, nil
}

// UpdateMarketData updates or inserts market data for a symbol
func (mds *MarketDataService) UpdateMarketData(md *MarketData) error {
	query := `
		INSERT INTO market_data (symbol, current_price, timestamp)
		VALUES (?, ?, NOW())
		ON DUPLICATE KEY UPDATE
		current_price = VALUES(current_price),
		timestamp = VALUES(timestamp)
	`

	_, err := mds.DB.Exec(query, md.Symbol, md.CurrentPrice)
	return err
}

// GetMarketDataForSymbols retrieves current prices for multiple symbols
func (mds *MarketDataService) GetMarketDataForSymbols(symbols []string) (map[string]float64, error) {
	query := `
		SELECT symbol, current_price
		FROM market_data
		WHERE symbol IN (?)
		AND timestamp >= DATE_SUB(NOW(), INTERVAL 5 MINUTE)
	`

	// Convert symbols slice to interface{} for the query
	args := make([]interface{}, len(symbols))
	for i, symbol := range symbols {
		args[i] = symbol
	}

	rows, err := mds.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	prices := make(map[string]float64)
	for rows.Next() {
		var symbol string
		var price float64
		if err := rows.Scan(&symbol, &price); err != nil {
			return nil, err
		}
		prices[symbol] = price
	}

	return prices, nil
}
