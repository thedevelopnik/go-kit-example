package todoservice

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/thedevelopnik/go-kit-example/complex/pb"
)

type Middleware func(Service) Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

func (mw loggingMiddleware) GetTodos(ctx context.Context, userID uint64) (todos []*pb.Todo, err error) {
	defer func() {
		mw.logger.Log("method", "GetTodos", "err", err)
	}()
	return mw.next.GetTodos(ctx, userID)
}
