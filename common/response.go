package common

import (
	"net/http"

	"github.com/iqbalrestu07/datting-apps-api/domain"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Response handle configuration for response body
type Response struct {
	Message string      `json:"message"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Token   string      `json:"token,omitempty"`
}

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// Meta hold configuration for meta
type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// APIResponse is a function for response api
func APIResponse(ctx echo.Context, message string, err error, data interface{}) error {

	jsonResponse := Response{
		Message: message,
		// Error:   err.Error(),
		Data: data,
	}
	if err != nil {
		jsonResponse.Error = err.Error()
		jsonResponse.Message = err.Error()
	}
	return ctx.JSON(getStatusCode(err), jsonResponse)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusBadRequest
	}
}
