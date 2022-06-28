package repository

import (
	DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type PlannerRepository struct {
	EntityManager persistence.EntityManagerInterface
	Table         string
	repository.PlannerRepositoryInterface
}

func (pr *PlannerRepository) FindOne(criteria *persistence.Criteria) (*DomainEntity.Planner, error) {
	entity, errorFindOne := pr.EntityManager.FindOne(pr.Table, criteria)

	if errorFindOne != nil {
		return nil, errorFindOne
	}
	//todo
	entityBsonM, _ := entity.(bson.M)
	result := DomainEntity.Planner{}
	bsonBytes, _ := bson.Marshal(entityBsonM)
	_ = bson.Unmarshal(bsonBytes, &result)

	return &result, nil
}

func (pr *PlannerRepository) FindAll(criteria *persistence.Criteria) ([]*DomainEntity.Planner, error) {
	var recipePlanners []*DomainEntity.Planner

	entities, errorFindAll := pr.EntityManager.FindAll(pr.Table, criteria)

	if errorFindAll != nil {
		return nil, errorFindAll
	}

	for _, entity := range entities {
		entityBsonM, _ := entity.(bson.M)
		result := DomainEntity.Planner{}
		bsonBytes, errorBSONBytesMarshal := bson.Marshal(entityBsonM)

		if errorBSONBytesMarshal != nil {
			return nil, errorBSONBytesMarshal
		}

		errorBSONBytesUnMarshal := bson.Unmarshal(bsonBytes, &result)

		if errorBSONBytesUnMarshal != nil {
			return nil, errorBSONBytesUnMarshal
		}

		recipePlanners = append(recipePlanners, &result)
	}

	return recipePlanners, nil
}

func (pr *PlannerRepository) InsertOne(entity *DomainEntity.Planner) (*DomainEntity.Planner, error) {
	_, errorInsertOne := pr.EntityManager.InsertOne(pr.Table, entity)

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (pr *PlannerRepository) InsertMany(entities []*DomainEntity.Planner) ([]*DomainEntity.Planner, error) {
	entitiesAsInterfaces := make([]interface{}, len(entities))

	for i, value := range entities {
		entitiesAsInterfaces[i] = &value
	}

	_, errorInsertMany := pr.EntityManager.InsertMany(pr.Table, entitiesAsInterfaces)

	if errorInsertMany != nil {
		return nil, errorInsertMany
	}

	return entities, nil
}

func (pr *PlannerRepository) UpdateOne(criteria *persistence.Criteria, entity *DomainEntity.Planner) (*DomainEntity.Planner, error) {
	_, errorInsertOne := pr.EntityManager.UpdateOne(pr.Table, criteria, &persistence.Wrapper{Set: *entity})

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (pr *PlannerRepository) UpdateMany(criteria *persistence.Criteria, entities []*DomainEntity.Planner) ([]*DomainEntity.Planner, error) {
	entitiesAsInterfaces := make([]interface{}, len(entities))

	for i, value := range entities {
		entitiesAsInterfaces[i] = &value
	}

	_, errorUpdateMany := pr.EntityManager.UpdateMany(pr.Table, criteria, &persistence.Wrapper{Mod: entitiesAsInterfaces})

	if errorUpdateMany != nil {
		return nil, errorUpdateMany
	}

	return entities, nil
}

func (pr *PlannerRepository) DeleteOne(criteria *persistence.Criteria) (bool, error) {
	return pr.EntityManager.DeleteOne(pr.Table, criteria)
}

func (pr *PlannerRepository) GetCriteria() *repository.CriteriaRepository {
	return &CriteriaRepository
}

type PlannerIntervalRepository struct {
	EntityManager persistence.EntityManagerInterface
	Table         string
	repository.PlannerIntervalRepositoryInterface
}

func (ur *PlannerIntervalRepository) FindOne(criteria *persistence.Criteria) (*DomainEntity.PlannerInterval, error) {
	entity, errorFindOne := ur.EntityManager.FindOne(ur.Table, criteria)

	if errorFindOne != nil {
		return nil, errorFindOne
	}
	//todo
	entityBsonM, _ := entity.(bson.M)
	result := DomainEntity.PlannerInterval{}
	bsonBytes, _ := bson.Marshal(entityBsonM)
	_ = bson.Unmarshal(bsonBytes, &result)

	return &result, nil
}

func (ur *PlannerIntervalRepository) FindAll(criteria *persistence.Criteria) ([]*DomainEntity.PlannerInterval, error) {
	var recipePlannerIntervals []*DomainEntity.PlannerInterval

	entities, errorFindAll := ur.EntityManager.FindAll(ur.Table, criteria)

	if errorFindAll != nil {
		return nil, errorFindAll
	}

	for _, entity := range entities {
		entityBsonM, _ := entity.(bson.M)
		result := DomainEntity.PlannerInterval{}
		bsonBytes, errorBSONBytesMarshal := bson.Marshal(entityBsonM)

		if errorBSONBytesMarshal != nil {
			return nil, errorBSONBytesMarshal
		}

		errorBSONBytesUnMarshal := bson.Unmarshal(bsonBytes, &result)

		if errorBSONBytesUnMarshal != nil {
			return nil, errorBSONBytesUnMarshal
		}

		recipePlannerIntervals = append(recipePlannerIntervals, &result)
	}

	return recipePlannerIntervals, nil
}

func (ur *PlannerIntervalRepository) InsertOne(entity *DomainEntity.PlannerInterval) (*DomainEntity.PlannerInterval, error) {
	_, errorInsertOne := ur.EntityManager.InsertOne(ur.Table, entity)

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *PlannerIntervalRepository) InsertMany(entities []*DomainEntity.PlannerInterval) ([]*DomainEntity.PlannerInterval, error) {
	entitiesAsInterfaces := make([]interface{}, len(entities))

	for i, value := range entities {
		entitiesAsInterfaces[i] = &value
	}

	_, errorInsertMany := ur.EntityManager.InsertMany(ur.Table, entitiesAsInterfaces)

	if errorInsertMany != nil {
		return nil, errorInsertMany
	}

	return entities, nil
}

func (ur *PlannerIntervalRepository) UpdateOne(criteria *persistence.Criteria, entity *DomainEntity.PlannerInterval) (*DomainEntity.PlannerInterval, error) {
	_, errorInsertOne := ur.EntityManager.UpdateOne(ur.Table, criteria, &persistence.Wrapper{Set: *entity})

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *PlannerIntervalRepository) UpdateMany(criteria *persistence.Criteria, entities []*DomainEntity.PlannerInterval) ([]*DomainEntity.PlannerInterval, error) {
	entitiesAsInterfaces := make([]interface{}, len(entities))

	for i, value := range entities {
		entitiesAsInterfaces[i] = &value
	}

	_, errorUpdateMany := ur.EntityManager.UpdateMany(ur.Table, criteria, &persistence.Wrapper{Mod: entitiesAsInterfaces})

	if errorUpdateMany != nil {
		return nil, errorUpdateMany
	}

	return entities, nil
}

func (ur *PlannerIntervalRepository) DeleteOne(criteria *persistence.Criteria) (bool, error) {
	return ur.EntityManager.DeleteOne(ur.Table, criteria)
}

func (ur *PlannerIntervalRepository) GetCriteria() *repository.CriteriaRepository {
	return &CriteriaRepository
}

type PlannerRecipeRepository struct {
	EntityManager persistence.EntityManagerInterface
	Table         string
	repository.PlannerRecipeRepositoryInterface
}

func (ur *PlannerRecipeRepository) FindOne(criteria *persistence.Criteria) (*DomainEntity.PlannerRecipe, error) {
	entity, errorFindOne := ur.EntityManager.FindOne(ur.Table, criteria)

	if errorFindOne != nil {
		return nil, errorFindOne
	}
	//todo
	entityBsonM, _ := entity.(bson.M)
	result := DomainEntity.PlannerRecipe{}
	bsonBytes, _ := bson.Marshal(entityBsonM)
	_ = bson.Unmarshal(bsonBytes, &result)

	return &result, nil
}

func (ur *PlannerRecipeRepository) FindAll(criteria *persistence.Criteria) ([]*DomainEntity.PlannerRecipe, error) {
	var recipePlannerRecipes []*DomainEntity.PlannerRecipe

	entities, errorFindAll := ur.EntityManager.FindAll(ur.Table, criteria)

	if errorFindAll != nil {
		return nil, errorFindAll
	}

	for _, entity := range entities {
		entityBsonM, _ := entity.(bson.M)
		result := DomainEntity.PlannerRecipe{}
		bsonBytes, errorBSONBytesMarshal := bson.Marshal(entityBsonM)

		if errorBSONBytesMarshal != nil {
			return nil, errorBSONBytesMarshal
		}

		errorBSONBytesUnMarshal := bson.Unmarshal(bsonBytes, &result)

		if errorBSONBytesUnMarshal != nil {
			return nil, errorBSONBytesUnMarshal
		}

		recipePlannerRecipes = append(recipePlannerRecipes, &result)
	}

	return recipePlannerRecipes, nil
}

func (ur *PlannerRecipeRepository) InsertOne(entity *DomainEntity.PlannerRecipe) (*DomainEntity.PlannerRecipe, error) {
	_, errorInsertOne := ur.EntityManager.InsertOne(ur.Table, entity)

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *PlannerRecipeRepository) InsertMany(entities []*DomainEntity.PlannerRecipe) ([]*DomainEntity.PlannerRecipe, error) {
	entitiesAsInterfaces := make([]interface{}, len(entities))

	for i, value := range entities {
		entitiesAsInterfaces[i] = &value
	}

	_, errorInsertMany := ur.EntityManager.InsertMany(ur.Table, entitiesAsInterfaces)

	if errorInsertMany != nil {
		return nil, errorInsertMany
	}

	return entities, nil
}

func (ur *PlannerRecipeRepository) UpdateOne(criteria *persistence.Criteria, entity *DomainEntity.PlannerRecipe) (*DomainEntity.PlannerRecipe, error) {
	_, errorInsertOne := ur.EntityManager.UpdateOne(ur.Table, criteria, &persistence.Wrapper{Set: *entity})

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *PlannerRecipeRepository) UpdateMany(criteria *persistence.Criteria, entities []*DomainEntity.PlannerRecipe) ([]*DomainEntity.PlannerRecipe, error) {
	entitiesAsInterfaces := make([]interface{}, len(entities))

	for i, value := range entities {
		entitiesAsInterfaces[i] = &value
	}

	_, errorUpdateMany := ur.EntityManager.UpdateMany(ur.Table, criteria, &persistence.Wrapper{Mod: entitiesAsInterfaces})

	if errorUpdateMany != nil {
		return nil, errorUpdateMany
	}

	return entities, nil
}

func (ur *PlannerRecipeRepository) DeleteOne(criteria *persistence.Criteria) (bool, error) {
	return ur.EntityManager.DeleteOne(ur.Table, criteria)
}

func (ur *PlannerRecipeRepository) GetCriteria() *repository.CriteriaRepository {
	return &CriteriaRepository
}
