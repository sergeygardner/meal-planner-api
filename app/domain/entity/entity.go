package entity

import (
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	"time"
)

type Unit struct {
	Id         uuid.UUID       `bson:"id" json:"id" validate:"required"`
	DateInsert time.Time       `bson:"date_insert" json:"date_insert" validate:"required"`
	DateUpdate time.Time       `bson:"date_update" json:"date_update" validate:"required"`
	Name       string          `bson:"name" json:"name" validate:"required,min=2,max=255"`
	Status     kind.UnitStatus `bson:"status" json:"status" validate:"required"`
}

type AltName struct {
	Id         uuid.UUID          `bson:"id" json:"id"`
	UserId     uuid.UUID          `bson:"user_id" json:"user_id"`
	EntityId   uuid.UUID          `bson:"entity_id" json:"entity_id"`
	DateInsert time.Time          `bson:"date_insert" json:"date_insert"`
	DateUpdate time.Time          `bson:"date_update" json:"date_update"`
	Name       string             `bson:"name" json:"name"`
	Status     kind.AltNameStatus `bson:"status" json:"status"`
}

type Picture struct {
	Id         uuid.UUID          `bson:"id" json:"id"`
	UserId     uuid.UUID          `bson:"user_id" json:"user_id"`
	EntityId   uuid.UUID          `bson:"entity_id" json:"entity_id"`
	DateInsert time.Time          `bson:"date_insert" json:"date_insert"`
	DateUpdate time.Time          `bson:"date_update" json:"date_update"`
	Name       string             `bson:"name" json:"name"`
	URL        string             `bson:"url" json:"url"`
	Width      int64              `bson:"width" json:"width"`
	Height     int64              `bson:"height" json:"height"`
	Size       int64              `bson:"size" json:"size"`
	Type       string             `bson:"type" json:"type"`
	Status     kind.PictureStatus `bson:"status" json:"status"`
}

type Category struct {
	Id         uuid.UUID           `bson:"id" json:"id"`
	UserId     uuid.UUID           `bson:"user_id" json:"user_id"`
	DateInsert time.Time           `bson:"date_insert" json:"date_insert"`
	DateUpdate time.Time           `bson:"date_update" json:"date_update"`
	Name       string              `bson:"name" json:"name"`
	Status     kind.CategoryStatus `bson:"status" json:"status"`
}

type Ingredient struct {
	Id         uuid.UUID             `bson:"id" json:"id"`
	UserId     uuid.UUID             `bson:"user_id" json:"user_id"`
	DateInsert time.Time             `bson:"date_insert" json:"date_insert"`
	DateUpdate time.Time             `bson:"date_update" json:"date_update"`
	Name       string                `bson:"name" json:"name"`
	Status     kind.IngredientStatus `bson:"status" json:"status"`
}
