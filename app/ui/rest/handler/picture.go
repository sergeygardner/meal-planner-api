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
	statusPictureDeleteSuccess = "the picture has been deleted successful"
	statusPictureDeleteError   = errors.New("the picture has not been deleted")
)

func PicturesInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	parentId, errorParentId := getParentId(chi.RouteContext(r.Context()), "picture_id")

	if errorParentId != nil {
		payload = RestService.Error400HandleService(w, errorParentId)
	} else {
		pictures, errorPicture := handler.PicturesInfo(&token.UserId, parentId, nil)
		if errorPicture != nil {
			payload = RestService.Error400HandleService(w, errorPicture)
		} else {
			if pictures == nil {
				pictures = []*DomainAggregate.Picture{}
			}
			payload = &response.PicturesInfo{Pictures: pictures}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func PictureCreate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	pictureUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromPictureUpdate(r.Body)
	parentId, errorParentId := getParentId(chi.RouteContext(r.Context()), "picture_id")

	if errorParentId != nil {
		payload = RestService.Error400HandleService(w, errorParentId)
	} else if errorJsonDecode != nil {
		payload = RestService.Error400HandleService(w, errorJsonDecode)
	} else {
		picture, errorAuthRegister := handler.PictureCreate(&token.UserId, parentId, &pictureUpdateDTO)

		if errorAuthRegister != nil {
			payload = RestService.Error400HandleService(w, errorAuthRegister)
		} else {
			payload = &response.PictureInfo{Picture: *picture}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func PictureInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	parentId, errorParentId := getParentId(chi.RouteContext(r.Context()), "picture_id")
	pictureId, errorPictureId := uuid.Parse(chi.URLParam(r, "picture_id"))

	if errorParentId != nil {
		payload = RestService.Error400HandleService(w, errorParentId)
	} else if errorPictureId != nil {
		payload = RestService.Error400HandleService(w, errorPictureId)
	} else {
		picture, errorPicture := handler.PictureInfo(&pictureId, &token.UserId, parentId, nil)

		if errorPicture != nil {
			payload = RestService.Error400HandleService(w, errorPicture)
		} else {
			payload = &response.PictureInfo{Picture: *picture}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func PictureUpdate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	parentId, errorParentId := getParentId(chi.RouteContext(r.Context()), "picture_id")
	pictureId, errorPictureId := uuid.Parse(chi.URLParam(r, "picture_id"))

	if errorParentId != nil {
		payload = RestService.Error400HandleService(w, errorParentId)
	} else if errorPictureId != nil {
		payload = RestService.Error400HandleService(w, errorPictureId)
	} else {
		pictureUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromPictureUpdate(r.Body)

		if errorJsonDecode != nil {
			payload = RestService.Error400HandleService(w, errorJsonDecode)
		} else {
			picture, errorPicture := handler.PictureUpdate(&pictureId, &token.UserId, parentId, &pictureUpdateDTO)

			if errorPicture != nil {
				payload = RestService.Error400HandleService(w, errorPicture)
			} else {
				payload = &response.PictureInfo{Picture: *picture}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func PictureDelete(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	parentId, errorParentId := getParentId(chi.RouteContext(r.Context()), "picture_id")
	pictureId, errorPictureId := uuid.Parse(chi.URLParam(r, "picture_id"))

	if errorParentId != nil {
		payload = RestService.Error400HandleService(w, errorParentId)
	} else if errorPictureId != nil {
		payload = RestService.Error400HandleService(w, errorPictureId)
	} else {
		pictureDeleteStatus, errorPictureDeleteStatus := handler.PictureDelete(&pictureId, &token.UserId, parentId)

		if errorPictureDeleteStatus != nil {
			payload = RestService.Error400HandleService(w, errorPictureDeleteStatus)
		} else if pictureDeleteStatus {
			payload = &response.PictureDelete{Message: statusPictureDeleteSuccess, Status: http.StatusOK}
		} else {
			payload = RestService.Error400HandleService(w, statusPictureDeleteError)
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}
