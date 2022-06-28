package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	UiService "github.com/sergeygardner/meal-planner-api/ui/service"
)

func getParentId(routerContext *chi.Context, exclude string) (*uuid.UUID, error) {
	return UiService.GetParentId(routerContext.URLParams.Keys, routerContext.URLParams.Values, exclude)
}
