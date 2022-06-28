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
	statusRecipeMeasureDeleteSuccess = "the recipe measure has been deleted successful"
	statusRecipeMeasureDeleteError   = errors.New("the recipe measure has not been deleted")
)

func RecipeMeasuresInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	ingredientId, errorIngredientId := uuid.Parse(chi.URLParam(r, "ingredient_id"))

	if errorIngredientId != nil {
		payload = RestService.Error400HandleService(w, errorIngredientId)
	} else {
		recipeMeasures, errorRecipeMeasure := handler.RecipeMeasuresInfo(&token.UserId, &ingredientId, nil)

		if errorRecipeMeasure != nil {
			payload = RestService.Error400HandleService(w, errorRecipeMeasure)
		} else {
			if recipeMeasures == nil {
				recipeMeasures = []*DomainAggregate.RecipeMeasure{}
			}
			payload = &response.RecipeMeasuresInfo{Measures: recipeMeasures}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func RecipeMeasureCreate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	ingredientId, errorIngredientId := uuid.Parse(chi.URLParam(r, "ingredient_id"))

	if errorIngredientId != nil {
		payload = RestService.Error400HandleService(w, errorIngredientId)
	} else {
		recipeMeasureUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromRecipeMeasureUpdate(r.Body)

		if errorJsonDecode != nil {
			payload = RestService.Error400HandleService(w, errorJsonDecode)
		} else {
			recipeMeasure, errorAuthRegister := handler.RecipeMeasureCreate(&token.UserId, &ingredientId, &recipeMeasureUpdateDTO)

			if errorAuthRegister != nil {
				payload = RestService.Error400HandleService(w, errorAuthRegister)
			} else {
				payload = &response.RecipeMeasureInfo{RecipeMeasure: *recipeMeasure}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func RecipeMeasureInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	ingredientId, errorIngredientId := uuid.Parse(chi.URLParam(r, "ingredient_id"))

	if errorIngredientId != nil {
		payload = RestService.Error400HandleService(w, errorIngredientId)
	} else {
		recipeMeasureId, errorRecipeMeasureId := uuid.Parse(chi.URLParam(r, "measure_id"))

		if errorRecipeMeasureId != nil {
			payload = RestService.Error400HandleService(w, errorRecipeMeasureId)
		} else {
			recipeMeasure, errorRecipeMeasure := handler.RecipeMeasureInfo(&recipeMeasureId, &token.UserId, &ingredientId, nil)

			if errorRecipeMeasure != nil {
				payload = RestService.Error400HandleService(w, errorRecipeMeasure)
			} else {
				payload = &response.RecipeMeasureInfo{RecipeMeasure: *recipeMeasure}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func RecipeMeasureUpdate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	ingredientId, errorIngredientId := uuid.Parse(chi.URLParam(r, "ingredient_id"))

	if errorIngredientId != nil {
		payload = RestService.Error400HandleService(w, errorIngredientId)
	} else {
		recipeMeasureId, errorRecipeMeasureId := uuid.Parse(chi.URLParam(r, "measure_id"))

		if errorRecipeMeasureId != nil {
			payload = RestService.Error400HandleService(w, errorRecipeMeasureId)
		} else {
			recipeMeasureUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromRecipeMeasureUpdate(r.Body)

			if errorJsonDecode != nil {
				payload = RestService.Error400HandleService(w, errorJsonDecode)
			} else {
				recipeMeasure, errorRecipeMeasure := handler.RecipeMeasureUpdate(&recipeMeasureId, &token.UserId, &ingredientId, &recipeMeasureUpdateDTO)

				if errorRecipeMeasure != nil {
					payload = RestService.Error400HandleService(w, errorRecipeMeasure)
				} else {
					payload = &response.RecipeMeasureInfo{RecipeMeasure: *recipeMeasure}
				}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func RecipeMeasureDelete(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	ingredientId, errorIngredientId := uuid.Parse(chi.URLParam(r, "ingredient_id"))

	if errorIngredientId != nil {
		payload = RestService.Error400HandleService(w, errorIngredientId)
	} else {
		recipeMeasureId, errorRecipeMeasureId := uuid.Parse(chi.URLParam(r, "measure_id"))

		if errorRecipeMeasureId != nil {
			payload = RestService.Error400HandleService(w, errorRecipeMeasureId)
		} else {
			recipeMeasureDeleteStatus, errorRecipeMeasureDeleteStatus := handler.RecipeMeasureDelete(&recipeMeasureId, &token.UserId, &ingredientId)

			if errorRecipeMeasureDeleteStatus != nil {
				payload = RestService.Error400HandleService(w, errorRecipeMeasureDeleteStatus)
			} else if recipeMeasureDeleteStatus {
				payload = &response.RecipeMeasureDelete{Message: statusRecipeMeasureDeleteSuccess, Status: http.StatusOK}
			} else {
				payload = RestService.Error400HandleService(w, statusRecipeMeasureDeleteError)
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}
