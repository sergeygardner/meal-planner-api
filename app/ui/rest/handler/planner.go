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
	statusPlannerDeleteSuccess = "the planner has been deleted successful"
	statusPlannerDeleteError   = errors.New("the planner has not been deleted")
)

func PlannersInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	planners, errorPlanners := handler.PlannersInfo(&token.UserId, nil)

	if errorPlanners != nil {
		payload = RestService.Error400HandleService(w, errorPlanners)
	} else {
		if planners == nil {
			planners = []*DomainAggregate.Planner{}
		}
		payload = &response.PlannersInfo{Planners: planners}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func PlannerCreate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	plannerUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromPlannerUpdate(r.Body)

	if errorJsonDecode != nil {
		payload = RestService.Error400HandleService(w, errorJsonDecode)
	} else {
		planner, errorPlanner := handler.PlannerCreate(&token.UserId, &plannerUpdateDTO)

		if errorPlanner != nil {
			payload = RestService.Error400HandleService(w, errorPlanner)
		} else {
			payload = &response.PlannerInfo{Planner: *planner}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func PlannerInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	plannerId, errorPlannerId := uuid.Parse(chi.URLParam(r, "planner_id"))

	if errorPlannerId != nil {
		payload = RestService.Error400HandleService(w, errorPlannerId)
	} else {
		planner, errorPlanner := handler.PlannerInfo(&plannerId, &token.UserId, nil)

		if errorPlanner != nil {
			payload = RestService.Error400HandleService(w, errorPlanner)
		} else {
			payload = &response.PlannerInfo{Planner: *planner}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func PlannerUpdate(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	plannerId, errorPlannerId := uuid.Parse(chi.URLParam(r, "planner_id"))

	if errorPlannerId != nil {
		payload = RestService.Error400HandleService(w, errorPlannerId)
	} else {
		plannerUpdateDTO, errorJsonDecode := DomainService.CreateEntityFromPlannerUpdate(r.Body)

		if errorJsonDecode != nil {
			payload = RestService.Error400HandleService(w, errorJsonDecode)
		} else {
			planner, errorPlanner := handler.PlannerUpdate(&plannerId, &token.UserId, &plannerUpdateDTO)

			if errorPlanner != nil {
				payload = RestService.Error400HandleService(w, errorPlanner)
			} else {
				payload = &response.PlannerInfo{Planner: *planner}
			}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func PlannerDelete(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	plannerId, errorPlannerId := uuid.Parse(chi.URLParam(r, "planner_id"))

	if errorPlannerId != nil {
		payload = RestService.Error400HandleService(w, errorPlannerId)
	} else {
		plannerDeleteStatus, errorPlannerDeleteStatus := handler.PlannerDelete(&plannerId, &token.UserId)

		if errorPlannerDeleteStatus != nil {
			payload = RestService.Error400HandleService(w, errorPlannerDeleteStatus)
		} else if plannerDeleteStatus {
			payload = &response.PlannerDelete{Message: statusPlannerDeleteSuccess, Status: http.StatusOK}
		} else {
			payload = RestService.Error400HandleService(w, statusPlannerDeleteError)
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}

func PlannerCalculateInfo(w http.ResponseWriter, r *http.Request) {
	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(r.Context())

	if errorExtractClaimsFromContext != nil {
		http.Error(w, errorExtractClaimsFromContext.Error(), 401)

		return
	}

	plannerId, errorPlannerId := uuid.Parse(chi.URLParam(r, "planner_id"))

	if errorPlannerId != nil {
		payload = RestService.Error400HandleService(w, errorPlannerId)
	} else {
		plannerCalculation, errorPlannerCalculation := handler.PlannerCalculate(&plannerId, &token.UserId)

		if errorPlannerCalculation != nil {
			payload = RestService.Error400HandleService(w, errorPlannerCalculation)
		} else {
			if plannerCalculation == nil {
				plannerCalculation = []*DomainAggregate.PlannerCalculation{}
			}
			payload = &response.PlannerCalculation{Overall: plannerCalculation}
		}
	}

	errorRender := RestService.Render(w, r, payload)

	if errorRender != nil {
		log.Panic(errorRender)
	}
}
