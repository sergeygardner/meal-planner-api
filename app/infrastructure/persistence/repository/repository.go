package repository

import (
	"github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence"
)

type AltNameRepositoryInterface interface {
	FindOne(criteria *persistence.Criteria) (*entity.AltName, error)
	FindAll(criteria *persistence.Criteria) ([]*entity.AltName, error)
	InsertOne(recipeAltName *entity.AltName) (*entity.AltName, error)
	InsertMany(recipeAltNames []*entity.AltName) ([]*entity.AltName, error)
	UpdateOne(criteria *persistence.Criteria, entity *entity.AltName) (*entity.AltName, error)
	UpdateMany(criteria *persistence.Criteria, entities []*entity.AltName) ([]*entity.AltName, error)
	DeleteOne(criteria *persistence.Criteria) (bool, error)
	GetCriteria() *CriteriaRepository
}

type PictureRepositoryInterface interface {
	FindOne(criteria *persistence.Criteria) (*entity.Picture, error)
	FindAll(criteria *persistence.Criteria) ([]*entity.Picture, error)
	InsertOne(recipePicture *entity.Picture) (*entity.Picture, error)
	InsertMany(recipePictures []*entity.Picture) ([]*entity.Picture, error)
	UpdateOne(criteria *persistence.Criteria, entity *entity.Picture) (*entity.Picture, error)
	UpdateMany(criteria *persistence.Criteria, entities []*entity.Picture) ([]*entity.Picture, error)
	DeleteOne(criteria *persistence.Criteria) (bool, error)
	GetCriteria() *CriteriaRepository
}
type UnitRepositoryInterface interface {
	FindOne(criteria *persistence.Criteria) (*entity.Unit, error)
	FindAll(criteria *persistence.Criteria) ([]*entity.Unit, error)
	InsertOne(recipeMeasure *entity.Unit) (*entity.Unit, error)
	InsertMany(recipeMeasures []*entity.Unit) ([]*entity.Unit, error)
	UpdateOne(criteria *persistence.Criteria, entity *entity.Unit) (*entity.Unit, error)
	UpdateMany(criteria *persistence.Criteria, entities []*entity.Unit) ([]*entity.Unit, error)
	DeleteOne(criteria *persistence.Criteria) (bool, error)
	GetCriteria() *CriteriaRepository
}

type CategoryRepositoryInterface interface {
	FindOne(criteria *persistence.Criteria) (*entity.Category, error)
	FindAll(criteria *persistence.Criteria) ([]*entity.Category, error)
	InsertOne(recipeMeasure *entity.Category) (*entity.Category, error)
	InsertMany(recipeMeasures []*entity.Category) ([]*entity.Category, error)
	UpdateOne(criteria *persistence.Criteria, entity *entity.Category) (*entity.Category, error)
	UpdateMany(criteria *persistence.Criteria, entities []*entity.Category) ([]*entity.Category, error)
	DeleteOne(criteria *persistence.Criteria) (bool, error)
	GetCriteria() *CriteriaRepository
}

type IngredientRepositoryInterface interface {
	FindOne(criteria *persistence.Criteria) (*entity.Ingredient, error)
	FindAll(criteria *persistence.Criteria) ([]*entity.Ingredient, error)
	InsertOne(recipeMeasure *entity.Ingredient) (*entity.Ingredient, error)
	InsertMany(recipeMeasures []*entity.Ingredient) ([]*entity.Ingredient, error)
	UpdateOne(criteria *persistence.Criteria, entity *entity.Ingredient) (*entity.Ingredient, error)
	UpdateMany(criteria *persistence.Criteria, entities []*entity.Ingredient) ([]*entity.Ingredient, error)
	DeleteOne(criteria *persistence.Criteria) (bool, error)
	GetCriteria() *CriteriaRepository
}
