package response

import (
	"github.com/sergeygardner/meal-planner-api/domain/aggregate"
	"net/http"
)

type CategoryInfo struct {
	aggregate.Category
	Response `json:",omitempty"`
}

func (ri *CategoryInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ri *CategoryInfo) GetStatus() int {
	return http.StatusOK
}

type CategoriesInfo struct {
	Categories []*aggregate.Category `json:"categories"`
	Response   `json:",omitempty"`
}

func (ri *CategoriesInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ri *CategoriesInfo) GetStatus() int {
	return http.StatusOK
}

type CategoryDelete struct {
	Message  string `json:"message"`
	Status   int    `json:"status"`
	Response `json:",omitempty"`
}

func (ud *CategoryDelete) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ud *CategoryDelete) GetStatus() int {
	return ud.Status
}
