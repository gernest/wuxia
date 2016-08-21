package base

import (
	"net/http"

	"github.com/gernest/wuxia/data"
	"github.com/gernest/wuxia/models"
	"github.com/gorilla/mux"
)

type Data map[string]interface{}

//Home renders home page
func Home(ctx *models.Context, w http.ResponseWriter, r *http.Request) {
	data := make(Data)
	data["config"] = ctx.Config
	err := ctx.HTML("home.html", data, w, http.StatusOK)
	if err != nil {
		// log this?
	}
}

func Handle(ctx *models.Context, h func(*models.Context, http.ResponseWriter,
	*http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		h(ctx, w, r)
	}
}

func App(ctx *models.Context) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", Handle(ctx, Home))
	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(data.HTTPAsset())))
	return r
}
