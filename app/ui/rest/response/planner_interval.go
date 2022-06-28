package response

import (
	"github.com/sergeygardner/meal-planner-api/domain/aggregate"
	"net/http"
)

type PlannerIntervalInfo struct {
	aggregate.PlannerInterval
	Response `json:",omitempty"`
}

func (ri *PlannerIntervalInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ri *PlannerIntervalInfo) GetStatus() int {
	return http.StatusOK
}

type PlannerIntervalsInfo struct {
	PlannerIntervals []*aggregate.PlannerInterval `json:"planner_intervals"`
	Response         `json:",omitempty"`
}

func (ri *PlannerIntervalsInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ri *PlannerIntervalsInfo) GetStatus() int {
	return http.StatusOK
}

type PlannerIntervalDelete struct {
	Message  string `json:"message"`
	Status   int    `json:"status"`
	Response `json:",omitempty"`
}

func (ud *PlannerIntervalDelete) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ud *PlannerIntervalDelete) GetStatus() int {
	return ud.Status
}
