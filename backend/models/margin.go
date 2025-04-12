package models

import (
	"database/sql"
	"time"
)

// Margin represents margin-related data for a client
type Margin struct {
	ID                int64     `json:"id"`
	ClientID          int64     `json:"client_id"`
	LoanAmount        float64   `json:"loan_amount"`
	InitialMargin     float64   `json:"initial_margin"`
	MaintenanceMargin float64   `json:"maintenance_margin"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// MarginStatus represents the current margin status for a client
type MarginStatus struct {
	PortfolioValue  float64 `json:"portfolio_value"`
	NetEquity       float64 `json:"net_equity"`
	MarginShortfall float64 `json:"margin_shortfall"`
	MarginCall      bool    `json:"margin_call"`
}

// MarginService handles database operations for margin data
type MarginService struct {
	DB *sql.DB
}

// GetMarginByClientID retrieves margin data for a specific client
func (ms *MarginService) GetMarginByClientID(clientID int64) (*Margin, error) {
	query := `
		SELECT id, client_id, loan_amount, initial_margin, maintenance_margin, created_at, updated_at
		FROM margins
		WHERE client_id = ?
	`

	var m Margin
	err := ms.DB.QueryRow(query, clientID).Scan(
		&m.ID,
		&m.ClientID,
		&m.LoanAmount,
		&m.InitialMargin,
		&m.MaintenanceMargin,
		&m.CreatedAt,
		&m.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// UpdateMargin updates margin data for a client
func (ms *MarginService) UpdateMargin(m *Margin) error {
	query := `
		INSERT INTO margins (client_id, loan_amount, initial_margin, maintenance_margin, created_at, updated_at)
		VALUES (?, ?, ?, ?, NOW(), NOW())
		ON DUPLICATE KEY UPDATE
		loan_amount = VALUES(loan_amount),
		initial_margin = VALUES(initial_margin),
		maintenance_margin = VALUES(maintenance_margin),
		updated_at = VALUES(updated_at)
	`

	_, err := ms.DB.Exec(query, m.ClientID, m.LoanAmount, m.InitialMargin, m.MaintenanceMargin)
	return err
}

// CalculateMarginStatus calculates the current margin status for a client
func (ms *MarginService) CalculateMarginStatus(clientID int64, positions []Position, marketPrices map[string]float64) (*MarginStatus, error) {
	margin, err := ms.GetMarginByClientID(clientID)
	if err != nil {
		return nil, err
	}
	if margin == nil {
		return nil, sql.ErrNoRows
	}

	// Calculate portfolio value
	var portfolioValue float64
	for _, position := range positions {
		if price, ok := marketPrices[position.Symbol]; ok {
			portfolioValue += float64(position.Quantity) * price
		}
	}

	// Calculate net equity
	netEquity := portfolioValue - margin.LoanAmount

	// Calculate margin shortfall
	requiredMargin := portfolioValue * margin.MaintenanceMargin
	marginShortfall := requiredMargin - netEquity

	// Determine if margin call is needed
	marginCall := marginShortfall > 0

	return &MarginStatus{
		PortfolioValue:  portfolioValue,
		NetEquity:       netEquity,
		MarginShortfall: marginShortfall,
		MarginCall:      marginCall,
	}, nil
}
