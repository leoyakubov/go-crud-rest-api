package controller

import (
	"go-crud-rest-api/server/dto"
	"go-crud-rest-api/server/errors"
	"go-crud-rest-api/server/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type TaskController struct {
	BaseController
}

func (tc *TaskController) AddTask(c echo.Context) error {
	td := &dto.TaskDto{}
	if err := c.Bind(td); err != nil {
		return c.JSON(http.StatusInternalServerError, errors.ResponseError{
			ErrorCodeId: 400,
			DevMessage:  err.Error(),
			UserMessage: errors.ERR_OCCURED,
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
		return c.JSON(http.StatusInternalServerError, errors.ResponseError{
			ErrorCodeId: 400,
			DevMessage:  err.Error(),
			UserMessage: errors.ERR_OCCURED,
		})
	}

	return c.JSON(http.StatusCreated, tm)

}

func (bc *BaseController) GetTaskById(c echo.Context) error {
	sid := c.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.ResponseError{
			ErrorCodeId: 400,
			DevMessage:  err.Error(),
			UserMessage: errors.INVALID_TASK_ID,
		})
	}

	ts, err := bc.PersistenceProvider.TaskRepo().FindOneById(id)

	if err == errors.ErrTaskNotFound {
		return c.JSON(http.StatusNotFound, errors.ResponseError{
			ErrorCodeId: 400,
			DevMessage:  err.Error(),
			UserMessage: errors.TASK_DOESNT_EXIST,
		})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.ResponseError{
			ErrorCodeId: 400,
			DevMessage:  err.Error(),
			UserMessage: errors.ERR_OCCURED,
		})
	}

	return c.JSON(http.StatusOK, ts)
}

func (bc *BaseController) GetAllTasks(c echo.Context) error {
	ts, err := bc.PersistenceProvider.TaskRepo().FindAll()

	if err == errors.ErrTaskNotFound {
		return c.JSON(http.StatusNotFound, errors.ResponseError{
			ErrorCodeId: 400,
			DevMessage:  err.Error(),
			UserMessage: errors.TASK_DOESNT_EXIST,
		})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.ResponseError{
			ErrorCodeId: 400,
			DevMessage:  err.Error(),
			UserMessage: errors.ERR_OCCURED,
		})
	}

	return c.JSON(http.StatusOK, ts)
}

func (bc *BaseController) UpdateTask(c echo.Context) error {
	sid := c.Param("id")
	id, err := strconv.Atoi(sid)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.ResponseError{
			ErrorCodeId: 400,
			DevMessage:  err.Error(),
			UserMessage: errors.ERR_OCCURED,
		})
	}

	td := &dto.TaskDto{}
	if err := c.Bind(td); err != nil {
		return c.JSON(http.StatusInternalServerError, errors.ResponseError{
			ErrorCodeId: 400,
			DevMessage:  err.Error(),
			UserMessage: errors.ERR_OCCURED,
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
		if err == errors.ErrTaskNotFound {
			return c.JSON(http.StatusNotFound, errors.ResponseError{
				ErrorCodeId: 400,
				DevMessage:  err.Error(),
				UserMessage: errors.TASK_DOESNT_EXIST,
			})
		}

		return c.JSON(http.StatusInternalServerError, errors.ResponseError{
			ErrorCodeId: 400,
			DevMessage:  err.Error(),
			UserMessage: errors.ERR_OCCURED,
		})
	}

	return c.JSON(http.StatusCreated, res)
}

func (bc *BaseController) DeleteTask(c echo.Context) error {
	sid := c.Param("id")
	id, err := strconv.Atoi(sid)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.ResponseError{
			ErrorCodeId: 400,
			DevMessage:  err.Error(),
			UserMessage: errors.ERR_OCCURED,
		})
	}

	err = bc.PersistenceProvider.TaskRepo().DeleteById(id)

	if err == errors.ErrTaskNotFound {
		return c.JSON(http.StatusNotFound, errors.ResponseError{
			ErrorCodeId: 400,
			DevMessage:  err.Error(),
			UserMessage: errors.TASK_DOESNT_EXIST,
		})

	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.ResponseError{
			ErrorCodeId: 400,
			DevMessage:  err.Error(),
			UserMessage: errors.ERR_OCCURED,
		})
	}

	return c.NoContent(http.StatusNoContent)
}
