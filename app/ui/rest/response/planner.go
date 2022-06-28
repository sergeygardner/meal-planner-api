package response

import (
	"github.com/sergeygardner/meal-planner-api/domain/aggregate"
	"net/http"
)

type PlannerInfo struct {
	aggregate.Planner
	Response `json:",omitempty"`
}

func (ri *PlannerInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ri *PlannerInfo) GetStatus() int {
	return http.StatusOK
}

type PlannersInfo struct {
	Planners []*aggregate.Planner `json:"planners"`
	Response `json:",omitempty"`
}

func (ri *PlannersInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ri *PlannersInfo) GetStatus() int {
	return http.StatusOK
}

type PlannerDelete struct {
	Message  string `json:"message"`
	Status   int    `json:"status"`
	Response `json:",omitempty"`
}

func (ud *PlannerDelete) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ud *PlannerDelete) GetStatus() int {
	return ud.Status
}

type PlannerCalculation struct {
	Overall  []*aggregate.PlannerCalculation `json:"overall"`
	Response `json:",omitempty"`
}

func (pc *PlannerCalculation) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (pc *PlannerCalculation) GetStatus() int {
	return http.StatusOK
}
