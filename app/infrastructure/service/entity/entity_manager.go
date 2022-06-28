package entity

import (
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence/mongodb"
	"github.com/sergeygardner/meal-planner-api/infrastructure/service/cache"
)

var entityManager persistence.EntityManagerInterface

func GetEntityManager() persistence.EntityManagerInterface {
	return entityManager
}

func SetEntityManager(DSN *persistence.DSN) {
	switch DSN.Type {
	case persistence.MongoType.String():
		entityManager = &mongodb.EntityManager{Database: DSN.DB, Type: persistence.MongoType, CacheManager: cache.GetCacheManager()}
	default:
		entityManager = &mongodb.EntityManager{Database: DSN.DB, Type: persistence.DefaultType, CacheManager: cache.GetCacheManager()}
	}

	entityManager.SetDSN(DSN)
}
