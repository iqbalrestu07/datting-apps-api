package handler

import (
	"errors"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/iqbalrestu07/datting-apps-api/common"
	"github.com/iqbalrestu07/datting-apps-api/domain"
)

// UserHandler  represent the httphandler for user
type UserHandler struct {
	userUsecase domain.UserUsecase
}

// NewUserHandler will initialize the users/ resources endpoint
func NewUserHandler(us domain.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: us,
	}
}

// FindAllAvailable will fetch the user based on given params
func (a *UserHandler) FindAllAvailable(c echo.Context) error {

	res, err := a.userUsecase.GetUserListSortedByInterest(c.Request().Context(), c.Get("user_id").(string))
	if err != nil {
		common.LogErrorWithLine(err)
		return common.APIResponse(c, "error when find users", err, nil)
	}

	return common.APIResponse(c, "successfully showing all users", err, res)
}

// FindByID will get user by given id
func (a *UserHandler) FindByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return common.APIResponse(c, "please check your input parameter", errors.New("id not found"), nil)
	}

	ctx := c.Request().Context()

	user, err := a.userUsecase.FindByID(ctx, id)
	if err != nil {
		return common.APIResponse(c, "error when find user by id", err, user)
	}

	return common.APIResponse(c, "success find user by id", err, user)
}

func isRequestValid(m *domain.User) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Update will update the user by given request body
func (ih *UserHandler) Update(c echo.Context) (err error) {
	var user domain.User
	err = c.Bind(&user)
	if err != nil {
		return common.APIResponse(c, "failed update user, check your payload", err, nil)
	}

	var ok bool
	if ok, err = isRequestValid(&user); !ok {
		return common.APIResponse(c, "failed update user, check your payload", err, nil)
	}

	ctx := c.Request().Context()
	userID, _ := uuid.FromString(c.Get("user_id").(string))
	user.ID = userID
	err = ih.userUsecase.Update(ctx, &user)
	if err != nil {
		return common.APIResponse(c, "failed update user", err, nil)
	}
	return common.APIResponse(c, "success update user", err, nil)
}
