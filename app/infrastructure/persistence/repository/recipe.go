package repository

import (
	"github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence"
)

type RecipeRepositoryInterface interface {
	FindOne(criteria *persistence.Criteria) (*entity.Recipe, error)
	FindAll(criteria *persistence.Criteria) ([]*entity.Recipe, error)
	InsertOne(recipe *entity.Recipe) (*entity.Recipe, error)
	InsertMany(recipes []entity.Recipe) ([]entity.Recipe, error)
	UpdateOne(criteria *persistence.Criteria, entity *entity.Recipe) (*entity.Recipe, error)
	UpdateMany(criteria *persistence.Criteria, entities []*entity.Recipe) ([]*entity.Recipe, error)
	DeleteOne(criteria *persistence.Criteria) (bool, error)
	GetCriteria() *CriteriaRepository
}

type RecipeCategoryRepositoryInterface interface {
	FindOne(criteria *persistence.Criteria) (*entity.RecipeCategory, error)
	FindAll(criteria *persistence.Criteria) ([]*entity.RecipeCategory, error)
	InsertOne(recipeCategory *entity.RecipeCategory) (*entity.RecipeCategory, error)
	InsertMany(recipeCategories []entity.RecipeCategory) ([]entity.RecipeCategory, error)
	UpdateOne(criteria *persistence.Criteria, entity *entity.RecipeCategory) (*entity.RecipeCategory, error)
	UpdateMany(criteria *persistence.Criteria, entities []*entity.RecipeCategory) ([]*entity.RecipeCategory, error)
	DeleteOne(criteria *persistence.Criteria) (bool, error)
	GetCriteria() *CriteriaRepository
}

type RecipeIngredientRepositoryInterface interface {
	FindOne(criteria *persistence.Criteria) (*entity.RecipeIngredient, error)
	FindAll(criteria *persistence.Criteria) ([]*entity.RecipeIngredient, error)
	InsertOne(recipeIngredient *entity.RecipeIngredient) (*entity.RecipeIngredient, error)
	InsertMany(recipeIngredients []entity.RecipeIngredient) ([]entity.RecipeIngredient, error)
	UpdateOne(criteria *persistence.Criteria, entity *entity.RecipeIngredient) (*entity.RecipeIngredient, error)
	UpdateMany(criteria *persistence.Criteria, entities []*entity.RecipeIngredient) ([]*entity.RecipeIngredient, error)
	DeleteOne(criteria *persistence.Criteria) (bool, error)
	GetCriteria() *CriteriaRepository
}

type RecipeProcessRepositoryInterface interface {
	FindOne(criteria *persistence.Criteria) (*entity.RecipeProcess, error)
	FindAll(criteria *persistence.Criteria) ([]*entity.RecipeProcess, error)
	InsertOne(recipeProcess *entity.RecipeProcess) (*entity.RecipeProcess, error)
	InsertMany(recipeProcesses []entity.RecipeProcess) ([]entity.RecipeProcess, error)
	UpdateOne(criteria *persistence.Criteria, entity *entity.RecipeProcess) (*entity.RecipeProcess, error)
	UpdateMany(criteria *persistence.Criteria, entities []*entity.RecipeProcess) ([]*entity.RecipeProcess, error)
	DeleteOne(criteria *persistence.Criteria) (bool, error)
	GetCriteria() *CriteriaRepository
}

type RecipeMeasureRepositoryInterface interface {
	FindOne(criteria *persistence.Criteria) (*entity.RecipeMeasure, error)
	FindAll(criteria *persistence.Criteria) ([]*entity.RecipeMeasure, error)
	InsertOne(recipeMeasure *entity.RecipeMeasure) (*entity.RecipeMeasure, error)
	InsertMany(recipeMeasures []*entity.RecipeMeasure) ([]*entity.RecipeMeasure, error)
	UpdateOne(criteria *persistence.Criteria, entity *entity.RecipeMeasure) (*entity.RecipeMeasure, error)
	UpdateMany(criteria *persistence.Criteria, entities []*entity.RecipeMeasure) ([]*entity.RecipeMeasure, error)
	DeleteOne(criteria *persistence.Criteria) (bool, error)
	GetCriteria() *CriteriaRepository
}
