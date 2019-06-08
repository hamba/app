package main

import (
	"log"
	"os"

	"github.com/hamba/cmd"
	"gopkg.in/urfave/cli.v2"
)

import _ "github.com/joho/godotenv/autoload"

const (
	flagSomeFlag = "some.flag"
)

var version = "¯\\_(ツ)_/¯"

var commands = []*cli.Command{
	{
		Name:  "server",
		Usage: "Run the ren HTTP server",
		Flags: cmd.Flags{
			&cli.StringFlag{
				Name:    flagSomeFlag,
				Usage:   "Some flag for application configuration.",
				EnvVars: []string{"SOME_FLAG"},
			},
		}.Merge(cmd.CommonFlags, cmd.ServerFlags),
		Action: runServer,
	},
}

func main() {
	app := &cli.App{
		Name:     "app",
		Version:  version,
		Commands: commands,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
