package services

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/minirisk/models"
)

// MarginAlertService handles margin calculations and alerts
type MarginAlertService struct {
	DB *sql.DB
}

// NewMarginAlertService creates a new MarginAlertService instance
func NewMarginAlertService(db *sql.DB) *MarginAlertService {
	return &MarginAlertService{DB: db}
}

// CheckMarginStatus checks margin status for all clients and sends alerts if needed
func (mas *MarginAlertService) CheckMarginStatus() error {
	// Get all clients with positions
	clients, err := mas.getClientsWithPositions()
	if err != nil {
		return fmt.Errorf("failed to get clients: %v", err)
	}

	// Check margin status for each client
	for _, clientID := range clients {
		status, err := mas.calculateClientMarginStatus(clientID)
		if err != nil {
			fmt.Printf("Failed to calculate margin status for client %d: %v\n", clientID, err)
			continue
		}

		// Send alert if margin call is needed
		if status.MarginCall {
			if err := mas.sendMarginCallAlert(clientID, status); err != nil {
				fmt.Printf("Failed to send margin call alert for client %d: %v\n", clientID, err)
			}
		}
	}

	return nil
}

// getClientsWithPositions retrieves all clients with active positions
func (mas *MarginAlertService) getClientsWithPositions() ([]int64, error) {
	query := "SELECT DISTINCT client_id FROM positions"
	rows, err := mas.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []int64
	for rows.Next() {
		var clientID int64
		if err := rows.Scan(&clientID); err != nil {
			return nil, err
		}
		clients = append(clients, clientID)
	}

	return clients, nil
}

// calculateClientMarginStatus calculates margin status for a specific client
func (mas *MarginAlertService) calculateClientMarginStatus(clientID int64) (*models.MarginStatus, error) {
	// Get client's positions
	positionService := &models.PositionService{DB: mas.DB}
	positions, err := positionService.GetPositionsByClientID(clientID)
	if err != nil {
		return nil, fmt.Errorf("failed to get positions: %v", err)
	}

	// Get current market prices
	var symbols []string
	for _, position := range positions {
		symbols = append(symbols, position.Symbol)
	}

	marketDataService := &models.MarketDataService{DB: mas.DB}
	marketPrices, err := marketDataService.GetMarketDataForSymbols(symbols)
	if err != nil {
		return nil, fmt.Errorf("failed to get market prices: %v", err)
	}

	// Calculate margin status
	marginService := &models.MarginService{DB: mas.DB}
	return marginService.CalculateMarginStatus(clientID, positions, marketPrices)
}

// sendMarginCallAlert sends a margin call alert for a client
func (mas *MarginAlertService) sendMarginCallAlert(clientID int64, status *models.MarginStatus) error {
	// In a real implementation, this would send an email, SMS, or other notification
	// For now, we'll just log the alert
	fmt.Printf("MARGIN CALL ALERT - Client ID: %d\n", clientID)
	fmt.Printf("Portfolio Value: $%.2f\n", status.PortfolioValue)
	fmt.Printf("Net Equity: $%.2f\n", status.NetEquity)
	fmt.Printf("Margin Shortfall: $%.2f\n", status.MarginShortfall)
	fmt.Printf("Time: %s\n", time.Now().Format(time.RFC3339))
	return nil
}

// StartMarginMonitoring begins the margin monitoring process
func (mas *MarginAlertService) StartMarginMonitoring(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			if err := mas.CheckMarginStatus(); err != nil {
				fmt.Printf("Error checking margin status: %v\n", err)
			}
		}
	}()
}
