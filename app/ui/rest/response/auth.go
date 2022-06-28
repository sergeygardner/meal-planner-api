package response

import (
	DomainResponse "github.com/sergeygardner/meal-planner-api/domain/response"
	"net/http"
)

type AuthToken struct {
	DomainResponse.AuthToken
	Response `json:",omitempty"`
}

func (at *AuthToken) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (at *AuthToken) GetStatus() int {
	return http.StatusOK
}

type AuthConfirmation struct {
	DomainResponse.AuthConfirmation
	Response `json:",omitempty"`
}

func (ac *AuthConfirmation) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (ac *AuthConfirmation) GetStatus() int {
	return ac.Status
}
