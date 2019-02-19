package tododatabase

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/thedevelopnik/go-kit-example/complex/pb"
)

type Middleware func(Querier) Querier

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Querier) Querier {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   Querier
}

func (mw loggingMiddleware) GetAllTodosByUserID(ctx context.Context, userID uint64) (todos []*pb.Todo, err error) {
	defer func() {
		mw.logger.Log("query", "GetAllTodosByUserID", "err", err)
	}()
	return mw.next.GetAllTodosByUserID(ctx, userID)
}
