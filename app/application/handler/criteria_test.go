package handler

import (
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/domain/dto"
	"github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPrepareUserRepositoryInsert(t *testing.T) {
	tests := []struct {
		Name            string
		UserRegisterDTO dto.UserRegisterDTO
		password        string
		MustBePanic     bool
		MustBeFault     bool
	}{
		{
			Name: "Test case with PrepareUserRepositoryInsert and correct data",
			UserRegisterDTO: dto.UserRegisterDTO{
				UserCredentialsDTO: dto.UserCredentialsDTO{
					Username: "usernameTest",
				},
				Name:       "NameTest",
				Surname:    "SurnameTest",
				MiddleName: "MiddleNameTest",
				Birthday:   time.Now().UTC(),
			},
			password:    "passwordTest",
			MustBePanic: false,
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				defer func() {
					if testCase.MustBePanic {
						assert.NotNil(t, recover())
					} else {
						assert.Nil(t, recover())
					}
				}()

				preparedUser := prepareUserRepositoryInsert(testCase.UserRegisterDTO, testCase.password)

				if testCase.MustBeFault {
					assert.Nil(t, preparedUser)
				} else {
					assert.NotNil(t, preparedUser)
					assert.NotEqual(t, uuid.Nil, preparedUser.Id)
					assert.Equal(t, testCase.UserRegisterDTO.Username, preparedUser.Username)
					assert.Equal(t, testCase.password, preparedUser.Password)
					assert.Equal(t, testCase.UserRegisterDTO.Name, preparedUser.Name)
					assert.Equal(t, testCase.UserRegisterDTO.Surname, preparedUser.Surname)
					assert.Equal(t, testCase.UserRegisterDTO.MiddleName, preparedUser.MiddleName)
					assert.Equal(t, testCase.UserRegisterDTO.Birthday, preparedUser.Birthday)
					assert.NotEqual(t, time.Date(0, 0, 0, 0, 0, 0, 0, time.Local), preparedUser.DateInsert)
					assert.NotEqual(t, time.Date(0, 0, 0, 0, 0, 0, 0, time.Local), preparedUser.DateUpdate)
					assert.Contains(t, []kind.UserStatus{kind.UserStatusRegister, kind.UserStatusNeedConfirmation, kind.UserStatusDisabled}, preparedUser.Status)
					assert.Contains(t, []bool{kind.UserInActive, kind.UserActive}, preparedUser.Active)
					assert.Subset(t, kind.UserRoles{kind.UserRoleCommon, kind.UserRoleAdmin}, preparedUser.Roles)
				}
			},
		)
	}
}

func TestPrepareUserConfirmationRepositoryInsert(t *testing.T) {
	tests := []struct {
		Name        string
		User        *entity.User
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name: "Test case with PrepareUserConfirmationRepositoryInsert and correct data",
			User: &entity.User{
				Id: uuid.New(),
			},
			MustBePanic: false,
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				defer func() {
					if testCase.MustBePanic {
						assert.NotNil(t, recover())
					} else {
						assert.Nil(t, recover())
					}
				}()

				preparedUserConfirmation := prepareUserConfirmationRepositoryInsert(*testCase.User)

				if testCase.MustBeFault {
					assert.Nil(t, preparedUserConfirmation)
				} else {
					assert.NotNil(t, preparedUserConfirmation)
					assert.NotEqual(t, uuid.Nil, preparedUserConfirmation.Id)
					assert.NotEqual(t, time.Date(0, 0, 0, 0, 0, 0, 0, time.Local), preparedUserConfirmation.DateInsert)
					assert.NotEqual(t, time.Date(0, 0, 0, 0, 0, 0, 0, time.Local), preparedUserConfirmation.DateUpdate)
					assert.Equal(t, testCase.User.Id, preparedUserConfirmation.UserId)
					assert.NotEmpty(t, preparedUserConfirmation.Value)
					assert.Contains(t, []bool{kind.UserConfirmationInActive, kind.UserConfirmationActive}, preparedUserConfirmation.Active)
				}
			},
		)
	}
}

func TestPrepareRecipeRepositoryInsert(t *testing.T) {
	tests := []struct {
		Name        string
		Recipe      *entity.Recipe
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name: "Test case with PrepareRecipeRepositoryInsert and correct data",
			Recipe: &entity.Recipe{
				UserId:      uuid.New(),
				DateInsert:  time.Now().UTC(),
				DateUpdate:  time.Now().UTC(),
				Name:        "Test case with correct data",
				Description: "Test case with correct data",
				Notes:       "Test case with correct data",
				Status:      kind.RecipeStatusUnPublished,
			},
			MustBePanic: false,
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				defer func() {
					if testCase.MustBePanic {
						assert.NotNil(t, recover())
					} else {
						assert.Nil(t, recover())
					}
				}()

				preparedRecipe := prepareRecipeRepositoryInsert(testCase.Recipe)

				if testCase.MustBeFault {
					assert.Nil(t, preparedRecipe)
				} else {
					assert.NotNil(t, preparedRecipe)
					assert.NotEqual(t, uuid.Nil, preparedRecipe.Id)
					assert.Equal(t, testCase.Recipe.UserId, preparedRecipe.UserId)
					assert.Equal(t, testCase.Recipe.DateInsert, preparedRecipe.DateInsert)
					assert.Equal(t, testCase.Recipe.DateUpdate, preparedRecipe.DateUpdate)
					assert.Equal(t, testCase.Recipe.Name, preparedRecipe.Name)
					assert.Equal(t, testCase.Recipe.Description, preparedRecipe.Description)
					assert.Equal(t, testCase.Recipe.Notes, preparedRecipe.Notes)
					assert.Equal(t, testCase.Recipe.Status, preparedRecipe.Status)
				}
			},
		)
	}
}

func TestPrepareRecipeCategoryRepositoryInsert(t *testing.T) {
	tests := []struct {
		Name           string
		RecipeCategory *entity.RecipeCategory
		MustBePanic    bool
		MustBeFault    bool
	}{
		{
			Name: "Test case with PrepareRecipeCategoryRepositoryInsert and correct data",
			RecipeCategory: &entity.RecipeCategory{
				UserId:     uuid.New(),
				EntityId:   uuid.New(),
				DeriveId:   uuid.New(),
				DateInsert: time.Now().UTC(),
				DateUpdate: time.Now().UTC(),
				Status:     kind.RecipeCategoryStatusUnPublished,
			},
			MustBePanic: false,
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				defer func() {
					if testCase.MustBePanic {
						assert.NotNil(t, recover())
					} else {
						assert.Nil(t, recover())
					}
				}()

				preparedRecipeCategory := prepareRecipeCategoryRepositoryInsert(testCase.RecipeCategory)

				if testCase.MustBeFault {
					assert.Nil(t, preparedRecipeCategory)
				} else {
					assert.NotNil(t, preparedRecipeCategory)
					assert.NotEqual(t, uuid.Nil, preparedRecipeCategory.Id)
					assert.Equal(t, testCase.RecipeCategory.UserId, preparedRecipeCategory.UserId)
					assert.Equal(t, testCase.RecipeCategory.EntityId, preparedRecipeCategory.EntityId)
					assert.Equal(t, testCase.RecipeCategory.DeriveId, preparedRecipeCategory.DeriveId)
					assert.Equal(t, testCase.RecipeCategory.DateInsert, preparedRecipeCategory.DateInsert)
					assert.Equal(t, testCase.RecipeCategory.DateUpdate, preparedRecipeCategory.DateUpdate)
					assert.Equal(t, testCase.RecipeCategory.Status, preparedRecipeCategory.Status)
				}
			},
		)
	}
}

func TestPrepareRecipeIngredientRepositoryInsert(t *testing.T) {
	tests := []struct {
		Name             string
		RecipeIngredient *entity.RecipeIngredient
		MustBePanic      bool
		MustBeFault      bool
	}{
		{
			Name: "Test case with PrepareRecipeIngredientRepositoryInsert and correct data",
			RecipeIngredient: &entity.RecipeIngredient{
				UserId:     uuid.New(),
				EntityId:   uuid.New(),
				DeriveId:   uuid.New(),
				DateInsert: time.Now().UTC(),
				DateUpdate: time.Now().UTC(),
				Name:       "Test case with correct data",
				Status:     kind.RecipeIngredientStatusUnPublished,
			},
			MustBePanic: false,
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				defer func() {
					if testCase.MustBePanic {
						assert.NotNil(t, recover())
					} else {
						assert.Nil(t, recover())
					}
				}()

				preparedRecipeIngredient := prepareRecipeIngredientRepositoryInsert(testCase.RecipeIngredient)

				if testCase.MustBeFault {
					assert.Nil(t, preparedRecipeIngredient)
				} else {
					assert.NotNil(t, preparedRecipeIngredient)
					assert.NotEqual(t, uuid.Nil, preparedRecipeIngredient.Id)
					assert.Equal(t, testCase.RecipeIngredient.UserId, preparedRecipeIngredient.UserId)
					assert.Equal(t, testCase.RecipeIngredient.EntityId, preparedRecipeIngredient.EntityId)
					assert.Equal(t, testCase.RecipeIngredient.DeriveId, preparedRecipeIngredient.DeriveId)
					assert.Equal(t, testCase.RecipeIngredient.DateInsert, preparedRecipeIngredient.DateInsert)
					assert.Equal(t, testCase.RecipeIngredient.DateUpdate, preparedRecipeIngredient.DateUpdate)
					assert.Equal(t, testCase.RecipeIngredient.Name, preparedRecipeIngredient.Name)
					assert.Equal(t, testCase.RecipeIngredient.Status, preparedRecipeIngredient.Status)
				}
			},
		)
	}
}

func TestPrepareRecipeProcessRepositoryInsert(t *testing.T) {
	tests := []struct {
		Name          string
		RecipeProcess *entity.RecipeProcess
		MustBePanic   bool
		MustBeFault   bool
	}{
		{
			Name: "Test case with PrepareRecipeProcessRepositoryInsert and correct data",
			RecipeProcess: &entity.RecipeProcess{
				UserId:      uuid.New(),
				EntityId:    uuid.New(),
				DateInsert:  time.Now().UTC(),
				DateUpdate:  time.Now().UTC(),
				Name:        "Test case with correct data",
				Description: "Test case with correct data",
				Notes:       "Test case with correct data",
				Status:      kind.RecipeProcessStatusUnPublished,
			},
			MustBePanic: false,
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				defer func() {
					if testCase.MustBePanic {
						assert.NotNil(t, recover())
					} else {
						assert.Nil(t, recover())
					}
				}()

				preparedRecipeProcess := prepareRecipeProcessRepositoryInsert(testCase.RecipeProcess)

				if testCase.MustBeFault {
					assert.Nil(t, preparedRecipeProcess)
				} else {
					assert.NotNil(t, preparedRecipeProcess)
					assert.NotEqual(t, uuid.Nil, preparedRecipeProcess.Id)
					assert.Equal(t, testCase.RecipeProcess.UserId, preparedRecipeProcess.UserId)
					assert.Equal(t, testCase.RecipeProcess.EntityId, preparedRecipeProcess.EntityId)
					assert.Equal(t, testCase.RecipeProcess.DateInsert, preparedRecipeProcess.DateInsert)
					assert.Equal(t, testCase.RecipeProcess.DateUpdate, preparedRecipeProcess.DateUpdate)
					assert.Equal(t, testCase.RecipeProcess.Name, preparedRecipeProcess.Name)
					assert.Equal(t, testCase.RecipeProcess.Description, preparedRecipeProcess.Description)
					assert.Equal(t, testCase.RecipeProcess.Notes, preparedRecipeProcess.Notes)
					assert.Equal(t, testCase.RecipeProcess.Status, preparedRecipeProcess.Status)
				}
			},
		)
	}
}

func TestPrepareRecipeMeasureRepositoryInsert(t *testing.T) {
	tests := []struct {
		Name          string
		RecipeMeasure *entity.RecipeMeasure
		MustBePanic   bool
		MustBeFault   bool
	}{
		{
			Name: "Test case with PrepareRecipeMeasureRepositoryInsert and correct data",
			RecipeMeasure: &entity.RecipeMeasure{
				UserId:     uuid.New(),
				EntityId:   uuid.New(),
				UnitId:     uuid.New(),
				DateInsert: time.Now().UTC(),
				DateUpdate: time.Now().UTC(),
				Status:     kind.RecipeMeasureStatusUnPublished,
			},
			MustBePanic: false,
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				defer func() {
					if testCase.MustBePanic {
						assert.NotNil(t, recover())
					} else {
						assert.Nil(t, recover())
					}
				}()

				preparedRecipeMeasure := prepareRecipeMeasureRepositoryInsert(testCase.RecipeMeasure)

				if testCase.MustBeFault {
					assert.Nil(t, preparedRecipeMeasure)
				} else {
					assert.NotNil(t, preparedRecipeMeasure)
					assert.NotEqual(t, uuid.Nil, preparedRecipeMeasure.Id)
					assert.Equal(t, testCase.RecipeMeasure.UserId, preparedRecipeMeasure.UserId)
					assert.Equal(t, testCase.RecipeMeasure.EntityId, preparedRecipeMeasure.EntityId)
					assert.Equal(t, testCase.RecipeMeasure.UnitId, preparedRecipeMeasure.UnitId)
					assert.Equal(t, testCase.RecipeMeasure.DateInsert, preparedRecipeMeasure.DateInsert)
					assert.Equal(t, testCase.RecipeMeasure.DateUpdate, preparedRecipeMeasure.DateUpdate)
					assert.Equal(t, testCase.RecipeMeasure.Status, preparedRecipeMeasure.Status)
				}
			},
		)
	}
}

func TestPrepareUnitRepositoryInsert(t *testing.T) {
	tests := []struct {
		Name        string
		Unit        *entity.Unit
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name: "Test case with PrepareUnitRepositoryInsert and correct data",
			Unit: &entity.Unit{
				DateInsert: time.Now().UTC(),
				DateUpdate: time.Now().UTC(),
				Name:       "Test case with correct data",
				Status:     kind.UnitStatusUnPublished,
			},
			MustBePanic: false,
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				defer func() {
					if testCase.MustBePanic {
						assert.NotNil(t, recover())
					} else {
						assert.Nil(t, recover())
					}
				}()

				preparedUnit := prepareUnitRepositoryInsert(testCase.Unit)

				if testCase.MustBeFault {
					assert.Nil(t, preparedUnit)
				} else {
					assert.NotNil(t, preparedUnit)
					assert.NotEqual(t, uuid.Nil, preparedUnit.Id)
					assert.Equal(t, testCase.Unit.DateInsert, preparedUnit.DateInsert)
					assert.Equal(t, testCase.Unit.DateUpdate, preparedUnit.DateUpdate)
					assert.Equal(t, testCase.Unit.Name, preparedUnit.Name)
					assert.Equal(t, testCase.Unit.Status, preparedUnit.Status)
				}
			},
		)
	}
}

func TestPrepareCategoryRepositoryInsert(t *testing.T) {
	tests := []struct {
		Name        string
		Category    *entity.Category
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name: "Test case with PrepareCategoryRepositoryInsert and correct data",
			Category: &entity.Category{
				UserId:     uuid.New(),
				DateInsert: time.Now().UTC(),
				DateUpdate: time.Now().UTC(),
				Name:       "Test case with correct data",
				Status:     kind.CategoryStatusUnPublished,
			},
			MustBePanic: false,
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				defer func() {
					if testCase.MustBePanic {
						assert.NotNil(t, recover())
					} else {
						assert.Nil(t, recover())
					}
				}()

				preparedCategory := prepareCategoryRepositoryInsert(testCase.Category)

				if testCase.MustBeFault {
					assert.Nil(t, preparedCategory)
				} else {
					assert.NotNil(t, preparedCategory)
					assert.NotEqual(t, uuid.Nil, preparedCategory.Id)
					assert.Equal(t, testCase.Category.UserId, preparedCategory.UserId)
					assert.Equal(t, testCase.Category.DateInsert, preparedCategory.DateInsert)
					assert.Equal(t, testCase.Category.DateUpdate, preparedCategory.DateUpdate)
					assert.Equal(t, testCase.Category.Name, preparedCategory.Name)
					assert.Equal(t, testCase.Category.Status, preparedCategory.Status)
				}
			},
		)
	}
}

func TestPrepareIngredientRepositoryInsert(t *testing.T) {
	tests := []struct {
		Name        string
		Ingredient  *entity.Ingredient
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name: "Test case with PrepareIngredientRepositoryInsert and correct data",
			Ingredient: &entity.Ingredient{
				UserId:     uuid.New(),
				DateInsert: time.Now().UTC(),
				DateUpdate: time.Now().UTC(),
				Name:       "Test case with correct data",
				Status:     kind.IngredientStatusUnPublished,
			},
			MustBePanic: false,
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				defer func() {
					if testCase.MustBePanic {
						assert.NotNil(t, recover())
					} else {
						assert.Nil(t, recover())
					}
				}()

				preparedIngredient := prepareIngredientRepositoryInsert(testCase.Ingredient)

				if testCase.MustBeFault {
					assert.Nil(t, preparedIngredient)
				} else {
					assert.NotNil(t, preparedIngredient)
					assert.NotEqual(t, uuid.Nil, preparedIngredient.Id)
					assert.Equal(t, testCase.Ingredient.UserId, preparedIngredient.UserId)
					assert.Equal(t, testCase.Ingredient.DateInsert, preparedIngredient.DateInsert)
					assert.Equal(t, testCase.Ingredient.DateUpdate, preparedIngredient.DateUpdate)
					assert.Equal(t, testCase.Ingredient.Name, preparedIngredient.Name)
					assert.Equal(t, testCase.Ingredient.Status, preparedIngredient.Status)
				}
			},
		)
	}
}

func TestPreparePictureRepositoryInsert(t *testing.T) {
	tests := []struct {
		Name        string
		Picture     *entity.Picture
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name: "Test case with PreparePictureRepositoryInsert and correct data",
			Picture: &entity.Picture{
				UserId:     uuid.New(),
				EntityId:   uuid.New(),
				DateInsert: time.Now().UTC(),
				DateUpdate: time.Now().UTC(),
				Name:       "Test case with correct data",
				URL:        "https://www.google.com/images/branding/googlelogo/2x/googlelogo_light_color_272x92dp.png",
				Width:      1024,
				Height:     768,
				Size:       786432,
				Type:       "image/png",
				Status:     kind.PictureStatusUnPublished,
			},
			MustBePanic: false,
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				defer func() {
					if testCase.MustBePanic {
						assert.NotNil(t, recover())
					} else {
						assert.Nil(t, recover())
					}
				}()

				preparedPicture := preparePictureRepositoryInsert(testCase.Picture)

				if testCase.MustBeFault {
					assert.Nil(t, preparedPicture)
				} else {
					assert.NotNil(t, preparedPicture)
					assert.NotEqual(t, uuid.Nil, preparedPicture.Id)
					assert.Equal(t, testCase.Picture.UserId, preparedPicture.UserId)
					assert.Equal(t, testCase.Picture.EntityId, preparedPicture.EntityId)
					assert.Equal(t, testCase.Picture.DateInsert, preparedPicture.DateInsert)
					assert.Equal(t, testCase.Picture.DateUpdate, preparedPicture.DateUpdate)
					assert.Equal(t, testCase.Picture.Name, preparedPicture.Name)
					assert.Equal(t, testCase.Picture.URL, preparedPicture.URL)
					assert.Equal(t, testCase.Picture.Width, preparedPicture.Width)
					assert.Equal(t, testCase.Picture.Height, preparedPicture.Height)
					assert.Equal(t, testCase.Picture.Size, preparedPicture.Size)
					assert.Equal(t, testCase.Picture.Type, preparedPicture.Type)
					assert.Equal(t, testCase.Picture.Status, preparedPicture.Status)
				}
			},
		)
	}
}

func TestPrepareAltNameRepositoryInsert(t *testing.T) {
	tests := []struct {
		Name        string
		AltName     *entity.AltName
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name: "Test case with PrepareAltNameRepositoryInsert and correct data",
			AltName: &entity.AltName{
				UserId:     uuid.New(),
				EntityId:   uuid.New(),
				DateInsert: time.Now().UTC(),
				DateUpdate: time.Now().UTC(),
				Name:       "Test case with correct data",
				Status:     kind.AltNameStatusUnPublished,
			},
			MustBePanic: false,
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				defer func() {
					if testCase.MustBePanic {
						assert.NotNil(t, recover())
					} else {
						assert.Nil(t, recover())
					}
				}()

				preparedAltName := prepareAltNameRepositoryInsert(testCase.AltName)

				if testCase.MustBeFault {
					assert.Nil(t, preparedAltName)
				} else {
					assert.NotNil(t, preparedAltName)
					assert.NotEqual(t, uuid.Nil, preparedAltName.Id)
					assert.Equal(t, testCase.AltName.UserId, preparedAltName.UserId)
					assert.Equal(t, testCase.AltName.EntityId, preparedAltName.EntityId)
					assert.Equal(t, testCase.AltName.DateInsert, preparedAltName.DateInsert)
					assert.Equal(t, testCase.AltName.DateUpdate, preparedAltName.DateUpdate)
					assert.Equal(t, testCase.AltName.Name, preparedAltName.Name)
					assert.Equal(t, testCase.AltName.Status, preparedAltName.Status)
				}
			},
		)
	}
}

func TestPreparePlannerRepositoryInsert(t *testing.T) {
	tests := []struct {
		Name        string
		Planner     *entity.Planner
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name: "Test case with PreparePlannerRepositoryInsert and correct data",
			Planner: &entity.Planner{
				UserId:     uuid.New(),
				DateInsert: time.Now().UTC(),
				DateUpdate: time.Now().UTC(),
				StartTime:  time.Now().UTC(),
				EndTime:    time.Now().UTC(),
				Name:       "Test case with correct data",
				Status:     kind.PlannerStatusInActive,
			},
			MustBePanic: false,
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				defer func() {
					if testCase.MustBePanic {
						assert.NotNil(t, recover())
					} else {
						assert.Nil(t, recover())
					}
				}()

				preparedPlanner := preparePlannerRepositoryInsert(testCase.Planner)

				if testCase.MustBeFault {
					assert.Nil(t, preparedPlanner)
				} else {
					assert.NotNil(t, preparedPlanner)
					assert.NotEqual(t, uuid.Nil, preparedPlanner.Id)
					assert.Equal(t, testCase.Planner.UserId, preparedPlanner.UserId)
					assert.Equal(t, testCase.Planner.DateInsert, preparedPlanner.DateInsert)
					assert.Equal(t, testCase.Planner.DateUpdate, preparedPlanner.DateUpdate)
					assert.Equal(t, testCase.Planner.StartTime, preparedPlanner.StartTime)
					assert.Equal(t, testCase.Planner.EndTime, preparedPlanner.EndTime)
					assert.Equal(t, testCase.Planner.Name, preparedPlanner.Name)
					assert.Equal(t, testCase.Planner.Status, preparedPlanner.Status)
				}
			},
		)
	}
}

func TestPreparePlannerIntervalRepositoryInsert(t *testing.T) {
	tests := []struct {
		Name            string
		PlannerInterval *entity.PlannerInterval
		MustBePanic     bool
		MustBeFault     bool
	}{
		{
			Name: "Test case with PreparePlannerIntervalRepositoryInsert and correct data",
			PlannerInterval: &entity.PlannerInterval{
				UserId:     uuid.New(),
				EntityId:   uuid.New(),
				DateInsert: time.Now().UTC(),
				DateUpdate: time.Now().UTC(),
				StartTime:  time.Now().UTC(),
				EndTime:    time.Now().UTC(),
				Name:       "Test case with correct data",
				Status:     kind.PlannerIntervalStatusInActive,
			},
			MustBePanic: false,
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				defer func() {
					if testCase.MustBePanic {
						assert.NotNil(t, recover())
					} else {
						assert.Nil(t, recover())
					}
				}()

				preparedPlannerInterval := preparePlannerIntervalRepositoryInsert(testCase.PlannerInterval)

				if testCase.MustBeFault {
					assert.Nil(t, preparedPlannerInterval)
				} else {
					assert.NotNil(t, preparedPlannerInterval)
					assert.NotEqual(t, uuid.Nil, preparedPlannerInterval.Id)
					assert.Equal(t, testCase.PlannerInterval.UserId, preparedPlannerInterval.UserId)
					assert.Equal(t, testCase.PlannerInterval.EntityId, preparedPlannerInterval.EntityId)
					assert.Equal(t, testCase.PlannerInterval.DateInsert, preparedPlannerInterval.DateInsert)
					assert.Equal(t, testCase.PlannerInterval.DateUpdate, preparedPlannerInterval.DateUpdate)
					assert.Equal(t, testCase.PlannerInterval.StartTime, preparedPlannerInterval.StartTime)
					assert.Equal(t, testCase.PlannerInterval.EndTime, preparedPlannerInterval.EndTime)
					assert.Equal(t, testCase.PlannerInterval.Name, preparedPlannerInterval.Name)
					assert.Equal(t, testCase.PlannerInterval.Status, preparedPlannerInterval.Status)
				}
			},
		)
	}
}

func TestPreparePlannerRecipeRepositoryInsert(t *testing.T) {
	tests := []struct {
		Name          string
		PlannerRecipe *entity.PlannerRecipe
		MustBePanic   bool
		MustBeFault   bool
	}{
		{
			Name: "Test case with PreparePlannerRecipeRepositoryInsert and correct data",
			PlannerRecipe: &entity.PlannerRecipe{
				UserId:     uuid.New(),
				EntityId:   uuid.New(),
				RecipeId:   uuid.New(),
				DateInsert: time.Now().UTC(),
				DateUpdate: time.Now().UTC(),
				Status:     kind.PlannerRecipeStatusInActive,
			},
			MustBePanic: false,
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				defer func() {
					if testCase.MustBePanic {
						assert.NotNil(t, recover())
					} else {
						assert.Nil(t, recover())
					}
				}()

				preparedPlannerRecipe := preparePlannerRecipeRepositoryInsert(testCase.PlannerRecipe)

				if testCase.MustBeFault {
					assert.Nil(t, preparedPlannerRecipe)
				} else {
					assert.NotNil(t, preparedPlannerRecipe)
					assert.NotEqual(t, uuid.Nil, preparedPlannerRecipe.Id)
					assert.Equal(t, testCase.PlannerRecipe.UserId, preparedPlannerRecipe.UserId)
					assert.Equal(t, testCase.PlannerRecipe.RecipeId, preparedPlannerRecipe.RecipeId)
					assert.Equal(t, testCase.PlannerRecipe.EntityId, preparedPlannerRecipe.EntityId)
					assert.Equal(t, testCase.PlannerRecipe.DateInsert, preparedPlannerRecipe.DateInsert)
					assert.Equal(t, testCase.PlannerRecipe.DateUpdate, preparedPlannerRecipe.DateUpdate)
					assert.Equal(t, testCase.PlannerRecipe.Status, preparedPlannerRecipe.Status)
				}
			},
		)
	}
}
