package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/morticius/accuracy/internal/repositories/file"
	"github.com/morticius/accuracy/internal/server/config"
	"github.com/morticius/accuracy/internal/server/middlewares"
	internalServices "github.com/morticius/accuracy/internal/server/services"
	externalServices "github.com/morticius/accuracy/pkg/services"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	btcService := externalServices.NewBTCService(config.Get().BtcRateEndpoint)
	ratesService := internalServices.NewRatesService(btcService)
	userRepository := file.NewUserFileRepository("data")
	jwtService := internalServices.NewJWTAuthService()
	authService := internalServices.NewAuthService(userRepository, jwtService)

	user := router.Group("/user")
	user.POST("create", authService.Create)
	user.POST("login", authService.SignIn)

	protected := router.Use(middlewares.AuthorizeJWT())

	protected.GET("btcRate", ratesService.GetRateBTCToUAHHandler)

	return router
}
