package main

import (
	"os"

	"github.com/gernest/wuxia/base"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := cli.App{}
	app.Name = "wuxia : The righteous static web solution"
	app.Usage = "create , build , deploy and scale your static websites"
	app.Commands = []*cli.Command{
		base.Server(),
	}
	app.Run(os.Args)
}
