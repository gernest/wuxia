package main

import (
	"net/http"
	"os"

	"github.com/codegangsta/cli"
	"github.com/gernest/wuxia/base"
	"github.com/gernest/wuxia/models"
	"github.com/gernest/wuxia/views"
)

func main() {
	app := cli.NewApp()
	app.Name = "wuxia : The righteous static web solution"
	app.Usage = "create , build , deploy and scale your static websites"
	app.Commands = []cli.Command{
		Server(),
	}
	app.Run(os.Args)
}

func Server() cli.Command {
	return cli.Command{
		Name:      "serve",
		ShortName: "s",
		Usage:     "start we server for wuxia",
		Action: func(c *cli.Context) error {
			m := &models.Context{}
			v, err := views.New("wuxia")
			if err != nil {
				return err
			}
			m.View = v
			h := base.App(m)
			return http.ListenAndServe(":8080", h)
		},
	}
}
