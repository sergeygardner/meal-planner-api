package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sergeygardner/meal-planner-api/infrastructure/cache"
	"strconv"
)

type CacheManager struct {
	context   context.Context
	cancel    context.CancelFunc
	Namespace string
	Type      cache.Type
	TTL       int64
	dsn       *cache.DSN
	driver    *redis.Ring
	cache.ManagerInterface
}

func (cm *CacheManager) SetDSN(DSN *cache.DSN) {
	cm.dsn = DSN
}

func (cm *CacheManager) GetDSN() *cache.DSN {
	return cm.dsn
}

func (cm *CacheManager) GetType() cache.Type {
	return cm.Type
}

func (cm *CacheManager) Set(key []byte, data any, ttl *int64) {

}

func (cm *CacheManager) Get(key []byte) (any, error) {
	return nil, nil
}

func (cm *CacheManager) Delete(key []byte) error {
	return nil
}

func (cm *CacheManager) Exists(key []byte) error {
	return nil
}

func (cm *CacheManager) GetSet(key []byte, data any, ttl *int64) (any, error) {
	return nil, nil
}

func (cm *CacheManager) SetDriver() {
	if cm.driver != nil {
		return
	}

	namespace, _ := strconv.Atoi(cm.dsn.Namespace)

	cm.driver = redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			cm.dsn.Host: cm.dsn.Port,
		},
		Username: cm.dsn.User,
		Password: cm.dsn.Password,
		DB:       namespace,
	})
}
