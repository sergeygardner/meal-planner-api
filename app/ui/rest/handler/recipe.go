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
	statusRecipeDeleteSuccess = "the recipe has been deleted successful"
	statusRecipeDeleteError   = errors.New("the recipe has not been deleted")
)

func RecipesInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	recipes, errorRecipe := handler.RecipesInfo(&token.UserId, nil)

	if errorRecipe != nil {
		payload = RestService.Error400HandleService(w, errorRecipe)
	} else {
		payload = &response.RecipesInfo{Recipes: recipes}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func RecipeCreate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	recipeUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromRecipeUpdate(r.Body)

	if errorJsonDecode != nil {
		payload = RestService.Error400HandleService(w, errorJsonDecode)
	} else {
		recipe, errorRecipeCreate := handler.RecipeCreate(&token.UserId, &recipeUpdateDTO)

		if errorRecipeCreate != nil {
			payload = RestService.Error400HandleService(w, errorRecipeCreate)
		} else {
			payload = &response.RecipeInfo{Recipe: *recipe}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func RecipeInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	recipeId, errorRecipeId := uuid.Parse(chi.URLParam(r, "recipe_id"))

	if errorRecipeId != nil {
		payload = RestService.Error400HandleService(w, errorRecipeId)
	} else {
		recipe, errorRecipe := handler.RecipeInfo(&recipeId, &token.UserId, nil)

		if errorRecipe != nil {
			payload = RestService.Error400HandleService(w, errorRecipe)
		} else {
			payload = &response.RecipeInfo{Recipe: *recipe}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func RecipeUpdate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	recipeId, errorRecipeId := uuid.Parse(chi.URLParam(r, "recipe_id"))

	if errorRecipeId != nil {
		payload = RestService.Error400HandleService(w, errorRecipeId)
	} else {
		recipeUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromRecipeUpdate(r.Body)

		if errorJsonDecode != nil {
			payload = RestService.Error400HandleService(w, errorJsonDecode)
		} else {
			recipe, errorRecipe := handler.RecipeUpdate(&recipeId, &token.UserId, &recipeUpdateDTO)

			if errorRecipe != nil {
				payload = RestService.Error400HandleService(w, errorRecipe)
			} else {
				payload = &response.RecipeInfo{Recipe: *recipe}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func RecipeDelete(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	recipeId, errorRecipeId := uuid.Parse(chi.URLParam(r, "recipe_id"))

	if errorRecipeId != nil {
		payload = RestService.Error400HandleService(w, errorRecipeId)
	} else {
		recipeDeleteStatus, errorRecipeDeleteStatus := handler.RecipeDelete(&recipeId, &token.UserId)

		if errorRecipeDeleteStatus != nil {
			payload = RestService.Error400HandleService(w, errorRecipeDeleteStatus)
		} else if recipeDeleteStatus {
			payload = &response.RecipeDelete{Message: statusRecipeDeleteSuccess, Status: http.StatusOK}
		} else {
			payload = RestService.Error400HandleService(w, statusRecipeDeleteError)
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}
