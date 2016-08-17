package api

import (
	"net/http"

	"github.com/gernest/wuxia/models"
)

type Data map[string]interface{}

//Home renders home page
func Home(ctx *models.Context, w http.ResponseWriter, r *http.Request) {
	data := make(Data)
	data["config"] = ctx.Config
	err := ctx.HTML("home.tpl", data, w, http.StatusOK)
	if err != nil {
		// log this?
	}
}
