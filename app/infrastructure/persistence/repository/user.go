package repository

import (
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence"
)

type UserRepositoryInterface interface {
	FindOne(criteria *persistence.Criteria) (*entity.User, error)
	FindAll(criteria *persistence.Criteria) ([]entity.User, error)
	InsertOne(user *entity.User) (*entity.User, error)
	InsertMany(users []entity.User) ([]entity.User, error)
	UpdateOne(criteria *persistence.Criteria, entity *entity.User) (*entity.User, error)
	UpdateMany(criteria *persistence.Criteria, entities []*entity.User) ([]*entity.User, error)
	DeleteOne(criteria *persistence.Criteria) (bool, error)
	GetCriteriaByUserId(id *uuid.UUID) *persistence.Criteria
	GetCriteriaByUsername(username string) *persistence.Criteria
}

type UserRoleRepositoryInterface interface {
	FindOne(criteria *persistence.Criteria) (*entity.UserRole, error)
	FindAll(criteria *persistence.Criteria) ([]entity.UserRole, error)
}

type UserToRoleRepositoryInterface interface {
	FindOne(criteria *persistence.Criteria) (*entity.UserToRole, error)
	FindAll(criteria *persistence.Criteria) ([]entity.UserToRole, error)
}

type UserConfirmationRepositoryInterface interface {
	FindOne(criteria *persistence.Criteria) (*entity.UserConfirmation, error)
	FindAll(criteria *persistence.Criteria) ([]entity.UserConfirmation, error)
	InsertOne(user *entity.UserConfirmation) (*entity.UserConfirmation, error)
	InsertMany(users []entity.UserConfirmation) ([]entity.UserConfirmation, error)
	UpdateOne(criteria *persistence.Criteria, entity *entity.UserConfirmation) (*entity.UserConfirmation, error)
	UpdateMany(criteria *persistence.Criteria, entities []*entity.UserConfirmation) ([]*entity.UserConfirmation, error)
	GetCriteriaByUserIdAndActive(user *entity.User) *persistence.Criteria
	GetCriteriaById(id *uuid.UUID) *persistence.Criteria
}
