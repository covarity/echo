package tcp

import (
	"github.com/covarity/echo/pkg/tcp/server"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {

	c := &cobra.Command{
		Use:   "tcp",
		Short: "create a tcp server",
		Run: func(c *cobra.Command, args []string) {
			server.TCPServer()
		},
	}
	return c
}
