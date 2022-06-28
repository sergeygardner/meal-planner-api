package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sergeygardner/meal-planner-api/application/handler"
	DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"
	DomainService "github.com/sergeygardner/meal-planner-api/domain/service"
	"github.com/sergeygardner/meal-planner-api/ui/rest/response"
	RestService "github.com/sergeygardner/meal-planner-api/ui/rest/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var (
	statusIngredientDeleteSuccess = "the ingredient has been deleted successful"
	statusIngredientDeleteError   = errors.New("the ingredient has not been deleted")
)

func IngredientsInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	ingredients, errorIngredients := handler.IngredientsInfo(&token.UserId, nil)

	if errorIngredients != nil {
		payload = RestService.Error400HandleService(w, errorIngredients)
	} else {
		if ingredients == nil {
			ingredients = []*DomainEntity.Ingredient{}
		}
		payload = &response.IngredientsInfo{Ingredients: ingredients}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func IngredientCreate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	ingredientUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromIngredientUpdate(r.Body)

	if errorJsonDecode != nil {
		payload = RestService.Error400HandleService(w, errorJsonDecode)
	} else {
		ingredient, errorIngredient := handler.IngredientCreate(&token.UserId, &ingredientUpdateDTO)

		if errorIngredient != nil {
			payload = RestService.Error400HandleService(w, errorIngredient)
		} else {
			payload = &response.IngredientInfo{Ingredient: *ingredient}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func IngredientInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	ingredientId, errorIngredientId := uuid.Parse(chi.URLParam(r, "ingredient_id"))

	if errorIngredientId != nil {
		payload = RestService.Error400HandleService(w, errorIngredientId)
	} else {
		ingredient, errorIngredient := handler.IngredientInfo(&ingredientId, &token.UserId, nil)

		if errorIngredient != nil {
			payload = RestService.Error400HandleService(w, errorIngredient)
		} else {
			payload = &response.IngredientInfo{Ingredient: *ingredient}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func IngredientUpdate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	ingredientId, errorIngredientId := uuid.Parse(chi.URLParam(r, "ingredient_id"))

	if errorIngredientId != nil {
		payload = RestService.Error400HandleService(w, errorIngredientId)
	} else {
		ingredientUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromIngredientUpdate(r.Body)

		if errorJsonDecode != nil {
			payload = RestService.Error400HandleService(w, errorJsonDecode)
		} else {
			ingredient, errorIngredient := handler.IngredientUpdate(&ingredientId, &token.UserId, &ingredientUpdateDTO)

			if errorIngredient != nil {
				payload = RestService.Error400HandleService(w, errorIngredient)
			} else {
				payload = &response.IngredientInfo{Ingredient: *ingredient}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func IngredientDelete(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	ingredientId, errorIngredientId := uuid.Parse(chi.URLParam(r, "ingredient_id"))

	if errorIngredientId != nil {
		payload = RestService.Error400HandleService(w, errorIngredientId)
	} else {
		ingredientDeleteStatus, errorIngredientDeleteStatus := handler.IngredientDelete(&ingredientId, &token.UserId)

		if errorIngredientDeleteStatus != nil {
			payload = RestService.Error400HandleService(w, errorIngredientDeleteStatus)
		} else if ingredientDeleteStatus {
			payload = &response.IngredientDelete{Message: statusIngredientDeleteSuccess, Status: http.StatusOK}
		} else {
			payload = RestService.Error400HandleService(w, statusIngredientDeleteError)
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}
