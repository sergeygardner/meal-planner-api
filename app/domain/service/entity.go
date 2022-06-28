package service

import (
	"encoding/json"
	"github.com/sergeygardner/meal-planner-api/domain/entity"
	"io"
)

func CreateEntityFromRecipeUpdate(data io.Reader) (entity.Recipe, error) {
	recipe := &entity.Recipe{}
	errorEntity := json.NewDecoder(data).Decode(&recipe)

	return *recipe, errorEntity
}

func CreateEntityFromRecipeCategoryUpdate(data io.Reader) (entity.RecipeCategory, error) {
	recipeCategory := &entity.RecipeCategory{}
	errorEntity := json.NewDecoder(data).Decode(&recipeCategory)

	return *recipeCategory, errorEntity
}

func CreateEntityFromRecipeIngredientUpdate(data io.Reader) (entity.RecipeIngredient, error) {
	recipeIngredient := &entity.RecipeIngredient{}
	errorEntity := json.NewDecoder(data).Decode(&recipeIngredient)

	return *recipeIngredient, errorEntity
}

func CreateEntityFromRecipeProcessUpdate(data io.Reader) (entity.RecipeProcess, error) {
	recipeProcess := &entity.RecipeProcess{}
	errorEntity := json.NewDecoder(data).Decode(&recipeProcess)

	return *recipeProcess, errorEntity
}

func CreateEntityFromPictureUpdate(data io.Reader) (entity.Picture, error) {
	recipePicture := &entity.Picture{}
	errorEntity := json.NewDecoder(data).Decode(&recipePicture)

	return *recipePicture, errorEntity
}

func CreateEntityFromRecipeMeasureUpdate(data io.Reader) (entity.RecipeMeasure, error) {
	recipeMeasure := &entity.RecipeMeasure{}
	errorEntity := json.NewDecoder(data).Decode(&recipeMeasure)

	return *recipeMeasure, errorEntity
}

func CreateEntityFromAltNameUpdate(data io.Reader) (entity.AltName, error) {
	recipeAltName := &entity.AltName{}
	errorEntity := json.NewDecoder(data).Decode(&recipeAltName)

	return *recipeAltName, errorEntity
}

func CreateEntityFromUnitUpdate(data io.Reader) (entity.Unit, error) {
	recipeUnit := &entity.Unit{}
	errorEntity := json.NewDecoder(data).Decode(&recipeUnit)

	return *recipeUnit, errorEntity
}

func CreateEntityFromCategoryUpdate(data io.Reader) (entity.Category, error) {
	recipeCategory := &entity.Category{}
	errorEntity := json.NewDecoder(data).Decode(&recipeCategory)

	return *recipeCategory, errorEntity
}

func CreateEntityFromIngredientUpdate(data io.Reader) (entity.Ingredient, error) {
	recipeIngredient := &entity.Ingredient{}
	errorEntity := json.NewDecoder(data).Decode(&recipeIngredient)

	return *recipeIngredient, errorEntity
}

func CreateEntityFromPlannerUpdate(data io.Reader) (entity.Planner, error) {
	recipePlanner := &entity.Planner{}
	errorEntity := json.NewDecoder(data).Decode(&recipePlanner)

	return *recipePlanner, errorEntity
}

func CreateEntityFromPlannerIntervalUpdate(data io.Reader) (entity.PlannerInterval, error) {
	recipePlannerInterval := &entity.PlannerInterval{}
	errorEntity := json.NewDecoder(data).Decode(&recipePlannerInterval)

	return *recipePlannerInterval, errorEntity
}

func CreateEntityFromPlannerRecipeUpdate(data io.Reader) (entity.PlannerRecipe, error) {
	recipePlannerRecipe := &entity.PlannerRecipe{}
	errorEntity := json.NewDecoder(data).Decode(&recipePlannerRecipe)

	return *recipePlannerRecipe, errorEntity
}
