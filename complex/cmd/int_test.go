package cmd

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/thedevelopnik/go-kit-example/complex/pb"
)

func TestHappyPath(t *testing.T) {
	go Execute()
	j, err := json.Marshal(struct{ userId uint64 }{userId: 1})
	if err != nil {
		t.Error(err)
	}
	res, err := http.Post("http://localhost:3030/gettodos", "application/json", bytes.NewBuffer(j))
	if err != nil {
		t.Error(err)
	}

	todoResponse := &pb.GetTodosResponse{}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status was equal to %d", res.StatusCode)
	}
	err = jsonpb.Unmarshal(res.Body, todoResponse)
	if err != nil {
		t.Error(err)
	}
	res.Body.Close()

	todo0 := *todoResponse.Todos[0]
	if todo0.Id != 1 {
		t.Error("todo id was not correct")
	}
	if strings.Compare(todo0.Name, "go to the grocery store") != 0 {
		t.Error("todo name was not correct")
	}
}
