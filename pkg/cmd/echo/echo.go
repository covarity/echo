package echo

import (
	"flag"

	"github.com/covarity/echo/pkg/cmd/cli/tcp"
	"github.com/covarity/echo/pkg/cmd/cli/version"
	"github.com/covarity/echo/pkg/cmd/server"
	"github.com/spf13/cobra"
	"k8s.io/klog"
)

func NewCommand(name string) *cobra.Command {

	// Load the config here so that we can extract features from it.
	// config, err := client.LoadConfig()
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "WARNING: Error reading config file: %v\n", err)
	// }
	c := &cobra.Command{
		Use:   name,
		Short: "Runs Networking interactions",
		Long:  `Can stimulate networking interactions for a number of protocols (TCP,HTTP,GRPC)`,
	}

	c.AddCommand(version.NewCommand())
	c.AddCommand(tcp.NewCommand())
	c.AddCommand(server.NewCommand())

	klog.InitFlags(flag.CommandLine)
	return c

}
