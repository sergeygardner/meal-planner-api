package cache

import (
	"github.com/sergeygardner/meal-planner-api/infrastructure/cache"
	"github.com/sergeygardner/meal-planner-api/infrastructure/cache/in_memory"
	"github.com/sergeygardner/meal-planner-api/infrastructure/cache/null"
	"github.com/sergeygardner/meal-planner-api/infrastructure/cache/redis"
)

var cacheManager cache.ManagerInterface

func GetCacheManager() cache.ManagerInterface {
	return cacheManager
}

func SetCacheManager(DSN *cache.DSN, ttl int64) {
	switch DSN.Type {
	case cache.RedisType.String():
		cacheManager = &redis.CacheManager{Namespace: DSN.Namespace, Type: cache.RedisType, TTL: ttl}
	case cache.InMemoryType.String():
		cacheManager = &in_memory.CacheManager{Namespace: DSN.Namespace, Type: cache.InMemoryType, TTL: ttl}
	case cache.NullType.String():
		cacheManager = &null.CacheManager{Namespace: DSN.Namespace, Type: cache.InMemoryType, TTL: ttl}
	default:
		cacheManager = &redis.CacheManager{Namespace: DSN.Namespace, Type: cache.DefaultType, TTL: ttl}
	}

	cacheManager.SetDSN(DSN)
	cacheManager.SetDriver()
}
