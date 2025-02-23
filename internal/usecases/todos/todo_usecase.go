package todos

import (
	todoModels "github.com/tofustream/gintodo/internal/domains/models/todos"
	todoRepositories "github.com/tofustream/gintodo/internal/infrastructures/todos"
)

type ITodoUsecase interface {
	FindAll() (*[]todoModels.Todo, error)
}

type TodoUsecase struct {
	repository todoRepositories.ITodoRepository
}

func NewTodoUsecase(repository todoRepositories.ITodoRepository) ITodoUsecase {
	return &TodoUsecase{repository: repository}
}

func (s *TodoUsecase) FindAll() (*[]todoModels.Todo, error) {
	return s.repository.FindAll()
}
