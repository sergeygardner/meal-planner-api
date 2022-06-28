package response

import (
	"github.com/sergeygardner/meal-planner-api/domain/aggregate"
	"net/http"
)

type RecipeInfo struct {
	aggregate.Recipe
	Response `json:",omitempty"`
}

func (ri *RecipeInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ri *RecipeInfo) GetStatus() int {
	return http.StatusOK
}

type RecipesInfo struct {
	Recipes  []*aggregate.Recipe `json:"recipes"`
	Response `json:",omitempty"`
}

func (ri *RecipesInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ri *RecipesInfo) GetStatus() int {
	return http.StatusOK
}

type RecipeDelete struct {
	Message  string `json:"message"`
	Status   int    `json:"status"`
	Response `json:",omitempty"`
}

func (ud *RecipeDelete) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ud *RecipeDelete) GetStatus() int {
	return ud.Status
}
