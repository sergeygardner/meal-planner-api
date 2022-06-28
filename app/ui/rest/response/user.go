package response

import (
	"github.com/sergeygardner/meal-planner-api/domain/entity"
	"net/http"
)

type UserInfo struct {
	entity.User
	Response `json:",omitempty"`
}

func (ui *UserInfo) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ui *UserInfo) GetStatus() int {
	return http.StatusOK
}

type UserDelete struct {
	Message  string `json:"message"`
	Status   int    `json:"status"`
	Response `json:",omitempty"`
}

func (ud *UserDelete) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ud *UserDelete) GetStatus() int {
	return ud.Status
}
