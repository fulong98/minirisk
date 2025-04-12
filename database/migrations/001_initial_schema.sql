-- Create margins table first (since positions references it)
CREATE TABLE IF NOT EXISTS margins (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    client_id BIGINT NOT NULL,
    loan_amount DECIMAL(20, 4) NOT NULL,
    initial_margin DECIMAL(5, 4) NOT NULL,
    maintenance_margin DECIMAL(5, 4) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY uk_client_id (client_id)
) ENGINE=InnoDB;

-- Create positions table
CREATE TABLE IF NOT EXISTS positions (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    client_id BIGINT NOT NULL,
    symbol VARCHAR(10) NOT NULL,
    quantity INT NOT NULL,
    cost_basis DECIMAL(20, 4) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_client_id (client_id),
    INDEX idx_symbol (symbol)
) ENGINE=InnoDB;

-- Create market_data table
CREATE TABLE IF NOT EXISTS market_data (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    symbol VARCHAR(10) NOT NULL,
    current_price DECIMAL(20, 4) NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY idx_symbol_timestamp (symbol, timestamp),
    INDEX idx_symbol (symbol)
) ENGINE=InnoDB;

-- Add foreign key constraints
ALTER TABLE positions
ADD CONSTRAINT fk_positions_client_id
FOREIGN KEY (client_id) REFERENCES margins(client_id)
ON DELETE CASCADE;

-- Add indexes for performance
CREATE INDEX idx_market_data_timestamp ON market_data(timestamp);
CREATE INDEX idx_positions_updated_at ON positions(updated_at);
CREATE INDEX idx_margins_updated_at ON margins(updated_at);