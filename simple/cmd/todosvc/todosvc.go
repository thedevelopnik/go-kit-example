package main

import (
	"net"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/oklog/oklog/pkg/group"

	"github.com/thedevelopnik/go-kit-example/simple/pkg/todoendpoint"
	"github.com/thedevelopnik/go-kit-example/simple/pkg/todoservice"
	"github.com/thedevelopnik/go-kit-example/simple/pkg/todotransport"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var (
		service     = todoservice.New(logger)
		endpoints   = todoendpoint.New(service, logger)
		httpHandler = todotransport.NewHTTPHandler(endpoints, logger)
	)

	var g group.Group
	{
		httpListener, err := net.Listen("tcp", ":3030")
		if err != nil {
			logger.Log("transport", "HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "HTTP", "addr", ":3030")
			return http.Serve(httpListener, httpHandler)
		}, func(error) {
			httpListener.Close()
		})
	}
	logger.Log("exit", g.Run())
}
