package response

import (
	"github.com/sergeygardner/meal-planner-api/domain/aggregate"
	"net/http"
)

type PlannerRecipeInfo struct {
	aggregate.PlannerRecipe
	Response `json:",omitempty"`
}

func (ri *PlannerRecipeInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ri *PlannerRecipeInfo) GetStatus() int {
	return http.StatusOK
}

type PlannerRecipesInfo struct {
	PlannerRecipes []*aggregate.PlannerRecipe `json:"planner_recipes"`
	Response       `json:",omitempty"`
}

func (ri *PlannerRecipesInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ri *PlannerRecipesInfo) GetStatus() int {
	return http.StatusOK
}

type PlannerRecipeDelete struct {
	Message  string `json:"message"`
	Status   int    `json:"status"`
	Response `json:",omitempty"`
}

func (ud *PlannerRecipeDelete) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ud *PlannerRecipeDelete) GetStatus() int {
	return ud.Status
}
