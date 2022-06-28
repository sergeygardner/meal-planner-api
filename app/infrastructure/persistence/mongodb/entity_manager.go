package mongodb

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sergeygardner/meal-planner-api/infrastructure/cache"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
)

var (
	errorDeleteOneEntity = errors.New("an error occurred while deleting one entity")
)

// EntityManager /**/
type EntityManager struct {
	client       *mongo.Client
	context      context.Context
	cancel       context.CancelFunc
	Database     string
	Type         persistence.Type
	CacheManager cache.ManagerInterface
	dsn          *persistence.DSN
	persistence.EntityManagerInterface
}

func (em *EntityManager) getConnection() *mongo.Client {
	if em.client == nil {
		em.setConnection()
	}

	return em.client
}

func (em *EntityManager) setConnection() {
	dsn := em.GetDSN()
	uri := dsn.DSN

	if uri == "" {
		uri = fmt.Sprintf(
			"mongodb://%s:%s@%s:%s/",
			dsn.User,
			dsn.Password,
			dsn.Host,
			dsn.Port,
		)
	}

	connectionContext, cancel := context.WithCancel(context.Background())
	logger := options.Logger()
	logger.SetComponentLevel(options.LogComponentAll, options.LogLevelDebug)
	client, errorConnect := mongo.Connect(
		connectionContext,
		options.Client().ApplyURI(uri),
		options.Client().SetLoggerOptions(logger),
	)

	if errorConnect != nil {
		panic("Can not connect to mongo db")
	}

	em.context = connectionContext
	em.client = client
	em.cancel = cancel
}

func (em *EntityManager) SetDSN(DSN *persistence.DSN) {
	em.dsn = DSN
}

func (em *EntityManager) GetDSN() *persistence.DSN {
	return em.dsn
}

func (em *EntityManager) GetType() persistence.Type {
	return em.Type
}

func (em *EntityManager) FindOne(table string, criteria *persistence.Criteria) (interface{}, error) {
	keyCache := cache.PrepareKey(table, criteria)
	bsonMResultCached, errorGet := em.CacheManager.Get(keyCache)

	if errorGet == nil {
		bsonMResultRestored, statusRestored := bsonMResultCached.(bson.M)

		if !statusRestored {
			return nil, errors.Wrapf(errorGet, "an error occurred while getting a result from the cache by provided data %p", criteria)
		} else {
			return bsonMResultRestored, nil
		}
	}

	bsonMResult := bson.M{}
	criteriaToBSONCriteria := em.convertCriteriaToBSONCriteria(criteria)
	errorFindOne := em.getConnection().Database(em.Database).Collection(table).FindOne(em.context, criteriaToBSONCriteria).Decode(&bsonMResult)

	if errorFindOne != nil {
		return nil, errors.Wrapf(errorFindOne, "an error occurred while getting a result from the database by provided data %p", criteria)
	}

	em.CacheManager.Set(keyCache, bsonMResult, nil)

	return bsonMResult, nil
}

func (em *EntityManager) FindAll(table string, criteria *persistence.Criteria) ([]interface{}, error) {
	var entities []interface{}

	keyCache := cache.PrepareKey(table, criteria)
	bsonMResultCached, errorGet := em.CacheManager.Get(keyCache)

	if errorGet == nil {
		bsonMResultRestored, statusRestored := bsonMResultCached.([]interface{})

		if !statusRestored {
			return nil, errors.Wrapf(errorGet, "an error occurred while getting a result from the cache by provided data %p", criteria)
		} else {
			for _, value := range bsonMResultRestored {
				valueRestored, statusValueRestored := value.(bson.M)

				if statusValueRestored {
					entities = append(entities, valueRestored)
				}
			}

			return entities, nil
		}
	}

	cursor, errorFind := em.getConnection().Database(em.Database).Collection(table).Find(em.context, em.convertCriteriaToBSONCriteria(criteria))

	if errorFind != nil {
		return nil, errors.Wrapf(errorFind, "an error occurred while getting results from the database by provided data %p", criteria)
	}

	for cursor.Next(context.TODO()) {
		entity := bson.M{}
		errorDecode := cursor.Decode(&entity)
		if errorDecode != nil {
			return nil, errors.Wrapf(errorDecode, "an error occurred while decoding a result from the database by provided data %p", criteria)
		}

		entities = append(entities, entity)
	}

	if errorCursor := cursor.Err(); errorCursor != nil {
		return nil, errors.Wrapf(errorCursor, "an error occurred while processing results from the database by provided data %p", criteria)
	}

	errorClose := cursor.Close(context.TODO())

	if errorClose != nil {
		return nil, errors.Wrap(errorClose, "an error occurred while closing a cursor")
	}

	em.CacheManager.Set(keyCache, entities, nil)

	return entities, nil
}

func (em *EntityManager) InsertOne(table string, entity interface{}) (interface{}, error) {
	_, errorInsertOne := em.getConnection().Database(em.Database).Collection(table).InsertOne(em.context, entity)

	if errorInsertOne != nil {
		return nil, errors.Wrapf(errorInsertOne, "an error occurred while inserting an entity to the database by provided data %s", entity)
	}

	return entity, nil
}

func (em *EntityManager) InsertMany(table string, entities []interface{}) ([]interface{}, error) {
	_, errorInsertMany := em.getConnection().Database(em.Database).Collection(table).InsertMany(em.context, entities)

	if errorInsertMany != nil {
		return nil, errors.Wrapf(errorInsertMany, "an error occurred while inserting entities to the database by provided data %s", entities)
	}

	return entities, nil
}

func (em *EntityManager) UpdateOne(table string, criteria *persistence.Criteria, wrapper *persistence.Wrapper) (interface{}, error) {
	_, errorUpdateOne := em.getConnection().Database(em.Database).Collection(table).UpdateOne(em.context, em.convertCriteriaToBSONCriteria(criteria), em.convertWrapperToBSONWrapper(wrapper))

	if errorUpdateOne != nil {
		return nil, errors.Wrapf(errorUpdateOne, "an error occurred while updating an entity in the database by provided data criteria=%p wrapper=%p", criteria, wrapper)
	}

	if wrapper.Set == nil {
		return nil, errors.New("wrapper doesn't have an entity")
	}

	return wrapper.Set, nil
}

func (em *EntityManager) UpdateMany(table string, criteria *persistence.Criteria, wrapper *persistence.Wrapper) ([]interface{}, error) {
	_, errorUpdateMany := em.getConnection().Database(em.Database).Collection(table).UpdateMany(em.context, em.convertCriteriaToBSONCriteria(criteria), em.convertWrapperToBSONWrapper(wrapper))

	if errorUpdateMany != nil {
		return nil, errors.Wrapf(errorUpdateMany, "an error occurred while updating entities in the database by provided data criteria=%p wrapper=%p", criteria, wrapper)
	}

	return []interface{}{}, nil
}

func (em *EntityManager) DeleteOne(table string, criteria *persistence.Criteria) (bool, error) {
	deleteResult, errorDeleteOne := em.getConnection().Database(em.Database).Collection(table).DeleteOne(em.context, em.convertCriteriaToBSONCriteria(criteria))

	if errorDeleteOne != nil {
		return false, errors.Wrapf(errorDeleteOne, "an error occurred while deleteing an entity in the database by provided data criteria=%p", criteria)
	} else if deleteResult.DeletedCount != 1 {
		return false, errorDeleteOneEntity
	}

	return true, nil
}

func (em *EntityManager) convertCriteriaToBSONCriteria(criteria *persistence.Criteria) bson.M {
	bsonCriteria := bson.M{}

	if criteria.Where != nil {
		for key, value := range criteria.Where {
			reflectValue := reflect.ValueOf(value)

			if reflectValue.Kind() == reflect.Slice {
				valueBSONA := bson.A{}

				for i := 0; i < reflectValue.Len(); i++ {
					valueBSONA = append(valueBSONA, reflectValue.Index(i).Interface())
				}

				bsonCriteria[key] = bson.D{{"$in", valueBSONA}}
			} else {
				bsonCriteria[key] = value
			}
		}
	}

	return bsonCriteria
}

func (em *EntityManager) convertWrapperToBSONWrapper(wrapper *persistence.Wrapper) bson.M {
	bsonWrapper := bson.M{}

	if wrapper.Set != nil {
		if wrapper.Set != nil {
			bsonWrapper["$set"] = wrapper.Set
		}

		if wrapper.Mod != nil {
			bsonWrapper["$mod"] = wrapper.Mod
		}
	}

	return bsonWrapper
}
