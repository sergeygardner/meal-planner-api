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
	statusAltNameDeleteSuccess = "the recipe alt name has been deleted successful"
	statusAltNameDeleteError   = errors.New("the recipe alt name has not been deleted")
)

func AltNamesInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	parentId, errorParentId := getParentId(chi.RouteContext(r.Context()), "alt_name_id")

	if errorParentId != nil {
		payload = RestService.Error400HandleService(w, errorParentId)
	} else {
		altNames, errorAltNames := handler.AltNamesInfo(&token.UserId, parentId, nil)

		if errorAltNames != nil {
			payload = RestService.Error400HandleService(w, errorAltNames)
		} else {
			if altNames == nil {
				altNames = []*DomainEntity.AltName{}
			}
			payload = &response.AltNamesInfo{AltNames: altNames}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func AltNameCreate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	parentId, errorParentId := getParentId(chi.RouteContext(r.Context()), "alt_name_id")
	altNameUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromAltNameUpdate(r.Body)

	if errorParentId != nil {
		payload = RestService.Error400HandleService(w, errorParentId)
	} else if errorJsonDecode != nil {
		payload = RestService.Error400HandleService(w, errorJsonDecode)
	} else {
		altName, errorAltName := handler.AltNameCreate(&token.UserId, parentId, &altNameUpdateDTO)

		if errorAltName != nil {
			payload = RestService.Error400HandleService(w, errorAltName)
		} else {
			payload = &response.AltNameInfo{AltName: *altName}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func AltNameInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	altNameId, errorAltNameId := uuid.Parse(chi.URLParam(r, "alt_name_id"))
	parentId, errorParentId := getParentId(chi.RouteContext(r.Context()), "alt_name_id")

	if errorParentId != nil {
		payload = RestService.Error400HandleService(w, errorParentId)
	} else if errorAltNameId != nil {
		payload = RestService.Error400HandleService(w, errorAltNameId)
	} else {
		altName, errorAltName := handler.AltNameInfo(&altNameId, &token.UserId, parentId, nil)

		if errorAltName != nil {
			payload = RestService.Error400HandleService(w, errorAltName)
		} else {
			payload = &response.AltNameInfo{AltName: *altName}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func AltNameUpdate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	altNameId, errorAltNameId := uuid.Parse(chi.URLParam(r, "alt_name_id"))
	parentId, errorParentId := getParentId(chi.RouteContext(r.Context()), "alt_name_id")

	if errorParentId != nil {
		payload = RestService.Error400HandleService(w, errorParentId)
	} else if errorAltNameId != nil {
		payload = RestService.Error400HandleService(w, errorAltNameId)
	} else {
		altNameUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromAltNameUpdate(r.Body)

		if errorJsonDecode != nil {
			payload = RestService.Error400HandleService(w, errorJsonDecode)
		} else {
			altName, errorAltName := handler.AltNameUpdate(&altNameId, &token.UserId, parentId, &altNameUpdateDTO)

			if errorAltName != nil {
				payload = RestService.Error400HandleService(w, errorAltName)
			} else {
				payload = &response.AltNameInfo{AltName: *altName}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func AltNameDelete(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	altNameId, errorAltNameId := uuid.Parse(chi.URLParam(r, "alt_name_id"))
	parentId, errorParentId := getParentId(chi.RouteContext(r.Context()), "alt_name_id")

	if errorParentId != nil {
		payload = RestService.Error400HandleService(w, errorParentId)
	} else if errorAltNameId != nil {
		payload = RestService.Error400HandleService(w, errorAltNameId)
	} else {
		altNameDeleteStatus, errorAltNameDeleteStatus := handler.AltNameDelete(&altNameId, &token.UserId, parentId)

		if errorAltNameDeleteStatus != nil {
			payload = RestService.Error400HandleService(w, errorAltNameDeleteStatus)
		} else if altNameDeleteStatus {
			payload = &response.AltNameDelete{Message: statusAltNameDeleteSuccess, Status: http.StatusOK}
		} else {
			payload = RestService.Error400HandleService(w, statusAltNameDeleteError)
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}
