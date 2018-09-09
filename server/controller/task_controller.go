package controller

import (
	"go-crud-rest-api/server/dto"
	"go-crud-rest-api/server/model"
	"go-crud-rest-api/server/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

const (
	TASK_OPTIONS = "Allowed: " +
		"Get all tasks - GET /tasks \n" +
		"Add Task - POST /tasks \n"

	TASK_BY_ID_OPTIONS = "Allowed: \n" +
		"Get task - GET /tasks/:id \n" +
		"Update Task - PUT /tasks/:id \n" +
		"Delete task - DELETE /tasks/:id \n"
)

type TaskController struct {
	BaseController
}

func (tc *TaskController) AddTask(c echo.Context) error {
	td := &dto.TaskDto{}
	if err := c.Bind(td); err != nil {
		return c.JSON(http.StatusInternalServerError, response.ResponseError{
			ErrorCodeId: 500,
			ServerError: err.Error(),
			UserMessage: response.ERR_OCCURED,
		})
	}

	//TODO add dto -> model transformer
	tm := &model.Task{
		Title:       td.Title,
		Description: td.Description,
		Priority:    td.Priority,
		CompletedAt: td.CompletedAt,
		IsCompleted: td.IsCompleted,
	}

	err := tc.PersistenceProvider.TaskRepo().Add(tm)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ResponseError{
			ErrorCodeId: 500,
			ServerError: err.Error(),
			UserMessage: response.ERR_OCCURED,
		})
	}

	return c.JSON(http.StatusCreated, tm)
}

func (bc *BaseController) GetTaskById(c echo.Context) error {
	sid := c.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ResponseError{
			ErrorCodeId: 500,
			ServerError: err.Error(),
			UserMessage: response.INVALID_TASK_ID,
		})
	}

	ts, err := bc.PersistenceProvider.TaskRepo().FindOneById(id)

	if err == response.ErrTaskNotFound {
		return c.JSON(http.StatusNotFound, response.ResponseError{
			ErrorCodeId: 500,
			ServerError: err.Error(),
			UserMessage: response.TASK_DOESNT_EXIST,
		})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ResponseError{
			ErrorCodeId: 500,
			ServerError: err.Error(),
			UserMessage: response.ERR_OCCURED,
		})
	}

	return c.JSON(http.StatusOK, ts)
}

func (bc *BaseController) GetAllTasks(c echo.Context) error {
	ts, err := bc.PersistenceProvider.TaskRepo().FindAll()

	if err == response.ErrTaskNotFound {
		return c.JSON(http.StatusNotFound, response.ResponseError{
			ErrorCodeId: 404,
			ServerError: err.Error(),
			UserMessage: response.TASK_DOESNT_EXIST,
		})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ResponseError{
			ErrorCodeId: 500,
			ServerError: err.Error(),
			UserMessage: response.ERR_OCCURED,
		})
	}

	return c.JSON(http.StatusOK, ts)
}

func (bc *BaseController) UpdateTask(c echo.Context) error {
	sid := c.Param("id")
	id, err := strconv.Atoi(sid)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ResponseError{
			ErrorCodeId: 500,
			ServerError: err.Error(),
			UserMessage: response.ERR_OCCURED,
		})
	}

	td := &dto.TaskDto{}
	if err := c.Bind(td); err != nil {
		return c.JSON(http.StatusInternalServerError, response.ResponseError{
			ErrorCodeId: 500,
			ServerError: err.Error(),
			UserMessage: response.ERR_OCCURED,
		})
	}

	tm := &model.Task{
		Title:       td.Title,
		Description: td.Description,
		Priority:    td.Priority,
		CompletedAt: td.CompletedAt,
		IsCompleted: td.IsCompleted,
	}

	res, err := bc.PersistenceProvider.TaskRepo().UpdateById(id, tm)

	if err != nil {
		if err == response.ErrTaskNotFound {
			return c.JSON(http.StatusNotFound, response.ResponseError{
				ErrorCodeId: 500,
				ServerError: err.Error(),
				UserMessage: response.TASK_DOESNT_EXIST,
			})
		}

		return c.JSON(http.StatusInternalServerError, response.ResponseError{
			ErrorCodeId: 500,
			ServerError: err.Error(),
			UserMessage: response.ERR_OCCURED,
		})
	}

	return c.JSON(http.StatusCreated, res)
}

func (bc *BaseController) DeleteTask(c echo.Context) error {
	sid := c.Param("id")
	id, err := strconv.Atoi(sid)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ResponseError{
			ErrorCodeId: 500,
			ServerError: err.Error(),
			UserMessage: response.ERR_OCCURED,
		})
	}

	err = bc.PersistenceProvider.TaskRepo().DeleteById(id)

	if err == response.ErrTaskNotFound {
		return c.JSON(http.StatusNotFound, response.ResponseError{
			ErrorCodeId: 404,
			ServerError: err.Error(),
			UserMessage: response.TASK_DOESNT_EXIST,
		})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ResponseError{
			ErrorCodeId: 500,
			ServerError: err.Error(),
			UserMessage: response.ERR_OCCURED,
		})
	}

	return c.NoContent(http.StatusNoContent)
}

func (bc *BaseController) TaskOptions(c echo.Context) error {
	c.Request().Header.Add("Allow", "GET,POST")
	//Send just some basic info about endpoint
	return c.JSON(http.StatusOK, TASK_OPTIONS)
}

func (bc *BaseController) TaskByIdOptions(c echo.Context) error {
	c.Request().Header.Add("Allow", "GET,PUT,DELETE")
	return c.JSON(http.StatusOK, TASK_BY_ID_OPTIONS)
}
