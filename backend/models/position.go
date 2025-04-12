package models

import (
	"database/sql"
	"time"
)

// Position represents a stock position in a client's portfolio
type Position struct {
	ID        int64     `json:"id"`
	ClientID  int64     `json:"client_id"`
	Symbol    string    `json:"symbol"`
	Quantity  int       `json:"quantity"`
	CostBasis float64   `json:"cost_basis"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PositionService handles database operations for positions
type PositionService struct {
	DB *sql.DB
}

// GetPositionsByClientID retrieves all positions for a specific client
func (ps *PositionService) GetPositionsByClientID(clientID int64) ([]Position, error) {
	query := `
		SELECT id, client_id, symbol, quantity, cost_basis, created_at, updated_at
		FROM positions
		WHERE client_id = ?
	`

	rows, err := ps.DB.Query(query, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var positions []Position
	for rows.Next() {
		var p Position
		err := rows.Scan(
			&p.ID,
			&p.ClientID,
			&p.Symbol,
			&p.Quantity,
			&p.CostBasis,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		positions = append(positions, p)
	}

	return positions, nil
}

// CreatePosition creates a new position for a client
func (ps *PositionService) CreatePosition(p *Position) error {
	query := `
		INSERT INTO positions (client_id, symbol, quantity, cost_basis, created_at, updated_at)
		VALUES (?, ?, ?, ?, NOW(), NOW())
	`

	result, err := ps.DB.Exec(query, p.ClientID, p.Symbol, p.Quantity, p.CostBasis)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	p.ID = id
	return nil
}

// UpdatePosition updates an existing position
func (ps *PositionService) UpdatePosition(p *Position) error {
	query := `
		UPDATE positions
		SET quantity = ?, cost_basis = ?, updated_at = NOW()
		WHERE id = ? AND client_id = ?
	`

	_, err := ps.DB.Exec(query, p.Quantity, p.CostBasis, p.ID, p.ClientID)
	return err
}

// DeletePosition deletes a position
func (ps *PositionService) DeletePosition(id, clientID int64) error {
	query := `
		DELETE FROM positions
		WHERE id = ? AND client_id = ?
	`

	_, err := ps.DB.Exec(query, id, clientID)
	return err
}
