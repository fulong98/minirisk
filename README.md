# Mini Risk Monitoring System

A real-time portfolio risk monitoring system that tracks client positions, calculates margin requirements, and provides dynamic risk insights.

## Project Overview

The Mini Risk Monitoring System is a full-stack application designed to monitor client portfolio risks in real-time. It fetches live stock market data, calculates margin requirements, and presents a dynamic dashboard for visual insights.

## Technology Stack

- **Backend**: Golang
  - RESTful APIs for data ingestion and margin calculations
  - Real-time market data processing
  - Concurrent processing capabilities

- **Frontend**: React
  - Responsive dashboard
  - Real-time data visualization
  - Interactive portfolio management

- **Database**: MySQL
  - Structured data storage
  - Efficient querying with proper indexing
  - Reliable data persistence

## Architecture

### Data Flow
1. Market data ingestion from external APIs
2. Position and margin calculations
3. Real-time dashboard updates
4. Alert generation for margin calls

### Database Schema
- Positions Table: Stock positions (symbol, quantity, cost basis, client_id)
- Market Data Table: Real-time market data (symbol, current_price, timestamp)
- Margin Table: Loan amounts and margin-related data per client

## API Endpoints

- `GET /api/market-data`: Current market prices
- `GET /api/positions/:clientId`: Client-specific portfolio data
- `GET /api/margin-status/:clientId`: Margin risk status and calculations

## Setup Instructions

### Prerequisites
- Go 1.21 or later
- Node.js 18 or later
- MySQL 8.0 or later
- Docker (optional)

### Backend Setup
1. Navigate to the backend directory
2. Install dependencies: `go mod download`
3. Configure environment variables
4. Run the server: `go run main.go`

### Frontend Setup
1. Navigate to the frontend directory
2. Install dependencies: `npm install`
3. Configure environment variables
4. Start the development server: `npm start`

### Database Setup
1. Create the MySQL database
2. Run the schema migrations
3. Configure connection settings

## Development

### Running Tests
- Backend: `go test ./...`
- Frontend: `npm test`

### Building for Production
- Backend: `go build`
- Frontend: `npm run build`

## Contributing

Please read CONTRIBUTING.md for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the LICENSE file for details.