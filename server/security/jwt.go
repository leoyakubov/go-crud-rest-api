package security

import (
	"fmt"
	"go-crud-rest-api/server/response"
	"net/http"
	"reflect"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	BEARER = "Bearer"
)

func JwtAuthHandler(config middleware.JWTConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			header := c.Request().Header.Get("Authorization")
			l := len(BEARER)

			if len(header) <= l+1 || header[:l] != BEARER {
				return c.JSON(http.StatusUnauthorized, response.ResponseError{
					ErrorCodeId: 401,
					ServerError: response.JWT_MISSING,
					UserMessage: response.INCORRECT_AUTH_TOKEN,
				})
			}

			auth := header[l+1:]
			token := new(jwt.Token)
			var err error

			keyFunc := func(t *jwt.Token) (interface{}, error) {
				// Check the signing method
				if t.Method.Alg() != config.SigningMethod {
					return nil, fmt.Errorf("Unexpected jwt signing method=%v", t.Header["alg"])
				}
				return config.SigningKey, nil
			}

			if _, ok := config.Claims.(jwt.MapClaims); ok {
				token, err = jwt.Parse(auth, keyFunc)
			} else {
				t := reflect.ValueOf(config.Claims).Type().Elem()
				claims := reflect.New(t).Interface().(jwt.Claims)
				token, err = jwt.ParseWithClaims(auth, claims, keyFunc)
			}

			if err == nil && token.Valid {
				// Store user information from token into context.
				c.Set(config.ContextKey, token)
				if config.SuccessHandler != nil {
					config.SuccessHandler(c)
				}

				return next(c)
			}

			return c.JSON(http.StatusForbidden, response.ResponseError{
				ErrorCodeId: 403,
				ServerError: err.Error(),
				UserMessage: response.FORBIDDEN,
			})
		}
	}
}
