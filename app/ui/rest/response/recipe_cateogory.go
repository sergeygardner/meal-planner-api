package response

import (
	"github.com/sergeygardner/meal-planner-api/domain/aggregate"
	"net/http"
)

type RecipeCategoryInfo struct {
	aggregate.RecipeCategory
	Response `json:",omitempty"`
}

func (ri *RecipeCategoryInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ri *RecipeCategoryInfo) GetStatus() int {
	return http.StatusOK
}

type RecipeCategoriesInfo struct {
	Categories []*aggregate.RecipeCategory `json:"categories"`
	Response   `json:",omitempty"`
}

func (ri *RecipeCategoriesInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ri *RecipeCategoriesInfo) GetStatus() int {
	return http.StatusOK
}

type RecipeCategoryDelete struct {
	Message  string `json:"message"`
	Status   int    `json:"status"`
	Response `json:",omitempty"`
}

func (ud *RecipeCategoryDelete) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ud *RecipeCategoryDelete) GetStatus() int {
	return ud.Status
}
