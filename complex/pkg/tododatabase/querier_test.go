package tododatabase

import (
	"context"
	"testing"
)

func TestGetAllTodosByUserID(t *testing.T) {
	db := CreateDB()
	querier := newBasicQuerier(db)
	ctx := context.Background()
	todos, err := querier.GetAllTodosByUserID(ctx, 1)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", todos)
}
