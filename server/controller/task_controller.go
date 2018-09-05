package controller

import (
	"fmt"
	"go-crud-rest-api/server/dto"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type TaskHandler struct {
	BaseController
}

//TODO add logging on echa action

func (bc *BaseController) AddTask(c echo.Context) error {
	d := &dto.TaskDto{
		ID: dto.Seq,
	}
	if err := c.Bind(d); err != nil {
		return err
	}
	dto.Tasks[d.ID] = d
	dto.Seq++
	return c.JSON(http.StatusCreated, d)
}

func (bc *BaseController) GetTaskById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	fmt.Println("id: ", id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, dto.Tasks[id])

}

func (bc *BaseController) GetAllTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, dto.Tasks)

}

func (bc *BaseController) UpdateTask(c echo.Context) error {
	d := new(dto.TaskDto)
	if err := c.Bind(d); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	dto.Tasks[id].Name = d.Name
	return c.JSON(http.StatusOK, dto.Tasks[id])
}

func (bc *BaseController) DeleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(dto.Tasks, id)
	return c.NoContent(http.StatusNoContent)
}
