package handler

import (
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/iqbalrestu07/datting-apps-api/common"
	"github.com/iqbalrestu07/datting-apps-api/domain"
	"github.com/iqbalrestu07/datting-apps-api/request"
)

// MatchHandler  represent the httphandler for match
type MatchHandler struct {
	matchUsecase domain.MatchUsecase
}

// NewMatchHandler will initialize the matchs/ resources endpoint
func NewMatchHandler(us domain.MatchUsecase) *MatchHandler {
	return &MatchHandler{
		matchUsecase: us,
	}
}

// FindAll will fetch the match based on given params
func (a *MatchHandler) FindAll(c echo.Context) error {
	var input request.MatchRequest
	if err := c.Bind(&input); err != nil {
		return common.APIResponse(c, "error when parsing request", err, nil)
	}

	res, err := a.matchUsecase.FindAll(c.Request().Context(), input)
	if err != nil {
		common.LogErrorWithLine(err)
		return common.APIResponse(c, "error when find matchs", err, nil)
	}

	return common.APIResponse(c, "successfully showing all matchs", err, res)
}

// Update will update the match by given request body
func (ih *MatchHandler) Match(c echo.Context) (err error) {

	var match domain.Match
	err = c.Bind(&match)
	if err != nil {
		return common.APIResponse(c, "failed create match, check your payload", err, nil)
	}

	if ok, err := isRequestValid(&match); !ok {
		return common.APIResponse(c, "failed create match, check your payload", err, nil)
	}
	match.UserID, _ = uuid.FromString(c.Get("user_id").(string))
	err = ih.matchUsecase.Match(c.Request().Context(), &match)
	if err != nil {
		return common.APIResponse(c, "failed create match", err, nil)
	}
	return common.APIResponse(c, "success create match", err, nil)
}

func isRequestValid(m *domain.Match) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
