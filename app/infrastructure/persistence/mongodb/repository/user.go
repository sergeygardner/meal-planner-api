package repository

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"
	DomainKind "github.com/sergeygardner/meal-planner-api/domain/kind"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence/repository"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	errorUserFindOneConvertToBSON = errors.New("an error occurred while converting data to specific type")
)

type UserRepository struct {
	EntityManager persistence.EntityManagerInterface
	Table         string
	repository.UserRepositoryInterface
}

type UserRoleRepository struct {
	EntityManager persistence.EntityManagerInterface
	Table         string
	repository.UserRoleRepositoryInterface
}

type UserToRoleRepository struct {
	EntityManager persistence.EntityManagerInterface
	Table         string
	repository.UserToRoleRepositoryInterface
}

type UserConfirmationRepository struct {
	EntityManager persistence.EntityManagerInterface
	Table         string
	repository.UserConfirmationRepositoryInterface
}

func (ur *UserRepository) FindOne(criteria *persistence.Criteria) (*DomainEntity.User, error) {
	entity, errorFindOne := ur.EntityManager.FindOne(ur.Table, criteria)

	if errorFindOne != nil {
		return nil, errorFindOne
	}

	entityBsonM, statusEntityBsonM := entity.(bson.M)

	if !statusEntityBsonM {
		return nil, errorUserFindOneConvertToBSON
	}

	result := DomainEntity.User{}
	bsonBytes, errorEntityBsonMMarshaled := bson.Marshal(entityBsonM)

	if errorEntityBsonMMarshaled != nil {
		return nil, errorEntityBsonMMarshaled
	}

	errorEntityBsonMUnMarshaled := bson.Unmarshal(bsonBytes, &result)

	if errorEntityBsonMUnMarshaled != nil {
		return nil, errorEntityBsonMUnMarshaled
	}

	return &result, nil
}

func (ur *UserRepository) FindAll(criteria *persistence.Criteria) ([]DomainEntity.User, error) {
	var users []DomainEntity.User

	entities, errorFindAll := ur.EntityManager.FindAll(ur.Table, criteria)

	if errorFindAll != nil {
		return nil, errorFindAll
	}

	for _, v := range entities {
		field, ok := v.(DomainEntity.User)
		if ok {
			users = append(users, field)
		}
	}

	return users, nil
}

func (ur *UserRepository) InsertOne(entity *DomainEntity.User) (*DomainEntity.User, error) {
	_, errorInsertOne := ur.EntityManager.InsertOne(ur.Table, entity)

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *UserRepository) InsertMany(entities []DomainEntity.User) ([]DomainEntity.User, error) {
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

func (ur *UserRepository) UpdateOne(criteria *persistence.Criteria, entity *DomainEntity.User) (*DomainEntity.User, error) {
	_, errorInsertOne := ur.EntityManager.UpdateOne(ur.Table, criteria, &persistence.Wrapper{Set: *entity})

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *UserRepository) UpdateMany(criteria *persistence.Criteria, entities []*DomainEntity.User) ([]*DomainEntity.User, error) {
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

func (ur *UserRepository) DeleteOne(criteria *persistence.Criteria) (bool, error) {
	return ur.EntityManager.DeleteOne(ur.Table, criteria)
}

func (ur *UserRepository) GetCriteriaByUserId(id *uuid.UUID) *persistence.Criteria {
	return &persistence.Criteria{
		Where: map[string]interface{}{
			"id": id,
		},
	}
}

func (ur *UserRepository) GetCriteriaByUsername(username string) *persistence.Criteria {
	return &persistence.Criteria{
		Where: map[string]interface{}{
			"userdto.userregisterdto.usercredentialsdto.username": username,
		},
	}
}

func (ur *UserRoleRepository) FindOne(criteria *persistence.Criteria) (*DomainEntity.UserRole, error) {
	entity, errorFindOne := ur.EntityManager.FindOne(ur.Table, criteria)

	if errorFindOne != nil {
		return nil, errorFindOne
	}

	result, ok := entity.(DomainEntity.UserRole)

	if !ok {
		//errorConvert := fmt.Errorf("error is occurred while converting interface '%s' to struct 'entity.UserRole'", entity)
		//
		//return nil, errorConvert
	}

	return &result, nil
}

func (ur *UserRoleRepository) FindAll(criteria *persistence.Criteria) ([]DomainEntity.UserRole, error) {
	var users []DomainEntity.UserRole

	entities, errorFindAll := ur.EntityManager.FindAll(ur.Table, criteria)

	if errorFindAll != nil {
		return nil, errorFindAll
	}

	for _, v := range entities {
		field, ok := v.(DomainEntity.UserRole)

		if ok {
			users = append(users, field)
		}
	}

	return users, nil
}

func (ur *UserToRoleRepository) FindOne(criteria *persistence.Criteria) (*DomainEntity.UserToRole, error) {
	entity, errorFindOne := ur.EntityManager.FindOne(ur.Table, criteria)

	if errorFindOne != nil {
		return nil, errorFindOne
	}

	result, ok := entity.(DomainEntity.UserToRole)

	if !ok {
		//errorConvert := fmt.Errorf("error is occurred while converting interface '%s' to struct 'entity.UserTpGroup'", entity)
		//
		//return nil, errorConvert
	}

	return &result, nil
}

func (ur *UserToRoleRepository) FindAll(criteria *persistence.Criteria) ([]DomainEntity.UserToRole, error) {
	var users []DomainEntity.UserToRole

	entities, errorFindAll := ur.EntityManager.FindAll(ur.Table, criteria)

	if errorFindAll != nil {
		return nil, errorFindAll
	}

	for _, v := range entities {
		field, ok := v.(DomainEntity.UserToRole)
		if ok {
			users = append(users, field)
		}
	}

	return users, nil
}

func (ur *UserConfirmationRepository) FindOne(criteria *persistence.Criteria) (*DomainEntity.UserConfirmation, error) {
	entity, errorFindOne := ur.EntityManager.FindOne(ur.Table, criteria)

	if errorFindOne != nil {
		return nil, errorFindOne
	}
	//todo
	entityBsonM, _ := entity.(bson.M)
	result := DomainEntity.UserConfirmation{}
	bsonBytes, _ := bson.Marshal(entityBsonM)
	_ = bson.Unmarshal(bsonBytes, &result)

	return &result, nil
}

func (ur *UserConfirmationRepository) FindAll(criteria *persistence.Criteria) ([]DomainEntity.UserConfirmation, error) {
	var userConfirmations []DomainEntity.UserConfirmation

	entities, errorFindAll := ur.EntityManager.FindAll(ur.Table, criteria)

	if errorFindAll != nil {
		return nil, errorFindAll
	}

	for _, entity := range entities {
		//todo
		entityBsonM, _ := entity.(bson.M)
		result := DomainEntity.UserConfirmation{}
		bsonBytes, _ := bson.Marshal(entityBsonM)
		_ = bson.Unmarshal(bsonBytes, &result)

		userConfirmations = append(userConfirmations, result)
	}

	return userConfirmations, nil
}

func (ur *UserConfirmationRepository) InsertOne(entity *DomainEntity.UserConfirmation) (*DomainEntity.UserConfirmation, error) {
	_, errorInsertOne := ur.EntityManager.InsertOne(ur.Table, entity)

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *UserConfirmationRepository) InsertMany(entities []DomainEntity.UserConfirmation) ([]DomainEntity.UserConfirmation, error) {
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

func (ur *UserConfirmationRepository) UpdateOne(criteria *persistence.Criteria, entity *DomainEntity.UserConfirmation) (*DomainEntity.UserConfirmation, error) {
	_, errorInsertOne := ur.EntityManager.UpdateOne(ur.Table, criteria, &persistence.Wrapper{Set: entity})

	if errorInsertOne != nil {
		return nil, errorInsertOne
	}

	return entity, nil
}

func (ur *UserConfirmationRepository) UpdateMany(criteria *persistence.Criteria, entities []*DomainEntity.UserConfirmation) ([]*DomainEntity.UserConfirmation, error) {
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

func (ur *UserConfirmationRepository) GetCriteriaByUserIdAndActive(user *DomainEntity.User) *persistence.Criteria {
	return &persistence.Criteria{
		Where: map[string]interface{}{
			"user_id": user.Id,
			"active":  DomainKind.UserConfirmationActive,
		},
	}
}

func (ur *UserConfirmationRepository) GetCriteriaById(id *uuid.UUID) *persistence.Criteria {
	return &persistence.Criteria{
		Where: map[string]interface{}{
			"id": id,
		},
	}
}
