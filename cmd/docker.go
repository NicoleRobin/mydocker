package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "docker"}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running docker run ...")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}
