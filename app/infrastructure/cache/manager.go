package cache

import (
	log "github.com/sirupsen/logrus"
	"reflect"
)

type Type struct {
	slug string
}

func (t Type) String() string {
	return t.slug
}

var (
	NullType     = Type{"null"}
	InMemoryType = Type{"inMemory"}
	RedisType    = Type{"redis"}
	DefaultType  = Type{"redis"}
)

type ManagerInterface interface {
	SetDriver()
	SetDSN(dsn *DSN)
	GetDSN() *DSN
	GetType() Type
	Set(key []byte, data any, ttl *int64)
	Get(key []byte) (any, error)
	Delete(key []byte) error
	Exists(key []byte) error
	GetSet(key []byte, data any, ttl *int64) (any, error)
}

type DSN struct {
	DSN       string
	Host      string
	Port      string
	User      string
	Password  string
	Namespace string
	Type      string
}

func PrepareKey(keys ...any) []byte {
	defer func() {
		if recoverValue := recover(); recoverValue != nil {
			log.Panic(recoverValue)
		}
	}()
	var key []byte

	for _, value := range keys {
		reflectValue := reflect.ValueOf(value)

		switch reflectValue.Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
			if reflectValue.IsValid() && !reflectValue.IsZero() && !reflectValue.IsNil() {
				reflectValueWithStringMethod := reflectValue.MethodByName("String")

				if reflectValueWithStringMethod.Kind() == reflect.Func {
					result := reflectValueWithStringMethod.Call([]reflect.Value{})

					for _, resultValue := range result {
						key = append(key, []byte(resultValue.Interface().(string))...)
					}
				}
			}
		case reflect.String:
			key = append(key, []byte(reflectValue.Interface().(string))...)
		default:
		}

	}

	return key
}
