package middleware

import (
	"github.com/alexmgriffiths/algotrading/lib/twelve/stocks"
	"github.com/alexmgriffiths/algotrading/lib/twelve/stocks/fundamentals"
	"github.com/gin-gonic/gin"
)

func StocksMiddleware() gin.HandlerFunc {
	stocks := stocks.NewStocks()
	fundamentals := fundamentals.NewFundamentals(stocks)
	return func(c *gin.Context) {
		c.Set("stocks", stocks)
		c.Set("fundamentals", fundamentals)
		c.Next()
	}
}
