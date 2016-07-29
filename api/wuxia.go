package api

import (
	"net/http"

	"github.com/gernest/wuxia/models"
	"github.com/oleiade/lane"
)

type Api struct {
	Ctx  *models.Context
	qeue *lane.PQueue
}

func (a *Api) Build(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ba, ok := ctx.Value(models.CtxBuild).(models.BuildArtifact)
	if !ok {
		// nothing to build
	}
	a.reportBuild(ba, w, r)
}

func (a *Api) Progress(w http.ResponseWriter, r http.Request) {
}

func (a *Api) reportBuild(ba models.BuildArtifact, w http.ResponseWriter, r *http.Request) {
}
