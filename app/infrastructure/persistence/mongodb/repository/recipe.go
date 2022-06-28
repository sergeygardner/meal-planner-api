package repository

import (
	DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type RecipeRepository struct {
	EntityManager persistence.EntityManagerInterface
	Table         string
	repository.RecipeRepositoryInterface
}

func (ur *RecipeRepository) FindOne(criteria *persistence.Criteria) (*DomainEntity.Recipe, error) {
	entity, errorFindOne := ur.EntityManager.FindOne(ur.Table, criteria)

	if errorFindOne != nil {
		return nil, errorFindOne
	}
	//todo
	entityBsonM, _ := entity.(bson.M)
	result := DomainEntity.Recipe{}
	bsonBytes, _ := bson.Marshal(entityBsonM)
	_ = bson.Unmarshal(bsonBytes, &result)

	return &result, nil
}

func (ur *RecipeRepository) FindAll(criteria *persistence.Criteria) ([]*DomainEntity.Recipe, error) {
	var recipes []*DomainEntity.Recipe

	entities, errorFindAll := ur.EntityManager.FindAll(ur.Table, criteria)

	if errorFindAll != nil {
		return nil, errorFindAll
	}

	for _, entity := range entities {
		entityBsonM, _ := entity.(bson.M)
		result := DomainEntity.Recipe{}
		bsonBytes, errorBSONBytesMarshal := bson.Marshal(entityBsonM)

		if errorBSONBytesMarshal != nil {
			return nil, errorBSONBytesMarshal
		}

		errorBSONBytesUnMarshal := bson.Unmarshal(bsonBytes, &result)

		if errorBSONBytesUnMarshal != nil {
			return nil, errorBSONBytesUnMarshal
		}

		recipes = append(recipes, &result)
	}

	return recipes, nil
}

func (ur *RecipeRepository) InsertOne(entity *DomainEntity.Recipe) (*DomainEntity.Recipe, error) {
	_, errorInsertOne := ur.EntityManager.InsertOne(ur.Table, entity)

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *RecipeRepository) InsertMany(entities []DomainEntity.Recipe) ([]DomainEntity.Recipe, error) {
	entitiesAsInterfaces := make([]interface{}, len(entities))

	for _, value := range entities {
		entitiesAsInterfaces = append(entitiesAsInterfaces, &value)
	}

	_, errorInsertMany := ur.EntityManager.InsertMany(ur.Table, entitiesAsInterfaces)

	if errorInsertMany != nil {
		return nil, errorInsertMany
	}

	return entities, nil
}

func (ur *RecipeRepository) UpdateOne(criteria *persistence.Criteria, entity *DomainEntity.Recipe) (*DomainEntity.Recipe, error) {
	_, errorInsertOne := ur.EntityManager.UpdateOne(ur.Table, criteria, &persistence.Wrapper{Set: *entity})

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *RecipeRepository) UpdateMany(criteria *persistence.Criteria, entities []*DomainEntity.Recipe) ([]*DomainEntity.Recipe, error) {
	entitiesAsInterfaces := make([]interface{}, len(entities))

	for _, value := range entities {
		entitiesAsInterfaces = append(entitiesAsInterfaces, &value)
	}

	_, errorInsertMany := ur.EntityManager.UpdateMany(ur.Table, criteria, &persistence.Wrapper{Set: entitiesAsInterfaces})

	if errorInsertMany != nil {
		return nil, errorInsertMany
	}

	return entities, nil
}

func (ur *RecipeRepository) DeleteOne(criteria *persistence.Criteria) (bool, error) {
	return ur.EntityManager.DeleteOne(ur.Table, criteria)
}

func (ur *RecipeRepository) GetCriteria() *repository.CriteriaRepository {
	return &CriteriaRepository
}

type RecipeCategoryRepository struct {
	EntityManager persistence.EntityManagerInterface
	Table         string
	repository.RecipeCategoryRepositoryInterface
}

func (ur *RecipeCategoryRepository) FindOne(criteria *persistence.Criteria) (*DomainEntity.RecipeCategory, error) {
	entity, errorFindOne := ur.EntityManager.FindOne(ur.Table, criteria)

	if errorFindOne != nil {
		return nil, errorFindOne
	}
	//todo
	entityBsonM, _ := entity.(bson.M)
	result := DomainEntity.RecipeCategory{}
	bsonBytes, _ := bson.Marshal(entityBsonM)
	_ = bson.Unmarshal(bsonBytes, &result)

	return &result, nil
}

func (ur *RecipeCategoryRepository) FindAll(criteria *persistence.Criteria) ([]*DomainEntity.RecipeCategory, error) {
	var recipeCategories []*DomainEntity.RecipeCategory

	entities, errorFindAll := ur.EntityManager.FindAll(ur.Table, criteria)

	if errorFindAll != nil {
		return nil, errorFindAll
	}

	for _, entity := range entities {
		entityBsonM, _ := entity.(bson.M)
		result := DomainEntity.RecipeCategory{}
		bsonBytes, errorBSONBytesMarshal := bson.Marshal(entityBsonM)

		if errorBSONBytesMarshal != nil {
			return nil, errorBSONBytesMarshal
		}

		errorBSONBytesUnMarshal := bson.Unmarshal(bsonBytes, &result)

		if errorBSONBytesUnMarshal != nil {
			return nil, errorBSONBytesUnMarshal
		}

		recipeCategories = append(recipeCategories, &result)
	}

	return recipeCategories, nil
}

func (ur *RecipeCategoryRepository) InsertOne(entity *DomainEntity.RecipeCategory) (*DomainEntity.RecipeCategory, error) {
	_, errorInsertOne := ur.EntityManager.InsertOne(ur.Table, entity)

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *RecipeCategoryRepository) InsertMany(entities []DomainEntity.RecipeCategory) ([]DomainEntity.RecipeCategory, error) {
	entitiesAsInterfaces := make([]interface{}, len(entities))

	for _, value := range entities {
		entitiesAsInterfaces = append(entitiesAsInterfaces, &value)
	}

	_, errorInsertMany := ur.EntityManager.InsertMany(ur.Table, entitiesAsInterfaces)

	if errorInsertMany != nil {
		return nil, errorInsertMany
	}

	return entities, nil
}

func (ur *RecipeCategoryRepository) UpdateOne(criteria *persistence.Criteria, entity *DomainEntity.RecipeCategory) (*DomainEntity.RecipeCategory, error) {
	_, errorInsertOne := ur.EntityManager.UpdateOne(ur.Table, criteria, &persistence.Wrapper{Set: *entity})

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *RecipeCategoryRepository) UpdateMany(criteria *persistence.Criteria, entities []*DomainEntity.RecipeCategory) ([]*DomainEntity.RecipeCategory, error) {
	entitiesAsInterfaces := make([]interface{}, len(entities))

	for _, value := range entities {
		entitiesAsInterfaces = append(entitiesAsInterfaces, &value)
	}

	_, errorInsertMany := ur.EntityManager.UpdateMany(ur.Table, criteria, &persistence.Wrapper{Set: entitiesAsInterfaces})

	if errorInsertMany != nil {
		return nil, errorInsertMany
	}

	return entities, nil
}

func (ur *RecipeCategoryRepository) DeleteOne(criteria *persistence.Criteria) (bool, error) {
	return ur.EntityManager.DeleteOne(ur.Table, criteria)
}

func (ur *RecipeCategoryRepository) GetCriteria() *repository.CriteriaRepository {
	return &CriteriaRepository
}

type RecipeIngredientRepository struct {
	EntityManager persistence.EntityManagerInterface
	Table         string
	repository.RecipeIngredientRepositoryInterface
}

func (ur *RecipeIngredientRepository) FindOne(criteria *persistence.Criteria) (*DomainEntity.RecipeIngredient, error) {
	entity, errorFindOne := ur.EntityManager.FindOne(ur.Table, criteria)

	if errorFindOne != nil {
		return nil, errorFindOne
	}
	//todo
	entityBsonM, _ := entity.(bson.M)
	result := DomainEntity.RecipeIngredient{}
	bsonBytes, _ := bson.Marshal(entityBsonM)
	_ = bson.Unmarshal(bsonBytes, &result)

	return &result, nil
}

func (ur *RecipeIngredientRepository) FindAll(criteria *persistence.Criteria) ([]*DomainEntity.RecipeIngredient, error) {
	var recipeIngredients []*DomainEntity.RecipeIngredient

	entities, errorFindAll := ur.EntityManager.FindAll(ur.Table, criteria)

	if errorFindAll != nil {
		return nil, errorFindAll
	}

	for _, entity := range entities {
		entityBsonM, _ := entity.(bson.M)
		result := DomainEntity.RecipeIngredient{}
		bsonBytes, errorBSONBytesMarshal := bson.Marshal(entityBsonM)

		if errorBSONBytesMarshal != nil {
			return nil, errorBSONBytesMarshal
		}

		errorBSONBytesUnMarshal := bson.Unmarshal(bsonBytes, &result)

		if errorBSONBytesUnMarshal != nil {
			return nil, errorBSONBytesUnMarshal
		}

		recipeIngredients = append(recipeIngredients, &result)
	}

	return recipeIngredients, nil
}

func (ur *RecipeIngredientRepository) InsertOne(entity *DomainEntity.RecipeIngredient) (*DomainEntity.RecipeIngredient, error) {
	_, errorInsertOne := ur.EntityManager.InsertOne(ur.Table, entity)

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *RecipeIngredientRepository) InsertMany(entities []DomainEntity.RecipeIngredient) ([]DomainEntity.RecipeIngredient, error) {
	entitiesAsInterfaces := make([]interface{}, len(entities))

	for _, value := range entities {
		entitiesAsInterfaces = append(entitiesAsInterfaces, &value)
	}

	_, errorInsertMany := ur.EntityManager.InsertMany(ur.Table, entitiesAsInterfaces)

	if errorInsertMany != nil {
		return nil, errorInsertMany
	}

	return entities, nil
}

func (ur *RecipeIngredientRepository) UpdateOne(criteria *persistence.Criteria, entity *DomainEntity.RecipeIngredient) (*DomainEntity.RecipeIngredient, error) {
	_, errorInsertOne := ur.EntityManager.UpdateOne(ur.Table, criteria, &persistence.Wrapper{Set: *entity})

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *RecipeIngredientRepository) UpdateMany(criteria *persistence.Criteria, entities []*DomainEntity.RecipeIngredient) ([]*DomainEntity.RecipeIngredient, error) {
	entitiesAsInterfaces := make([]interface{}, len(entities))

	for _, value := range entities {
		entitiesAsInterfaces = append(entitiesAsInterfaces, &value)
	}

	_, errorInsertMany := ur.EntityManager.UpdateMany(ur.Table, criteria, &persistence.Wrapper{Set: entitiesAsInterfaces})

	if errorInsertMany != nil {
		return nil, errorInsertMany
	}

	return entities, nil
}

func (ur *RecipeIngredientRepository) DeleteOne(criteria *persistence.Criteria) (bool, error) {
	return ur.EntityManager.DeleteOne(ur.Table, criteria)
}

func (ur *RecipeIngredientRepository) GetCriteria() *repository.CriteriaRepository {
	return &CriteriaRepository
}

type RecipeProcessRepository struct {
	EntityManager persistence.EntityManagerInterface
	Table         string
	repository.RecipeProcessRepositoryInterface
}

func (ur *RecipeProcessRepository) FindOne(criteria *persistence.Criteria) (*DomainEntity.RecipeProcess, error) {
	entity, errorFindOne := ur.EntityManager.FindOne(ur.Table, criteria)

	if errorFindOne != nil {
		return nil, errorFindOne
	}
	//todo
	entityBsonM, _ := entity.(bson.M)
	result := DomainEntity.RecipeProcess{}
	bsonBytes, _ := bson.Marshal(entityBsonM)
	_ = bson.Unmarshal(bsonBytes, &result)

	return &result, nil
}

func (ur *RecipeProcessRepository) FindAll(criteria *persistence.Criteria) ([]*DomainEntity.RecipeProcess, error) {
	var recipeProcesses []*DomainEntity.RecipeProcess

	entities, errorFindAll := ur.EntityManager.FindAll(ur.Table, criteria)

	if errorFindAll != nil {
		return nil, errorFindAll
	}

	for _, entity := range entities {
		entityBsonM, _ := entity.(bson.M)
		result := DomainEntity.RecipeProcess{}
		bsonBytes, errorBSONBytesMarshal := bson.Marshal(entityBsonM)

		if errorBSONBytesMarshal != nil {
			return nil, errorBSONBytesMarshal
		}

		errorBSONBytesUnMarshal := bson.Unmarshal(bsonBytes, &result)

		if errorBSONBytesUnMarshal != nil {
			return nil, errorBSONBytesUnMarshal
		}

		recipeProcesses = append(recipeProcesses, &result)
	}

	return recipeProcesses, nil
}

func (ur *RecipeProcessRepository) InsertOne(entity *DomainEntity.RecipeProcess) (*DomainEntity.RecipeProcess, error) {
	_, errorInsertOne := ur.EntityManager.InsertOne(ur.Table, entity)

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *RecipeProcessRepository) InsertMany(entities []DomainEntity.RecipeProcess) ([]DomainEntity.RecipeProcess, error) {
	entitiesAsInterfaces := make([]interface{}, len(entities))

	for _, value := range entities {
		entitiesAsInterfaces = append(entitiesAsInterfaces, &value)
	}

	_, errorInsertMany := ur.EntityManager.InsertMany(ur.Table, entitiesAsInterfaces)

	if errorInsertMany != nil {
		return nil, errorInsertMany
	}

	return entities, nil
}

func (ur *RecipeProcessRepository) UpdateOne(criteria *persistence.Criteria, entity *DomainEntity.RecipeProcess) (*DomainEntity.RecipeProcess, error) {
	_, errorInsertOne := ur.EntityManager.UpdateOne(ur.Table, criteria, &persistence.Wrapper{Set: *entity})

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *RecipeProcessRepository) UpdateMany(criteria *persistence.Criteria, entities []*DomainEntity.RecipeProcess) ([]*DomainEntity.RecipeProcess, error) {
	entitiesAsInterfaces := make([]interface{}, len(entities))

	for _, value := range entities {
		entitiesAsInterfaces = append(entitiesAsInterfaces, &value)
	}

	_, errorInsertMany := ur.EntityManager.UpdateMany(ur.Table, criteria, &persistence.Wrapper{Set: entitiesAsInterfaces})

	if errorInsertMany != nil {
		return nil, errorInsertMany
	}

	return entities, nil
}

func (ur *RecipeProcessRepository) DeleteOne(criteria *persistence.Criteria) (bool, error) {
	return ur.EntityManager.DeleteOne(ur.Table, criteria)
}

func (ur *RecipeProcessRepository) GetCriteria() *repository.CriteriaRepository {
	return &CriteriaRepository
}

type RecipeMeasureRepository struct {
	EntityManager persistence.EntityManagerInterface
	Table         string
	repository.RecipeMeasureRepositoryInterface
}

func (ur *RecipeMeasureRepository) FindOne(criteria *persistence.Criteria) (*DomainEntity.RecipeMeasure, error) {
	entity, errorFindOne := ur.EntityManager.FindOne(ur.Table, criteria)

	if errorFindOne != nil {
		return nil, errorFindOne
	}
	//todo
	entityBsonM, _ := entity.(bson.M)
	result := DomainEntity.RecipeMeasure{}
	bsonBytes, _ := bson.Marshal(entityBsonM)
	_ = bson.Unmarshal(bsonBytes, &result)

	return &result, nil
}

func (ur *RecipeMeasureRepository) FindAll(criteria *persistence.Criteria) ([]*DomainEntity.RecipeMeasure, error) {
	var recipeMeasures []*DomainEntity.RecipeMeasure

	entities, errorFindAll := ur.EntityManager.FindAll(ur.Table, criteria)

	if errorFindAll != nil {
		return nil, errorFindAll
	}

	for _, entity := range entities {
		entityBsonM, _ := entity.(bson.M)
		result := DomainEntity.RecipeMeasure{}
		bsonBytes, errorBSONBytesMarshal := bson.Marshal(entityBsonM)

		if errorBSONBytesMarshal != nil {
			return nil, errorBSONBytesMarshal
		}

		errorBSONBytesUnMarshal := bson.Unmarshal(bsonBytes, &result)

		if errorBSONBytesUnMarshal != nil {
			return nil, errorBSONBytesUnMarshal
		}

		recipeMeasures = append(recipeMeasures, &result)
	}

	return recipeMeasures, nil
}

func (ur *RecipeMeasureRepository) InsertOne(entity *DomainEntity.RecipeMeasure) (*DomainEntity.RecipeMeasure, error) {
	_, errorInsertOne := ur.EntityManager.InsertOne(ur.Table, entity)

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *RecipeMeasureRepository) InsertMany(entities []*DomainEntity.RecipeMeasure) ([]*DomainEntity.RecipeMeasure, error) {
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

func (ur *RecipeMeasureRepository) UpdateOne(criteria *persistence.Criteria, entity *DomainEntity.RecipeMeasure) (*DomainEntity.RecipeMeasure, error) {
	_, errorInsertOne := ur.EntityManager.UpdateOne(ur.Table, criteria, &persistence.Wrapper{Set: *entity})

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *RecipeMeasureRepository) UpdateMany(criteria *persistence.Criteria, entities []*DomainEntity.RecipeMeasure) ([]*DomainEntity.RecipeMeasure, error) {
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

func (ur *RecipeMeasureRepository) DeleteOne(criteria *persistence.Criteria) (bool, error) {
	return ur.EntityManager.DeleteOne(ur.Table, criteria)
}

func (ur *RecipeMeasureRepository) GetCriteria() *repository.CriteriaRepository {
	return &CriteriaRepository
}
