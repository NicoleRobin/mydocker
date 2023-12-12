package docker

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "mydocker is a simple docker implemention"}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A self-sufficient runtime for containers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running docker init ...")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(runCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}
