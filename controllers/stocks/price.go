package stocks

import (
	"github.com/alexmgriffiths/algotrading/lib/twelve/stocks"
	"github.com/gin-gonic/gin"
)

func Price(ctx *gin.Context) {
	stocks := ctx.MustGet("stocks").(*stocks.Stocks)
	stockData, err := stocks.GetStockPrice(ctx.Param("symbol"))
	if err != nil {
		ctx.AbortWithStatusJSON(200, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, gin.H{"price": stockData.Price})
}
