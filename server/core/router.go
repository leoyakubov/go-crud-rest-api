package core

import (
	"go-crud-rest-api/server/controller"
)

func (server *Server) setRoutes() {
	server.Logger.Infoln("Initializing API...")
	//defer server.Logger.Infoln("Initializing API - done")

	api := server.Webserver.Group("/api")

	// Init Task API
	taskHandler := &controller.TaskController{}
	taskHandler.SetConfig(server.Config)
	taskHandler.SetPersistenceProvider(server.PersistenceProvider)

	api.POST("/tasks", taskHandler.AddTask)
	api.GET("/tasks/:id", taskHandler.GetTaskById)
	api.GET("/tasks", taskHandler.GetAllTasks)
	api.PUT("/tasks/:id", taskHandler.UpdateTask)
	api.DELETE("/tasks/:id", taskHandler.DeleteTask)

}
