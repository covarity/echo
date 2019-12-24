package main

import (
	"os"
	"path/filepath"
	"k8s.io/klog"
	"github.com/covarity/echo/pkg/cmd"
	"github.com/covarity/echo/pkg/cmd/echo"
)

func main() {
	defer klog.Flush()

	baseName := filepath.Base(os.Args[0])

	err := echo.NewCommand(baseName).Execute()
	cmd.CheckError(err)
}
