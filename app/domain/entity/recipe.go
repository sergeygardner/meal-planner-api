package entity

import (
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	"time"
)

type Recipe struct {
	Id          uuid.UUID         `bson:"id" json:"id"`
	UserId      uuid.UUID         `bson:"user_id" json:"user_id"`
	DateInsert  time.Time         `bson:"date_insert" json:"date_insert"`
	DateUpdate  time.Time         `bson:"date_update" json:"date_update"`
	Name        string            `bson:"name" json:"name"`
	Description string            `bson:"description" json:"description"`
	Notes       string            `bson:"notes" json:"notes"`
	Status      kind.RecipeStatus `bson:"status" json:"status"`
}

type RecipeCategory struct {
	Id         uuid.UUID                 `bson:"id" json:"id"`
	UserId     uuid.UUID                 `bson:"user_id" json:"user_id"`
	EntityId   uuid.UUID                 `bson:"entity_id" json:"entity_id"`
	DeriveId   uuid.UUID                 `bson:"derive_id" json:"derive_id"`
	DateInsert time.Time                 `bson:"date_insert" json:"date_insert"`
	DateUpdate time.Time                 `bson:"date_update" json:"date_update"`
	Status     kind.RecipeCategoryStatus `bson:"status" json:"status"`
}

type RecipeIngredient struct {
	Id         uuid.UUID                   `bson:"id" json:"id"`
	UserId     uuid.UUID                   `bson:"user_id" json:"user_id"`
	EntityId   uuid.UUID                   `bson:"entity_id" json:"entity_id"`
	DeriveId   uuid.UUID                   `bson:"derive_id" json:"derive_id"`
	DateInsert time.Time                   `bson:"date_insert" json:"date_insert"`
	DateUpdate time.Time                   `bson:"date_update" json:"date_update"`
	Name       string                      `bson:"name" json:"name"`
	Status     kind.RecipeIngredientStatus `bson:"status" json:"status"`
}

type RecipeMeasure struct {
	Id         uuid.UUID                `bson:"id" json:"id"`
	UserId     uuid.UUID                `bson:"user_id" json:"user_id"`
	EntityId   uuid.UUID                `bson:"entity_id" json:"entity_id"`
	UnitId     uuid.UUID                `bson:"unit_id" json:"unit_id"`
	DateInsert time.Time                `bson:"date_insert" json:"date_insert"`
	DateUpdate time.Time                `bson:"date_update" json:"date_update"`
	Value      int64                    `bson:"value" json:"value"`
	Status     kind.RecipeMeasureStatus `bson:"status" json:"status"`
}

type RecipeProcess struct {
	Id          uuid.UUID                `bson:"id" json:"id"`
	UserId      uuid.UUID                `bson:"user_id" json:"user_id"`
	EntityId    uuid.UUID                `bson:"entity_id" json:"entity_id"`
	DateInsert  time.Time                `bson:"date_insert" json:"date_insert"`
	DateUpdate  time.Time                `bson:"date_update" json:"date_update"`
	Name        string                   `bson:"name" json:"name"`
	Description string                   `bson:"description" json:"description"`
	Notes       string                   `bson:"notes" json:"notes"`
	Status      kind.RecipeProcessStatus `bson:"status" json:"status"`
}
