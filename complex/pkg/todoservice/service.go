package todoservice

import (
	"context"

	"github.com/go-kit/kit/log"

	"github.com/thedevelopnik/go-kit-example/complex/pb"
	"github.com/thedevelopnik/go-kit-example/complex/pkg/tododatabase"
)

type Service interface {
	GetTodos(ctx context.Context, userID uint64) ([]*pb.Todo, error)
}

func New(querier tododatabase.Querier, logger log.Logger) Service {
	var svc Service
	{
		svc = NewBasicService(querier)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

func NewBasicService(querier tododatabase.Querier) Service {
	return basicService{
		querier: querier,
	}
}

type basicService struct {
	querier tododatabase.Querier
}

func (s basicService) GetTodos(ctx context.Context, userID uint64) ([]*pb.Todo, error) {
	todos, err := s.querier.GetAllTodosByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return todos, nil
}
