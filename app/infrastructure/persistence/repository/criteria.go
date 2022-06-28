package repository

import (
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence"
)

type CriteriaRepository struct {
	GetCriteriaById       func(id *uuid.UUID, criteria *persistence.Criteria) *persistence.Criteria
	GetCriteriaByIds      func(id []*uuid.UUID, criteria *persistence.Criteria) *persistence.Criteria
	GetCriteriaByEntityId func(id *uuid.UUID, criteria *persistence.Criteria) *persistence.Criteria
	GetCriteriaByDeriveId func(id *uuid.UUID, criteria *persistence.Criteria) *persistence.Criteria
	GetCriteriaByUnitId   func(id *uuid.UUID, criteria *persistence.Criteria) *persistence.Criteria
	GetCriteriaByRecipeId func(id *uuid.UUID, criteria *persistence.Criteria) *persistence.Criteria
	GetCriteriaByUserId   func(id *uuid.UUID, criteria *persistence.Criteria) *persistence.Criteria
	GetCriteriaByName     func(name *string, criteria *persistence.Criteria) *persistence.Criteria
}
