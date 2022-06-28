package handler

import (
	Response2 "github.com/sergeygardner/meal-planner-api/ui/rest/response"
	"github.com/sergeygardner/meal-planner-api/ui/rest/service"
	"net/http"
)

var payload Response2.Response

func Index(w http.ResponseWriter, r *http.Request) {
	_ = service.Render(w, r, &Response2.Index{Description: "Description"})
}
