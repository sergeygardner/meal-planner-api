package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sergeygardner/meal-planner-api/application/handler"
	DomainService "github.com/sergeygardner/meal-planner-api/domain/service"
	"github.com/sergeygardner/meal-planner-api/ui/rest/response"
	RestService "github.com/sergeygardner/meal-planner-api/ui/rest/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var (
	statusRecipeProcessDeleteSuccess = "the recipe process has been deleted successful"
	statusRecipeProcessDeleteError   = errors.New("the recipe process has not been deleted")
)

func RecipeProcessesInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	recipeId, errorRecipeId := uuid.Parse(chi.URLParam(r, "recipe_id"))

	if errorRecipeId != nil {
		payload = RestService.Error400HandleService(w, errorRecipeId)
	} else {
		recipeProcessesAggregate, errorRecipeProcessAggregate := handler.RecipeProcessesInfo(&token.UserId, &recipeId, nil)

		if errorRecipeProcessAggregate != nil {
			payload = RestService.Error400HandleService(w, errorRecipeProcessAggregate)
		} else {
			payload = &response.RecipeProcessAggregatesInfo{Processes: recipeProcessesAggregate}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func RecipeProcessCreate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	recipeId, errorRecipeId := uuid.Parse(chi.URLParam(r, "recipe_id"))

	if errorRecipeId != nil {
		payload = RestService.Error400HandleService(w, errorRecipeId)
	} else {
		recipeProcessUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromRecipeProcessUpdate(r.Body)

		if errorJsonDecode != nil {
			payload = RestService.Error400HandleService(w, errorJsonDecode)
		} else {
			recipeProcessAggregate, errorAuthRegister := handler.RecipeProcessCreate(&token.UserId, &recipeId, &recipeProcessUpdateDTO)

			if errorAuthRegister != nil {
				payload = RestService.Error400HandleService(w, errorAuthRegister)
			} else {
				payload = &response.RecipeProcessAggregateInfo{RecipeProcess: recipeProcessAggregate}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func RecipeProcessInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	recipeId, errorRecipeId := uuid.Parse(chi.URLParam(r, "recipe_id"))

	if errorRecipeId != nil {
		payload = RestService.Error400HandleService(w, errorRecipeId)
	} else {
		recipeProcessId, errorRecipeProcessId := uuid.Parse(chi.URLParam(r, "process_id"))

		if errorRecipeProcessId != nil {
			payload = RestService.Error400HandleService(w, errorRecipeProcessId)
		} else {
			recipeProcessAggregate, errorRecipeProcessAggregate := handler.RecipeProcessInfo(&recipeProcessId, &token.UserId, &recipeId, nil)

			if errorRecipeProcessAggregate != nil {
				payload = RestService.Error400HandleService(w, errorRecipeProcessAggregate)
			} else {
				payload = &response.RecipeProcessAggregateInfo{RecipeProcess: recipeProcessAggregate}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func RecipeProcessUpdate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	recipeId, errorRecipeId := uuid.Parse(chi.URLParam(r, "recipe_id"))

	if errorRecipeId != nil {
		payload = RestService.Error400HandleService(w, errorRecipeId)
	} else {
		recipeProcessId, errorRecipeProcessId := uuid.Parse(chi.URLParam(r, "process_id"))

		if errorRecipeProcessId != nil {
			payload = RestService.Error400HandleService(w, errorRecipeProcessId)
		} else {
			recipeProcessUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromRecipeProcessUpdate(r.Body)

			if errorJsonDecode != nil {
				payload = RestService.Error400HandleService(w, errorJsonDecode)
			} else {
				recipeProcess, errorRecipeProcess := handler.RecipeProcessUpdate(&recipeProcessId, &token.UserId, &recipeId, &recipeProcessUpdateDTO)

				if errorRecipeProcess != nil {
					payload = RestService.Error400HandleService(w, errorRecipeProcess)
				} else {
					payload = &response.RecipeProcessAggregateInfo{RecipeProcess: recipeProcess}
				}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func RecipeProcessDelete(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	recipeId, errorRecipeId := uuid.Parse(chi.URLParam(r, "recipe_id"))

	if errorRecipeId != nil {
		payload = RestService.Error400HandleService(w, errorRecipeId)
	} else {
		recipeProcessId, errorRecipeProcessId := uuid.Parse(chi.URLParam(r, "process_id"))

		if errorRecipeProcessId != nil {
			payload = RestService.Error400HandleService(w, errorRecipeProcessId)
		} else {
			recipeProcessDeleteStatus, errorRecipeProcessDeleteStatus := handler.RecipeProcessDelete(&recipeProcessId, &token.UserId, &recipeId)

			if errorRecipeProcessDeleteStatus != nil {
				payload = RestService.Error400HandleService(w, errorRecipeProcessDeleteStatus)
			} else if recipeProcessDeleteStatus {
				payload = &response.RecipeProcessDelete{Message: statusRecipeProcessDeleteSuccess, Status: http.StatusOK}
			} else {
				payload = RestService.Error400HandleService(w, statusRecipeProcessDeleteError)
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}
