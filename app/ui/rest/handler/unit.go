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
	statusUnitDeleteSuccess = "the unit has been deleted successful"
	statusUnitDeleteError   = errors.New("the unit has not been deleted")
)

func UnitsInfo(w http.ResponseWriter, r *http.Request) {
	units, errorUnit := handler.UnitsInfo(nil)

	if errorUnit != nil {
		payload = RestService.Error400HandleService(w, errorUnit)
	} else {
		if units == nil {
			units = []*DomainEntity.Unit{}
		}
		payload = &response.UnitsInfo{Units: units}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func UnitCreate(w http.ResponseWriter, r *http.Request) {
	unitUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromUnitUpdate(r.Body)

	if errorJsonDecode != nil {
		payload = RestService.Error400HandleService(w, errorJsonDecode)
	} else {
		unit, errorAuthRegister := handler.UnitCreate(&unitUpdateDTO)

		if errorAuthRegister != nil {
			payload = RestService.Error400HandleService(w, errorAuthRegister)
		} else {
			payload = &response.UnitInfo{Unit: *unit}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func UnitInfo(w http.ResponseWriter, r *http.Request) {
	unitId, errorUnitId := uuid.Parse(chi.URLParam(r, "unit_id"))

	if errorUnitId != nil {
		payload = RestService.Error400HandleService(w, errorUnitId)
	} else {
		unit, errorUnit := handler.UnitInfo(&unitId, nil)

		if errorUnit != nil {
			payload = RestService.Error400HandleService(w, errorUnit)
		} else {
			payload = &response.UnitInfo{Unit: *unit}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func UnitUpdate(w http.ResponseWriter, r *http.Request) {
	unitId, errorUnitId := uuid.Parse(chi.URLParam(r, "unit_id"))

	if errorUnitId != nil {
		payload = RestService.Error400HandleService(w, errorUnitId)
	} else {
		unitUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromUnitUpdate(r.Body)

		if errorJsonDecode != nil {
			payload = RestService.Error400HandleService(w, errorJsonDecode)
		} else {
			unit, errorUnit := handler.UnitUpdate(&unitId, &unitUpdateDTO)

			if errorUnit != nil {
				payload = RestService.Error400HandleService(w, errorUnit)
			} else {
				payload = &response.UnitInfo{Unit: *unit}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func UnitDelete(w http.ResponseWriter, r *http.Request) {
	unitId, errorUnitId := uuid.Parse(chi.URLParam(r, "unit_id"))

	if errorUnitId != nil {
		payload = RestService.Error400HandleService(w, errorUnitId)
	} else {
		unitDeleteStatus, errorUnitDeleteStatus := handler.UnitDelete(&unitId)

		if errorUnitDeleteStatus != nil {
			payload = RestService.Error400HandleService(w, errorUnitDeleteStatus)
		} else if unitDeleteStatus {
			payload = &response.UnitDelete{Message: statusUnitDeleteSuccess, Status: http.StatusOK}
		} else {
			payload = RestService.Error400HandleService(w, statusUnitDeleteError)
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}
