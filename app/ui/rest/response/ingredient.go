package response

import (
	"github.com/sergeygardner/meal-planner-api/domain/entity"
	"net/http"
)

type IngredientInfo struct {
	entity.Ingredient
	Response `json:",omitempty"`
}

func (ri *IngredientInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ri *IngredientInfo) GetStatus() int {
	return http.StatusOK
}

type IngredientsInfo struct {
	Ingredients []*entity.Ingredient `json:"ingredients"`
	Response    `json:",omitempty"`
}

func (ri *IngredientsInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ri *IngredientsInfo) GetStatus() int {
	return http.StatusOK
}

type IngredientDelete struct {
	Message  string `json:"message"`
	Status   int    `json:"status"`
	Response `json:",omitempty"`
}

func (ud *IngredientDelete) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ud *IngredientDelete) GetStatus() int {
	return ud.Status
}
