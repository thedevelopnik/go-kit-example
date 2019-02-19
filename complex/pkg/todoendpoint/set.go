package todoendpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	"github.com/thedevelopnik/go-kit-example/complex/pb"
	"github.com/thedevelopnik/go-kit-example/complex/pkg/todoservice"
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

func MakeGetTodosEndpoint(s todoservice.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.GetTodosRequest)
		todos, err := s.GetTodos(ctx, req.UserId)
		if err != nil {
			return nil, err
		}

		resp := &pb.GetTodosResponse{
			Todos: todos,
		}
		return resp, nil
	}
}
