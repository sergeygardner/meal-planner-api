package response

import (
	"github.com/sergeygardner/meal-planner-api/domain/aggregate"
	"net/http"
)

type RecipeMeasureInfo struct {
	aggregate.RecipeMeasure
	Response `json:",omitempty"`
}

func (ri *RecipeMeasureInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ri *RecipeMeasureInfo) GetStatus() int {
	return http.StatusOK
}

type RecipeMeasuresInfo struct {
	Measures []*aggregate.RecipeMeasure `json:"measures"`
	Response `json:",omitempty"`
}

func (ri *RecipeMeasuresInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ri *RecipeMeasuresInfo) GetStatus() int {
	return http.StatusOK
}

type RecipeMeasureDelete struct {
	Message  string `json:"message"`
	Status   int    `json:"status"`
	Response `json:",omitempty"`
}

func (ud *RecipeMeasureDelete) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ud *RecipeMeasureDelete) GetStatus() int {
	return ud.Status
}
