package entity

import (
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	"time"
)

type Planner struct {
	Id         uuid.UUID          `bson:"id" json:"id"`
	UserId     uuid.UUID          `bson:"user_id" json:"user_id"`
	DateInsert time.Time          `bson:"date_insert" json:"date_insert"`
	DateUpdate time.Time          `bson:"date_update" json:"date_update"`
	StartTime  time.Time          `bson:"start_time" json:"start_time"`
	EndTime    time.Time          `bson:"end_time" json:"end_time"`
	Name       string             `bson:"name" json:"name"`
	Status     kind.PlannerStatus `bson:"status" json:"status"`
}

type PlannerInterval struct {
	Id         uuid.UUID                  `bson:"id" json:"id"`
	UserId     uuid.UUID                  `bson:"user_id" json:"user_id"`
	EntityId   uuid.UUID                  `bson:"entity_id" json:"entity_id"`
	DateInsert time.Time                  `bson:"date_insert" json:"date_insert"`
	DateUpdate time.Time                  `bson:"date_update" json:"date_update"`
	StartTime  time.Time                  `bson:"start_time" json:"start_time"`
	EndTime    time.Time                  `bson:"end_time" json:"end_time"`
	Name       string                     `bson:"name" json:"name"`
	Status     kind.PlannerIntervalStatus `bson:"status" json:"status"`
}

type PlannerRecipe struct {
	Id         uuid.UUID                `bson:"id" json:"id"`
	UserId     uuid.UUID                `bson:"user_id" json:"user_id"`
	EntityId   uuid.UUID                `bson:"entity_id" json:"entity_id"`
	RecipeId   uuid.UUID                `bson:"recipe_id" json:"recipe_id"`
	DateInsert time.Time                `bson:"date_insert" json:"date_insert"`
	DateUpdate time.Time                `bson:"date_update" json:"date_update"`
	Status     kind.PlannerRecipeStatus `bson:"status" json:"status"`
}
