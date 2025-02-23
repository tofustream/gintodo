package main

import (
	"github.com/gin-gonic/gin"
	todoControllers "github.com/tofustream/gintodo/internal/controllers/todos"
	todoModels "github.com/tofustream/gintodo/internal/domains/models/todos"
	todoRepositories "github.com/tofustream/gintodo/internal/infrastructures/todos"
	todoUsecases "github.com/tofustream/gintodo/internal/usecases/todos"
)

func main() {
	todos := []todoModels.Todo{
		{ID: 1, Title: "Todo 1", Description: "This is Todo 1.", IsCompleted: false},
		{ID: 2, Title: "Todo 2", Description: "This is Todo 2.", IsCompleted: true},
		{ID: 3, Title: "Todo 3", Description: "This is Todo 3.", IsCompleted: false},
	}

	todoRepository := todoRepositories.NewInMemoryTodoRepository(todos)
	todoUsecase := todoUsecases.NewTodoUsecase(todoRepository)
	todoController := todoControllers.NewTodoController(todoUsecase)

	r := gin.Default()
	r.GET("/todos", todoController.FindAll)
	r.Run("localhost:8080")
}
