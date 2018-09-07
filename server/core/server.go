package core

import (
	"fmt"
	"go-crud-rest-api/server/config"
	"go-crud-rest-api/server/repository"
	"go-crud-rest-api/server/security"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server struct {
	Config              *config.Config
	Logger              *logrus.Logger
	Webserver           *echo.Echo
	DB                  *gorm.DB
	PersistenceProvider repository.IPersistenceProvider
}

func BuildServer() *Server {
	return &Server{
		Config: config.NewConfig(),
	}
}

func (server *Server) Init(path string, logger *logrus.Logger) error {
	logger.Infoln("Initializing server...")
	defer logger.Infoln("Initializing server - done")

	if err := server.loadConfig(path); err != nil {
		logger.Fatal(err)
	}

	if err := server.initLogger(logger); err != nil {
		logger.Fatal(err)
	}

	if err := server.initDb(); err != nil {
		logger.Fatal(err)
	}

	if err := server.initPersistenceProvider(); err != nil {
		logger.Fatal(err)
	}

	if err := server.initWebserver(); err != nil {
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
	logger.Infoln("Initializing logger...")
	defer logger.Infoln("Initializing logger - done")

	file, err := os.OpenFile(server.Config.ServerLogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		return fmt.Errorf("Error openning log file %v", err)
	}

	logger.SetOutput(file)

	server.Logger = logger
	return nil
}

func (server *Server) initWebserver() error {
	server.Logger.Infoln("Initializing web engine...")
	defer server.Logger.Infoln("Initializing web engine - done")

	server.Webserver = echo.New()
	server.setMiddleware()
	server.setRoutes()

	return nil
}

func (server *Server) setMiddleware() error {
	//TODO set logrus output to .log file
	server.Webserver.Use(middleware.Recover())

	server.Webserver.Use(security.CORS())

	return nil
}

func (server *Server) initDb() error {
	server.Logger.Infoln("Initializing database...")
	defer server.Logger.Infoln("Initializing database - done")

	db, err := repository.NewDb(server.Config)
	if err != nil {
		return err
	}

	server.DB = db

	return nil
}
func (server *Server) initPersistenceProvider() error {
	server.PersistenceProvider = repository.NewPersistenceProvider(server.DB)

	return nil
}

func (server *Server) Run() error {
	if err := server.Webserver.Start(":4000"); err != nil {
		return err
	}

	return nil
}
