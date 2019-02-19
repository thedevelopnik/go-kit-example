package todoservice

import (
	"context"
	"errors"
	"testing"

	"github.com/thedevelopnik/go-kit-example/complex/pb"
)

type mockQuerier struct{}

func (m mockQuerier) GetAllTodosByUserID(ctx context.Context, userID uint64) (todos []*pb.Todo, err error) {
	if userID != 1 {
		return nil, errors.New("ya done screwed up")
	}
	todo1 := &pb.Todo{
		Id:   1,
		Name: "go to the grocery store",
	}
	todo2 := &pb.Todo{
		Id:   2,
		Name: "vacuum",
	}
	todo3 := &pb.Todo{
		Id:   3,
		Name: "learn go",
	}
	return []*pb.Todo{todo1, todo2, todo3}, nil
}

func TestGetTodos(t *testing.T) {
	querier := &mockQuerier{}
	svc := NewBasicService(querier)
	ctx := context.Background()
	todos, err := svc.GetTodos(ctx, 1)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", todos)
}
