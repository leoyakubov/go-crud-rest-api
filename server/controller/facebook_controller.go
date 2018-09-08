package controller

import (
	"go-crud-rest-api/server/response"
	"net/http"

	"github.com/labstack/echo"
	"github.com/markbates/goth/gothic"
)

type FacebookController struct {
	BaseController
}

func (fc *FacebookController) FacebookAuth(c echo.Context) error {
	gothic.BeginAuthHandler(c.Response().Writer, c.Request())

	return nil
}

func (fc *FacebookController) FacebookCallback(c echo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response().Writer, c.Request())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ResponseError{
			ErrorCodeId: 500,
			ServerError: err.Error(),
			UserMessage: response.ERR_OCCURED,
		})

	}

	return c.JSON(http.StatusOK, map[string]string{
		"UserAccessToken": user.AccessToken,
		"UserName":        user.Name,
		"UserEmail":       user.Email,
	})
}
