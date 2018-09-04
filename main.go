package main

import (
	_ "database/sql"
	"fmt"
	"net/http"

	"github.com/JonathanMH/goClacks/echo"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	e.Use(goClacks.Terrify) // optional ;)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Route => handler
	e.GET("/", func(c echo.Context) error {

		return c.JSON(http.StatusOK, "Hi!")
	})

	e.GET("/id/:id", func(c echo.Context) error {
		requestedID := c.Param("id")
		fmt.Println(requestedID)
		return c.JSON(http.StatusOK, requestedID)
	})

	e.Logger.Fatal(e.Start(":4000"))
}
