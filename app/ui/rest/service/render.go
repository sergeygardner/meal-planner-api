package service

import (
	"github.com/go-chi/render"
	"github.com/sergeygardner/meal-planner-api/ui/rest/response"
	"net/http"
)

func Render(w http.ResponseWriter, r *http.Request, payload response.Response) error {
	w.WriteHeader(payload.GetStatus())

	return render.Render(w, r, payload)
}
