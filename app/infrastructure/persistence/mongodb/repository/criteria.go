package repository

import (
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence"
	PersistenceRepository "github.com/sergeygardner/meal-planner-api/infrastructure/persistence/repository"
)

var CriteriaRepository = PersistenceRepository.CriteriaRepository{
	GetCriteriaById: func(id *uuid.UUID, criteria *persistence.Criteria) *persistence.Criteria {
		if criteria == nil {
			criteria = &persistence.Criteria{
				Where: map[string]interface{}{},
			}
		} else if criteria.Where == nil {
			criteria.Where = map[string]interface{}{}
		}

		criteria.Where["id"] = id

		return criteria
	},
	GetCriteriaByIds: func(id []*uuid.UUID, criteria *persistence.Criteria) *persistence.Criteria {
		if criteria == nil {
			criteria = &persistence.Criteria{
				Where: map[string]interface{}{},
			}
		} else if criteria.Where == nil {
			criteria.Where = map[string]interface{}{}
		}

		criteria.Where["id"] = id

		return criteria
	},
	GetCriteriaByEntityId: func(id *uuid.UUID, criteria *persistence.Criteria) *persistence.Criteria {
		if criteria == nil {
			criteria = &persistence.Criteria{
				Where: map[string]interface{}{},
			}
		} else if criteria.Where == nil {
			criteria.Where = map[string]interface{}{}
		}

		criteria.Where["entity_id"] = id

		return criteria
	},
	GetCriteriaByDeriveId: func(id *uuid.UUID, criteria *persistence.Criteria) *persistence.Criteria {
		if criteria == nil {
			criteria = &persistence.Criteria{
				Where: map[string]interface{}{},
			}
		} else if criteria.Where == nil {
			criteria.Where = map[string]interface{}{}
		}

		criteria.Where["derive_id"] = id

		return criteria
	},
	GetCriteriaByUnitId: func(id *uuid.UUID, criteria *persistence.Criteria) *persistence.Criteria {
		if criteria == nil {
			criteria = &persistence.Criteria{
				Where: map[string]interface{}{},
			}
		} else if criteria.Where == nil {
			criteria.Where = map[string]interface{}{}
		}

		criteria.Where["unit_id"] = id

		return criteria
	},
	GetCriteriaByRecipeId: func(id *uuid.UUID, criteria *persistence.Criteria) *persistence.Criteria {
		if criteria == nil {
			criteria = &persistence.Criteria{
				Where: map[string]interface{}{},
			}
		} else if criteria.Where == nil {
			criteria.Where = map[string]interface{}{}
		}

		criteria.Where["recipe_id"] = id

		return criteria
	},
	GetCriteriaByUserId: func(id *uuid.UUID, criteria *persistence.Criteria) *persistence.Criteria {
		if criteria == nil {
			criteria = &persistence.Criteria{
				Where: map[string]interface{}{},
			}
		} else if criteria.Where == nil {
			criteria.Where = map[string]interface{}{}
		}

		criteria.Where["user_id"] = id

		return criteria
	},
	GetCriteriaByName: func(name *string, criteria *persistence.Criteria) *persistence.Criteria {
		if criteria == nil {
			criteria = &persistence.Criteria{
				Where: map[string]interface{}{},
			}
		} else if criteria.Where == nil {
			criteria.Where = map[string]interface{}{}
		}

		criteria.Where["name"] = name

		return criteria
	},
}
