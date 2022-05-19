package route

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func ProvideGinEngine() *gin.Engine {
	return gin.Default()
}

var Module = fx.Option(
	fx.Provide(ProvideGinEngine),
)
