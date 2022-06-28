package null

import (
	"github.com/sergeygardner/meal-planner-api/infrastructure/cache"
)

type CacheManager struct {
	Namespace string
	Type      cache.Type
	TTL       int64
	dsn       *cache.DSN
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

func (cm *CacheManager) Set(_ []byte, _ any, _ *int64) {}

func (cm *CacheManager) Get(_ []byte) (any, error) {
	return nil, nil
}

func (cm *CacheManager) Delete(_ []byte) error {
	return nil
}

func (cm *CacheManager) Exists(_ []byte) error {
	return nil
}

func (cm *CacheManager) GetSet(_ []byte, data any, _ *int64) (any, error) {
	return data, nil
}

func (cm *CacheManager) SetDriver() {}
