package tcp

import (
	"github.com/covarity/echo/pkg/cmd/agent"
	"github.com/covarity/echo/pkg/tcp/client"
	"github.com/spf13/cobra"
)

var daemon bool // enable daemon mode

func NewCommand() *cobra.Command {

	c := &cobra.Command{
		Use:   "tcp",
		Short: "interact with a TCP endpoint",
		Run: func(c *cobra.Command, args []string) {
			if daemon {
				println("running daemon mode")
				agent.RunServer()
			} else {
				client.TCPClient()
			}
		},
	}
	c.PersistentFlags().BoolVar(&daemon, "d", false, "run in daemon mode (managment service)")
	return c
}
