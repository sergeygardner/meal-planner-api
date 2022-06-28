package response

import (
	"net/http"
)

type Version struct {
	Version  string `json:"version"`
	Response `json:",omitempty"`
}

func (v *Version) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (v *Version) GetStatus() int {
	return http.StatusOK
}
