package main

import (
	"github.com/urfave/cli"
	"gitlab.azbit.cn/web/golang-framework/cmd/server"
	"gitlab.azbit.cn/web/golang-framework/cmd/tool"
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
