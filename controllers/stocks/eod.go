package stocks

import (
	"github.com/alexmgriffiths/algotrading/lib/twelve/stocks"
	"github.com/gin-gonic/gin"
)

func EOD(ctx *gin.Context) {
	stocks := ctx.MustGet("stocks").(*stocks.Stocks)
	eod, err := stocks.GetEOD(ctx.Param("symbol"))
	if err != nil {
		ctx.AbortWithStatusJSON(200, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, eod)
}
