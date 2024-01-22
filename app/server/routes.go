package server

import (
	"github.com/iqbalrestu07/datting-apps-api/app/server/middleware"
	auth "github.com/iqbalrestu07/datting-apps-api/service/auth/handler"
	match "github.com/iqbalrestu07/datting-apps-api/service/match/handler"

	premium "github.com/iqbalrestu07/datting-apps-api/service/premium_feature/handler"

	upload "github.com/iqbalrestu07/datting-apps-api/service/upload/handler"
	user "github.com/iqbalrestu07/datting-apps-api/service/user/handler"

	"github.com/labstack/echo/v4"
)

func PrivateRoutes(
	e *echo.Echo,
	userHandler *user.UserHandler,
	matchHandler *match.MatchHandler,
	premiumFeatureHandler *premium.PremiumFeatureHandler,
	uploadHandler *upload.UploadHandler,

) {
	v1 := e.Group("/api/v1")
	v1.Use(middleware.Authorize(middleware.NewJWTAuthService()))

	// user routes
	v1.GET("/users", userHandler.FindAllAvailable)
	v1.GET("/users/:id", userHandler.FindByID)
	v1.PUT("/users", userHandler.Update)

	// match routes
	v1.POST("/match", matchHandler.Match)
	v1.GET("/match", matchHandler.FindAll)

	// premium feature
	v1.POST("/subscribe", premiumFeatureHandler.Subscribe)
	// upload routes
	v1.POST("/upload", uploadHandler.Upload)

}

func PublicRoutes(
	e *echo.Echo,
	authhandler *auth.AuthHandler) {
	e.POST("/register", authhandler.Register)
	e.POST("/login", authhandler.Login)

}
