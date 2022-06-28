package service

import (
	DomainResponse "github.com/sergeygardner/meal-planner-api/domain/response"
	"github.com/sergeygardner/meal-planner-api/ui/rest/response"
)

func MakeResponseAuthToken(domainAuthToken *DomainResponse.AuthToken) *response.AuthToken {
	return &response.AuthToken{
		AuthToken: *domainAuthToken,
	}
}
