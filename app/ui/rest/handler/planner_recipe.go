package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sergeygardner/meal-planner-api/application/handler"
	DomainAggregate "github.com/sergeygardner/meal-planner-api/domain/aggregate"
	DomainService "github.com/sergeygardner/meal-planner-api/domain/service"
	"github.com/sergeygardner/meal-planner-api/ui/rest/response"
	RestService "github.com/sergeygardner/meal-planner-api/ui/rest/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var (
	statusPlannerRecipeDeleteSuccess = "the planner recipe has been deleted successful"
	statusPlannerRecipeDeleteError   = errors.New("the planner recipe has not been deleted")
)

func PlannerRecipesInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	intervalId, errorIntervalId := uuid.Parse(chi.URLParam(r, "interval_id"))

	if errorIntervalId != nil {
		payload = RestService.Error400HandleService(w, errorIntervalId)
	} else {
		plannerRecipes, errorPlannerRecipes := handler.PlannerRecipesInfo(&token.UserId, &intervalId, nil)

		if errorPlannerRecipes != nil {
			payload = RestService.Error400HandleService(w, errorPlannerRecipes)
		} else {
			if plannerRecipes == nil {
				plannerRecipes = []*DomainAggregate.PlannerRecipe{}
			}
			payload = &response.PlannerRecipesInfo{PlannerRecipes: plannerRecipes}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func PlannerRecipeCreate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	intervalId, errorIntervalId := uuid.Parse(chi.URLParam(r, "interval_id"))

	if errorIntervalId != nil {
		payload = RestService.Error400HandleService(w, errorIntervalId)
	} else {
		plannerRecipeUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromPlannerRecipeUpdate(r.Body)

		if errorJsonDecode != nil {
			payload = RestService.Error400HandleService(w, errorJsonDecode)
		} else {
			plannerRecipe, errorPlannerRecipe := handler.PlannerRecipeCreate(&token.UserId, &intervalId, &plannerRecipeUpdateDTO)

			if errorPlannerRecipe != nil {
				payload = RestService.Error400HandleService(w, errorPlannerRecipe)
			} else {
				payload = &response.PlannerRecipeInfo{PlannerRecipe: *plannerRecipe}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func PlannerRecipeInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	intervalId, errorIntervalId := uuid.Parse(chi.URLParam(r, "interval_id"))

	if errorIntervalId != nil {
		payload = RestService.Error400HandleService(w, errorIntervalId)
	} else {
		plannerRecipeId, errorPlannerRecipeId := uuid.Parse(chi.URLParam(r, "recipe_id"))

		if errorPlannerRecipeId != nil {
			payload = RestService.Error400HandleService(w, errorPlannerRecipeId)
		} else {
			plannerRecipe, errorPlannerRecipe := handler.PlannerRecipeInfo(&plannerRecipeId, &token.UserId, &intervalId, nil)

			if errorPlannerRecipe != nil {
				payload = RestService.Error400HandleService(w, errorPlannerRecipe)
			} else {
				payload = &response.PlannerRecipeInfo{PlannerRecipe: *plannerRecipe}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func PlannerRecipeUpdate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	intervalId, errorIntervalId := uuid.Parse(chi.URLParam(r, "interval_id"))

	if errorIntervalId != nil {
		payload = RestService.Error400HandleService(w, errorIntervalId)
	} else {
		plannerRecipeId, errorPlannerRecipeId := uuid.Parse(chi.URLParam(r, "recipe_id"))

		if errorPlannerRecipeId != nil {
			payload = RestService.Error400HandleService(w, errorPlannerRecipeId)
		} else {
			plannerRecipeUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromPlannerRecipeUpdate(r.Body)

			if errorJsonDecode != nil {
				payload = RestService.Error400HandleService(w, errorJsonDecode)
			} else {
				plannerRecipe, errorPlannerRecipe := handler.PlannerRecipeUpdate(&plannerRecipeId, &token.UserId, &intervalId, &plannerRecipeUpdateDTO)

				if errorPlannerRecipe != nil {
					payload = RestService.Error400HandleService(w, errorPlannerRecipe)
				} else {
					payload = &response.PlannerRecipeInfo{PlannerRecipe: *plannerRecipe}
				}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func PlannerRecipeDelete(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	intervalId, errorIntervalId := uuid.Parse(chi.URLParam(r, "interval_id"))

	if errorIntervalId != nil {
		payload = RestService.Error400HandleService(w, errorIntervalId)
	} else {
		plannerRecipeId, errorPlannerRecipeId := uuid.Parse(chi.URLParam(r, "recipe_id"))

		if errorPlannerRecipeId != nil {
			payload = RestService.Error400HandleService(w, errorPlannerRecipeId)
		} else {
			plannerRecipeDeleteStatus, errorPlannerRecipeDeleteStatus := handler.PlannerRecipeDelete(&plannerRecipeId, &token.UserId, &intervalId)

			if errorPlannerRecipeDeleteStatus != nil {
				payload = RestService.Error400HandleService(w, errorPlannerRecipeDeleteStatus)
			} else if plannerRecipeDeleteStatus {
				payload = &response.PlannerRecipeDelete{Message: statusPlannerRecipeDeleteSuccess, Status: http.StatusOK}
			} else {
				payload = RestService.Error400HandleService(w, statusPlannerRecipeDeleteError)
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}
