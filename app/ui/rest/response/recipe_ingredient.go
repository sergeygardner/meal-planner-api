package response

import (
	"github.com/sergeygardner/meal-planner-api/domain/aggregate"
	"net/http"
)

type RecipeIngredientInfo struct {
	aggregate.RecipeIngredient
	Response `json:",omitempty"`
}

func (ri *RecipeIngredientInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ri *RecipeIngredientInfo) GetStatus() int {
	return http.StatusOK
}

type RecipeIngredientsInfo struct {
	Ingredients []*aggregate.RecipeIngredient `json:"ingredients"`
	Response    `json:",omitempty"`
}

func (ri *RecipeIngredientsInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ri *RecipeIngredientsInfo) GetStatus() int {
	return http.StatusOK
}

type RecipeIngredientDelete struct {
	Message  string `json:"message"`
	Status   int    `json:"status"`
	Response `json:",omitempty"`
}

func (ud *RecipeIngredientDelete) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ud *RecipeIngredientDelete) GetStatus() int {
	return ud.Status
}
