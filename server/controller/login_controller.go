package controller

import (
	"github.com/leoyakubov/go-crud-rest-api/server/response"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

const (
	DEMO_LOGIN    string = "demouser"
	DEMO_PASSWORD string = "demopassword"
	DEMO_USERNAME string = "Demo User"
	DEMO_USER_ID  string = "666"
)

type AuthController struct {
	BaseController
}

func (ah *AuthController) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	//Here we should get user from db and check credentials
	if username == DEMO_LOGIN && password == DEMO_PASSWORD {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = DEMO_USERNAME
		claims["userId"] = DEMO_USER_ID
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(ah.Config.JwtSecret))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseError{
				ErrorCodeId: 500,
				ServerError: err.Error(),
				UserMessage: response.ERR_OCCURED,
			})
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return c.JSON(http.StatusUnauthorized, response.ResponseError{
		ErrorCodeId: 401,
		ServerError: response.USERNAME_PASS_INCORRECT,
		UserMessage: response.USERNAME_PASS_INCORRECT,
	})
}

func (bc *BaseController) LoginOptions(c echo.Context) error {
	c.Request().Header.Add("Allow", "POST")
	//Send just some basic info about endpoint
	return c.JSON(http.StatusOK, "Allowed: POST /login")
}
