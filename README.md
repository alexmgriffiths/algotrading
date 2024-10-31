# AlgoTrading

AlgoTrading is a Golang-based API for retrieving stock market trading tools and data, powered by the [Twelve Data API](https://twelvedata.com/).

## Features

- Fetch stock prices, dividends, earnings, and end-of-day data.
- Modular setup to facilitate middleware and controller expansion.

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/alexmgriffiths/algotrading.git
2. Install dependencies
    ```bash
    go mod tidy
3. Set up environment variables
   Add your API key to the `.env.sample` file and rename the file to `.env`

## Usage

### Running the server

    ```bash
    go run ./cmd/main.go


The server listens on port 8090 by default.

## API Routes

* Get stock `/api/v1/stocks/SYMBOL`
* Get stock EOD price `/api/v1/stocks/SYMBOL/eod`
* Get stock dividends `/api/v1/stocks/SYMBOL/dividends`
* Get stock splits `/api/v1/stocks/SYMBOL/splits`
* Get stock earnings `/api/v1/stocks/SYMBOL/earnings`

### Examples

* http://localhost:8090/api/v1/stocks/AAPL
* http://localhost:8090/api/v1/stocks/AAPL/eod
* http://localhost:8090/api/v1/stocks/AAPL/dividends
* http://localhost:8090/api/v1/stocks/AAPL/splits
* http://localhost:8090/api/v1/stocks/AAPL/earnings