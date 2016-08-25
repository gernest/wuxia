package base

import (
	"net/http"

	"github.com/gernest/wuxia/models"
	"github.com/gernest/wuxia/views"
	"gopkg.in/urfave/cli.v2"
)

func Server() *cli.Command {
	return &cli.Command{
		Name:    "serve",
		Aliases: []string{"s"},
		Usage:   "start we server for wuxia",
		Action: func(c *cli.Context) error {
			m := &models.Context{}
			v, err := views.New("wuxia")
			if err != nil {
				return err
			}
			m.View = v
			h := App(m)
			return http.ListenAndServe(":8080", h)
		},
	}
}
