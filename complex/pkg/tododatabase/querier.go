package tododatabase

import (
	"context"

	"github.com/go-kit/kit/log"
	memdb "github.com/hashicorp/go-memdb"

	"github.com/thedevelopnik/go-kit-example/complex/pb"
)

type Querier interface {
	GetAllTodosByUserID(ctx context.Context, userID uint64) (todos []*pb.Todo, err error)
}

func New(db *memdb.MemDB, logger log.Logger) Querier {
	var q Querier
	{
		q = newBasicQuerier(db)
		q = LoggingMiddleware(logger)(q)
	}
	return q
}

func newBasicQuerier(db *memdb.MemDB) Querier {
	return basicQuerier{
		db: db,
	}
}

type basicQuerier struct {
	db *memdb.MemDB
}

func (q basicQuerier) GetAllTodosByUserID(ctx context.Context, userID uint64) (todos []*pb.Todo, err error) {
	txn := q.db.Txn(false)
	defer txn.Abort()
	todoData, err := txn.Get("todo", "id")
	if err != nil {
		return nil, err
	}
	for todo := todoData.Next(); todo != nil; todo = todoData.Next() {
		todos = append(todos, todo.(*pb.Todo))
	}
	return todos, nil
}
