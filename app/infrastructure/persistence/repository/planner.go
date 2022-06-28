package repository

import (
	"github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence"
)

type PlannerRepositoryInterface interface {
	FindOne(criteria *persistence.Criteria) (*entity.Planner, error)
	FindAll(criteria *persistence.Criteria) ([]*entity.Planner, error)
	InsertOne(planner *entity.Planner) (*entity.Planner, error)
	InsertMany(planners []*entity.Planner) ([]*entity.Planner, error)
	UpdateOne(criteria *persistence.Criteria, entity *entity.Planner) (*entity.Planner, error)
	UpdateMany(criteria *persistence.Criteria, entities []*entity.Planner) ([]*entity.Planner, error)
	DeleteOne(criteria *persistence.Criteria) (bool, error)
	GetCriteria() *CriteriaRepository
}

type PlannerIntervalRepositoryInterface interface {
	FindOne(criteria *persistence.Criteria) (*entity.PlannerInterval, error)
	FindAll(criteria *persistence.Criteria) ([]*entity.PlannerInterval, error)
	InsertOne(plannerInterval *entity.PlannerInterval) (*entity.PlannerInterval, error)
	InsertMany(plannerIntervals []*entity.PlannerInterval) ([]*entity.PlannerInterval, error)
	UpdateOne(criteria *persistence.Criteria, entity *entity.PlannerInterval) (*entity.PlannerInterval, error)
	UpdateMany(criteria *persistence.Criteria, entities []*entity.PlannerInterval) ([]*entity.PlannerInterval, error)
	DeleteOne(criteria *persistence.Criteria) (bool, error)
	GetCriteria() *CriteriaRepository
}

type PlannerRecipeRepositoryInterface interface {
	FindOne(criteria *persistence.Criteria) (*entity.PlannerRecipe, error)
	FindAll(criteria *persistence.Criteria) ([]*entity.PlannerRecipe, error)
	InsertOne(plannerRecipe *entity.PlannerRecipe) (*entity.PlannerRecipe, error)
	InsertMany(plannerRecipes []*entity.PlannerRecipe) ([]*entity.PlannerRecipe, error)
	UpdateOne(criteria *persistence.Criteria, entity *entity.PlannerRecipe) (*entity.PlannerRecipe, error)
	UpdateMany(criteria *persistence.Criteria, entities []*entity.PlannerRecipe) ([]*entity.PlannerRecipe, error)
	DeleteOne(criteria *persistence.Criteria) (bool, error)
	GetCriteria() *CriteriaRepository
}
