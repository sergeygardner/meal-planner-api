package handler

import (
	"github.com/sergeygardner/meal-planner-api/ui/rest/response"
	"github.com/sergeygardner/meal-planner-api/ui/rest/service"
	"net/http"
)

var (
	version  = "v1"
	versions = []string{version}
)

func API(w http.ResponseWriter, r *http.Request) {
	_ = service.Render(w, r, &response.API{Version: version, Versions: versions})
}

//func APIPost(w http.ResponseWriter, r *http.Request) {
//	_ = render.Render(w, r, &response.API{Version: version, Versions: versions})
//}
