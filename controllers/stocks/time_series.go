package stocks

import (
	"github.com/alexmgriffiths/algotrading/lib/twelve/stocks"
	"github.com/gin-gonic/gin"
)

func TimeSeries(ctx *gin.Context) {
	stocks := ctx.MustGet("stocks").(*stocks.Stocks)
	timeseries, err := stocks.GetTimeSeries(ctx.Param("symbol"), ctx.Param("interval"))
	if err != nil {
		ctx.AbortWithStatusJSON(200, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, timeseries)
}
