package response

import (
	"net/http"
)

type Index struct {
	Description string `json:"description"`
	Response    `json:",omitempty"`
}

func (i *Index) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (i *Index) GetStatus() int {
	return http.StatusOK
}
