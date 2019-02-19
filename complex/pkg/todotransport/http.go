package todotransport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/thedevelopnik/go-kit-example/complex/pb"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/thedevelopnik/go-kit-example/complex/pkg/todoendpoint"
)

func NewHTTPHandler(endpoints todoendpoint.Set, logger log.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerErrorLogger(logger),
	}

	m := http.NewServeMux()
	m.Handle("/gettodos", accessControl(httptransport.NewServer(
		endpoints.GetTodosEndpoint,
		decodeHttpGetTodosRequest,
		newHttpProtoMessageResponseEncoder(logger),
		options...,
	)))
	return m
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

func err2code(err error) int {
	return 400
}

type errorWrapper struct {
	Error string `json:"error"`
}

func decodeHttpGetTodosRequest(_ context.Context, r *http.Request) (interface{}, error) {
	todoRequest := &pb.GetTodosRequest{}

	if err := jsonpb.Unmarshal(r.Body, todoRequest); err != nil {
		return nil, err
	}

	return todoRequest, nil
}

func newHttpProtoMessageResponseEncoder(logger log.Logger) func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if res, valid := response.(proto.Message); valid {
			m := jsonpb.Marshaler{}
			return m.Marshal(w, res)
		}

		logger.Log(logger.Log("Encoding WARNING", "Could not JSON encode response because it is not a proto.Message"))
		return errors.New("could not json encode response")
	}
}
