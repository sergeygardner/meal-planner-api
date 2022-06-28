package service

import (
	"context"
	"github.com/sergeygardner/meal-planner-api/domain/model"
	UiService "github.com/sergeygardner/meal-planner-api/ui/service"
)

func ExtractClaimsFromContext(ctx context.Context) (*model.Token, error) {
	return UiService.ExtractClaimsFromContext(ctx)
}
