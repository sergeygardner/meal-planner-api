package service

import (
	"fmt"
	"github.com/sergeygardner/meal-planner-api/infrastructure/cache"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence"
	ServiceCache "github.com/sergeygardner/meal-planner-api/infrastructure/service/cache"
	ServiceEntity "github.com/sergeygardner/meal-planner-api/infrastructure/service/entity"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var (
	cacheTTLConverted  int64
	errorParseCacheTTL error
)

func PreparePersistence() {
	persistenceDSN := &persistence.DSN{}
	dbDSN, dbDSNOk := os.LookupEnv("DB_DSN")
	dbHost, dbHostOk := os.LookupEnv("DB_HOST")
	dbPort, dbPortOk := os.LookupEnv("DB_PORT")
	dbUser, dbUserOk := os.LookupEnv("DB_USER")
	dbPassword, dbPasswordOk := os.LookupEnv("DB_PASSWORD")
	dbName, dbNameOk := os.LookupEnv("DB_NAME")
	dbType, dbTypeOk := os.LookupEnv("DB_TYPE")

	if dbDSNOk && dbTypeOk {
		persistenceDSN.DSN = dbDSN
		persistenceDSN.Type = dbType
	} else if dbHostOk && dbPortOk && dbUserOk && dbPasswordOk && dbNameOk && dbTypeOk {
		persistenceDSN.Host = dbHost
		persistenceDSN.Port = dbPort
		persistenceDSN.User = dbUser
		persistenceDSN.Password = dbPassword
		persistenceDSN.DB = dbName
		persistenceDSN.Type = dbType
	} else {
		errorTitle, _ := fmt.Printf("the db environments are not found DSN (%s) or DB_HOST (%s) or DB_PORT (%s) or DB_USER (%s) or DB_PASSWORD (%s) or DB_NAME (%s) or DB_TYPE (%s)", dbDSN, dbHost, dbPort, dbUser, dbPassword, dbName, dbType)

		log.Panic(errorTitle)

		panic(errorTitle)
	}

	ServiceEntity.SetEntityManager(persistenceDSN)
}
func PrepareCache() {
	dsn := &cache.DSN{}
	cacheDSN, cacheDSNOk := os.LookupEnv("CACHE_DSN")
	cacheHost, cacheHostOk := os.LookupEnv("CACHE_HOST")
	cachePort, cachePortOk := os.LookupEnv("CACHE_PORT")
	cacheUser, cacheUserOk := os.LookupEnv("CACHE_USER")
	cachePassword, cachePasswordOk := os.LookupEnv("CACHE_PASSWORD")
	cacheNamespace, cacheNamespaceOk := os.LookupEnv("CACHE_NAMESPACE")
	cacheType, cacheTypeOk := os.LookupEnv("CACHE_TYPE")
	cacheTTL, cacheTTLOk := os.LookupEnv("CACHE_TTL")

	if !cacheTTLOk {
		cacheTTLConverted = int64(8600)
	} else {
		cacheTTLConverted, errorParseCacheTTL = strconv.ParseInt(cacheTTL, 10, 64)

		if errorParseCacheTTL != nil {
			panic("an error occurred while parsing the CACHE_TTL variable")
		}
	}

	if cacheDSNOk && cacheTypeOk {
		dsn.DSN = cacheDSN
		dsn.Type = cacheType
	} else if cacheHostOk && cachePortOk && cacheUserOk && cachePasswordOk && cacheNamespaceOk && cacheTypeOk {
		dsn.Host = cacheHost
		dsn.Port = cachePort
		dsn.User = cacheUser
		dsn.Password = cachePassword
		dsn.Namespace = cacheNamespace
		dsn.Type = cacheType
	} else {
		errorTitle, _ := fmt.Printf("the cache environments are not found DSN (%s) or CACHE_HOST (%s) or CACHE_PORT (%s) or CACHE_USER (%s) or CACHE_PASSWORD (%s) or CACHE_NAMESPACE (%s) or CACHE_TYPE (%s)", cacheDSN, cacheHost, cachePort, cacheUser, cachePassword, cacheNamespace, cacheType)

		log.Panic(errorTitle)

		panic(errorTitle)
	}

	ServiceCache.SetCacheManager(dsn, cacheTTLConverted)
}
