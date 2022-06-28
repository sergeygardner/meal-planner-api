package in_memory

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sergeygardner/meal-planner-api/infrastructure/cache"
	"time"
)

var (
	errorCachedItemDoesNotExist = errors.New("an error occurred while getting a cache item cause one does not exist.")
	errorCachedItemIsExpired    = errors.New("an error occurred while getting a cache item cause a TTL of one is expired.")
)

type CacheManager struct {
	context          context.Context
	cancel           context.CancelFunc
	Namespace        string
	Type             cache.Type
	TTL              int64
	dsn              *cache.DSN
	cachedNamespaces *cachedNamespaces
	cache.ManagerInterface
}

type cachedNamespaces map[string]*cachedItems
type cachedItems map[string]*cachedItem
type cachedItem struct {
	TTL   time.Time
	Value any
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
	cachedNamespace := cm.setCachedNamespace(cm.Namespace)

	if ttl == nil {
		ttl = &cm.TTL
	}

	(*cachedNamespace)[string(key)] = &cachedItem{TTL: time.Now().UTC().Add(time.Duration(*ttl) * time.Minute), Value: data}
}

func (cm *CacheManager) Get(key []byte) (any, error) {
	cachedNamespace := cm.setCachedNamespace(cm.Namespace)

	cachedNamespaceValue, cachedNamespaceValueOk := (*cachedNamespace)[string(key)]

	if !cachedNamespaceValueOk {
		return nil, errorCachedItemDoesNotExist
	}

	if cachedNamespaceValue.TTL.Before(time.Now().UTC()) {
		return nil, errorCachedItemIsExpired
	}

	return cachedNamespaceValue.Value, nil
}

func (cm *CacheManager) Delete(key []byte) error {
	errorExisting := cm.Exists(key)

	if errorExisting != nil {
		return errorExisting
	}

	cachedNamespace := cm.setCachedNamespace(cm.Namespace)

	delete(*cachedNamespace, string(key))

	return nil
}

func (cm *CacheManager) Exists(key []byte) error {
	cachedNamespace := cm.setCachedNamespace(cm.Namespace)
	stringedKey := string(key)
	_, cachedNamespaceValueOk := (*cachedNamespace)[stringedKey]

	if !cachedNamespaceValueOk {
		return errorCachedItemDoesNotExist
	}

	return nil
}

func (cm *CacheManager) GetSet(key []byte, data any, ttl *int64) (any, error) {
	cm.Set(key, data, ttl)

	return cm.Get(key)
}

func (cm *CacheManager) SetDriver() {
	if cm.cachedNamespaces == nil {
		cm.cachedNamespaces = &cachedNamespaces{}
	}
}

func (cm *CacheManager) setCachedNamespace(namespace string) *cachedItems {
	namespaceData, namespaceDataOk := (*cm.cachedNamespaces)[namespace]

	if !namespaceDataOk {
		namespaceData = &cachedItems{}

		(*cm.cachedNamespaces)[namespace] = namespaceData
	}

	return namespaceData
}
