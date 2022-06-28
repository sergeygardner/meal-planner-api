package handler

import (
	"github.com/sergeygardner/meal-planner-api/application/handler"
	DomainService "github.com/sergeygardner/meal-planner-api/domain/service"
	RestResponse "github.com/sergeygardner/meal-planner-api/ui/rest/response"
	RestService "github.com/sergeygardner/meal-planner-api/ui/rest/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func AuthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func AuthCredentials(w http.ResponseWriter, r *http.Request) {
	authCredentialsDTO, errorJsonDecode := DomainService.CreateDTOFromAuthCredentials(r.Body)

	if errorJsonDecode != nil {
		payload = RestService.Error400HandleService(w, errorJsonDecode)
	} else {
		authConfirmation, _, errorAuthCredentials := handler.AuthCredentials(authCredentialsDTO)

		if errorAuthCredentials != nil {
			payload = RestService.Error400HandleService(w, errorAuthCredentials)
		} else {
			payload = &RestResponse.AuthConfirmation{AuthConfirmation: *authConfirmation}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func AuthConfirmation(w http.ResponseWriter, r *http.Request) {
	authConfirmationDTO, errorJsonDecode := DomainService.CreateDTOFromAuthConfirmation(r.Body)

	if errorJsonDecode != nil {
		payload = RestService.Error400HandleService(w, errorJsonDecode)
	} else {
		authToken, errorAuthConfirmation := handler.AuthConfirmation(authConfirmationDTO)

		if errorAuthConfirmation != nil {
			payload = RestService.Error400HandleService(w, errorAuthConfirmation)
		} else {
			payload = RestService.MakeResponseAuthToken(authToken)
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func AuthRefresh(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		payload = RestService.Error400HandleService(w, errorExtractClaimsFromContext)
	} else {
		authToken, errorAuthToken := handler.AuthTokenByUserId(&token.UserId)

		if errorAuthToken != nil {
			payload = RestService.Error400HandleService(w, errorAuthToken)
		} else {
			payload = RestService.MakeResponseAuthToken(authToken)
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func AuthRegister(w http.ResponseWriter, r *http.Request) {
	userRegisterDTO, errorJsonDecode := DomainService.CreateDTOFromUserRegister(r.Body)

	if errorJsonDecode != nil {
		payload = RestService.Error400HandleService(w, errorJsonDecode)
	} else {
		user, errorAuthRegister := handler.AuthRegister(userRegisterDTO)

		if errorAuthRegister != nil {
			payload = RestService.Error400HandleService(w, errorAuthRegister)
		} else {
			payload = &RestResponse.UserInfo{User: *user}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}
