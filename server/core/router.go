package core

import (
	"go-crud-rest-api/server/controller"
)

func (server *Server) SetRoutes() {
	server.Logger.Infoln("initializing routing and handlers...")
	defer server.Logger.Infoln("initializing routing and handlers - done!")

	// routes for tasks CRUD operations
	tasks := server.Webserver.Group("/api")

	//Add Task CRUD API
	taskHandler := &controller.TaskHandler{}
	tasks.POST("/tasks", taskHandler.AddTask)
	tasks.GET("/tasks/:id", taskHandler.GetTaskById)
	tasks.GET("/tasks", taskHandler.GetAllTasks)
	tasks.PUT("/tasks/:id", taskHandler.UpdateTask)
	tasks.DELETE("/tasks/:id", taskHandler.DeleteTask)

}
