package todos

import "github.com/tofustream/gintodo/internal/domains/models/todos"

type ITodoRepository interface {
	FindAll() (*[]todos.Todo, error)
}

type InMemoryTodoRepository struct {
	todos []todos.Todo
}

func NewInMemoryTodoRepository(todos []todos.Todo) ITodoRepository {
	return &InMemoryTodoRepository{todos: todos}
}

func (r *InMemoryTodoRepository) FindAll() (*[]todos.Todo, error) {
	return &r.todos, nil
}
