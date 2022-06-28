package repository

import (
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence"
	MongoDBRepository "github.com/sergeygardner/meal-planner-api/infrastructure/persistence/mongodb/repository"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence/repository"
	"github.com/sergeygardner/meal-planner-api/infrastructure/service/entity"
)

var (
	factory FactoryRepositoryInterface
)

type FactoryRepositoryInterface interface {
	GetUserRepository() repository.UserRepositoryInterface
	GetUserRoleRepository() repository.UserRoleRepositoryInterface
	GetUserToRoleRepository() repository.UserToRoleRepositoryInterface
	GetUserConfirmationRepository() repository.UserConfirmationRepositoryInterface
	GetRecipeRepository() repository.RecipeRepositoryInterface
	GetRecipeCategoryRepository() repository.RecipeCategoryRepositoryInterface
	GetRecipeIngredientRepository() repository.RecipeIngredientRepositoryInterface
	GetRecipeProcessRepository() repository.RecipeProcessRepositoryInterface
	GetRecipeMeasureRepository() repository.RecipeMeasureRepositoryInterface
	GetUnitRepository() repository.UnitRepositoryInterface
	GetCategoryRepository() repository.CategoryRepositoryInterface
	GetIngredientRepository() repository.IngredientRepositoryInterface
	GetPictureRepository() repository.PictureRepositoryInterface
	GetAltNameRepository() repository.AltNameRepositoryInterface
	GetPlannerRepository() repository.PlannerRepositoryInterface
	GetPlannerIntervalRepository() repository.PlannerIntervalRepositoryInterface
	GetPlannerRecipeRepository() repository.PlannerRecipeRepositoryInterface
}

type FactoryRepository struct {
	userRepository             repository.UserRepositoryInterface
	userGroupRepository        repository.UserRoleRepositoryInterface
	userToGroupRepository      repository.UserToRoleRepositoryInterface
	userConfirmationRepository repository.UserConfirmationRepositoryInterface
	recipeRepository           repository.RecipeRepositoryInterface
	recipeCategoryRepository   repository.RecipeCategoryRepositoryInterface
	recipeIngredientRepository repository.RecipeIngredientRepositoryInterface
	recipeProcessRepository    repository.RecipeProcessRepositoryInterface
	recipeMeasureRepository    repository.RecipeMeasureRepositoryInterface
	unitRepository             repository.UnitRepositoryInterface
	categoryRepository         repository.CategoryRepositoryInterface
	ingredientRepository       repository.IngredientRepositoryInterface
	pictureRepository          repository.PictureRepositoryInterface
	altNameRepository          repository.AltNameRepositoryInterface
	plannerRepository          repository.PlannerRepositoryInterface
	plannerIntervalRepository  repository.PlannerIntervalRepositoryInterface
	plannerRecipeRepository    repository.PlannerRecipeRepositoryInterface
	FactoryRepositoryInterface
}

func (f *FactoryRepository) GetUserRepository() repository.UserRepositoryInterface {
	if f.userRepository == nil {
		entityManager := entity.GetEntityManager()

		switch entityManager.GetType() {
		case persistence.MongoType:
			f.userRepository = &MongoDBRepository.UserRepository{Table: "user", EntityManager: entity.GetEntityManager()}
		default:
			f.userRepository = &MongoDBRepository.UserRepository{Table: "user", EntityManager: entity.GetEntityManager()}
		}
	}

	return f.userRepository
}

func (f *FactoryRepository) GetUserRoleRepository() repository.UserRoleRepositoryInterface {
	if f.userGroupRepository == nil {
		entityManager := entity.GetEntityManager()

		switch entityManager.GetType() {
		case persistence.MongoType:
			f.userGroupRepository = &MongoDBRepository.UserRoleRepository{Table: "user_group", EntityManager: entity.GetEntityManager()}
		default:
			f.userGroupRepository = &MongoDBRepository.UserRoleRepository{Table: "user_group", EntityManager: entity.GetEntityManager()}
		}
	}

	return f.userGroupRepository
}

func (f *FactoryRepository) GetUserToRoleRepository() repository.UserToRoleRepositoryInterface {
	if f.userToGroupRepository == nil {
		entityManager := entity.GetEntityManager()

		switch entityManager.GetType() {
		case persistence.MongoType:
			f.userToGroupRepository = &MongoDBRepository.UserToRoleRepository{Table: "user_to_group", EntityManager: entity.GetEntityManager()}
		default:
			f.userToGroupRepository = &MongoDBRepository.UserToRoleRepository{Table: "user_to_group", EntityManager: entity.GetEntityManager()}
		}
	}

	return f.userToGroupRepository
}

func (f *FactoryRepository) GetUserConfirmationRepository() repository.UserConfirmationRepositoryInterface {
	if f.userConfirmationRepository == nil {
		entityManager := entity.GetEntityManager()

		switch entityManager.GetType() {
		case persistence.MongoType:
			f.userConfirmationRepository = &MongoDBRepository.UserConfirmationRepository{Table: "user_confirmation", EntityManager: entity.GetEntityManager()}
		default:
			f.userConfirmationRepository = &MongoDBRepository.UserConfirmationRepository{Table: "user_confirmation", EntityManager: entity.GetEntityManager()}
		}
	}

	return f.userConfirmationRepository
}

func (f *FactoryRepository) GetRecipeRepository() repository.RecipeRepositoryInterface {
	if f.recipeRepository == nil {
		entityManager := entity.GetEntityManager()

		switch entityManager.GetType() {
		case persistence.MongoType:
			f.recipeRepository = &MongoDBRepository.RecipeRepository{Table: "recipe", EntityManager: entity.GetEntityManager()}
		default:
			f.recipeRepository = &MongoDBRepository.RecipeRepository{Table: "recipe", EntityManager: entity.GetEntityManager()}
		}
	}

	return f.recipeRepository
}

func (f *FactoryRepository) GetRecipeCategoryRepository() repository.RecipeCategoryRepositoryInterface {
	if f.recipeCategoryRepository == nil {
		entityManager := entity.GetEntityManager()

		switch entityManager.GetType() {
		case persistence.MongoType:
			f.recipeCategoryRepository = &MongoDBRepository.RecipeCategoryRepository{Table: "recipe_category", EntityManager: entity.GetEntityManager()}
		default:
			f.recipeCategoryRepository = &MongoDBRepository.RecipeCategoryRepository{Table: "recipe_category", EntityManager: entity.GetEntityManager()}
		}
	}

	return f.recipeCategoryRepository
}

func (f *FactoryRepository) GetRecipeIngredientRepository() repository.RecipeIngredientRepositoryInterface {
	if f.recipeIngredientRepository == nil {
		entityManager := entity.GetEntityManager()

		switch entityManager.GetType() {
		case persistence.MongoType:
			f.recipeIngredientRepository = &MongoDBRepository.RecipeIngredientRepository{Table: "recipe_ingredient", EntityManager: entity.GetEntityManager()}
		default:
			f.recipeIngredientRepository = &MongoDBRepository.RecipeIngredientRepository{Table: "recipe_ingredient", EntityManager: entity.GetEntityManager()}
		}
	}

	return f.recipeIngredientRepository
}

func (f *FactoryRepository) GetRecipeProcessRepository() repository.RecipeProcessRepositoryInterface {
	if f.recipeProcessRepository == nil {
		entityManager := entity.GetEntityManager()

		switch entityManager.GetType() {
		case persistence.MongoType:
			f.recipeProcessRepository = &MongoDBRepository.RecipeProcessRepository{Table: "recipe_process", EntityManager: entity.GetEntityManager()}
		default:
			f.recipeProcessRepository = &MongoDBRepository.RecipeProcessRepository{Table: "recipe_process", EntityManager: entity.GetEntityManager()}
		}
	}

	return f.recipeProcessRepository
}

func (f *FactoryRepository) GetAltNameRepository() repository.AltNameRepositoryInterface {
	if f.altNameRepository == nil {
		entityManager := entity.GetEntityManager()

		switch entityManager.GetType() {
		case persistence.MongoType:
			f.altNameRepository = &MongoDBRepository.AltNameRepository{Table: "alt_name", EntityManager: entity.GetEntityManager()}
		default:
			f.altNameRepository = &MongoDBRepository.AltNameRepository{Table: "alt_name", EntityManager: entity.GetEntityManager()}
		}
	}

	return f.altNameRepository
}

func (f *FactoryRepository) GetPictureRepository() repository.PictureRepositoryInterface {
	if f.pictureRepository == nil {
		entityManager := entity.GetEntityManager()

		switch entityManager.GetType() {
		case persistence.MongoType:
			f.pictureRepository = &MongoDBRepository.PictureRepository{Table: "picture", EntityManager: entity.GetEntityManager()}
		default:
			f.pictureRepository = &MongoDBRepository.PictureRepository{Table: "picture", EntityManager: entity.GetEntityManager()}
		}
	}

	return f.pictureRepository
}

func (f *FactoryRepository) GetRecipeMeasureRepository() repository.RecipeMeasureRepositoryInterface {
	if f.recipeMeasureRepository == nil {
		entityManager := entity.GetEntityManager()

		switch entityManager.GetType() {
		case persistence.MongoType:
			f.recipeMeasureRepository = &MongoDBRepository.RecipeMeasureRepository{Table: "recipe_measure", EntityManager: entity.GetEntityManager()}
		default:
			f.recipeMeasureRepository = &MongoDBRepository.RecipeMeasureRepository{Table: "recipe_measure", EntityManager: entity.GetEntityManager()}
		}
	}

	return f.recipeMeasureRepository
}

func (f *FactoryRepository) GetUnitRepository() repository.UnitRepositoryInterface {
	if f.unitRepository == nil {
		entityManager := entity.GetEntityManager()

		switch entityManager.GetType() {
		case persistence.MongoType:
			f.unitRepository = &MongoDBRepository.UnitRepository{Table: "unit", EntityManager: entity.GetEntityManager()}
		default:
			f.unitRepository = &MongoDBRepository.UnitRepository{Table: "unit", EntityManager: entity.GetEntityManager()}
		}
	}

	return f.unitRepository
}

func (f *FactoryRepository) GetCategoryRepository() repository.CategoryRepositoryInterface {
	if f.categoryRepository == nil {
		entityManager := entity.GetEntityManager()

		switch entityManager.GetType() {
		case persistence.MongoType:
			f.categoryRepository = &MongoDBRepository.CategoryRepository{Table: "category", EntityManager: entity.GetEntityManager()}
		default:
			f.categoryRepository = &MongoDBRepository.CategoryRepository{Table: "category", EntityManager: entity.GetEntityManager()}
		}
	}

	return f.categoryRepository
}

func (f *FactoryRepository) GetIngredientRepository() repository.IngredientRepositoryInterface {
	if f.ingredientRepository == nil {
		entityManager := entity.GetEntityManager()

		switch entityManager.GetType() {
		case persistence.MongoType:
			f.ingredientRepository = &MongoDBRepository.IngredientRepository{Table: "ingredient", EntityManager: entity.GetEntityManager()}
		default:
			f.ingredientRepository = &MongoDBRepository.IngredientRepository{Table: "ingredient", EntityManager: entity.GetEntityManager()}
		}
	}

	return f.ingredientRepository
}

func (f *FactoryRepository) GetPlannerRepository() repository.PlannerRepositoryInterface {
	if f.plannerRepository == nil {
		entityManager := entity.GetEntityManager()

		switch entityManager.GetType() {
		case persistence.MongoType:
			f.plannerRepository = &MongoDBRepository.PlannerRepository{Table: "planner", EntityManager: entity.GetEntityManager()}
		default:
			f.plannerRepository = &MongoDBRepository.PlannerRepository{Table: "planner", EntityManager: entity.GetEntityManager()}
		}
	}

	return f.plannerRepository
}

func (f *FactoryRepository) GetPlannerIntervalRepository() repository.PlannerIntervalRepositoryInterface {
	if f.plannerIntervalRepository == nil {
		entityManager := entity.GetEntityManager()

		switch entityManager.GetType() {
		case persistence.MongoType:
			f.plannerIntervalRepository = &MongoDBRepository.PlannerIntervalRepository{Table: "planner_interval", EntityManager: entity.GetEntityManager()}
		default:
			f.plannerIntervalRepository = &MongoDBRepository.PlannerIntervalRepository{Table: "planner_interval", EntityManager: entity.GetEntityManager()}
		}
	}

	return f.plannerIntervalRepository
}

func (f *FactoryRepository) GetPlannerRecipeRepository() repository.PlannerRecipeRepositoryInterface {
	if f.plannerRecipeRepository == nil {
		entityManager := entity.GetEntityManager()

		switch entityManager.GetType() {
		case persistence.MongoType:
			f.plannerRecipeRepository = &MongoDBRepository.PlannerRecipeRepository{Table: "planner_recipe", EntityManager: entity.GetEntityManager()}
		default:
			f.plannerRecipeRepository = &MongoDBRepository.PlannerRecipeRepository{Table: "planner_recipe", EntityManager: entity.GetEntityManager()}
		}
	}

	return f.plannerRecipeRepository
}

func GetFactoryRepository() FactoryRepositoryInterface {
	if factory == nil {
		factory = &FactoryRepository{}
	}

	return factory
}
