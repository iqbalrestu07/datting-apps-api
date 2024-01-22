package handler

import (
	"fmt"
	"net/http"

	"github.com/iqbalrestu07/datting-apps-api/common"
	"github.com/iqbalrestu07/datting-apps-api/domain"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authUsecase domain.AuthUsecase
}

// NewAuthHandler will initialize the auth/ resources endpoint
func NewAuthHandler(e *echo.Echo, au domain.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: au,
	}
}

// Register will used to register new user
func (a *AuthHandler) Register(c echo.Context) (err error) {
	var auth domain.Auth
	err = c.Bind(&auth)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = a.authUsecase.Register(c.Request().Context(), auth)
	if err != nil {
		common.LogErrorWithLine(err)
		return common.APIResponse(c, "register failed", err, nil)
	}
	fmt.Println("lewat")

	return common.APIResponse(c, "successfully register", err, "success")
}

// Login used for logged in to the apps
func (a *AuthHandler) Login(c echo.Context) error {
	var auth domain.Auth
	err := c.Bind(&auth)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	token, err := a.authUsecase.Login(c.Request().Context(), auth)
	if err != nil {
		common.LogErrorWithLine(err)
		return common.APIResponse(c, "error to log in, check your credential", err, nil)
	}

	return common.APIResponse(c, "login success", err, token)
}
