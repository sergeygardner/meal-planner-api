package response

import (
	"github.com/sergeygardner/meal-planner-api/domain/aggregate"
	"net/http"
)

type PictureInfo struct {
	aggregate.Picture
	Response `json:",omitempty"`
}

func (ri *PictureInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ri *PictureInfo) GetStatus() int {
	return http.StatusOK
}

type PicturesInfo struct {
	Pictures []*aggregate.Picture `json:"pictures"`
	Response `json:",omitempty"`
}

func (ri *PicturesInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ri *PicturesInfo) GetStatus() int {
	return http.StatusOK
}

type PictureDelete struct {
	Message  string `json:"message"`
	Status   int    `json:"status"`
	Response `json:",omitempty"`
}

func (ud *PictureDelete) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ud *PictureDelete) GetStatus() int {
	return ud.Status
}
