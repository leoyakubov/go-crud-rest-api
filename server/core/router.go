package core

import (
	"go-crud-rest-api/server/controller"
	"go-crud-rest-api/server/security"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"
)

func (server *Server) setRoutes() {
	server.Logger.Infoln("Initializing API...")

	//Task controller
	taskHandler := &controller.TaskController{}
	taskHandler.SetConfig(server.Config)
	taskHandler.SetPersistenceProvider(server.PersistenceProvider)

	//Custom JWT config
	jwtConf := middleware.JWTConfig{
		Claims:        jwt.MapClaims{},
		SigningKey:    []byte(server.Config.JwtSecret),
		SigningMethod: jwt.SigningMethodHS256.Name,
	}

	//Task API
	api := server.Webserver.Group("/api")
	api.Use(security.JwtAuthHandler(jwtConf))
	api.POST("/tasks", taskHandler.AddTask)
	api.GET("/tasks/:id", taskHandler.GetTaskById)
	api.GET("/tasks", taskHandler.GetAllTasks)
	api.PUT("/tasks/:id", taskHandler.UpdateTask)
	api.DELETE("/tasks/:id", taskHandler.DeleteTask)

	//Security controllers
	authHandler := &controller.AuthController{}
	authHandler.SetConfig(server.Config)
	fbHandler := &controller.FacebookController{}

	//Auth API
	auth := server.Webserver.Group("/auth")
	auth.POST("/login", authHandler.Login)
	auth.GET("/facebook", fbHandler.FacebookAuth)
	auth.GET("/facebook/callback", fbHandler.FacebookCallback)
}
