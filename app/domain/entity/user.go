package entity

import (
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/domain/dto"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	"time"
)

const (
	UserConfirmationValueMin = 100000
	UserConfirmationValueMax = 999999
)

type User struct {
	Id uuid.UUID `bson:"id" json:"id"`
	dto.UserDTO
}

type UserRole struct {
	Id         uuid.UUID           `bson:"id" json:"id"`
	DateInsert time.Time           `bson:"date_insert" json:"date_insert"`
	DateUpdate time.Time           `bson:"date_update" json:"date_update"`
	Name       string              `bson:"name" json:"name"`
	Code       kind.UserRole       `bson:"code" json:"code"`
	Status     kind.UserRoleStatus `bson:"status" json:"status"`
}

type UserToRole struct {
	Id         uuid.UUID             `bson:"id" json:"id"`
	UserId     uuid.UUID             `bson:"user_id" json:"user_id"`
	RoleId     uuid.UUID             `bson:"role_id" json:"role_id"`
	DateInsert time.Time             `bson:"date_insert" json:"date_insert"`
	DateUpdate time.Time             `bson:"date_update" json:"date_update"`
	Rights     []kind.UserRight      `bson:"rights" json:"rights"`
	Status     kind.UserToRoleStatus `bson:"status" json:"status"`
}

type UserConfirmation struct {
	Id         uuid.UUID `bson:"id" json:"id"`
	DateInsert time.Time `bson:"date_insert" json:"date_insert"`
	DateUpdate time.Time `bson:"date_update" json:"date_update"`
	UserId     uuid.UUID `bson:"user_id" json:"user_id"`
	Value      string    `bson:"value" json:"value"`
	Active     bool      `bson:"active" json:"active"`
}
