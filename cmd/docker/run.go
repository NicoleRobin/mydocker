package docker

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Create a docker container",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running docker run ...")
	},
}

var flagInteractive bool
var flagTty bool

func init() {
	runCmd.PersistentFlags().BoolVarP(&flagInteractive, "interaction", "i", false, "Keep STDIN open even if not attached")
	runCmd.PersistentFlags().BoolVarP(&flagTty, "tty", "t", false, "Allocate a pseudo-TTY")
}
