package core

import (
	"fmt"
	"go-crud-rest-api/server/config"
	"os"

	"github.com/Sirupsen/logrus"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server struct {
	Config    *config.Config
	Logger    *logrus.Logger
	Webserver *echo.Echo
}

func BuildServer() *Server {
	return &Server{
		Config: config.NewConfig(),
	}
}

func (server *Server) Init(path string, logger *logrus.Logger) error {
	if err := server.loadConfig(path); err != nil {
		logger.Fatal(err)
	}

	if err := server.initLogger(logger); err != nil {
		logger.Fatal(err)
	}

	if err := server.initWebServer(); err != nil {
		logger.Fatal(err)
	}

	return nil
}

func (server *Server) loadConfig(path string) error {
	if err := server.Config.Read(path); err != nil {
		return err
	}

	return nil
}

func (server *Server) initLogger(logger *logrus.Logger) error {
	file, err := os.OpenFile(server.Config.ServerLogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		return fmt.Errorf("Error openning log file %v", err)
	}

	logger.SetOutput(file)
	//logger.SetFormatter(&logger.Formatter.Format.

	server.Logger = logger
	return nil
}

func (server *Server) initWebServer() error {
	// Echo instance
	server.Webserver = echo.New()
	//server.Webserver.Use(echologrus.NewWithNameAndLogger("web", a.logger))
	server.SetRoutes()
	server.seMiddleware()

	return nil
}

func (server *Server) seMiddleware() error {
	//server.Webserver.Use(middleware.Logger())
	server.Webserver.Use(middleware.Recover())

	server.Webserver.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	return nil
}

func (server *Server) initHTTPLogger() error {
	/*
		file, err := os.OpenFile(s.Config.Server.HttpLogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

		if err != nil {
			return fmt.Errorf("Error openning http log file %v", err)
		}

		gin.DefaultWriter = file
	*/

	return nil
}

func (server *Server) Run() error {
	fmt.Println(server.Webserver)
	if err := server.Webserver.Start(":4000"); err != nil {
		return err
	}

	return nil
}
