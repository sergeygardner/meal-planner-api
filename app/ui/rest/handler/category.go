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
	statusCategoryDeleteSuccess = "the category has been deleted successful"
	statusCategoryDeleteError   = errors.New("the category has not been deleted")
)

func CategoriesInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	categories, errorCategories := handler.CategoriesInfo(&token.UserId, nil)

	if errorCategories != nil {
		payload = RestService.Error400HandleService(w, errorCategories)
	} else {
		if categories == nil {
			categories = []*DomainAggregate.Category{}
		}
		payload = &response.CategoriesInfo{Categories: categories}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func CategoryCreate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	categoryUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromCategoryUpdate(r.Body)

	if errorJsonDecode != nil {
		payload = RestService.Error400HandleService(w, errorJsonDecode)
	} else {
		category, errorCategory := handler.CategoryCreate(&token.UserId, &categoryUpdateDTO)

		if errorCategory != nil {
			payload = RestService.Error400HandleService(w, errorCategory)
		} else {
			payload = &response.CategoryInfo{Category: *category}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func CategoryInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	categoryId, errorCategoryId := uuid.Parse(chi.URLParam(r, "category_id"))

	if errorCategoryId != nil {
		payload = RestService.Error400HandleService(w, errorCategoryId)
	} else {
		category, errorCategory := handler.CategoryInfo(&categoryId, &token.UserId, nil)

		if errorCategory != nil {
			payload = RestService.Error400HandleService(w, errorCategory)
		} else {
			payload = &response.CategoryInfo{Category: *category}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func CategoryUpdate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	categoryId, errorCategoryId := uuid.Parse(chi.URLParam(r, "category_id"))

	if errorCategoryId != nil {
		payload = RestService.Error400HandleService(w, errorCategoryId)
	} else {
		categoryUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromCategoryUpdate(r.Body)

		if errorJsonDecode != nil {
			payload = RestService.Error400HandleService(w, errorJsonDecode)
		} else {
			category, errorCategory := handler.CategoryUpdate(&categoryId, &token.UserId, &categoryUpdateDTO)

			if errorCategory != nil {
				payload = RestService.Error400HandleService(w, errorCategory)
			} else {
				payload = &response.CategoryInfo{Category: *category}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func CategoryDelete(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	categoryId, errorCategoryId := uuid.Parse(chi.URLParam(r, "category_id"))

	if errorCategoryId != nil {
		payload = RestService.Error400HandleService(w, errorCategoryId)
	} else {
		categoryDeleteStatus, errorCategoryDeleteStatus := handler.CategoryDelete(&categoryId, &token.UserId)

		if errorCategoryDeleteStatus != nil {
			payload = RestService.Error400HandleService(w, errorCategoryDeleteStatus)
		} else if categoryDeleteStatus {
			payload = &response.CategoryDelete{Message: statusCategoryDeleteSuccess, Status: http.StatusOK}
		} else {
			payload = RestService.Error400HandleService(w, statusCategoryDeleteError)
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}
