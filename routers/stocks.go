package routers

import (
	"github.com/alexmgriffiths/algotrading/controllers/stocks"
	"github.com/alexmgriffiths/algotrading/controllers/stocks/fundamentals"
	"github.com/alexmgriffiths/algotrading/middleware"
	"github.com/gin-gonic/gin"
)

func StocksRouter(engine *gin.RouterGroup) *gin.RouterGroup {
	stocksRouter := engine.Group("stocks")
	{
		stocksRouter.Use(middleware.StocksMiddleware())
		stocksRouter.GET("/version", stocks.Version)

		// Fundamentals
		stocksRouter.GET("/:symbol", fundamentals.Profile)
		stocksRouter.GET("/:symbol/timeseries/:interval", stocks.TimeSeries)
		stocksRouter.GET("/:symbol/eod", stocks.EOD)
		stocksRouter.GET("/:symbol/dividends", fundamentals.Dividends)
		stocksRouter.GET("/:symbol/splits", fundamentals.Splits)
		stocksRouter.GET("/:symbol/earnings", fundamentals.Earnings)
	}
	return stocksRouter
}
