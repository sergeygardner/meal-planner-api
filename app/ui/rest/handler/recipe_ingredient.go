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
	statusRecipeIngredientDeleteSuccess = "the recipe ingredient has been deleted successful"
	statusRecipeIngredientDeleteError   = errors.New("the recipe ingredient has not been deleted")
)

func RecipeIngredientsInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	recipeId, errorRecipeId := uuid.Parse(chi.URLParam(r, "recipe_id"))

	if errorRecipeId != nil {
		payload = RestService.Error400HandleService(w, errorRecipeId)
	} else {
		recipeIngredients, errorRecipeIngredient := handler.RecipeIngredientsInfo(&token.UserId, &recipeId, nil)

		if errorRecipeIngredient != nil {
			payload = RestService.Error400HandleService(w, errorRecipeIngredient)
		} else {
			if recipeIngredients == nil {
				recipeIngredients = []*DomainAggregate.RecipeIngredient{}
			}
			payload = &response.RecipeIngredientsInfo{Ingredients: recipeIngredients}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func RecipeIngredientCreate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	recipeId, errorRecipeId := uuid.Parse(chi.URLParam(r, "recipe_id"))

	if errorRecipeId != nil {
		payload = RestService.Error400HandleService(w, errorRecipeId)
	} else {
		recipeIngredientUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromRecipeIngredientUpdate(r.Body)

		if errorJsonDecode != nil {
			payload = RestService.Error400HandleService(w, errorJsonDecode)
		} else {
			recipeIngredient, errorAuthRegister := handler.RecipeIngredientCreate(&token.UserId, &recipeId, &recipeIngredientUpdateDTO)

			if errorAuthRegister != nil {
				payload = RestService.Error400HandleService(w, errorAuthRegister)
			} else {
				payload = &response.RecipeIngredientInfo{RecipeIngredient: *recipeIngredient}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func RecipeIngredientInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	recipeId, errorRecipeId := uuid.Parse(chi.URLParam(r, "recipe_id"))

	if errorRecipeId != nil {
		payload = RestService.Error400HandleService(w, errorRecipeId)
	} else {
		recipeIngredientId, errorRecipeIngredientId := uuid.Parse(chi.URLParam(r, "ingredient_id"))

		if errorRecipeIngredientId != nil {
			payload = RestService.Error400HandleService(w, errorRecipeIngredientId)
		} else {
			recipeIngredient, errorRecipeIngredient := handler.RecipeIngredientInfo(&recipeIngredientId, &token.UserId, &recipeId, nil)

			if errorRecipeIngredient != nil {
				payload = RestService.Error400HandleService(w, errorRecipeIngredient)
			} else {
				payload = &response.RecipeIngredientInfo{RecipeIngredient: *recipeIngredient}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func RecipeIngredientUpdate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	recipeId, errorRecipeId := uuid.Parse(chi.URLParam(r, "recipe_id"))

	if errorRecipeId != nil {
		payload = RestService.Error400HandleService(w, errorRecipeId)
	} else {
		recipeIngredientId, errorRecipeIngredientId := uuid.Parse(chi.URLParam(r, "ingredient_id"))

		if errorRecipeIngredientId != nil {
			payload = RestService.Error400HandleService(w, errorRecipeIngredientId)
		} else {
			recipeIngredientUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromRecipeIngredientUpdate(r.Body)

			if errorJsonDecode != nil {
				payload = RestService.Error400HandleService(w, errorJsonDecode)
			} else {
				recipeIngredient, errorRecipeIngredient := handler.RecipeIngredientUpdate(&recipeIngredientId, &token.UserId, &recipeId, &recipeIngredientUpdateDTO)

				if errorRecipeIngredient != nil {
					payload = RestService.Error400HandleService(w, errorRecipeIngredient)
				} else {
					payload = &response.RecipeIngredientInfo{RecipeIngredient: *recipeIngredient}
				}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func RecipeIngredientDelete(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	recipeId, errorRecipeId := uuid.Parse(chi.URLParam(r, "recipe_id"))

	if errorRecipeId != nil {
		payload = RestService.Error400HandleService(w, errorRecipeId)
	} else {
		recipeIngredientId, errorRecipeIngredientId := uuid.Parse(chi.URLParam(r, "ingredient_id"))

		if errorRecipeIngredientId != nil {
			payload = RestService.Error400HandleService(w, errorRecipeIngredientId)
		} else {
			recipeIngredientDeleteStatus, errorRecipeIngredientDeleteStatus := handler.RecipeIngredientDelete(&recipeIngredientId, &token.UserId, &recipeId)

			if errorRecipeIngredientDeleteStatus != nil {
				payload = RestService.Error400HandleService(w, errorRecipeIngredientDeleteStatus)
			} else if recipeIngredientDeleteStatus {
				payload = &response.RecipeIngredientDelete{Message: statusRecipeIngredientDeleteSuccess, Status: http.StatusOK}
			} else {
				payload = RestService.Error400HandleService(w, statusRecipeIngredientDeleteError)
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}
