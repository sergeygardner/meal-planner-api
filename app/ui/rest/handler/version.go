package handler

import (
	"github.com/sergeygardner/meal-planner-api/ui/rest/response"
	"github.com/sergeygardner/meal-planner-api/ui/rest/service"
	"net/http"
)

func Version(w http.ResponseWriter, r *http.Request) {
	_ = service.Render(w, r, &response.Version{Version: version})
}
