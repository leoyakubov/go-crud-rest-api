package core

import (
	"fmt"
	"go-crud-rest-api/server/config"
	"go-crud-rest-api/server/repository"
	"go-crud-rest-api/server/security"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	LOGS_DIR = "logs"
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
	logger.Infoln("Starting server...")
	defer logger.Infoln("Server started")

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
	if _, err := os.Stat(LOGS_DIR); os.IsNotExist(err) {
		err = os.MkdirAll(LOGS_DIR, 0755)
		if err != nil {
			panic(err)
		}
	}

	file, err := os.OpenFile(server.Config.ServerLogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		return fmt.Errorf("Error openning log file %v", err)
	}

	logger.SetOutput(file)

	server.Logger = logger

	return nil
}

func (server *Server) initDb() error {
	server.Logger.Infoln("Connecting to the database...")

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

func (server *Server) initWebserver() error {
	server.Webserver = echo.New()

	server.setMiddleware()
	server.initOAuth()
	server.setRoutes()

	return nil
}

func (server *Server) setMiddleware() error {
	server.Webserver.Use(config.CustomLoggingHandler("web", server.Logger))
	server.Webserver.Use(middleware.Recover())
	server.Webserver.Use(security.CORS())

	return nil
}

func (server *Server) initOAuth() {
	goth.UseProviders(facebook.New(
		server.Config.FbAppId,
		server.Config.FbSecret,
		server.Config.FbCallbackURL),
	)

	gothic.GetProviderName = func(req *http.Request) (string, error) {
		return "facebook", nil
	}

	//NOTE We should inject our own cookie store, gothic returns
	// "no SESSION_SECRET environment variable is set." with default one
	gothic.Store = sessions.NewCookieStore([]byte(server.Config.FbSecret))
}

func (server *Server) Run() error {
	if err := server.Webserver.Start(server.Config.ListenAddress); err != nil {
		return err
	}

	return nil
}
