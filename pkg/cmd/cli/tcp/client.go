package tcp

import (
	"github.com/covarity/echo/pkg/tcp/client"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {

	c := &cobra.Command{
		Use:   "tcp",
		Short: "interact with a TCP endpoint",
		Run: func(c *cobra.Command, args []string) {
			client.TCPClient()
		},
	}
	return c
}
