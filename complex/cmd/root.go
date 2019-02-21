package cmd

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/gospotcheck/go-foundation/asgard/conf"
	"github.com/oklog/oklog/pkg/group"
	"github.com/spf13/cobra"
	"github.com/thedevelopnik/go-kit-example/complex/pkg/tododatabase"
	"github.com/thedevelopnik/go-kit-example/complex/pkg/todoendpoint"
	"github.com/thedevelopnik/go-kit-example/complex/pkg/todoservice"
	"github.com/thedevelopnik/go-kit-example/complex/pkg/todotransport"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Todo serves up a SPA and provides an API for retriving todos",
	Run:   run,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	config := &todoservice.Config{}
	{
		if err := conf.LoadConfigWithCobra(config, cmd); err != nil {
			logger.Log("config", "failed to load", "err", err)
		}
	}

	db := tododatabase.CreateDB()

	var (
		querier     = tododatabase.New(db, logger)
		service     = todoservice.New(querier, logger)
		endpoints   = todoendpoint.New(service, logger)
		httpHandler = todotransport.NewHTTPHandler(endpoints, logger)
	)

	var g group.Group
	{
		httpPort := fmt.Sprintf(":%d", config.HTTPPort)
		httpListener, err := net.Listen("tcp", httpPort)
		if err != nil {
			logger.Log("transport", "HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "HTTP", "addr", httpPort)
			return http.Serve(httpListener, httpHandler)
		}, func(error) {
			httpListener.Close()
		})
	}
	logger.Log("exit", g.Run())
}
