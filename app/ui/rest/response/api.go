package response

import (
	"net/http"
)

type API struct {
	Version  string   `json:"version"`
	Versions []string `json:"versions"`
	Response `json:",omitempty"`
}

func (a *API) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (a *API) GetStatus() int {
	return http.StatusOK
}
