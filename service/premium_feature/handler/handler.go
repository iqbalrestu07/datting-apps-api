package handler

import (
	"net/http"

	"github.com/iqbalrestu07/datting-apps-api/common"
	"github.com/iqbalrestu07/datting-apps-api/domain"
	"github.com/labstack/echo/v4"
)

type PremiumFeatureHandler struct {
	premiumFeatureUsecase domain.PremiumFeatureUsecase
}

// NewPremiumFeatureHandler will initialize the premiumFeature resources endpoint
func NewPremiumFeatureHandler(au domain.PremiumFeatureUsecase) *PremiumFeatureHandler {
	return &PremiumFeatureHandler{
		premiumFeatureUsecase: au,
	}
}

// Subscribe will used to register new user
func (a *PremiumFeatureHandler) Subscribe(c echo.Context) (err error) {
	var premiumFeature domain.PremiumFeature
	err = c.Bind(&premiumFeature)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = a.premiumFeatureUsecase.Subscribe(c.Request().Context(), premiumFeature)
	if err != nil {
		common.LogErrorWithLine(err)
		return common.APIResponse(c, "error when subscribe feature", err, nil)
	}

	return common.APIResponse(c, "successfully subscribe feature", err, "success")
}
