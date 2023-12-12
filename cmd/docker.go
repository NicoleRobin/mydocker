package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "mydocker is a simple docker implemention"}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running docker init ...")
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Create a docker container",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running docker run ...")
	},
}

func init() {
	runCmd.PersistentFlags().StringP("interactive", "i", "", "Keep STDIN open even if not attached")
	runCmd.PersistentFlags().StringP("tty", "t", "", "Allocate a pseudo-TTY")

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(runCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}
