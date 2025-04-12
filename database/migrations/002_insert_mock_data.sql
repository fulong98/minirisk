-- Insert sample client margin accounts
INSERT INTO margins (client_id, loan_amount, initial_margin, maintenance_margin) VALUES
(1001, 100000.0000, 0.5000, 0.2500),
(1002, 250000.0000, 0.4500, 0.2000),
(1003, 75000.0000, 0.5500, 0.3000),
(1004, 500000.0000, 0.4000, 0.1800),
(1005, 150000.0000, 0.4800, 0.2200),
(1006, 300000.0000, 0.4200, 0.1900),
(1007, 125000.0000, 0.5200, 0.2700),
(1008, 200000.0000, 0.4700, 0.2300),
(1009, 180000.0000, 0.5100, 0.2600),
(1010, 400000.0000, 0.4300, 0.2100);

-- Insert sample positions
INSERT INTO positions (client_id, symbol, quantity, cost_basis) VALUES
-- Client 1001
(1001, 'AAPL', 500, 165.7500),
(1001, 'MSFT', 300, 340.2200),
(1001, 'GOOGL', 100, 142.3600),

-- Client 1002
(1002, 'AMZN', 200, 178.5000),
(1002, 'NVDA', 400, 820.3000),
(1002, 'TSLA', 150, 245.9800),

-- Client 1003
(1003, 'META', 250, 455.6700),
(1003, 'JPM', 300, 188.4200),

-- Client 1004
(1004, 'AAPL', 1000, 162.4500),
(1004, 'MSFT', 800, 330.8700),
(1004, 'AMZN', 500, 172.3400),
(1004, 'NVDA', 600, 810.2500),

-- Client 1005
(1005, 'GOOGL', 300, 144.7800),
(1005, 'TSLA', 200, 249.7600),

-- Client 1006
(1006, 'META', 400, 448.9200),
(1006, 'JPM', 500, 185.6300),
(1006, 'AMD', 800, 154.2400),

-- Client 1007
(1007, 'INTC', 700, 44.3200),
(1007, 'PFE', 600, 28.5600),

-- Client 1008
(1008, 'KO', 900, 62.4300),
(1008, 'PEP', 400, 170.8700),
(1008, 'JNJ', 300, 145.6500),

-- Client 1009
(1009, 'DIS', 500, 103.4200),
(1009, 'NFLX', 150, 605.7800),

-- Client 1010
(1010, 'AAPL', 800, 167.2300),
(1010, 'MSFT', 600, 339.4600),
(1010, 'GOOGL', 400, 143.8900),
(1010, 'AMZN', 300, 176.5200);

-- Insert current market prices
INSERT INTO market_data (symbol, current_price) VALUES
('AAPL', 172.8500),
('MSFT', 345.6700),
('GOOGL', 145.9800),
('AMZN', 182.3400),
('NVDA', 840.7600),
('TSLA', 235.4200),
('META', 475.2300),
('JPM', 192.4500),
('AMD', 162.8700),
('INTC', 42.7600),
('PFE', 27.3400),
('KO', 64.5600),
('PEP', 174.2300),
('JNJ', 148.7600),
('DIS', 107.5600),
('NFLX', 622.3400);
