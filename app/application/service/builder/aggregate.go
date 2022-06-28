package builder

import (
	"context"
	"github.com/google/uuid"
	DomainAggregate "github.com/sergeygardner/meal-planner-api/domain/aggregate"
	DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence/repository"
	InfrastructureService "github.com/sergeygardner/meal-planner-api/infrastructure/service/repository"
	"sync"
)

var (
	factoryRepository              = InfrastructureService.GetFactoryRepository()
	errorBuildingRecipe            error
	errorBuildingRecipeCategories  error
	errorBuildingRecipeIngredients error
	errorBuildingRecipeProcesses   error
	errorBuildingPictures          error
	errorBuildingRecipeMeasures    error
	errorBuildingAltNames          error
	errorBuildingUnits             error
	errorBuildingCategories        error
	errorBuildingIngredients       error
	errorBuildingPlanners          error
	errorBuildingPlannerIntervals  error
	errorBuildingPlannerRecipes    error
)

type recipeComposite struct {
	Id       *uuid.UUID
	UserId   *uuid.UUID
	Entities *[]*DomainAggregate.Recipe
	Criteria *persistence.Criteria
}

type recipeCategoryComposite struct {
	Id       *uuid.UUID
	UserId   *uuid.UUID
	EntityId *uuid.UUID
	Entities *[]*DomainAggregate.RecipeCategory
	Criteria *persistence.Criteria
}

type recipeIngredientComposite struct {
	Id       *uuid.UUID
	UserId   *uuid.UUID
	EntityId *uuid.UUID
	Entities *[]*DomainAggregate.RecipeIngredient
	Criteria *persistence.Criteria
}

type recipeProcessComposite struct {
	Id       *uuid.UUID
	UserId   *uuid.UUID
	EntityId *uuid.UUID
	Entities *[]*DomainAggregate.RecipeProcess
	Criteria *persistence.Criteria
}

type pictureComposite struct {
	Id       *uuid.UUID
	UserId   *uuid.UUID
	EntityId *uuid.UUID
	Entities *[]*DomainAggregate.Picture
	Criteria *persistence.Criteria
}

type recipeMeasureComposite struct {
	Id       *uuid.UUID
	UserId   *uuid.UUID
	EntityId *uuid.UUID
	Entities *[]*DomainAggregate.RecipeMeasure
	Criteria *persistence.Criteria
}

type altNameComposite struct {
	Id       *uuid.UUID
	UserId   *uuid.UUID
	EntityId *uuid.UUID
	Entities *[]*DomainEntity.AltName
	Criteria *persistence.Criteria
}

type unitComposite struct {
	Id       *uuid.UUID
	Entity   **DomainEntity.Unit
	Entities *[]*DomainEntity.Unit
	Criteria *persistence.Criteria
}

type categoryComposite struct {
	Id       *uuid.UUID
	UserId   *uuid.UUID
	Entity   **DomainAggregate.Category
	Entities *[]*DomainAggregate.Category
	Criteria *persistence.Criteria
}

type ingredientComposite struct {
	Id       *uuid.UUID
	UserId   *uuid.UUID
	Entity   **DomainEntity.Ingredient
	Entities *[]*DomainEntity.Ingredient
	Criteria *persistence.Criteria
}

type plannerComposite struct {
	Id       *uuid.UUID
	UserId   *uuid.UUID
	Entities *[]*DomainAggregate.Planner
	Criteria *persistence.Criteria
}

type plannerIntervalComposite struct {
	Id       *uuid.UUID
	UserId   *uuid.UUID
	EntityId *uuid.UUID
	Entity   **DomainAggregate.PlannerInterval
	Entities *[]*DomainAggregate.PlannerInterval
	Criteria *persistence.Criteria
}
type plannerRecipeComposite struct {
	Id       *uuid.UUID
	UserId   *uuid.UUID
	EntityId *uuid.UUID
	Entities *[]*DomainAggregate.PlannerRecipe
	Criteria *persistence.Criteria
}

func BuildRecipesAggregate(
	id *uuid.UUID,
	userId *uuid.UUID,
	criteria *persistence.Criteria,
) ([]*DomainAggregate.Recipe, error) {
	var recipeAggregates []*DomainAggregate.Recipe

	channelRecipe := make(chan *recipeComposite)
	channelRecipeCategory := make(chan *recipeCategoryComposite)
	channelRecipeIngredient := make(chan *recipeIngredientComposite)
	channelRecipeProcess := make(chan *recipeProcessComposite)
	channelRecipeMeasure := make(chan *recipeMeasureComposite)
	channelCategory := make(chan *categoryComposite)
	channelIngredient := make(chan *ingredientComposite)
	channelPicture := make(chan *pictureComposite)
	channelUnit := make(chan *unitComposite)
	channelAltName := make(chan *altNameComposite)
	waitGroup := &sync.WaitGroup{}
	aggregationContext, aggregationCancel := context.WithCancel(context.TODO())

	waitGroup.Add(1)

	defer aggregationCancel()

	go buildRecipeAggregate(aggregationContext, waitGroup, channelRecipe, channelRecipeCategory, channelRecipeIngredient, channelRecipeProcess, channelPicture, channelAltName)
	go buildRecipeCategoriesAggregate(aggregationContext, waitGroup, channelRecipeCategory, channelCategory)
	go buildRecipeIngredientsAggregate(aggregationContext, waitGroup, channelRecipeIngredient, channelRecipeMeasure, channelIngredient, channelPicture, channelAltName)
	go buildRecipeProcessesAggregate(aggregationContext, waitGroup, channelRecipeProcess, channelPicture, channelAltName)
	go buildRecipeMeasuresAggregate(aggregationContext, waitGroup, channelRecipeMeasure, channelUnit, channelAltName)
	go buildCategoryAggregate(aggregationContext, waitGroup, channelCategory, channelPicture, channelAltName)
	go buildIngredientEntities(aggregationContext, waitGroup, channelIngredient)
	go buildPicturesAggregate(aggregationContext, waitGroup, channelPicture, channelAltName)
	go buildUnitEntities(aggregationContext, waitGroup, channelUnit)
	go buildAltNames(aggregationContext, waitGroup, channelAltName)

	channelRecipe <- &recipeComposite{Entities: &recipeAggregates, Id: id, UserId: userId, Criteria: criteria}

	waitGroup.Wait()

	close(channelRecipe)
	close(channelRecipeCategory)
	close(channelRecipeIngredient)
	close(channelRecipeProcess)
	close(channelCategory)
	close(channelIngredient)
	close(channelPicture)
	close(channelRecipeMeasure)
	close(channelUnit)
	close(channelAltName)

	return recipeAggregates, errorBuildingRecipe
}

func buildRecipeAggregate(
	aggregationContext context.Context,
	parentWaitGroup *sync.WaitGroup,
	channelRecipe chan *recipeComposite,
	channelRecipeCategory chan *recipeCategoryComposite,
	channelRecipeIngredient chan *recipeIngredientComposite,
	channelRecipeProcess chan *recipeProcessComposite,
	channelPicture chan *pictureComposite,
	channelAltName chan *altNameComposite,
) {
	var (
		recipeEntities      []*DomainEntity.Recipe
		errorRecipeEntities error
	)
	recipeRepository := factoryRepository.GetRecipeRepository()
	recipeRepositoryCriteria := recipeRepository.GetCriteria()

	for {
		select {
		case <-aggregationContext.Done():
			return
		case recipeCompositeItem := <-channelRecipe:
			if recipeCompositeItem == nil {
				continue
			}
			recipeEntities, errorRecipeEntities = recipeRepository.FindAll(
				composeCriteria(
					recipeCompositeItem.Id,
					recipeCompositeItem.UserId,
					nil,
					recipeCompositeItem.Criteria,
					recipeRepositoryCriteria,
				),
			)

			if errorRecipeEntities != nil || len(recipeEntities) == 0 {
				errorBuildingRecipe = errorRecipeEntities
			} else {
				for _, recipeEntity := range recipeEntities {
					recipeAggregate := &DomainAggregate.Recipe{Entity: recipeEntity}
					*recipeCompositeItem.Entities = append(*recipeCompositeItem.Entities, recipeAggregate)

					parentWaitGroup.Add(5)

					channelRecipeCategory <- &recipeCategoryComposite{Entities: &recipeAggregate.Categories, UserId: &recipeEntity.UserId, EntityId: &recipeEntity.Id}
					channelRecipeIngredient <- &recipeIngredientComposite{Entities: &recipeAggregate.Ingredients, UserId: &recipeEntity.UserId, EntityId: &recipeEntity.Id}
					channelRecipeProcess <- &recipeProcessComposite{Entities: &recipeAggregate.Processes, UserId: &recipeEntity.UserId, EntityId: &recipeEntity.Id}
					channelPicture <- &pictureComposite{Entities: &recipeAggregate.Pictures, UserId: &recipeEntity.UserId, EntityId: &recipeEntity.Id}
					channelAltName <- &altNameComposite{Entities: &recipeAggregate.AltNames, UserId: &recipeEntity.UserId, EntityId: &recipeEntity.Id}
				}
			}

			parentWaitGroup.Done()

		}
	}
}

func BuildRecipeCategoriesAggregate(
	id *uuid.UUID,
	userId *uuid.UUID,
	entityId *uuid.UUID,
	criteria *persistence.Criteria,
) ([]*DomainAggregate.RecipeCategory, error) {
	var recipeCategoriesAggregate []*DomainAggregate.RecipeCategory
	channelRecipeCategory := make(chan *recipeCategoryComposite)
	channelCategory := make(chan *categoryComposite)
	channelPicture := make(chan *pictureComposite)
	channelAltName := make(chan *altNameComposite)
	waitGroup := &sync.WaitGroup{}
	aggregationContext := context.TODO()

	waitGroup.Add(1)

	go buildRecipeCategoriesAggregate(aggregationContext, waitGroup, channelRecipeCategory, channelCategory)
	go buildCategoryAggregate(aggregationContext, waitGroup, channelCategory, channelPicture, channelAltName)
	go buildPicturesAggregate(aggregationContext, waitGroup, channelPicture, channelAltName)
	go buildAltNames(aggregationContext, waitGroup, channelAltName)

	channelRecipeCategory <- &recipeCategoryComposite{Entities: &recipeCategoriesAggregate, Id: id, UserId: userId, EntityId: entityId, Criteria: criteria}

	waitGroup.Wait()

	close(channelRecipeCategory)
	close(channelCategory)
	close(channelPicture)
	close(channelAltName)

	return recipeCategoriesAggregate, errorBuildingRecipeCategories
}

func buildRecipeCategoriesAggregate(
	aggregationContext context.Context,
	parentWaitGroup *sync.WaitGroup,
	channelRecipeCategory chan *recipeCategoryComposite,
	channelCategory chan *categoryComposite,
) {
	var (
		recipeCategoryEntities      []*DomainEntity.RecipeCategory
		errorRecipeCategoryEntities error
	)
	recipeCategoryRepository := factoryRepository.GetRecipeCategoryRepository()
	recipeCategoryRepositoryCriteria := recipeCategoryRepository.GetCriteria()

	for {
		select {
		case <-aggregationContext.Done():
			return
		case recipeCategoryCompositeItem := <-channelRecipeCategory:
			if recipeCategoryCompositeItem == nil {
				continue
			}
			recipeCategoryEntities, errorRecipeCategoryEntities = recipeCategoryRepository.FindAll(
				composeCriteria(
					recipeCategoryCompositeItem.Id,
					recipeCategoryCompositeItem.UserId,
					recipeCategoryCompositeItem.EntityId,
					recipeCategoryCompositeItem.Criteria,
					recipeCategoryRepositoryCriteria,
				),
			)

			if errorRecipeCategoryEntities != nil || len(recipeCategoryEntities) == 0 {
				errorBuildingRecipe = errorRecipeCategoryEntities
			} else {
				for _, recipeCategoryEntity := range recipeCategoryEntities {
					recipeCategoryAggregate := &DomainAggregate.RecipeCategory{Entity: recipeCategoryEntity}
					*recipeCategoryCompositeItem.Entities = append(*recipeCategoryCompositeItem.Entities, recipeCategoryAggregate)

					parentWaitGroup.Add(1)

					channelCategory <- &categoryComposite{Entity: &recipeCategoryAggregate.Derive, Id: &recipeCategoryEntity.DeriveId, UserId: &recipeCategoryEntity.UserId}
				}
			}

			parentWaitGroup.Done()

		}
	}
}

func BuildRecipeIngredientsAggregate(
	id *uuid.UUID,
	userId *uuid.UUID,
	entityId *uuid.UUID,
	criteria *persistence.Criteria,
) ([]*DomainAggregate.RecipeIngredient, error) {
	var recipeIngredientsAggregate []*DomainAggregate.RecipeIngredient
	channelRecipeIngredient := make(chan *recipeIngredientComposite)
	channelRecipeMeasure := make(chan *recipeMeasureComposite)
	channelIngredient := make(chan *ingredientComposite)
	channelPicture := make(chan *pictureComposite)
	channelUnit := make(chan *unitComposite)
	channelAltName := make(chan *altNameComposite)
	waitGroup := &sync.WaitGroup{}
	aggregationContext := context.TODO()

	waitGroup.Add(1)

	go buildRecipeIngredientsAggregate(aggregationContext, waitGroup, channelRecipeIngredient, channelRecipeMeasure, channelIngredient, channelPicture, channelAltName)
	go buildRecipeMeasuresAggregate(aggregationContext, waitGroup, channelRecipeMeasure, channelUnit, channelAltName)
	go buildIngredientEntities(aggregationContext, waitGroup, channelIngredient)
	go buildPicturesAggregate(aggregationContext, waitGroup, channelPicture, channelAltName)
	go buildUnitEntities(aggregationContext, waitGroup, channelUnit)
	go buildAltNames(aggregationContext, waitGroup, channelAltName)

	channelRecipeIngredient <- &recipeIngredientComposite{Entities: &recipeIngredientsAggregate, Id: id, UserId: userId, EntityId: entityId, Criteria: criteria}

	waitGroup.Wait()

	close(channelRecipeIngredient)
	close(channelRecipeMeasure)
	close(channelIngredient)
	close(channelPicture)
	close(channelUnit)
	close(channelAltName)

	return recipeIngredientsAggregate, errorBuildingRecipeIngredients
}

func buildRecipeIngredientsAggregate(
	aggregationContext context.Context,
	parentWaitGroup *sync.WaitGroup,
	channelRecipeIngredient chan *recipeIngredientComposite,
	channelRecipeMeasure chan *recipeMeasureComposite,
	channelIngredient chan *ingredientComposite,
	channelPicture chan *pictureComposite,
	channelAltName chan *altNameComposite,
) {
	var (
		recipeIngredientEntities      []*DomainEntity.RecipeIngredient
		errorRecipeIngredientEntities error
	)
	recipeIngredientRepository := factoryRepository.GetRecipeIngredientRepository()
	recipeIngredientRepositoryCriteria := recipeIngredientRepository.GetCriteria()

	for {
		select {
		case <-aggregationContext.Done():
			return
		case recipeIngredientCompositeItem := <-channelRecipeIngredient:
			if recipeIngredientCompositeItem == nil {
				continue
			}
			recipeIngredientEntities, errorRecipeIngredientEntities = recipeIngredientRepository.FindAll(
				composeCriteria(
					recipeIngredientCompositeItem.Id,
					recipeIngredientCompositeItem.UserId,
					recipeIngredientCompositeItem.EntityId,
					recipeIngredientCompositeItem.Criteria,
					recipeIngredientRepositoryCriteria,
				),
			)

			if errorRecipeIngredientEntities != nil || len(recipeIngredientEntities) == 0 {
				errorBuildingRecipe = errorRecipeIngredientEntities
			} else {
				for _, recipeIngredientEntity := range recipeIngredientEntities {
					recipeIngredientAggregate := &DomainAggregate.RecipeIngredient{Entity: recipeIngredientEntity}
					*recipeIngredientCompositeItem.Entities = append(*recipeIngredientCompositeItem.Entities, recipeIngredientAggregate)

					parentWaitGroup.Add(4)

					channelRecipeMeasure <- &recipeMeasureComposite{Entities: &recipeIngredientAggregate.Measures, UserId: &recipeIngredientEntity.UserId, EntityId: &recipeIngredientEntity.Id}
					channelIngredient <- &ingredientComposite{Entity: &recipeIngredientAggregate.Derive, Id: &recipeIngredientEntity.DeriveId, UserId: &recipeIngredientEntity.UserId}
					channelPicture <- &pictureComposite{Entities: &recipeIngredientAggregate.Pictures, UserId: &recipeIngredientEntity.UserId, EntityId: &recipeIngredientEntity.Id}
					channelAltName <- &altNameComposite{Entities: &recipeIngredientAggregate.AltNames, UserId: &recipeIngredientEntity.UserId, EntityId: &recipeIngredientEntity.Id}
				}
			}

			parentWaitGroup.Done()

		}
	}
}

func BuildRecipeProcessesAggregate(
	id *uuid.UUID,
	userId *uuid.UUID,
	entityId *uuid.UUID,
	criteria *persistence.Criteria,
) ([]*DomainAggregate.RecipeProcess, error) {
	var recipeProcessesAggregate []*DomainAggregate.RecipeProcess
	channelRecipeProcess := make(chan *recipeProcessComposite)
	channelPicture := make(chan *pictureComposite)
	channelAltName := make(chan *altNameComposite)
	waitGroup := &sync.WaitGroup{}
	aggregationContext := context.TODO()

	waitGroup.Add(1)

	go buildRecipeProcessesAggregate(aggregationContext, waitGroup, channelRecipeProcess, channelPicture, channelAltName)
	go buildPicturesAggregate(aggregationContext, waitGroup, channelPicture, channelAltName)
	go buildAltNames(aggregationContext, waitGroup, channelAltName)

	channelRecipeProcess <- &recipeProcessComposite{Entities: &recipeProcessesAggregate, Id: id, UserId: userId, EntityId: entityId, Criteria: criteria}

	waitGroup.Wait()

	close(channelRecipeProcess)
	close(channelPicture)
	close(channelAltName)

	return recipeProcessesAggregate, errorBuildingRecipeProcesses
}

func buildRecipeProcessesAggregate(
	aggregationContext context.Context,
	parentWaitGroup *sync.WaitGroup,
	channelRecipeProcess chan *recipeProcessComposite,
	channelPicture chan *pictureComposite,
	channelAltName chan *altNameComposite,
) {
	var (
		recipeProcessEntities      []*DomainEntity.RecipeProcess
		errorRecipeProcessEntities error
	)
	recipeProcessRepository := factoryRepository.GetRecipeProcessRepository()
	recipeProcessRepositoryCriteria := recipeProcessRepository.GetCriteria()

	for {
		select {
		case <-aggregationContext.Done():
			return
		case recipeProcessCompositeItem := <-channelRecipeProcess:
			if recipeProcessCompositeItem == nil {
				continue
			}
			recipeProcessEntities, errorRecipeProcessEntities = recipeProcessRepository.FindAll(
				composeCriteria(
					recipeProcessCompositeItem.Id,
					recipeProcessCompositeItem.UserId,
					recipeProcessCompositeItem.EntityId,
					recipeProcessCompositeItem.Criteria,
					recipeProcessRepositoryCriteria,
				),
			)

			if errorRecipeProcessEntities != nil || len(recipeProcessEntities) == 0 {
				errorBuildingRecipe = errorRecipeProcessEntities
			} else {
				for _, recipeProcessEntity := range recipeProcessEntities {
					recipeProcessAggregate := &DomainAggregate.RecipeProcess{Entity: recipeProcessEntity}
					*recipeProcessCompositeItem.Entities = append(*recipeProcessCompositeItem.Entities, recipeProcessAggregate)

					parentWaitGroup.Add(2)

					channelPicture <- &pictureComposite{Entities: &recipeProcessAggregate.Pictures, UserId: &recipeProcessEntity.UserId, EntityId: &recipeProcessEntity.Id}
					channelAltName <- &altNameComposite{Entities: &recipeProcessAggregate.AltNames, UserId: &recipeProcessEntity.UserId, EntityId: &recipeProcessEntity.Id}
				}
			}

			parentWaitGroup.Done()

		}
	}
}

func BuildPicturesAggregate(
	id *uuid.UUID,
	userId *uuid.UUID,
	entityId *uuid.UUID,
	criteria *persistence.Criteria,
) ([]*DomainAggregate.Picture, error) {
	var picturesAggregate []*DomainAggregate.Picture
	channelPicture := make(chan *pictureComposite)
	channelAltName := make(chan *altNameComposite)
	waitGroup := &sync.WaitGroup{}
	aggregationContext := context.TODO()

	waitGroup.Add(1)

	go buildPicturesAggregate(aggregationContext, waitGroup, channelPicture, channelAltName)
	go buildAltNames(aggregationContext, waitGroup, channelAltName)

	channelPicture <- &pictureComposite{Entities: &picturesAggregate, Id: id, UserId: userId, EntityId: entityId, Criteria: criteria}

	waitGroup.Wait()

	close(channelPicture)
	close(channelAltName)

	return picturesAggregate, errorBuildingPictures
}

func buildPicturesAggregate(
	aggregationContext context.Context,
	parentWaitGroup *sync.WaitGroup,
	channelPicture chan *pictureComposite,
	channelAltName chan *altNameComposite,
) {
	var (
		pictureEntities      []*DomainEntity.Picture
		errorPictureEntities error
	)
	pictureRepository := factoryRepository.GetPictureRepository()
	pictureRepositoryCriteria := pictureRepository.GetCriteria()

	for {
		select {
		case <-aggregationContext.Done():
			return
		case pictureCompositeItem := <-channelPicture:
			if pictureCompositeItem == nil {
				continue
			}
			pictureEntities, errorPictureEntities = pictureRepository.FindAll(
				composeCriteria(
					pictureCompositeItem.Id,
					pictureCompositeItem.UserId,
					pictureCompositeItem.EntityId,
					pictureCompositeItem.Criteria,
					pictureRepositoryCriteria,
				),
			)

			if errorPictureEntities != nil || len(pictureEntities) == 0 {
				errorBuildingRecipe = errorPictureEntities
			} else {
				for _, pictureEntity := range pictureEntities {
					pictureAggregate := &DomainAggregate.Picture{Entity: pictureEntity}
					*pictureCompositeItem.Entities = append(*pictureCompositeItem.Entities, pictureAggregate)

					parentWaitGroup.Add(1)

					channelAltName <- &altNameComposite{Entities: &pictureAggregate.AltNames, UserId: &pictureEntity.UserId, EntityId: &pictureEntity.Id}
				}
			}

			parentWaitGroup.Done()

		}
	}
}

func BuildRecipeMeasuresAggregate(
	id *uuid.UUID,
	userId *uuid.UUID,
	entityId *uuid.UUID,
	criteria *persistence.Criteria,
) ([]*DomainAggregate.RecipeMeasure, error) {
	var recipeMeasuresAggregate []*DomainAggregate.RecipeMeasure
	channelRecipeMeasure := make(chan *recipeMeasureComposite)
	channelUnit := make(chan *unitComposite)
	channelAltName := make(chan *altNameComposite)
	waitGroup := &sync.WaitGroup{}
	aggregationContext := context.TODO()

	waitGroup.Add(1)

	go buildRecipeMeasuresAggregate(aggregationContext, waitGroup, channelRecipeMeasure, channelUnit, channelAltName)
	go buildUnitEntities(aggregationContext, waitGroup, channelUnit)
	go buildAltNames(aggregationContext, waitGroup, channelAltName)

	channelRecipeMeasure <- &recipeMeasureComposite{Entities: &recipeMeasuresAggregate, Id: id, UserId: userId, EntityId: entityId, Criteria: criteria}

	waitGroup.Wait()

	close(channelRecipeMeasure)
	close(channelUnit)
	close(channelAltName)

	return recipeMeasuresAggregate, errorBuildingRecipeMeasures
}

func buildRecipeMeasuresAggregate(
	aggregationContext context.Context,
	parentWaitGroup *sync.WaitGroup,
	channelRecipeMeasure chan *recipeMeasureComposite,
	channelUnit chan *unitComposite,
	channelAltName chan *altNameComposite,
) {
	var (
		recipeMeasureEntities      []*DomainEntity.RecipeMeasure
		errorRecipeMeasureEntities error
	)
	recipeMeasureRepository := factoryRepository.GetRecipeMeasureRepository()
	recipeMeasureRepositoryCriteria := recipeMeasureRepository.GetCriteria()

	for {
		select {
		case <-aggregationContext.Done():
			return
		case recipeMeasureCompositeItem := <-channelRecipeMeasure:
			if recipeMeasureCompositeItem == nil {
				continue
			}
			recipeMeasureEntities, errorRecipeMeasureEntities = recipeMeasureRepository.FindAll(
				composeCriteria(
					recipeMeasureCompositeItem.Id,
					recipeMeasureCompositeItem.UserId,
					recipeMeasureCompositeItem.EntityId,
					recipeMeasureCompositeItem.Criteria,
					recipeMeasureRepositoryCriteria,
				),
			)

			if errorRecipeMeasureEntities != nil || len(recipeMeasureEntities) == 0 {
				errorBuildingRecipe = errorRecipeMeasureEntities
			} else {
				for _, recipeMeasureEntity := range recipeMeasureEntities {
					recipeMeasureAggregate := &DomainAggregate.RecipeMeasure{Entity: recipeMeasureEntity}
					*recipeMeasureCompositeItem.Entities = append(*recipeMeasureCompositeItem.Entities, recipeMeasureAggregate)

					parentWaitGroup.Add(2)

					channelUnit <- &unitComposite{Entity: &recipeMeasureAggregate.Unit, Id: &recipeMeasureEntity.UnitId}
					channelAltName <- &altNameComposite{Entities: &recipeMeasureAggregate.AltNames, UserId: &recipeMeasureEntity.UserId, EntityId: &recipeMeasureEntity.Id}
				}
			}

			parentWaitGroup.Done()

		}
	}
}

func BuildUnitEntities(
	id *uuid.UUID,
	criteria *persistence.Criteria,
) ([]*DomainEntity.Unit, error) {
	var unitEntities []*DomainEntity.Unit
	channelUnit := make(chan *unitComposite)
	waitGroup := &sync.WaitGroup{}
	aggregationContext := context.TODO()

	waitGroup.Add(1)

	go buildUnitEntities(aggregationContext, waitGroup, channelUnit)

	channelUnit <- &unitComposite{Entities: &unitEntities, Id: id, Criteria: criteria}

	waitGroup.Wait()

	close(channelUnit)

	return unitEntities, errorBuildingUnits
}

func buildUnitEntities(
	aggregationContext context.Context,
	parentWaitGroup *sync.WaitGroup,
	channelUnit chan *unitComposite,
) {
	var (
		unitEntities      []*DomainEntity.Unit
		errorUnitEntities error
	)
	unitRepository := factoryRepository.GetUnitRepository()
	unitRepositoryCriteria := unitRepository.GetCriteria()

	for {
		select {
		case <-aggregationContext.Done():
			return
		case unitCompositeItem := <-channelUnit:
			if unitCompositeItem == nil {
				continue
			}
			unitEntities, errorUnitEntities = unitRepository.FindAll(
				composeCriteria(
					unitCompositeItem.Id,
					nil,
					nil,
					unitCompositeItem.Criteria,
					unitRepositoryCriteria,
				),
			)

			if errorUnitEntities != nil || len(unitEntities) == 0 {
				errorBuildingRecipe = errorUnitEntities
			} else {
				for _, unitEntity := range unitEntities {
					if unitCompositeItem.Entities != nil {
						*unitCompositeItem.Entities = append(*unitCompositeItem.Entities, unitEntity)
					} else if unitCompositeItem.Entity != nil {
						*unitCompositeItem.Entity = unitEntity
					}
				}
			}

			parentWaitGroup.Done()

		}
	}
}

func BuildCategoryAggregate(
	id *uuid.UUID,
	userId *uuid.UUID,
	criteria *persistence.Criteria,
) ([]*DomainAggregate.Category, error) {
	var categoryEntities []*DomainAggregate.Category
	channelCategory := make(chan *categoryComposite)
	channelPicture := make(chan *pictureComposite)
	channelAltName := make(chan *altNameComposite)
	waitGroup := &sync.WaitGroup{}
	aggregationContext := context.TODO()

	waitGroup.Add(1)

	go buildCategoryAggregate(aggregationContext, waitGroup, channelCategory, channelPicture, channelAltName)
	go buildPicturesAggregate(aggregationContext, waitGroup, channelPicture, channelAltName)
	go buildAltNames(aggregationContext, waitGroup, channelAltName)

	channelCategory <- &categoryComposite{Entities: &categoryEntities, Id: id, UserId: userId, Criteria: criteria}

	waitGroup.Wait()

	close(channelCategory)
	close(channelPicture)
	close(channelAltName)

	return categoryEntities, errorBuildingCategories
}

func buildCategoryAggregate(
	aggregationContext context.Context,
	parentWaitGroup *sync.WaitGroup,
	channelCategory chan *categoryComposite,
	channelPicture chan *pictureComposite,
	channelAltName chan *altNameComposite,
) {
	var (
		categoryEntities      []*DomainEntity.Category
		errorCategoryEntities error
	)
	categoryRepository := factoryRepository.GetCategoryRepository()
	categoryRepositoryCriteria := categoryRepository.GetCriteria()

	for {
		select {
		case <-aggregationContext.Done():
			return
		case categoryCompositeItem := <-channelCategory:
			if categoryCompositeItem == nil {
				continue
			}
			categoryEntities, errorCategoryEntities = categoryRepository.FindAll(
				composeCriteria(
					categoryCompositeItem.Id,
					categoryCompositeItem.UserId,
					nil,
					categoryCompositeItem.Criteria,
					categoryRepositoryCriteria,
				),
			)

			if errorCategoryEntities != nil || len(categoryEntities) == 0 {
				errorBuildingRecipe = errorCategoryEntities
			} else {
				for _, categoryEntity := range categoryEntities {
					recipeCategoryAggregate := &DomainAggregate.Category{Entity: categoryEntity}
					if categoryCompositeItem.Entities != nil {
						*categoryCompositeItem.Entities = append(*categoryCompositeItem.Entities, recipeCategoryAggregate)
					} else if categoryCompositeItem.Entity != nil {
						*categoryCompositeItem.Entity = recipeCategoryAggregate
					}

					parentWaitGroup.Add(2)

					channelPicture <- &pictureComposite{Entities: &recipeCategoryAggregate.Pictures, UserId: &categoryEntity.UserId, EntityId: &categoryEntity.Id}
					channelAltName <- &altNameComposite{Entities: &recipeCategoryAggregate.AltNames, UserId: &categoryEntity.UserId, EntityId: &categoryEntity.Id}
				}
			}

			parentWaitGroup.Done()

		}
	}
}

func BuildIngredientEntities(
	id *uuid.UUID,
	userId *uuid.UUID,
	criteria *persistence.Criteria,
) ([]*DomainEntity.Ingredient, error) {
	var ingredientEntities []*DomainEntity.Ingredient
	channelIngredient := make(chan *ingredientComposite)
	waitGroup := &sync.WaitGroup{}
	aggregationContext := context.TODO()

	waitGroup.Add(1)

	go buildIngredientEntities(aggregationContext, waitGroup, channelIngredient)

	channelIngredient <- &ingredientComposite{Entities: &ingredientEntities, Id: id, UserId: userId, Criteria: criteria}

	waitGroup.Wait()

	close(channelIngredient)

	return ingredientEntities, errorBuildingIngredients
}

func buildIngredientEntities(
	aggregationContext context.Context,
	parentWaitGroup *sync.WaitGroup,
	channelIngredient chan *ingredientComposite,
) {
	var (
		ingredientEntities      []*DomainEntity.Ingredient
		errorIngredientEntities error
	)
	ingredientRepository := factoryRepository.GetIngredientRepository()
	ingredientRepositoryCriteria := ingredientRepository.GetCriteria()

	for {
		select {
		case <-aggregationContext.Done():
			return
		case ingredientCompositeItem := <-channelIngredient:
			if ingredientCompositeItem == nil {
				continue
			}
			ingredientEntities, errorIngredientEntities = ingredientRepository.FindAll(
				composeCriteria(
					ingredientCompositeItem.Id,
					ingredientCompositeItem.UserId,
					nil,
					ingredientCompositeItem.Criteria,
					ingredientRepositoryCriteria,
				),
			)

			if errorIngredientEntities != nil || len(ingredientEntities) == 0 {
				errorBuildingRecipe = errorIngredientEntities
			} else {
				for _, ingredientEntity := range ingredientEntities {
					if ingredientCompositeItem.Entities != nil {
						*ingredientCompositeItem.Entities = append(*ingredientCompositeItem.Entities, ingredientEntity)
					} else if ingredientCompositeItem.Entity != nil {
						*ingredientCompositeItem.Entity = ingredientEntity
					}
				}
			}

			parentWaitGroup.Done()

		}
	}
}

func BuildAltNameEntities(
	id *uuid.UUID,
	userId *uuid.UUID,
	entityId *uuid.UUID,
	criteria *persistence.Criteria,
) ([]*DomainEntity.AltName, error) {
	channelAltName := make(chan *altNameComposite)
	waitGroup := &sync.WaitGroup{}
	aggregationContext := context.TODO()
	var recipeAltNameEntities []*DomainEntity.AltName

	waitGroup.Add(1)

	go buildAltNames(aggregationContext, waitGroup, channelAltName)

	channelAltName <- &altNameComposite{Entities: &recipeAltNameEntities, Id: id, UserId: userId, EntityId: entityId, Criteria: criteria}

	waitGroup.Wait()

	close(channelAltName)

	return recipeAltNameEntities, errorBuildingAltNames
}

func buildAltNames(
	aggregationContext context.Context,
	parentWaitGroup *sync.WaitGroup,
	channelAltName chan *altNameComposite,
) {
	var (
		recipeAltNameEntities []*DomainEntity.AltName
		errorAltNameEntities  error
	)
	recipeAltNameRepository := factoryRepository.GetAltNameRepository()
	recipeAltNameRepositoryCriteria := recipeAltNameRepository.GetCriteria()

	for {
		select {
		case <-aggregationContext.Done():
			return
		case altNameCompositeItem := <-channelAltName:
			if altNameCompositeItem == nil {
				continue
			}
			recipeAltNameEntities, errorAltNameEntities = recipeAltNameRepository.FindAll(
				composeCriteria(
					altNameCompositeItem.Id,
					altNameCompositeItem.UserId,
					altNameCompositeItem.EntityId,
					altNameCompositeItem.Criteria,
					recipeAltNameRepositoryCriteria,
				),
			)
			if errorAltNameEntities != nil || len(recipeAltNameEntities) == 0 {
				errorBuildingRecipe = errorAltNameEntities
			} else {
				for _, recipeAltNameEntity := range recipeAltNameEntities {
					*altNameCompositeItem.Entities = append(*altNameCompositeItem.Entities, recipeAltNameEntity)
				}
			}

			parentWaitGroup.Done()

		}
	}
}

func BuildPlannersAggregate(
	id *uuid.UUID,
	userId *uuid.UUID,
	criteria *persistence.Criteria,
) ([]*DomainAggregate.Planner, error) {
	var recipeIngredientsAggregate []*DomainAggregate.Planner
	channelPlanner := make(chan *plannerComposite)
	channelPlannerInterval := make(chan *plannerIntervalComposite)
	channelPlannerRecipe := make(chan *plannerRecipeComposite)
	waitGroup := &sync.WaitGroup{}
	aggregationContext := context.TODO()

	waitGroup.Add(1)

	go buildPlannersAggregate(aggregationContext, waitGroup, channelPlanner, channelPlannerInterval)
	go buildPlannerIntervals(aggregationContext, waitGroup, channelPlannerInterval, channelPlannerRecipe)
	go buildPlannerRecipes(aggregationContext, waitGroup, channelPlannerRecipe)

	channelPlanner <- &plannerComposite{Entities: &recipeIngredientsAggregate, Id: id, UserId: userId, Criteria: criteria}

	waitGroup.Wait()

	close(channelPlanner)
	close(channelPlannerInterval)
	close(channelPlannerRecipe)

	return recipeIngredientsAggregate, errorBuildingPlanners
}

func buildPlannersAggregate(
	aggregationContext context.Context,
	parentWaitGroup *sync.WaitGroup,
	channelPlanner chan *plannerComposite,
	channelPlannerInterval chan *plannerIntervalComposite,
) {
	var (
		plannerEntities      []*DomainEntity.Planner
		errorPlannerEntities error
	)
	plannerRepository := factoryRepository.GetPlannerRepository()
	plannerRepositoryCriteria := plannerRepository.GetCriteria()

	for {
		select {
		case <-aggregationContext.Done():
			return
		case plannerCompositeItem := <-channelPlanner:
			if plannerCompositeItem == nil {
				continue
			}
			plannerEntities, errorPlannerEntities = plannerRepository.FindAll(
				composeCriteria(
					plannerCompositeItem.Id,
					plannerCompositeItem.UserId,
					nil,
					plannerCompositeItem.Criteria,
					plannerRepositoryCriteria,
				),
			)

			if errorPlannerEntities != nil || len(plannerEntities) == 0 {
				errorBuildingRecipe = errorPlannerEntities
			} else {
				for _, plannerEntity := range plannerEntities {
					plannerAggregate := &DomainAggregate.Planner{Entity: plannerEntity}
					*plannerCompositeItem.Entities = append(*plannerCompositeItem.Entities, plannerAggregate)

					parentWaitGroup.Add(1)

					channelPlannerInterval <- &plannerIntervalComposite{Entities: &plannerAggregate.Intervals, UserId: &plannerEntity.UserId, EntityId: &plannerEntity.Id}
				}
			}

			parentWaitGroup.Done()

		}
	}
}

func BuildPlannerIntervalAggregates(
	id *uuid.UUID,
	userId *uuid.UUID,
	entityId *uuid.UUID,
	criteria *persistence.Criteria,
) ([]*DomainAggregate.PlannerInterval, error) {
	channelPlannerInterval := make(chan *plannerIntervalComposite)
	channelPlannerRecipe := make(chan *plannerRecipeComposite)
	waitGroup := &sync.WaitGroup{}
	aggregationContext := context.TODO()
	var recipePlannerIntervalEntities []*DomainAggregate.PlannerInterval

	waitGroup.Add(1)

	go buildPlannerIntervals(aggregationContext, waitGroup, channelPlannerInterval, channelPlannerRecipe)
	go buildPlannerRecipes(aggregationContext, waitGroup, channelPlannerRecipe)

	channelPlannerInterval <- &plannerIntervalComposite{Entities: &recipePlannerIntervalEntities, Id: id, UserId: userId, EntityId: entityId, Criteria: criteria}

	waitGroup.Wait()

	close(channelPlannerInterval)
	close(channelPlannerRecipe)

	return recipePlannerIntervalEntities, errorBuildingPlannerIntervals
}

func buildPlannerIntervals(
	aggregationContext context.Context,
	parentWaitGroup *sync.WaitGroup,
	channelPlannerInterval chan *plannerIntervalComposite,
	channelPlannerRecipe chan *plannerRecipeComposite,
) {
	var (
		plannerIntervalEntities      []*DomainEntity.PlannerInterval
		errorPlannerIntervalEntities error
	)
	plannerIntervalRepository := factoryRepository.GetPlannerIntervalRepository()
	plannerIntervalRepositoryCriteria := plannerIntervalRepository.GetCriteria()

	for {
		select {
		case <-aggregationContext.Done():
			return
		case plannerIntervalCompositeItem := <-channelPlannerInterval:
			if plannerIntervalCompositeItem == nil {
				continue
			}
			plannerIntervalEntities, errorPlannerIntervalEntities = plannerIntervalRepository.FindAll(
				composeCriteria(
					plannerIntervalCompositeItem.Id,
					plannerIntervalCompositeItem.UserId,
					plannerIntervalCompositeItem.EntityId,
					plannerIntervalCompositeItem.Criteria,
					plannerIntervalRepositoryCriteria,
				),
			)
			if errorPlannerIntervalEntities != nil || len(plannerIntervalEntities) == 0 {
				errorBuildingRecipe = errorPlannerIntervalEntities
			} else {
				for _, plannerIntervalEntity := range plannerIntervalEntities {
					if plannerIntervalCompositeItem.Entities != nil {
						plannerIntervalAggregate := &DomainAggregate.PlannerInterval{Entity: plannerIntervalEntity}
						*plannerIntervalCompositeItem.Entities = append(*plannerIntervalCompositeItem.Entities, plannerIntervalAggregate)

						parentWaitGroup.Add(1)

						channelPlannerRecipe <- &plannerRecipeComposite{Entities: &plannerIntervalAggregate.Recipes, UserId: &plannerIntervalEntity.UserId, EntityId: &plannerIntervalEntity.Id}
					}
				}
			}

			parentWaitGroup.Done()
		}
	}
}

func BuildPlannerRecipeAggregates(
	id *uuid.UUID,
	userId *uuid.UUID,
	entityId *uuid.UUID,
	criteria *persistence.Criteria,
) ([]*DomainAggregate.PlannerRecipe, error) {
	channelPlannerRecipe := make(chan *plannerRecipeComposite)
	waitGroup := &sync.WaitGroup{}
	aggregationContext := context.TODO()
	waitGroup.Add(1)
	var recipePlannerRecipeEntities []*DomainAggregate.PlannerRecipe

	go buildPlannerRecipes(aggregationContext, waitGroup, channelPlannerRecipe)

	channelPlannerRecipe <- &plannerRecipeComposite{Entities: &recipePlannerRecipeEntities, Id: id, UserId: userId, EntityId: entityId, Criteria: criteria}

	waitGroup.Wait()

	close(channelPlannerRecipe)

	return recipePlannerRecipeEntities, errorBuildingPlannerRecipes
}

func buildPlannerRecipes(
	aggregationContext context.Context,
	parentWaitGroup *sync.WaitGroup,
	channelPlannerRecipe chan *plannerRecipeComposite,
) {
	var (
		plannerRecipeEntities      []*DomainEntity.PlannerRecipe
		errorPlannerRecipeEntities error
	)
	plannerRecipeRepository := factoryRepository.GetPlannerRecipeRepository()
	plannerRecipeRepositoryCriteria := plannerRecipeRepository.GetCriteria()

	for {
		select {
		case <-aggregationContext.Done():
			return
		case plannerRecipeCompositeItem := <-channelPlannerRecipe:
			if plannerRecipeCompositeItem == nil {
				continue
			}
			plannerRecipeEntities, errorPlannerRecipeEntities = plannerRecipeRepository.FindAll(
				composeCriteria(
					plannerRecipeCompositeItem.Id,
					plannerRecipeCompositeItem.UserId,
					plannerRecipeCompositeItem.EntityId,
					plannerRecipeCompositeItem.Criteria,
					plannerRecipeRepositoryCriteria,
				),
			)
			if errorPlannerRecipeEntities != nil || len(plannerRecipeEntities) == 0 {
				errorBuildingRecipe = errorPlannerRecipeEntities
			} else {
				for _, plannerRecipeEntity := range plannerRecipeEntities {
					recipesAggregate, errorRecipesAggregate := BuildRecipesAggregate(&plannerRecipeEntity.RecipeId, &plannerRecipeEntity.UserId, nil)
					if errorRecipesAggregate == nil && len(recipesAggregate) == 1 {
						plannerRecipeAggregate := &DomainAggregate.PlannerRecipe{Entity: plannerRecipeEntity, Recipe: recipesAggregate[0]}
						*plannerRecipeCompositeItem.Entities = append(*plannerRecipeCompositeItem.Entities, plannerRecipeAggregate)
					}

				}
			}

			parentWaitGroup.Done()

		}
	}
}

func composeCriteria(
	id *uuid.UUID,
	userId *uuid.UUID,
	entityId *uuid.UUID,
	criteria *persistence.Criteria,
	criteriaRepository *repository.CriteriaRepository) *persistence.Criteria {

	if id != nil {
		criteria = criteriaRepository.GetCriteriaById(
			id,
			criteria,
		)
	}

	criteria = criteriaRepository.GetCriteriaByUserId(userId, criteria)

	if entityId != nil {
		criteria = criteriaRepository.GetCriteriaByEntityId(
			entityId,
			criteria,
		)
	}

	return criteria
}
