package main

import (
	"github.com/urfave/cli"
)

var (
	initCommand = cli.Command{}
	runCommand  = cli.Command{}
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		initCommand,
		runCommand,
	}
}
