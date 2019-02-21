package todoservice

import (
	"context"

	"github.com/go-kit/kit/log"
)

type Todo struct {
	Name string `json:"name"`
}

type Service interface {
	GetTodos(ctx context.Context) ([]Todo, error)
}

func New(logger log.Logger) Service {
	var svc Service
	{
		svc = NewBasicService()
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

func NewBasicService() Service {
	return basicService{}
}

type basicService struct{}

func (s basicService) GetTodos(_ context.Context) ([]Todo, error) {
	todo1 := Todo{
		Name: "go to the grocery store",
	}
	todo2 := Todo{
		Name: "vacuum",
	}
	todo3 := Todo{
		Name: "learn go",
	}

	todos := []Todo{todo1, todo2, todo3}

	return todos, nil
}
