package main

import (
	"github.com/KoteiIto/go-todo/handler"
	"github.com/KoteiIto/go-todo/repository"
	"github.com/KoteiIto/go-todo/service"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	memoryRepository := repository.NewMemoryRepository()
	serviceImpl := service.NewServiceImpl(10, memoryRepository)
	handler := handler.NewHandler(serviceImpl)

	e.GET("/list", handler.GetTodoList)
	e.GET("/list/:page", handler.GetTodoList)
	e.GET("/:id", handler.GetTodo)
	e.POST("/create", handler.CreateTodo)
	e.POST("/:id/update", handler.UpdateTodo)
	e.POST("/:id/delete", handler.DeleteTodo)

	e.Logger.Fatal(e.Start(":1323"))
}
