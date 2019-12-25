package server

import (
	"github.com/covarity/echo/pkg/cmd/server/tcp"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {

	// Load the config here so that we can extract features from it.
	// config, err := client.LoadConfig()
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "WARNING: Error reading config file: %v\n", err)
	// }
	c := &cobra.Command{
		Use:   "serve",
		Short: "Runs Networking interactions",
		Long:  `Can stimulate networking interactions for a number of protocols (TCP,HTTP,GRPC)`,
	}

	c.AddCommand(tcp.NewCommand())

	return c

}
