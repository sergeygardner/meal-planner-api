package response

import (
	"github.com/sergeygardner/meal-planner-api/domain/aggregate"
	"net/http"
)

type RecipeProcessAggregateInfo struct {
	*aggregate.RecipeProcess `json:",omitempty"`
	Response                 `json:",omitempty"`
}

func (rpa *RecipeProcessAggregateInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (rpa *RecipeProcessAggregateInfo) GetStatus() int {
	return http.StatusOK
}

type RecipeProcessAggregatesInfo struct {
	Processes []*aggregate.RecipeProcess `json:"processes"`
	Response  `json:",omitempty"`
}

func (rpa *RecipeProcessAggregatesInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (rpa *RecipeProcessAggregatesInfo) GetStatus() int {
	return http.StatusOK
}

type RecipeProcessDelete struct {
	Message  string `json:"message"`
	Status   int    `json:"status"`
	Response `json:",omitempty"`
}

func (rpd *RecipeProcessDelete) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (rpd *RecipeProcessDelete) GetStatus() int {
	return rpd.Status
}
