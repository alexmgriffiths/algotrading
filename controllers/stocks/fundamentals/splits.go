package fundamentals

import (
	"github.com/alexmgriffiths/algotrading/lib/twelve/stocks/fundamentals"
	"github.com/gin-gonic/gin"
)

func Splits(ctx *gin.Context) {
	stocks := ctx.MustGet("fundamentals").(*fundamentals.Fundamentals)
	splits, err := stocks.GetSplits(ctx.Param("symbol"))
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, splits)
}
