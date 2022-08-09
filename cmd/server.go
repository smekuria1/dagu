package main

import (
	"os"

	"github.com/urfave/cli/v2"
	"github.com/yohamta/dagu/internal/admin"
	"github.com/yohamta/dagu/internal/utils"
)

func newServerCommand() *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "dagu server",
		Action: func(c *cli.Context) error {
			return startServer(globalConfig)
		},
	}
}

func startServer(cfg *admin.Config) error {
	server := admin.NewServer(cfg)

	listenSignals(func(sig os.Signal) {
		server.Shutdown()
	})

	err := server.Serve()
	utils.LogErr("running server", err)
	return err
}
