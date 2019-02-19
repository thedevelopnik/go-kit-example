package todoendpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	"github.com/thedevelopnik/go-kit-example/simple/pkg/todoservice"
)

type Set struct {
	GetTodosEndpoint endpoint.Endpoint
}

func New(svc todoservice.Service, logger log.Logger) Set {
	var getTodosEndpoint endpoint.Endpoint
	{
		getTodosEndpoint = MakeGetTodosEndpoint(svc)
		getTodosEndpoint = LoggingMiddleware(log.With(logger, "method", "GetTodos"))(getTodosEndpoint)
	}
	return Set{
		GetTodosEndpoint: getTodosEndpoint,
	}
}

func (s Set) GetTodos(ctx context.Context) ([3]todoservice.Todo, error) {
	resp, err := s.GetTodosEndpoint(ctx, nil)
	if err != nil {
		return resp.(GetTodosResponse).V, err
	}
	response := resp.(GetTodosResponse)
	return response.V, response.Err
}

func MakeGetTodosEndpoint(s todoservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		todos, err := s.GetTodos(ctx)
		return GetTodosResponse{V: todos, Err: err}, nil
	}
}

type GetTodosResponse struct {
	V   [3]todoservice.Todo `json:"v"`
	Err error               `json:"-"`
}
