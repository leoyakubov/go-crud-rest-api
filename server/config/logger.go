package config

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// Middleware for custom logging
func CustomLoggingHandler(name string, logger *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			entry := logger.WithFields(logrus.Fields{
				"request": c.Request().RequestURI,
				"method":  c.Request().Method,
				"remote":  c.Request().RemoteAddr,
			})

			if reqID := c.Request().Header.Get("X-Request-Id"); reqID != "" {
				entry = entry.WithField("request_id", reqID)
			}

			entry.Info("<")

			if err := next(c); err != nil {
				c.Error(err)
			}

			latency := time.Since(start)

			entry.WithFields(logrus.Fields{
				"status":                                c.Response().Status,
				"text_status":                           http.StatusText(c.Response().Status),
				"took":                                  latency,
				fmt.Sprintf("measure#%s.latency", name): latency.Nanoseconds(),
			}).Info(">")

			return nil
		}
	}
}
