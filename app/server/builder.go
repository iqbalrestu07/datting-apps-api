package server

import (
	authH "github.com/iqbalrestu07/datting-apps-api/service/auth/handler"
	authUc "github.com/iqbalrestu07/datting-apps-api/service/auth/usecase"
	matchH "github.com/iqbalrestu07/datting-apps-api/service/match/handler"
	matchRepo "github.com/iqbalrestu07/datting-apps-api/service/match/repository"
	matchUc "github.com/iqbalrestu07/datting-apps-api/service/match/usecase"
	premiumH "github.com/iqbalrestu07/datting-apps-api/service/premium_feature/handler"
	premiumUc "github.com/iqbalrestu07/datting-apps-api/service/premium_feature/usecase"
	uploadH "github.com/iqbalrestu07/datting-apps-api/service/upload/handler"
	uploadRepo "github.com/iqbalrestu07/datting-apps-api/service/upload/repository"
	uploadUc "github.com/iqbalrestu07/datting-apps-api/service/upload/usecase"
	userH "github.com/iqbalrestu07/datting-apps-api/service/user/handler"
	userRepo "github.com/iqbalrestu07/datting-apps-api/service/user/repository"
	userUc "github.com/iqbalrestu07/datting-apps-api/service/user/usecase"
	userInterestRepo "github.com/iqbalrestu07/datting-apps-api/service/user_interest/repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BuildRoutes(e *echo.Echo, db *gorm.DB) {
	userRepository := userRepo.NewUserRepository(db)
	userInterestRepository := userInterestRepo.NewUserInterestRepository(db)
	matchRepository := matchRepo.NewMatchRepository(db)
	userUC := userUc.NewUserUsecase(userRepository, userInterestRepository, matchRepository)
	authUsecase := authUc.NewAuthUsecase(userRepository)
	matchUsecase := matchUc.NewMatchUsecase(matchRepository, userRepository)
	userhandler := userH.NewUserHandler(userUC)
	authHandler := authH.NewAuthHandler(e, authUsecase)
	matchHandler := matchH.NewMatchHandler(matchUsecase)
	premiumUsecase := premiumUc.NewPremiumFeatureUsecase(userRepository)
	premiumHandler := premiumH.NewPremiumFeatureHandler(premiumUsecase)
	uploadRepository := uploadRepo.NewUploadRepository(db)
	uploadUsecase := uploadUc.NewUploadUsecase(uploadRepository)
	uploadHandler := uploadH.NewUploadHandler(uploadUsecase)

	PrivateRoutes(e, userhandler, matchHandler, premiumHandler, uploadHandler)
	PublicRoutes(e, authHandler)
}
