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
	statusRecipeCategoryDeleteSuccess = "the recipe category has been deleted successful"
	statusRecipeCategoryDeleteError   = errors.New("the recipe category has not been deleted")
)

func RecipeCategoriesInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	recipeId, errorRecipeId := uuid.Parse(chi.URLParam(r, "recipe_id"))

	if errorRecipeId != nil {
		payload = RestService.Error400HandleService(w, errorRecipeId)
	} else {
		recipeCategories, errorRecipeCategory := handler.RecipeCategoriesInfo(&token.UserId, &recipeId, nil)

		if errorRecipeCategory != nil {
			payload = RestService.Error400HandleService(w, errorRecipeCategory)
		} else {
			if recipeCategories == nil {
				recipeCategories = []*DomainAggregate.RecipeCategory{}
			}
			payload = &response.RecipeCategoriesInfo{Categories: recipeCategories}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func RecipeCategoryCreate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	recipeId, errorRecipeId := uuid.Parse(chi.URLParam(r, "recipe_id"))

	if errorRecipeId != nil {
		payload = RestService.Error400HandleService(w, errorRecipeId)
	} else {
		recipeCategoryUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromRecipeCategoryUpdate(r.Body)

		if errorJsonDecode != nil {
			payload = RestService.Error400HandleService(w, errorJsonDecode)
		} else {

			recipeCategory, errorRecipeCategory := handler.RecipeCategoryCreate(&token.UserId, &recipeId, &recipeCategoryUpdateDTO)

			if errorRecipeCategory != nil {
				payload = RestService.Error400HandleService(w, errorRecipeCategory)
			} else {
				payload = &response.RecipeCategoryInfo{RecipeCategory: *recipeCategory}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func RecipeCategoryInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	recipeId, errorRecipeId := uuid.Parse(chi.URLParam(r, "recipe_id"))

	if errorRecipeId != nil {
		payload = RestService.Error400HandleService(w, errorRecipeId)
	} else {
		recipeCategoryId, errorRecipeCategoryId := uuid.Parse(chi.URLParam(r, "category_id"))

		if errorRecipeCategoryId != nil {
			payload = RestService.Error400HandleService(w, errorRecipeCategoryId)
		} else {
			recipeCategory, errorRecipeCategory := handler.RecipeCategoryInfo(&recipeCategoryId, &token.UserId, &recipeId, nil)

			if errorRecipeCategory != nil {
				payload = RestService.Error400HandleService(w, errorRecipeCategory)
			} else {
				payload = &response.RecipeCategoryInfo{RecipeCategory: *recipeCategory}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func RecipeCategoryUpdate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	recipeId, errorRecipeId := uuid.Parse(chi.URLParam(r, "recipe_id"))

	if errorRecipeId != nil {
		payload = RestService.Error400HandleService(w, errorRecipeId)
	} else {
		recipeCategoryId, errorRecipeCategoryId := uuid.Parse(chi.URLParam(r, "category_id"))

		if errorRecipeCategoryId != nil {
			payload = RestService.Error400HandleService(w, errorRecipeCategoryId)
		} else {
			recipeCategoryUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromRecipeCategoryUpdate(r.Body)

			if errorJsonDecode != nil {
				payload = RestService.Error400HandleService(w, errorJsonDecode)
			} else {
				recipeCategory, errorRecipeCategory := handler.RecipeCategoryUpdate(&recipeCategoryId, &token.UserId, &recipeId, &recipeCategoryUpdateDTO)

				if errorRecipeCategory != nil {
					payload = RestService.Error400HandleService(w, errorRecipeCategory)
				} else {
					payload = &response.RecipeCategoryInfo{RecipeCategory: *recipeCategory}
				}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func RecipeCategoryDelete(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	recipeId, errorRecipeId := uuid.Parse(chi.URLParam(r, "recipe_id"))

	if errorRecipeId != nil {
		payload = RestService.Error400HandleService(w, errorRecipeId)
	} else {
		recipeCategoryId, errorRecipeCategoryId := uuid.Parse(chi.URLParam(r, "category_id"))

		if errorRecipeCategoryId != nil {
			payload = RestService.Error400HandleService(w, errorRecipeCategoryId)
		} else {
			recipeCategoryDeleteStatus, errorRecipeCategoryDeleteStatus := handler.RecipeCategoryDelete(&recipeCategoryId, &token.UserId, &recipeId)

			if errorRecipeCategoryDeleteStatus != nil {
				payload = RestService.Error400HandleService(w, errorRecipeCategoryDeleteStatus)
			} else if recipeCategoryDeleteStatus {
				payload = &response.RecipeCategoryDelete{Message: statusRecipeCategoryDeleteSuccess, Status: http.StatusOK}
			} else {
				payload = RestService.Error400HandleService(w, statusRecipeCategoryDeleteError)
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}
