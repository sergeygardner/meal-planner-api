package persistence

import (
	"fmt"
	"reflect"
)

type Type struct {
	slug string
}

func (t Type) String() string {
	return t.slug
}

var (
	MongoType   = Type{"mongo"}
	DefaultType = Type{"mongo"}
)

type DSN struct {
	DSN      string
	Host     string
	Port     string
	User     string
	Password string
	DB       string
	Type     string
}

type Criteria struct {
	Where  map[string]interface{}
	Order  map[string]interface{}
	Limit  int
	Offset int
}

func (c Criteria) String() string {
	var (
		where  map[string]interface{}
		order  map[string]interface{}
		limit  int
		offset int
	)

	if reflect.ValueOf(c.Where).IsZero() {
		where = map[string]interface{}{}
	}

	if reflect.ValueOf(c.Order).IsZero() {
		order = map[string]interface{}{}
	}

	if reflect.ValueOf(c.Limit).IsZero() {
		limit = 0
	}

	if reflect.ValueOf(c.Offset).IsZero() {
		offset = 0
	}

	return fmt.Sprintf(
		"%v%v%d%d",
		where,
		order,
		offset,
		limit,
	)
}

type Wrapper struct {
	Set any
	Mod any
}

type EntityManagerInterface interface {
	SetDSN(DSN *DSN)
	GetDSN() *DSN
	GetType() Type
	FindOne(table string, criteria *Criteria) (interface{}, error)
	FindAll(table string, criteria *Criteria) ([]interface{}, error)
	InsertOne(table string, entity interface{}) (interface{}, error)
	InsertMany(table string, entities []interface{}) ([]interface{}, error)
	UpdateOne(table string, criteria *Criteria, wrapper *Wrapper) (interface{}, error)
	UpdateMany(table string, criteria *Criteria, wrapper *Wrapper) ([]interface{}, error)
	DeleteOne(table string, criteria *Criteria) (bool, error)
}
