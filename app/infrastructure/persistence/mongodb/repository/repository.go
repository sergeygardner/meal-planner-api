package repository

import (
	DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type AltNameRepository struct {
	EntityManager persistence.EntityManagerInterface
	Table         string
	repository.AltNameRepositoryInterface
}

func (ur *AltNameRepository) FindOne(criteria *persistence.Criteria) (*DomainEntity.AltName, error) {
	entity, errorFindOne := ur.EntityManager.FindOne(ur.Table, criteria)

	if errorFindOne != nil {
		return nil, errorFindOne
	}
	//todo
	entityBsonM, _ := entity.(bson.M)
	result := DomainEntity.AltName{}
	bsonBytes, _ := bson.Marshal(entityBsonM)
	_ = bson.Unmarshal(bsonBytes, &result)

	return &result, nil
}

func (ur *AltNameRepository) FindAll(criteria *persistence.Criteria) ([]*DomainEntity.AltName, error) {
	var recipeAltNames []*DomainEntity.AltName

	entities, errorFindAll := ur.EntityManager.FindAll(ur.Table, criteria)

	if errorFindAll != nil {
		return nil, errorFindAll
	}

	for _, entity := range entities {
		entityBsonM, _ := entity.(bson.M)
		result := DomainEntity.AltName{}
		bsonBytes, errorBSONBytesMarshal := bson.Marshal(entityBsonM)

		if errorBSONBytesMarshal != nil {
			return nil, errorBSONBytesMarshal
		}

		errorBSONBytesUnMarshal := bson.Unmarshal(bsonBytes, &result)

		if errorBSONBytesUnMarshal != nil {
			return nil, errorBSONBytesUnMarshal
		}

		recipeAltNames = append(recipeAltNames, &result)
	}

	return recipeAltNames, nil
}

func (ur *AltNameRepository) InsertOne(entity *DomainEntity.AltName) (*DomainEntity.AltName, error) {
	_, errorInsertOne := ur.EntityManager.InsertOne(ur.Table, entity)

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *AltNameRepository) InsertMany(entities []*DomainEntity.AltName) ([]*DomainEntity.AltName, error) {
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

func (ur *AltNameRepository) UpdateOne(criteria *persistence.Criteria, entity *DomainEntity.AltName) (*DomainEntity.AltName, error) {
	_, errorInsertOne := ur.EntityManager.UpdateOne(ur.Table, criteria, &persistence.Wrapper{Set: *entity})

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *AltNameRepository) UpdateMany(criteria *persistence.Criteria, entities []*DomainEntity.AltName) ([]*DomainEntity.AltName, error) {
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

func (ur *AltNameRepository) DeleteOne(criteria *persistence.Criteria) (bool, error) {
	return ur.EntityManager.DeleteOne(ur.Table, criteria)
}

func (ur *AltNameRepository) GetCriteria() *repository.CriteriaRepository {
	return &CriteriaRepository
}

type PictureRepository struct {
	EntityManager persistence.EntityManagerInterface
	Table         string
	repository.PictureRepositoryInterface
}

func (ur *PictureRepository) FindOne(criteria *persistence.Criteria) (*DomainEntity.Picture, error) {
	entity, errorFindOne := ur.EntityManager.FindOne(ur.Table, criteria)

	if errorFindOne != nil {
		return nil, errorFindOne
	}
	//todo
	entityBsonM, _ := entity.(bson.M)
	result := DomainEntity.Picture{}
	bsonBytes, _ := bson.Marshal(entityBsonM)
	_ = bson.Unmarshal(bsonBytes, &result)

	return &result, nil
}

func (ur *PictureRepository) FindAll(criteria *persistence.Criteria) ([]*DomainEntity.Picture, error) {
	var recipePictures []*DomainEntity.Picture

	entities, errorFindAll := ur.EntityManager.FindAll(ur.Table, criteria)

	if errorFindAll != nil {
		return nil, errorFindAll
	}

	for _, entity := range entities {
		entityBsonM, _ := entity.(bson.M)
		result := DomainEntity.Picture{}
		bsonBytes, errorBSONBytesMarshal := bson.Marshal(entityBsonM)

		if errorBSONBytesMarshal != nil {
			return nil, errorBSONBytesMarshal
		}

		errorBSONBytesUnMarshal := bson.Unmarshal(bsonBytes, &result)

		if errorBSONBytesUnMarshal != nil {
			return nil, errorBSONBytesUnMarshal
		}

		recipePictures = append(recipePictures, &result)
	}

	return recipePictures, nil
}

func (ur *PictureRepository) InsertOne(entity *DomainEntity.Picture) (*DomainEntity.Picture, error) {
	_, errorInsertOne := ur.EntityManager.InsertOne(ur.Table, entity)

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *PictureRepository) InsertMany(entities []*DomainEntity.Picture) ([]*DomainEntity.Picture, error) {
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

func (ur *PictureRepository) UpdateOne(criteria *persistence.Criteria, entity *DomainEntity.Picture) (*DomainEntity.Picture, error) {
	_, errorInsertOne := ur.EntityManager.UpdateOne(ur.Table, criteria, &persistence.Wrapper{Set: *entity})

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *PictureRepository) UpdateMany(criteria *persistence.Criteria, entities []*DomainEntity.Picture) ([]*DomainEntity.Picture, error) {
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

func (ur *PictureRepository) DeleteOne(criteria *persistence.Criteria) (bool, error) {
	return ur.EntityManager.DeleteOne(ur.Table, criteria)
}

func (ur *PictureRepository) GetCriteria() *repository.CriteriaRepository {
	return &CriteriaRepository
}

type UnitRepository struct {
	EntityManager persistence.EntityManagerInterface
	Table         string
	repository.UnitRepositoryInterface
}

func (ur *UnitRepository) FindOne(criteria *persistence.Criteria) (*DomainEntity.Unit, error) {
	entity, errorFindOne := ur.EntityManager.FindOne(ur.Table, criteria)

	if errorFindOne != nil {
		return nil, errorFindOne
	}
	//todo
	entityBsonM, _ := entity.(bson.M)
	result := DomainEntity.Unit{}
	bsonBytes, _ := bson.Marshal(entityBsonM)
	_ = bson.Unmarshal(bsonBytes, &result)

	return &result, nil
}

func (ur *UnitRepository) FindAll(criteria *persistence.Criteria) ([]*DomainEntity.Unit, error) {
	var recipeMeasures []*DomainEntity.Unit

	entities, errorFindAll := ur.EntityManager.FindAll(ur.Table, criteria)

	if errorFindAll != nil {
		return nil, errorFindAll
	}

	for _, entity := range entities {
		entityBsonM, _ := entity.(bson.M)
		result := DomainEntity.Unit{}
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

func (ur *UnitRepository) InsertOne(entity *DomainEntity.Unit) (*DomainEntity.Unit, error) {
	_, errorInsertOne := ur.EntityManager.InsertOne(ur.Table, entity)

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *UnitRepository) InsertMany(entities []*DomainEntity.Unit) ([]*DomainEntity.Unit, error) {
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

func (ur *UnitRepository) UpdateOne(criteria *persistence.Criteria, entity *DomainEntity.Unit) (*DomainEntity.Unit, error) {
	_, errorInsertOne := ur.EntityManager.UpdateOne(ur.Table, criteria, &persistence.Wrapper{Set: *entity})

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *UnitRepository) UpdateMany(criteria *persistence.Criteria, entities []*DomainEntity.Unit) ([]*DomainEntity.Unit, error) {
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

func (ur *UnitRepository) DeleteOne(criteria *persistence.Criteria) (bool, error) {
	return ur.EntityManager.DeleteOne(ur.Table, criteria)
}

func (ur *UnitRepository) GetCriteria() *repository.CriteriaRepository {
	return &CriteriaRepository
}

type CategoryRepository struct {
	EntityManager persistence.EntityManagerInterface
	Table         string
	repository.CategoryRepositoryInterface
}

func (ur *CategoryRepository) FindOne(criteria *persistence.Criteria) (*DomainEntity.Category, error) {
	entity, errorFindOne := ur.EntityManager.FindOne(ur.Table, criteria)

	if errorFindOne != nil {
		return nil, errorFindOne
	}
	//todo
	entityBsonM, _ := entity.(bson.M)
	result := DomainEntity.Category{}
	bsonBytes, _ := bson.Marshal(entityBsonM)
	_ = bson.Unmarshal(bsonBytes, &result)

	return &result, nil
}

func (ur *CategoryRepository) FindAll(criteria *persistence.Criteria) ([]*DomainEntity.Category, error) {
	var recipeMeasures []*DomainEntity.Category

	entities, errorFindAll := ur.EntityManager.FindAll(ur.Table, criteria)

	if errorFindAll != nil {
		return nil, errorFindAll
	}

	for _, entity := range entities {
		entityBsonM, _ := entity.(bson.M)
		result := DomainEntity.Category{}
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

func (ur *CategoryRepository) InsertOne(entity *DomainEntity.Category) (*DomainEntity.Category, error) {
	_, errorInsertOne := ur.EntityManager.InsertOne(ur.Table, entity)

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *CategoryRepository) InsertMany(entities []*DomainEntity.Category) ([]*DomainEntity.Category, error) {
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

func (ur *CategoryRepository) UpdateOne(criteria *persistence.Criteria, entity *DomainEntity.Category) (*DomainEntity.Category, error) {
	_, errorInsertOne := ur.EntityManager.UpdateOne(ur.Table, criteria, &persistence.Wrapper{Set: *entity})

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *CategoryRepository) UpdateMany(criteria *persistence.Criteria, entities []*DomainEntity.Category) ([]*DomainEntity.Category, error) {
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

func (ur *CategoryRepository) DeleteOne(criteria *persistence.Criteria) (bool, error) {
	return ur.EntityManager.DeleteOne(ur.Table, criteria)
}

func (ur *CategoryRepository) GetCriteria() *repository.CriteriaRepository {
	return &CriteriaRepository
}

type IngredientRepository struct {
	EntityManager persistence.EntityManagerInterface
	Table         string
	repository.IngredientRepositoryInterface
}

func (ur *IngredientRepository) FindOne(criteria *persistence.Criteria) (*DomainEntity.Ingredient, error) {
	entity, errorFindOne := ur.EntityManager.FindOne(ur.Table, criteria)

	if errorFindOne != nil {
		return nil, errorFindOne
	}
	//todo
	entityBsonM, _ := entity.(bson.M)
	result := DomainEntity.Ingredient{}
	bsonBytes, _ := bson.Marshal(entityBsonM)
	_ = bson.Unmarshal(bsonBytes, &result)

	return &result, nil
}

func (ur *IngredientRepository) FindAll(criteria *persistence.Criteria) ([]*DomainEntity.Ingredient, error) {
	var recipeMeasures []*DomainEntity.Ingredient

	entities, errorFindAll := ur.EntityManager.FindAll(ur.Table, criteria)

	if errorFindAll != nil {
		return nil, errorFindAll
	}

	for _, entity := range entities {
		entityBsonM, _ := entity.(bson.M)
		result := DomainEntity.Ingredient{}
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

func (ur *IngredientRepository) InsertOne(entity *DomainEntity.Ingredient) (*DomainEntity.Ingredient, error) {
	_, errorInsertOne := ur.EntityManager.InsertOne(ur.Table, entity)

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *IngredientRepository) InsertMany(entities []*DomainEntity.Ingredient) ([]*DomainEntity.Ingredient, error) {
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

func (ur *IngredientRepository) UpdateOne(criteria *persistence.Criteria, entity *DomainEntity.Ingredient) (*DomainEntity.Ingredient, error) {
	_, errorInsertOne := ur.EntityManager.UpdateOne(ur.Table, criteria, &persistence.Wrapper{Set: *entity})

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *IngredientRepository) UpdateMany(criteria *persistence.Criteria, entities []*DomainEntity.Ingredient) ([]*DomainEntity.Ingredient, error) {
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

func (ur *IngredientRepository) DeleteOne(criteria *persistence.Criteria) (bool, error) {
	return ur.EntityManager.DeleteOne(ur.Table, criteria)
}

func (ur *IngredientRepository) GetCriteria() *repository.CriteriaRepository {
	return &CriteriaRepository
}
