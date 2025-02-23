package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tofustream/gintodo/internal/usecases/todos"
)

type ITodoController interface {
	FindAll(ctx *gin.Context)
}

type TodoController struct {
	usecase todos.ITodoUsecase
}

func NewTodoController(usecase todos.ITodoUsecase) ITodoController {
	return &TodoController{usecase: usecase}
}

func (c *TodoController) FindAll(ctx *gin.Context) {
	todos, err := c.usecase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": todos})
}
