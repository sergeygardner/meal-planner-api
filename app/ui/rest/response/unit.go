package response

import (
	"github.com/sergeygardner/meal-planner-api/domain/entity"
	"net/http"
)

type UnitInfo struct {
	entity.Unit
	Response `json:",omitempty"`
}

func (ri *UnitInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ri *UnitInfo) GetStatus() int {
	return http.StatusOK
}

type UnitsInfo struct {
	Units    []*entity.Unit `json:"units"`
	Response `json:",omitempty"`
}

func (ri *UnitsInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ri *UnitsInfo) GetStatus() int {
	return http.StatusOK
}

type UnitDelete struct {
	Message  string `json:"message"`
	Status   int    `json:"status"`
	Response `json:",omitempty"`
}

func (ud *UnitDelete) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ud *UnitDelete) GetStatus() int {
	return ud.Status
}
