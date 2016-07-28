package api

import (
	"net/http"

	"github.com/gernest/wuxia/models"
)

type Api struct {
	Ctx *models.Context
}

func (a *Api) Build(w http.ResponseWriter, r http.Request) {
}

func (a *Api) Brogress(w http.ResponseWriter, r http.Request) {
}
