package handler

import (
	"github.com/pkg/errors"
	"github.com/sergeygardner/meal-planner-api/application/handler"
	DomainService "github.com/sergeygardner/meal-planner-api/domain/service"
	"github.com/sergeygardner/meal-planner-api/ui/rest/response"
	"github.com/sergeygardner/meal-planner-api/ui/rest/service"
	RestService "github.com/sergeygardner/meal-planner-api/ui/rest/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var (
	statusUserDeleteSuccess = "the user has been deleted successful"
	statusUserDeleteError   = errors.New("the user has not been deleted")
)

func UserInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		payload = RestService.Error400HandleService(w, errorExtractClaimsFromContext)
	} else {
		user, errorUser := handler.UserInfo(&token.UserId)

		if errorUser != nil {
			payload = service.Error400HandleService(w, errorUser)
		} else {
			payload = &response.UserInfo{User: *user}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		payload = RestService.Error400HandleService(w, errorExtractClaimsFromContext)
	} else {
		userUpdateDTO, errorJsonDecode := DomainService.CreateDTOFromUserUpdate(r.Body)

		if errorJsonDecode != nil {
			payload = service.Error400HandleService(w, errorJsonDecode)
		} else {
			user, errorUser := handler.UserUpdate(&token.UserId, &userUpdateDTO)

			if errorUser != nil {
				payload = service.Error400HandleService(w, errorUser)
			} else {
				payload = &response.UserInfo{User: *user}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		payload = RestService.Error400HandleService(w, errorExtractClaimsFromContext)
	} else {
		userDeleteStatus, errorUserDeleteStatus := handler.UserDelete(&token.UserId)

		if errorUserDeleteStatus != nil {
			payload = service.Error400HandleService(w, errorUserDeleteStatus)
		} else if userDeleteStatus {
			payload = &response.UserDelete{Message: statusUserDeleteSuccess, Status: http.StatusOK}
		} else {
			payload = service.Error400HandleService(w, statusUserDeleteError)
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}
