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
	statusPlannerIntervalDeleteSuccess = "the planner interval has been deleted successful"
	statusPlannerIntervalDeleteError   = errors.New("the planner interval has not been deleted")
)

func PlannerIntervalsInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	plannerId, errorPlannerId := uuid.Parse(chi.URLParam(r, "planner_id"))

	if errorPlannerId != nil {
		payload = RestService.Error400HandleService(w, errorPlannerId)
	} else {
		plannerIntervals, errorPlannerIntervals := handler.PlannerIntervalsInfo(&token.UserId, &plannerId, nil)

		if errorPlannerIntervals != nil {
			payload = RestService.Error400HandleService(w, errorPlannerIntervals)
		} else {
			if plannerIntervals == nil {
				plannerIntervals = []*DomainAggregate.PlannerInterval{}
			}
			payload = &response.PlannerIntervalsInfo{PlannerIntervals: plannerIntervals}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func PlannerIntervalCreate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	plannerId, errorPlannerId := uuid.Parse(chi.URLParam(r, "planner_id"))

	if errorPlannerId != nil {
		payload = RestService.Error400HandleService(w, errorPlannerId)
	} else {
		plannerIntervalUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromPlannerIntervalUpdate(r.Body)

		if errorJsonDecode != nil {
			payload = RestService.Error400HandleService(w, errorJsonDecode)
		} else {
			plannerInterval, errorPlannerInterval := handler.PlannerIntervalCreate(&token.UserId, &plannerId, &plannerIntervalUpdateDTO)

			if errorPlannerInterval != nil {
				payload = RestService.Error400HandleService(w, errorPlannerInterval)
			} else {
				payload = &response.PlannerIntervalInfo{PlannerInterval: *plannerInterval}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func PlannerIntervalInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	plannerId, errorPlannerId := uuid.Parse(chi.URLParam(r, "planner_id"))

	if errorPlannerId != nil {
		payload = RestService.Error400HandleService(w, errorPlannerId)
	} else {
		plannerIntervalId, errorPlannerIntervalId := uuid.Parse(chi.URLParam(r, "interval_id"))

		if errorPlannerIntervalId != nil {
			payload = RestService.Error400HandleService(w, errorPlannerIntervalId)
		} else {
			plannerInterval, errorPlannerInterval := handler.PlannerIntervalInfo(&plannerIntervalId, &token.UserId, &plannerId, nil)

			if errorPlannerInterval != nil {
				payload = RestService.Error400HandleService(w, errorPlannerInterval)
			} else {
				payload = &response.PlannerIntervalInfo{PlannerInterval: *plannerInterval}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func PlannerIntervalUpdate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	plannerId, errorPlannerId := uuid.Parse(chi.URLParam(r, "planner_id"))

	if errorPlannerId != nil {
		payload = RestService.Error400HandleService(w, errorPlannerId)
	} else {
		plannerIntervalId, errorPlannerIntervalId := uuid.Parse(chi.URLParam(r, "interval_id"))

		if errorPlannerIntervalId != nil {
			payload = RestService.Error400HandleService(w, errorPlannerIntervalId)
		} else {
			plannerIntervalUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromPlannerIntervalUpdate(r.Body)

			if errorJsonDecode != nil {
				payload = RestService.Error400HandleService(w, errorJsonDecode)
			} else {
				plannerInterval, errorPlannerInterval := handler.PlannerIntervalUpdate(&plannerIntervalId, &token.UserId, &plannerId, &plannerIntervalUpdateDTO)

				if errorPlannerInterval != nil {
					payload = RestService.Error400HandleService(w, errorPlannerInterval)
				} else {
					payload = &response.PlannerIntervalInfo{PlannerInterval: *plannerInterval}
				}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func PlannerIntervalDelete(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	plannerId, errorPlannerId := uuid.Parse(chi.URLParam(r, "planner_id"))

	if errorPlannerId != nil {
		payload = RestService.Error400HandleService(w, errorPlannerId)
	} else {
		plannerIntervalId, errorPlannerIntervalId := uuid.Parse(chi.URLParam(r, "interval_id"))

		if errorPlannerIntervalId != nil {
			payload = RestService.Error400HandleService(w, errorPlannerIntervalId)
		} else {
			plannerIntervalDeleteStatus, errorPlannerIntervalDeleteStatus := handler.PlannerIntervalDelete(&plannerIntervalId, &token.UserId, &plannerId)

			if errorPlannerIntervalDeleteStatus != nil {
				payload = RestService.Error400HandleService(w, errorPlannerIntervalDeleteStatus)
			} else if plannerIntervalDeleteStatus {
				payload = &response.PlannerIntervalDelete{Message: statusPlannerIntervalDeleteSuccess, Status: http.StatusOK}
			} else {
				payload = RestService.Error400HandleService(w, statusPlannerIntervalDeleteError)
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}
