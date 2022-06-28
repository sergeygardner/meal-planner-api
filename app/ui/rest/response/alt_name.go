package response

import (
	"github.com/sergeygardner/meal-planner-api/domain/entity"
	"net/http"
)

type AltNameInfo struct {
	entity.AltName
	Response `json:",omitempty"`
}

func (ri *AltNameInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ri *AltNameInfo) GetStatus() int {
	return http.StatusOK
}

type AltNamesInfo struct {
	AltNames []*entity.AltName `json:"alt_names"`
	Response `json:",omitempty"`
}

func (ri *AltNamesInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ri *AltNamesInfo) GetStatus() int {
	return http.StatusOK
}

type AltNameDelete struct {
	Message  string `json:"message"`
	Status   int    `json:"status"`
	Response `json:",omitempty"`
}

func (ud *AltNameDelete) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ud *AltNameDelete) GetStatus() int {
	return ud.Status
}
