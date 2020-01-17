package main

import (
	"github.com/urfave/cli"
	"golang-framework/cmd/server"
	"golang-framework/cmd/tool"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "golang-framework"
	app.Commands = []cli.Command{
		server.Server,
		tool.InitDB,
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
