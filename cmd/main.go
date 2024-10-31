package main

import (
	"net/http"

	"github.com/alexmgriffiths/algotrading/routers"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func setupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	public := r.Group("/api/v1")
	{
		public.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		routers.StocksRouter(public)
	}
	return r
}

func main() {
	router := setupRouter()
	router.Run("0.0.0.0:8090")
	// fmt.Println("Boo!")
	// t := stocks.NewStocks()
	// price, err := t.GetStockPrice("AAPL")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Printf("Stock price for Apple: %s \n", price.Price)
	// marketMovers, mmErr := t.GetMarketMovers()
	// if mmErr != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println(marketMovers)
}
