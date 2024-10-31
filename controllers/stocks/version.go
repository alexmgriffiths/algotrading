package stocks

import "github.com/gin-gonic/gin"

func Version(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"version": "1.0.0"})
}
