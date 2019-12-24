package version

import (
	"fmt"
	"github.com/covarity/echo/pkg/buildinfo"
	"github.com/spf13/cobra"
	"io"
	"os"
)

func NewCommand() *cobra.Command {

	c := &cobra.Command{
		Use:   "version",
		Short: "Print the velero version and associated image",
		Run: func(c *cobra.Command, args []string) {
			printVersion(os.Stdout)
		},
	}
	return c
}

func printVersion(w io.Writer) {
	fmt.Fprintln(w, "Client:")
	fmt.Fprintf(w, "\tVersion: %s\n", buildinfo.Version)
	fmt.Fprintf(w, "\tGit commit: %s\n", buildinfo.FormattedGitSHA())
}
