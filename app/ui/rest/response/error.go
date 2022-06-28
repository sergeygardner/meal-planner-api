package response

import (
	"github.com/google/uuid"
	"net/http"
)

type Error struct {
	Id       uuid.UUID `json:"id"`
	Code     string    `json:"code"`
	Status   int       `json:"status"`
	Line     int       `json:"line"`
	Message  string    `json:"message"`
	Trace    string    `json:"trace"`
	Response `json:",omitempty"`
}

func (e *Error) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (e *Error) GetStatus() int {
	return e.Status
}
