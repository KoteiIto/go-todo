package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/KoteiIto/go-todo/entity"

	"github.com/KoteiIto/go-todo/service"
	"github.com/labstack/echo"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

type createTodoBody struct {
	Title       string    `body:"title"`
	Description string    `body:"description"`
	Expire      time.Time `body:"expire"`
}

type updateTodoBody struct {
	ID int `body:"id"`
	createTodoBody
}

func (h *Handler) GetTodoList(c echo.Context) error {
	var page int
	maybePage := c.Param("page")
	if v, err := strconv.Atoi(maybePage); err == nil {
		page = v
	} else {
		page = 1
	}

	todoList, err := h.service.GetList(page)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, todoList)
}

func (h *Handler) GetTodo(c echo.Context) error {
	var id int
	maybeID := c.Param("id")
	if v, err := strconv.Atoi(maybeID); err == nil {
		id = v
	}

	todo, err := h.service.Get(id)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, todo)
}

func (h *Handler) CreateTodo(c echo.Context) error {
	body := new(createTodoBody)
	if err := c.Bind(body); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	todo := &entity.Todo{
		Title:       body.Title,
		Description: body.Description,
		Expire:      body.Expire,
	}

	err := h.service.Create(todo)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, nil)
}

func (h *Handler) UpdateTodo(c echo.Context) error {
	body := new(updateTodoBody)
	if err := c.Bind(body); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	todo := &entity.Todo{
		ID:          body.ID,
		Title:       body.Title,
		Description: body.Description,
		Expire:      body.Expire,
	}

	err := h.service.Update(todo)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, nil)
}

func (h *Handler) DeleteTodo(c echo.Context) error {
	var id int
	maybeID := c.Param("id")
	if v, err := strconv.Atoi(maybeID); err == nil {
		id = v
	}

	err := h.service.Delete(id)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, nil)
}
