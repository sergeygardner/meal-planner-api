package aggregate

import (
	"encoding/json"
	"github.com/google/uuid"
	DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
	"time"
)

type planner struct {
	Id         uuid.UUID
	UserId     uuid.UUID
	DateInsert time.Time
	DateUpdate time.Time
	StartTime  time.Time
	EndTime    time.Time
	Name       string
	Status     kind.PlannerStatus
}

type plannerInterval struct {
	Id         uuid.UUID
	UserId     uuid.UUID
	EntityId   uuid.UUID
	DateInsert time.Time
	DateUpdate time.Time
	StartTime  time.Time
	EndTime    time.Time
	Name       string
	Status     kind.PlannerIntervalStatus
}

type plannerRecipe struct {
	Id         uuid.UUID
	UserId     uuid.UUID
	EntityId   uuid.UUID
	RecipeId   uuid.UUID
	DateInsert time.Time
	DateUpdate time.Time
	Status     kind.PlannerRecipeStatus
}

func TestPlanner(t *testing.T) {
	tests := []struct {
		name      string
		json      string
		Entity    planner
		Intervals []struct {
			Entity  plannerInterval
			Recipes []struct {
				Entity plannerRecipe
				Recipe testRecipeAggregate
			}
		}
	}{
		{
			name: "Test case with active planner properties",
			json: "{\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000100\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"start_time\":\"2000-01-11T00:00:00Z\",\"end_time\":\"2000-01-17T23:59:59Z\",\"name\":\"Planner\",\"status\":\"active\"},\"intervals\":[{\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000003\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000100\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"start_time\":\"2000-01-11T00:00:00Z\",\"end_time\":\"2000-01-17T23:59:59Z\",\"name\":\"PlannerInterval\",\"status\":\"active\"},\"recipes\":[{\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000003\",\"recipe_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"status\":\"active\"},\"recipe\":{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000005\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"categories\":[{\"derive\":{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000006\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000007\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000007\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Category\",\"status\":\"published\"},\"pictures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000008\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000009\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000009\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000007\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"published\"}}]},\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000010\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"derive_id\":\"00000000-0000-0000-0000-000000000006\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"status\":\"published\"}}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000004\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Recipe\",\"description\":\"Description\",\"notes\":\"Notes\",\"status\":\"published\"},\"ingredients\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000011\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000013\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"derive\":{\"id\":\"00000000-0000-0000-0000-000000000012\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Ingredient\",\"status\":\"published\"},\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000013\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"derive_id\":\"00000000-0000-0000-0000-000000000012\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"RecipeIngredient\",\"status\":\"published\"},\"measures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000014\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000015\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000015\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000013\",\"unit_id\":\"00000000-0000-0000-0000-000000000016\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"value\":42,\"status\":\"published\"},\"unit\":{\"id\":\"00000000-0000-0000-0000-000000000016\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Unit\",\"status\":\"published\"}}],\"pictures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000017\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000018\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000018\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000013\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"published\"}}]}],\"processes\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000019\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000020\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000020\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"RecipeProcess\",\"description\":\"Description\",\"notes\":\"Notes\",\"status\":\"published\"},\"pictures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000021\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000022\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000022\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000020\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"published\"}}]}],\"pictures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000023\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000024\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000024\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"published\"}}]}}]}]}\n",
			Entity: planner{
				Id:         uuid.MustParse("00000000-0000-0000-0000-000000000100"),
				UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
				DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
				StartTime:  time.Date(2000, time.January, 11, 0, 0, 0, 0, time.UTC),
				EndTime:    time.Date(2000, time.January, 17, 23, 59, 59, 0, time.UTC),
				Name:       "Planner",
				Status:     kind.PlannerStatusActive,
			},
			Intervals: []struct {
				Entity  plannerInterval
				Recipes []struct {
					Entity plannerRecipe
					Recipe testRecipeAggregate
				}
			}{
				{
					Entity: plannerInterval{
						Id:         uuid.MustParse("00000000-0000-0000-0000-000000000003"),
						UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
						EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000100"),
						DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
						DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
						StartTime:  time.Date(2000, time.January, 11, 0, 0, 0, 0, time.UTC),
						EndTime:    time.Date(2000, time.January, 17, 23, 59, 59, 0, time.UTC),
						Name:       "PlannerInterval",
						Status:     kind.PlannerIntervalStatusActive,
					},
					Recipes: []struct {
						Entity plannerRecipe
						Recipe testRecipeAggregate
					}{
						{
							Entity: plannerRecipe{
								Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
								UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
								EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000003"),
								RecipeId:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
								DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
								DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
								Status:     kind.PlannerRecipeStatusActive,
							},
							Recipe: struct {
								AltNames   []testAltName
								Categories []struct {
									Derive struct {
										AltNames []testAltName
										Entity   testCategory
										Pictures []testPicture
									}
									Entity testRecipeCategory
								}
								Entity      testRecipe
								Ingredients []struct {
									AltNames []testAltName
									Derive   testIngredient
									Entity   testRecipeIngredient
									Measures []struct {
										AltNames []testAltName
										Entity   testRecipeMeasure
										Unit     testUnit
									}
									Pictures []testPicture
								}
								Processes []struct {
									AltNames []testAltName
									Entity   testRecipeProcess
									Pictures []testPicture
								}
								Pictures []testPicture
							}{
								AltNames: []testAltName{
									{
										Id:         uuid.MustParse("00000000-0000-0000-0000-000000000005"),
										UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
										EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
										DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
										DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
										Name:       "AltName",
										Status:     kind.AltNameStatusPublished,
									},
								},
								Categories: []struct {
									Derive struct {
										AltNames []testAltName
										Entity   testCategory
										Pictures []testPicture
									}
									Entity testRecipeCategory
								}{
									{
										Derive: struct {
											AltNames []testAltName
											Entity   testCategory
											Pictures []testPicture
										}{
											AltNames: []testAltName{
												{
													Id:         uuid.MustParse("00000000-0000-0000-0000-000000000006"),
													UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
													EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000007"),
													DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
													DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
													Name:       "AltName",
													Status:     kind.AltNameStatusPublished,
												},
											},
											Entity: testCategory{
												Id:         uuid.MustParse("00000000-0000-0000-0000-000000000007"),
												UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
												DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).UTC(),
												DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC).UTC(),
												Name:       "Category",
												Status:     kind.CategoryStatusPublished,
											},
											Pictures: []testPicture{
												{
													AltNames: []testAltName{
														{
															Id:         uuid.MustParse("00000000-0000-0000-0000-000000000008"),
															UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
															EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000009"),
															DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
															DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
															Name:       "AltName",
															Status:     kind.AltNameStatusPublished,
														},
													},
													Entity: struct {
														Id         uuid.UUID
														UserId     uuid.UUID
														EntityId   uuid.UUID
														DateInsert time.Time
														DateUpdate time.Time
														Name       string
														URL        string
														Width      int64
														Height     int64
														Size       int64
														Type       string
														Status     kind.PictureStatus
													}{
														Id:         uuid.MustParse("00000000-0000-0000-0000-000000000009"),
														UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
														EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000007"),
														DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
														DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
														Name:       "Picture",
														URL:        "https://google.com/doodle.png",
														Width:      512,
														Height:     512,
														Size:       1024,
														Type:       "image/png",
														Status:     kind.PictureStatusPublished,
													},
												},
											},
										},
										Entity: testRecipeCategory{
											Id:         uuid.MustParse("00000000-0000-0000-0000-000000000010"),
											UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
											EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
											DeriveId:   uuid.MustParse("00000000-0000-0000-0000-000000000006"),
											DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).UTC(),
											DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC).UTC(),
											Status:     kind.RecipeCategoryStatusPublished,
										},
									},
								},
								Entity: testRecipe{
									Id:          uuid.MustParse("00000000-0000-0000-0000-000000000004"),
									UserId:      uuid.MustParse("00000000-0000-0000-0000-000000000002"),
									DateInsert:  time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).UTC(),
									DateUpdate:  time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC).UTC(),
									Name:        "Recipe",
									Description: "Description",
									Notes:       "Notes",
									Status:      kind.RecipeStatusPublished,
								},
								Ingredients: []struct {
									AltNames []testAltName
									Derive   testIngredient
									Entity   testRecipeIngredient
									Measures []struct {
										AltNames []testAltName
										Entity   testRecipeMeasure
										Unit     testUnit
									}
									Pictures []testPicture
								}{
									{
										AltNames: []testAltName{
											{
												Id:         uuid.MustParse("00000000-0000-0000-0000-000000000011"),
												UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
												EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000013"),
												DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
												DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
												Name:       "AltName",
												Status:     kind.AltNameStatusPublished,
											},
										},
										Derive: struct {
											Id         uuid.UUID
											UserId     uuid.UUID
											DateInsert time.Time
											DateUpdate time.Time
											Name       string
											Status     kind.IngredientStatus
										}{
											Id:         uuid.MustParse("00000000-0000-0000-0000-000000000012"),
											UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
											DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).UTC(),
											DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC).UTC(),
											Name:       "Ingredient",
											Status:     kind.IngredientStatusPublished,
										},
										Entity: struct {
											Id         uuid.UUID
											UserId     uuid.UUID
											EntityId   uuid.UUID
											DeriveId   uuid.UUID
											DateInsert time.Time
											DateUpdate time.Time
											Name       string
											Status     kind.RecipeIngredientStatus
										}{
											Id:         uuid.MustParse("00000000-0000-0000-0000-000000000013"),
											UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
											EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
											DeriveId:   uuid.MustParse("00000000-0000-0000-0000-000000000012"),
											DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).UTC(),
											DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC).UTC(),
											Name:       "RecipeIngredient",
											Status:     kind.RecipeIngredientStatusPublished,
										},
										Measures: []struct {
											AltNames []testAltName
											Entity   testRecipeMeasure
											Unit     testUnit
										}{
											{
												AltNames: []testAltName{
													{
														Id:         uuid.MustParse("00000000-0000-0000-0000-000000000014"),
														UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
														EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000015"),
														DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
														DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
														Name:       "AltName",
														Status:     kind.AltNameStatusPublished,
													},
												},
												Entity: testRecipeMeasure{
													Id:         uuid.MustParse("00000000-0000-0000-0000-000000000015"),
													UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
													EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000013"),
													UnitId:     uuid.MustParse("00000000-0000-0000-0000-000000000016"),
													DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).UTC(),
													DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC).UTC(),
													Value:      42,
													Status:     kind.RecipeMeasureStatusPublished,
												},
												Unit: testUnit{
													Id:         uuid.MustParse("00000000-0000-0000-0000-000000000016"),
													DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
													DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
													Name:       "Unit",
													Status:     kind.UnitStatusPublished,
												},
											},
										},
										Pictures: []testPicture{
											{
												AltNames: []testAltName{
													{
														Id:         uuid.MustParse("00000000-0000-0000-0000-000000000017"),
														UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
														EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000018"),
														DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
														DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
														Name:       "AltName",
														Status:     kind.AltNameStatusPublished,
													},
												},
												Entity: struct {
													Id         uuid.UUID
													UserId     uuid.UUID
													EntityId   uuid.UUID
													DateInsert time.Time
													DateUpdate time.Time
													Name       string
													URL        string
													Width      int64
													Height     int64
													Size       int64
													Type       string
													Status     kind.PictureStatus
												}{
													Id:         uuid.MustParse("00000000-0000-0000-0000-000000000018"),
													UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
													EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000013"),
													DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
													DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
													Name:       "Picture",
													URL:        "https://google.com/doodle.png",
													Width:      512,
													Height:     512,
													Size:       1024,
													Type:       "image/png",
													Status:     kind.PictureStatusPublished,
												},
											},
										},
									},
								},
								Processes: []struct {
									AltNames []testAltName
									Entity   testRecipeProcess
									Pictures []testPicture
								}{
									{
										AltNames: []testAltName{
											{
												Id:         uuid.MustParse("00000000-0000-0000-0000-000000000019"),
												UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
												EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000020"),
												DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
												DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
												Name:       "AltName",
												Status:     kind.AltNameStatusPublished,
											},
										},
										Entity: testRecipeProcess{
											Id:          uuid.MustParse("00000000-0000-0000-0000-000000000020"),
											UserId:      uuid.MustParse("00000000-0000-0000-0000-000000000002"),
											EntityId:    uuid.MustParse("00000000-0000-0000-0000-000000000004"),
											DateInsert:  time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).UTC(),
											DateUpdate:  time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC).UTC(),
											Name:        "RecipeProcess",
											Description: "Description",
											Notes:       "Notes",
											Status:      kind.RecipeProcessStatusPublished,
										},
										Pictures: []testPicture{
											{
												AltNames: []testAltName{
													{
														Id:         uuid.MustParse("00000000-0000-0000-0000-000000000021"),
														UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
														EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000022"),
														DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
														DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
														Name:       "AltName",
														Status:     kind.AltNameStatusPublished,
													},
												},
												Entity: struct {
													Id         uuid.UUID
													UserId     uuid.UUID
													EntityId   uuid.UUID
													DateInsert time.Time
													DateUpdate time.Time
													Name       string
													URL        string
													Width      int64
													Height     int64
													Size       int64
													Type       string
													Status     kind.PictureStatus
												}{
													Id:         uuid.MustParse("00000000-0000-0000-0000-000000000022"),
													UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
													EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000020"),
													DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
													DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
													Name:       "Picture",
													URL:        "https://google.com/doodle.png",
													Width:      512,
													Height:     512,
													Size:       1024,
													Type:       "image/png",
													Status:     kind.PictureStatusPublished,
												},
											},
										},
									},
								},
								Pictures: []testPicture{
									{
										AltNames: []testAltName{
											{
												Id:         uuid.MustParse("00000000-0000-0000-0000-000000000023"),
												UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
												EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000024"),
												DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
												DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
												Name:       "AltName",
												Status:     kind.AltNameStatusPublished,
											},
										},
										Entity: struct {
											Id         uuid.UUID
											UserId     uuid.UUID
											EntityId   uuid.UUID
											DateInsert time.Time
											DateUpdate time.Time
											Name       string
											URL        string
											Width      int64
											Height     int64
											Size       int64
											Type       string
											Status     kind.PictureStatus
										}{
											Id:         uuid.MustParse("00000000-0000-0000-0000-000000000024"),
											UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
											EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
											DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
											DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
											Name:       "Picture",
											URL:        "https://google.com/doodle.png",
											Width:      512,
											Height:     512,
											Size:       1024,
											Type:       "image/png",
											Status:     kind.PictureStatusPublished,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				plannerAggregate := Planner{
					Entity: &DomainEntity.Planner{
						Id:         testCase.Entity.Id,
						UserId:     testCase.Entity.UserId,
						DateInsert: testCase.Entity.DateInsert,
						DateUpdate: testCase.Entity.DateUpdate,
						StartTime:  testCase.Entity.StartTime,
						EndTime:    testCase.Entity.EndTime,
						Name:       testCase.Entity.Name,
						Status:     testCase.Entity.Status,
					},
					Intervals: []*PlannerInterval{{
						Entity: &DomainEntity.PlannerInterval{
							Id:         testCase.Intervals[0].Entity.Id,
							UserId:     testCase.Intervals[0].Entity.UserId,
							EntityId:   testCase.Intervals[0].Entity.EntityId,
							DateInsert: testCase.Intervals[0].Entity.DateInsert,
							DateUpdate: testCase.Intervals[0].Entity.DateUpdate,
							StartTime:  testCase.Intervals[0].Entity.StartTime,
							EndTime:    testCase.Intervals[0].Entity.EndTime,
							Name:       testCase.Intervals[0].Entity.Name,
							Status:     testCase.Intervals[0].Entity.Status,
						},
						Recipes: []*PlannerRecipe{
							{
								Entity: &DomainEntity.PlannerRecipe{
									Id:         testCase.Intervals[0].Recipes[0].Entity.Id,
									UserId:     testCase.Intervals[0].Recipes[0].Entity.UserId,
									EntityId:   testCase.Intervals[0].Recipes[0].Entity.EntityId,
									RecipeId:   testCase.Intervals[0].Recipes[0].Entity.RecipeId,
									DateInsert: testCase.Intervals[0].Recipes[0].Entity.DateInsert,
									DateUpdate: testCase.Intervals[0].Recipes[0].Entity.DateUpdate,
									Status:     testCase.Intervals[0].Recipes[0].Entity.Status,
								},
								Recipe: &Recipe{
									AltNames: []*DomainEntity.AltName{
										{
											Id:         testCase.Intervals[0].Recipes[0].Recipe.AltNames[0].Id,
											UserId:     testCase.Intervals[0].Recipes[0].Recipe.AltNames[0].UserId,
											EntityId:   testCase.Intervals[0].Recipes[0].Recipe.AltNames[0].EntityId,
											DateInsert: testCase.Intervals[0].Recipes[0].Recipe.AltNames[0].DateInsert,
											DateUpdate: testCase.Intervals[0].Recipes[0].Recipe.AltNames[0].DateUpdate,
											Name:       testCase.Intervals[0].Recipes[0].Recipe.AltNames[0].Name,
											Status:     testCase.Intervals[0].Recipes[0].Recipe.AltNames[0].Status,
										},
									},
									Categories: []*RecipeCategory{
										{
											Derive: &Category{
												AltNames: []*DomainEntity.AltName{
													{
														Id:         testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.AltNames[0].Id,
														UserId:     testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.AltNames[0].UserId,
														EntityId:   testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.AltNames[0].EntityId,
														DateInsert: testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.AltNames[0].DateInsert,
														DateUpdate: testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.AltNames[0].DateUpdate,
														Name:       testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.AltNames[0].Name,
														Status:     testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.AltNames[0].Status,
													},
												},
												Entity: &DomainEntity.Category{
													Id:         testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Entity.Id,
													UserId:     testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Entity.UserId,
													DateInsert: testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Entity.DateInsert,
													DateUpdate: testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Entity.DateUpdate,
													Name:       testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Entity.Name,
													Status:     testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Entity.Status,
												},
												Pictures: []*Picture{
													{
														AltNames: []*DomainEntity.AltName{
															{
																Id:         testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Id,
																UserId:     testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].UserId,
																EntityId:   testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].EntityId,
																DateInsert: testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].DateInsert,
																DateUpdate: testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].DateUpdate,
																Name:       testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Name,
																Status:     testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Status,
															},
														},
														Entity: &DomainEntity.Picture{
															Id:         testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Id,
															UserId:     testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.UserId,
															EntityId:   testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.EntityId,
															DateInsert: testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.DateInsert,
															DateUpdate: testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.DateUpdate,
															Name:       testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Name,
															URL:        testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.URL,
															Width:      testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Width,
															Height:     testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Height,
															Size:       testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Size,
															Type:       testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Type,
															Status:     testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Status,
														},
													},
												},
											},
											Entity: &DomainEntity.RecipeCategory{
												Id:         testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Entity.Id,
												UserId:     testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Entity.UserId,
												EntityId:   testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Entity.EntityId,
												DeriveId:   testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Entity.DeriveId,
												DateInsert: testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Entity.DateInsert,
												DateUpdate: testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Entity.DateUpdate,
												Status:     testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Entity.Status,
											},
										},
									},
									Entity: &DomainEntity.Recipe{
										Id:          testCase.Intervals[0].Recipes[0].Recipe.Entity.Id,
										UserId:      testCase.Intervals[0].Recipes[0].Recipe.Entity.UserId,
										DateInsert:  testCase.Intervals[0].Recipes[0].Recipe.Entity.DateInsert,
										DateUpdate:  testCase.Intervals[0].Recipes[0].Recipe.Entity.DateUpdate,
										Name:        testCase.Intervals[0].Recipes[0].Recipe.Entity.Name,
										Description: testCase.Intervals[0].Recipes[0].Recipe.Entity.Description,
										Notes:       testCase.Intervals[0].Recipes[0].Recipe.Entity.Notes,
										Status:      testCase.Intervals[0].Recipes[0].Recipe.Entity.Status,
									},
									Ingredients: []*RecipeIngredient{
										{
											AltNames: []*DomainEntity.AltName{
												{
													Id:         testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].AltNames[0].Id,
													UserId:     testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].AltNames[0].UserId,
													EntityId:   testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].AltNames[0].EntityId,
													DateInsert: testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].AltNames[0].DateInsert,
													DateUpdate: testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].AltNames[0].DateUpdate,
													Name:       testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].AltNames[0].Name,
													Status:     testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].AltNames[0].Status,
												},
											},
											Derive: &DomainEntity.Ingredient{
												Id:         testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Derive.Id,
												UserId:     testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Derive.UserId,
												DateInsert: testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Derive.DateInsert,
												DateUpdate: testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Derive.DateUpdate,
												Name:       testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Derive.Name,
												Status:     testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Derive.Status,
											},
											Entity: &DomainEntity.RecipeIngredient{
												Id:         testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Entity.Id,
												UserId:     testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Entity.UserId,
												EntityId:   testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Entity.EntityId,
												DeriveId:   testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Entity.DeriveId,
												DateInsert: testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Entity.DateInsert,
												DateUpdate: testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Entity.DateUpdate,
												Name:       testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Entity.Name,
												Status:     testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Entity.Status,
											},
											Measures: []*RecipeMeasure{
												{
													AltNames: []*DomainEntity.AltName{
														{
															Id:         testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].Id,
															UserId:     testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].UserId,
															EntityId:   testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].EntityId,
															DateInsert: testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].DateInsert,
															DateUpdate: testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].DateUpdate,
															Name:       testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].Name,
															Status:     testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].Status,
														},
													},
													Entity: &DomainEntity.RecipeMeasure{
														Id:         testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.Id,
														UserId:     testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.UserId,
														EntityId:   testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.EntityId,
														UnitId:     testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.UnitId,
														DateInsert: testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.DateInsert,
														DateUpdate: testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.DateUpdate,
														Value:      testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.Value,
														Status:     testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.Status,
													},
													Unit: &DomainEntity.Unit{
														Id:         testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.Id,
														DateInsert: testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.DateInsert,
														DateUpdate: testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.DateUpdate,
														Name:       testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.Name,
														Status:     testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.Status,
													},
												},
											},
											Pictures: []*Picture{
												{
													AltNames: []*DomainEntity.AltName{
														{
															Id:         testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].Id,
															UserId:     testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].UserId,
															EntityId:   testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].EntityId,
															DateInsert: testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].DateInsert,
															DateUpdate: testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].DateUpdate,
															Name:       testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].Name,
															Status:     testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].Status,
														},
													},
													Entity: &DomainEntity.Picture{
														Id:         testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Id,
														UserId:     testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.UserId,
														EntityId:   testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.EntityId,
														DateInsert: testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.DateInsert,
														DateUpdate: testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.DateUpdate,
														Name:       testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Name,
														URL:        testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.URL,
														Width:      testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Width,
														Height:     testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Height,
														Size:       testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Size,
														Type:       testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Type,
														Status:     testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Status,
													},
												},
											},
										},
									},
									Processes: []*RecipeProcess{
										{
											AltNames: []*DomainEntity.AltName{
												{
													Id:         testCase.Intervals[0].Recipes[0].Recipe.Processes[0].AltNames[0].Id,
													UserId:     testCase.Intervals[0].Recipes[0].Recipe.Processes[0].AltNames[0].UserId,
													EntityId:   testCase.Intervals[0].Recipes[0].Recipe.Processes[0].AltNames[0].EntityId,
													DateInsert: testCase.Intervals[0].Recipes[0].Recipe.Processes[0].AltNames[0].DateInsert,
													DateUpdate: testCase.Intervals[0].Recipes[0].Recipe.Processes[0].AltNames[0].DateUpdate,
													Name:       testCase.Intervals[0].Recipes[0].Recipe.Processes[0].AltNames[0].Name,
													Status:     testCase.Intervals[0].Recipes[0].Recipe.Processes[0].AltNames[0].Status,
												},
											},
											Entity: &DomainEntity.RecipeProcess{
												Id:          testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.Id,
												UserId:      testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.UserId,
												EntityId:    testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.EntityId,
												DateInsert:  testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.DateInsert,
												DateUpdate:  testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.DateUpdate,
												Name:        testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.Name,
												Description: testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.Description,
												Notes:       testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.Notes,
												Status:      testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.Status,
											},
											Pictures: []*Picture{
												{
													AltNames: []*DomainEntity.AltName{
														{
															Id:         testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].Id,
															UserId:     testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].UserId,
															EntityId:   testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].EntityId,
															DateInsert: testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].DateInsert,
															DateUpdate: testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].DateUpdate,
															Name:       testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].Name,
															Status:     testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].Status,
														},
													},
													Entity: &DomainEntity.Picture{
														Id:         testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Id,
														UserId:     testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.UserId,
														EntityId:   testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.EntityId,
														DateInsert: testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.DateInsert,
														DateUpdate: testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.DateUpdate,
														Name:       testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Name,
														URL:        testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.URL,
														Width:      testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Width,
														Height:     testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Height,
														Size:       testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Size,
														Type:       testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Type,
														Status:     testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Status,
													},
												},
											},
										},
									},
									Pictures: []*Picture{
										{
											AltNames: []*DomainEntity.AltName{
												{
													Id:         testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].AltNames[0].Id,
													UserId:     testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].AltNames[0].UserId,
													EntityId:   testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].AltNames[0].EntityId,
													DateInsert: testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].AltNames[0].DateInsert,
													DateUpdate: testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].AltNames[0].DateUpdate,
													Name:       testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].AltNames[0].Name,
													Status:     testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].AltNames[0].Status,
												},
											},
											Entity: &DomainEntity.Picture{
												Id:         testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.Id,
												UserId:     testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.UserId,
												EntityId:   testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.EntityId,
												DateInsert: testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.DateInsert,
												DateUpdate: testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.DateUpdate,
												Name:       testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.Name,
												URL:        testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.URL,
												Width:      testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.Width,
												Height:     testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.Height,
												Size:       testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.Size,
												Type:       testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.Type,
												Status:     testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.Status,
											},
										},
									},
								},
							},
						},
					},
					},
				}

				assert.Equal(t, testCase.Entity.Id, plannerAggregate.Entity.Id)
				assert.Equal(t, testCase.Entity.UserId, plannerAggregate.Entity.UserId)
				assert.Equal(t, testCase.Entity.DateInsert, plannerAggregate.Entity.DateInsert)
				assert.Equal(t, testCase.Entity.DateUpdate, plannerAggregate.Entity.DateUpdate)
				assert.Equal(t, testCase.Entity.StartTime, plannerAggregate.Entity.StartTime)
				assert.Equal(t, testCase.Entity.EndTime, plannerAggregate.Entity.EndTime)
				assert.Equal(t, testCase.Entity.Name, plannerAggregate.Entity.Name)
				assert.Equal(t, testCase.Entity.Status, plannerAggregate.Entity.Status)
				assert.Equal(t, testCase.Intervals[0].Entity.Id, plannerAggregate.Intervals[0].Entity.Id)
				assert.Equal(t, testCase.Intervals[0].Entity.UserId, plannerAggregate.Intervals[0].Entity.UserId)
				assert.Equal(t, testCase.Intervals[0].Entity.EntityId, plannerAggregate.Intervals[0].Entity.EntityId)
				assert.Equal(t, testCase.Intervals[0].Entity.DateInsert, plannerAggregate.Intervals[0].Entity.DateInsert)
				assert.Equal(t, testCase.Intervals[0].Entity.DateUpdate, plannerAggregate.Intervals[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Entity.StartTime, plannerAggregate.Intervals[0].Entity.StartTime)
				assert.Equal(t, testCase.Intervals[0].Entity.EndTime, plannerAggregate.Intervals[0].Entity.EndTime)
				assert.Equal(t, testCase.Intervals[0].Entity.Name, plannerAggregate.Intervals[0].Entity.Name)
				assert.Equal(t, testCase.Intervals[0].Entity.Status, plannerAggregate.Intervals[0].Entity.Status)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Entity.Id, plannerAggregate.Intervals[0].Recipes[0].Entity.Id)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Entity.UserId, plannerAggregate.Intervals[0].Recipes[0].Entity.UserId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Entity.EntityId, plannerAggregate.Intervals[0].Recipes[0].Entity.EntityId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Entity.RecipeId, plannerAggregate.Intervals[0].Recipes[0].Entity.RecipeId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Entity.DateInsert, plannerAggregate.Intervals[0].Recipes[0].Entity.DateInsert)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Entity.DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Entity.Status, plannerAggregate.Intervals[0].Recipes[0].Entity.Status)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.AltNames[0].Id, plannerAggregate.Intervals[0].Recipes[0].Recipe.AltNames[0].Id)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.AltNames[0].UserId, plannerAggregate.Intervals[0].Recipes[0].Recipe.AltNames[0].UserId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.AltNames[0].EntityId, plannerAggregate.Intervals[0].Recipes[0].Recipe.AltNames[0].EntityId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.AltNames[0].DateInsert, plannerAggregate.Intervals[0].Recipes[0].Recipe.AltNames[0].DateInsert)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.AltNames[0].DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.AltNames[0].Name, plannerAggregate.Intervals[0].Recipes[0].Recipe.AltNames[0].Name)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.AltNames[0].Status, plannerAggregate.Intervals[0].Recipes[0].Recipe.AltNames[0].Status)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.AltNames[0].Id, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.AltNames[0].Id)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.AltNames[0].UserId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.AltNames[0].UserId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.AltNames[0].EntityId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.AltNames[0].EntityId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.AltNames[0].DateInsert, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.AltNames[0].DateInsert)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.AltNames[0].DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.AltNames[0].Name, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.AltNames[0].Name)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.AltNames[0].Status, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.AltNames[0].Status)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Entity.Id, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Entity.Id)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Entity.UserId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Entity.UserId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Entity.DateInsert, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Entity.DateInsert)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Entity.DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Entity.DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Entity.Name, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Entity.Name)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Entity.Status, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Entity.Status)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Id, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Id)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].UserId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].EntityId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].DateInsert, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Name, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Name)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Status, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Status)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Id, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Id)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.UserId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.UserId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.EntityId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.EntityId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.DateInsert, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.DateInsert)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Name, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Name)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.URL, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.URL)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Width, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Width)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Height, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Height)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Size, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Size)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Type, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Type)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Status, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Status)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Entity.Id, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Entity.Id)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Entity.UserId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Entity.UserId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Entity.EntityId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Entity.EntityId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Entity.DeriveId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Entity.DeriveId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Entity.DateInsert, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Entity.DateInsert)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Entity.DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Categories[0].Entity.Status, plannerAggregate.Intervals[0].Recipes[0].Recipe.Categories[0].Entity.Status)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Entity.Id, plannerAggregate.Intervals[0].Recipes[0].Recipe.Entity.Id)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Entity.UserId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Entity.UserId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Entity.DateInsert, plannerAggregate.Intervals[0].Recipes[0].Recipe.Entity.DateInsert)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Entity.DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.Entity.DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Entity.Name, plannerAggregate.Intervals[0].Recipes[0].Recipe.Entity.Name)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Entity.Description, plannerAggregate.Intervals[0].Recipes[0].Recipe.Entity.Description)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Entity.Notes, plannerAggregate.Intervals[0].Recipes[0].Recipe.Entity.Notes)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Entity.Status, plannerAggregate.Intervals[0].Recipes[0].Recipe.Entity.Status)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].AltNames[0].Id, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].AltNames[0].Id)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].AltNames[0].UserId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].AltNames[0].EntityId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].AltNames[0].DateInsert, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].AltNames[0].DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].AltNames[0].Name, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].AltNames[0].Name)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].AltNames[0].Status, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].AltNames[0].Status)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Derive.Id, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Derive.Id)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Derive.UserId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Derive.UserId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Derive.DateInsert, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Derive.DateInsert)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Derive.DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Derive.DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Derive.Name, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Derive.Name)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Derive.Status, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Derive.Status)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Entity.Id, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Entity.Id)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Entity.UserId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Entity.UserId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Entity.EntityId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Entity.EntityId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Entity.DeriveId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Entity.DeriveId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Entity.DateInsert, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Entity.DateInsert)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Entity.DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Entity.Name, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Entity.Name)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Entity.Status, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Entity.Status)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].Id, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].Id)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].UserId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].EntityId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].DateInsert, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].Name, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].Name)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].Status, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].Status)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.Id, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.Id)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.UserId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.UserId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.EntityId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.EntityId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.DateInsert, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.DateInsert)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.Value, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.Value)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.Status, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.Status)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.Id, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.Id)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.DateInsert, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.DateInsert)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.Name, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.Name)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.Status, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.Status)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].Id, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].Id)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].UserId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].EntityId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].DateInsert, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].Name, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].Name)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].Status, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].Status)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Id, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Id)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.UserId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.UserId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.EntityId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.EntityId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.DateInsert, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.DateInsert)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Name, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Name)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.URL, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.URL)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Width, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Width)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Height, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Height)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Size, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Size)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Type, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Type)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Status, plannerAggregate.Intervals[0].Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Status)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].AltNames[0].Id, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].AltNames[0].Id)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].AltNames[0].UserId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].AltNames[0].EntityId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].AltNames[0].DateInsert, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].AltNames[0].DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].AltNames[0].Name, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].AltNames[0].Name)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].AltNames[0].Status, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].AltNames[0].Status)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.Id, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.Id)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.UserId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.UserId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.EntityId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.EntityId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.DateInsert, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.DateInsert)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.Name, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.Name)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.Description, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.Description)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.Notes, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.Notes)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.Status, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Entity.Status)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].Id, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].Id)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].UserId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].EntityId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].DateInsert, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].Name, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].Name)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].Status, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].Status)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Id, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Id)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.UserId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.UserId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.EntityId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.EntityId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.DateInsert, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.DateInsert)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Name, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Name)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.URL, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.URL)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Width, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Width)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Height, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Height)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Size, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Size)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Type, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Type)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Status, plannerAggregate.Intervals[0].Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Status)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].AltNames[0].Id, plannerAggregate.Intervals[0].Recipes[0].Recipe.Pictures[0].AltNames[0].Id)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].AltNames[0].UserId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Pictures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].AltNames[0].EntityId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Pictures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].AltNames[0].DateInsert, plannerAggregate.Intervals[0].Recipes[0].Recipe.Pictures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].AltNames[0].DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.Pictures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].AltNames[0].Name, plannerAggregate.Intervals[0].Recipes[0].Recipe.Pictures[0].AltNames[0].Name)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].AltNames[0].Status, plannerAggregate.Intervals[0].Recipes[0].Recipe.Pictures[0].AltNames[0].Status)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.Id, plannerAggregate.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.Id)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.UserId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.UserId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.EntityId, plannerAggregate.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.EntityId)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.DateInsert, plannerAggregate.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.DateInsert)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.DateUpdate, plannerAggregate.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.Name, plannerAggregate.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.Name)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.URL, plannerAggregate.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.URL)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.Width, plannerAggregate.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.Width)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.Height, plannerAggregate.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.Height)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.Size, plannerAggregate.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.Size)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.Type, plannerAggregate.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.Type)
				assert.Equal(t, testCase.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.Status, plannerAggregate.Intervals[0].Recipes[0].Recipe.Pictures[0].Entity.Status)

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(plannerAggregate)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())

				reflectPlannerRecipeAggregate := reflect.ValueOf(plannerAggregate)

				for i := 0; i < reflectPlannerRecipeAggregate.NumField(); i++ {
					assert.False(t, reflectPlannerRecipeAggregate.Field(i).IsZero())
				}
			},
		)
	}
}
func TestPlannerInterval(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		Entity  plannerInterval
		Recipes []struct {
			Entity plannerRecipe
			Recipe testRecipeAggregate
		}
	}{
		{
			name: "Test case with active planner interval properties",
			json: "{\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000003\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000100\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"start_time\":\"2000-01-11T00:00:00Z\",\"end_time\":\"2000-01-17T23:59:59Z\",\"name\":\"PlannerInterval\",\"status\":\"active\"},\"recipes\":[{\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000003\",\"recipe_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"status\":\"active\"},\"recipe\":{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000005\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"categories\":[{\"derive\":{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000006\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000007\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000007\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Category\",\"status\":\"published\"},\"pictures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000008\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000009\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000009\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000007\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"published\"}}]},\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000010\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"derive_id\":\"00000000-0000-0000-0000-000000000006\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"status\":\"published\"}}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000004\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Recipe\",\"description\":\"Description\",\"notes\":\"Notes\",\"status\":\"published\"},\"ingredients\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000011\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000013\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"derive\":{\"id\":\"00000000-0000-0000-0000-000000000012\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Ingredient\",\"status\":\"published\"},\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000013\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"derive_id\":\"00000000-0000-0000-0000-000000000012\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"RecipeIngredient\",\"status\":\"published\"},\"measures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000014\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000015\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000015\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000013\",\"unit_id\":\"00000000-0000-0000-0000-000000000016\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"value\":42,\"status\":\"published\"},\"unit\":{\"id\":\"00000000-0000-0000-0000-000000000016\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Unit\",\"status\":\"published\"}}],\"pictures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000017\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000018\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000018\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000013\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"published\"}}]}],\"processes\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000019\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000020\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000020\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"RecipeProcess\",\"description\":\"Description\",\"notes\":\"Notes\",\"status\":\"published\"},\"pictures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000021\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000022\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000022\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000020\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"published\"}}]}],\"pictures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000023\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000024\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000024\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"published\"}}]}}]}\n",
			Entity: plannerInterval{
				Id:         uuid.MustParse("00000000-0000-0000-0000-000000000003"),
				UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000100"),
				DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
				DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
				StartTime:  time.Date(2000, time.January, 11, 0, 0, 0, 0, time.UTC),
				EndTime:    time.Date(2000, time.January, 17, 23, 59, 59, 0, time.UTC),
				Name:       "PlannerInterval",
				Status:     kind.PlannerIntervalStatusActive,
			},
			Recipes: []struct {
				Entity plannerRecipe
				Recipe testRecipeAggregate
			}{
				{
					Entity: plannerRecipe{
						Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
						UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
						EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000003"),
						RecipeId:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
						DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
						DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
						Status:     kind.PlannerRecipeStatusActive,
					},
					Recipe: struct {
						AltNames   []testAltName
						Categories []struct {
							Derive struct {
								AltNames []testAltName
								Entity   testCategory
								Pictures []testPicture
							}
							Entity testRecipeCategory
						}
						Entity      testRecipe
						Ingredients []struct {
							AltNames []testAltName
							Derive   testIngredient
							Entity   testRecipeIngredient
							Measures []struct {
								AltNames []testAltName
								Entity   testRecipeMeasure
								Unit     testUnit
							}
							Pictures []testPicture
						}
						Processes []struct {
							AltNames []testAltName
							Entity   testRecipeProcess
							Pictures []testPicture
						}
						Pictures []testPicture
					}{
						AltNames: []testAltName{
							{
								Id:         uuid.MustParse("00000000-0000-0000-0000-000000000005"),
								UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
								EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
								DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
								DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
								Name:       "AltName",
								Status:     kind.AltNameStatusPublished,
							},
						},
						Categories: []struct {
							Derive struct {
								AltNames []testAltName
								Entity   testCategory
								Pictures []testPicture
							}
							Entity testRecipeCategory
						}{
							{
								Derive: struct {
									AltNames []testAltName
									Entity   testCategory
									Pictures []testPicture
								}{
									AltNames: []testAltName{
										{
											Id:         uuid.MustParse("00000000-0000-0000-0000-000000000006"),
											UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
											EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000007"),
											DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
											DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
											Name:       "AltName",
											Status:     kind.AltNameStatusPublished,
										},
									},
									Entity: testCategory{
										Id:         uuid.MustParse("00000000-0000-0000-0000-000000000007"),
										UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
										DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).UTC(),
										DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC).UTC(),
										Name:       "Category",
										Status:     kind.CategoryStatusPublished,
									},
									Pictures: []testPicture{
										{
											AltNames: []testAltName{
												{
													Id:         uuid.MustParse("00000000-0000-0000-0000-000000000008"),
													UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
													EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000009"),
													DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
													DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
													Name:       "AltName",
													Status:     kind.AltNameStatusPublished,
												},
											},
											Entity: struct {
												Id         uuid.UUID
												UserId     uuid.UUID
												EntityId   uuid.UUID
												DateInsert time.Time
												DateUpdate time.Time
												Name       string
												URL        string
												Width      int64
												Height     int64
												Size       int64
												Type       string
												Status     kind.PictureStatus
											}{
												Id:         uuid.MustParse("00000000-0000-0000-0000-000000000009"),
												UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
												EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000007"),
												DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
												DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
												Name:       "Picture",
												URL:        "https://google.com/doodle.png",
												Width:      512,
												Height:     512,
												Size:       1024,
												Type:       "image/png",
												Status:     kind.PictureStatusPublished,
											},
										},
									},
								},
								Entity: testRecipeCategory{
									Id:         uuid.MustParse("00000000-0000-0000-0000-000000000010"),
									UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
									EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
									DeriveId:   uuid.MustParse("00000000-0000-0000-0000-000000000006"),
									DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).UTC(),
									DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC).UTC(),
									Status:     kind.RecipeCategoryStatusPublished,
								},
							},
						},
						Entity: testRecipe{
							Id:          uuid.MustParse("00000000-0000-0000-0000-000000000004"),
							UserId:      uuid.MustParse("00000000-0000-0000-0000-000000000002"),
							DateInsert:  time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).UTC(),
							DateUpdate:  time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC).UTC(),
							Name:        "Recipe",
							Description: "Description",
							Notes:       "Notes",
							Status:      kind.RecipeStatusPublished,
						},
						Ingredients: []struct {
							AltNames []testAltName
							Derive   testIngredient
							Entity   testRecipeIngredient
							Measures []struct {
								AltNames []testAltName
								Entity   testRecipeMeasure
								Unit     testUnit
							}
							Pictures []testPicture
						}{
							{
								AltNames: []testAltName{
									{
										Id:         uuid.MustParse("00000000-0000-0000-0000-000000000011"),
										UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
										EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000013"),
										DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
										DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
										Name:       "AltName",
										Status:     kind.AltNameStatusPublished,
									},
								},
								Derive: struct {
									Id         uuid.UUID
									UserId     uuid.UUID
									DateInsert time.Time
									DateUpdate time.Time
									Name       string
									Status     kind.IngredientStatus
								}{
									Id:         uuid.MustParse("00000000-0000-0000-0000-000000000012"),
									UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
									DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).UTC(),
									DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC).UTC(),
									Name:       "Ingredient",
									Status:     kind.IngredientStatusPublished,
								},
								Entity: struct {
									Id         uuid.UUID
									UserId     uuid.UUID
									EntityId   uuid.UUID
									DeriveId   uuid.UUID
									DateInsert time.Time
									DateUpdate time.Time
									Name       string
									Status     kind.RecipeIngredientStatus
								}{
									Id:         uuid.MustParse("00000000-0000-0000-0000-000000000013"),
									UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
									EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
									DeriveId:   uuid.MustParse("00000000-0000-0000-0000-000000000012"),
									DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).UTC(),
									DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC).UTC(),
									Name:       "RecipeIngredient",
									Status:     kind.RecipeIngredientStatusPublished,
								},
								Measures: []struct {
									AltNames []testAltName
									Entity   testRecipeMeasure
									Unit     testUnit
								}{
									{
										AltNames: []testAltName{
											{
												Id:         uuid.MustParse("00000000-0000-0000-0000-000000000014"),
												UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
												EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000015"),
												DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
												DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
												Name:       "AltName",
												Status:     kind.AltNameStatusPublished,
											},
										},
										Entity: testRecipeMeasure{
											Id:         uuid.MustParse("00000000-0000-0000-0000-000000000015"),
											UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
											EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000013"),
											UnitId:     uuid.MustParse("00000000-0000-0000-0000-000000000016"),
											DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).UTC(),
											DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC).UTC(),
											Value:      42,
											Status:     kind.RecipeMeasureStatusPublished,
										},
										Unit: testUnit{
											Id:         uuid.MustParse("00000000-0000-0000-0000-000000000016"),
											DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
											DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
											Name:       "Unit",
											Status:     kind.UnitStatusPublished,
										},
									},
								},
								Pictures: []testPicture{
									{
										AltNames: []testAltName{
											{
												Id:         uuid.MustParse("00000000-0000-0000-0000-000000000017"),
												UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
												EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000018"),
												DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
												DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
												Name:       "AltName",
												Status:     kind.AltNameStatusPublished,
											},
										},
										Entity: struct {
											Id         uuid.UUID
											UserId     uuid.UUID
											EntityId   uuid.UUID
											DateInsert time.Time
											DateUpdate time.Time
											Name       string
											URL        string
											Width      int64
											Height     int64
											Size       int64
											Type       string
											Status     kind.PictureStatus
										}{
											Id:         uuid.MustParse("00000000-0000-0000-0000-000000000018"),
											UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
											EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000013"),
											DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
											DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
											Name:       "Picture",
											URL:        "https://google.com/doodle.png",
											Width:      512,
											Height:     512,
											Size:       1024,
											Type:       "image/png",
											Status:     kind.PictureStatusPublished,
										},
									},
								},
							},
						},
						Processes: []struct {
							AltNames []testAltName
							Entity   testRecipeProcess
							Pictures []testPicture
						}{
							{
								AltNames: []testAltName{
									{
										Id:         uuid.MustParse("00000000-0000-0000-0000-000000000019"),
										UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
										EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000020"),
										DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
										DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
										Name:       "AltName",
										Status:     kind.AltNameStatusPublished,
									},
								},
								Entity: testRecipeProcess{
									Id:          uuid.MustParse("00000000-0000-0000-0000-000000000020"),
									UserId:      uuid.MustParse("00000000-0000-0000-0000-000000000002"),
									EntityId:    uuid.MustParse("00000000-0000-0000-0000-000000000004"),
									DateInsert:  time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).UTC(),
									DateUpdate:  time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC).UTC(),
									Name:        "RecipeProcess",
									Description: "Description",
									Notes:       "Notes",
									Status:      kind.RecipeProcessStatusPublished,
								},
								Pictures: []testPicture{
									{
										AltNames: []testAltName{
											{
												Id:         uuid.MustParse("00000000-0000-0000-0000-000000000021"),
												UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
												EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000022"),
												DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
												DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
												Name:       "AltName",
												Status:     kind.AltNameStatusPublished,
											},
										},
										Entity: struct {
											Id         uuid.UUID
											UserId     uuid.UUID
											EntityId   uuid.UUID
											DateInsert time.Time
											DateUpdate time.Time
											Name       string
											URL        string
											Width      int64
											Height     int64
											Size       int64
											Type       string
											Status     kind.PictureStatus
										}{
											Id:         uuid.MustParse("00000000-0000-0000-0000-000000000022"),
											UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
											EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000020"),
											DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
											DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
											Name:       "Picture",
											URL:        "https://google.com/doodle.png",
											Width:      512,
											Height:     512,
											Size:       1024,
											Type:       "image/png",
											Status:     kind.PictureStatusPublished,
										},
									},
								},
							},
						},
						Pictures: []testPicture{
							{
								AltNames: []testAltName{
									{
										Id:         uuid.MustParse("00000000-0000-0000-0000-000000000023"),
										UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
										EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000024"),
										DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
										DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
										Name:       "AltName",
										Status:     kind.AltNameStatusPublished,
									},
								},
								Entity: struct {
									Id         uuid.UUID
									UserId     uuid.UUID
									EntityId   uuid.UUID
									DateInsert time.Time
									DateUpdate time.Time
									Name       string
									URL        string
									Width      int64
									Height     int64
									Size       int64
									Type       string
									Status     kind.PictureStatus
								}{
									Id:         uuid.MustParse("00000000-0000-0000-0000-000000000024"),
									UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
									EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
									DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
									DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
									Name:       "Picture",
									URL:        "https://google.com/doodle.png",
									Width:      512,
									Height:     512,
									Size:       1024,
									Type:       "image/png",
									Status:     kind.PictureStatusPublished,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				plannerIntervalAggregate := PlannerInterval{
					Entity: &DomainEntity.PlannerInterval{
						Id:         testCase.Entity.Id,
						UserId:     testCase.Entity.UserId,
						EntityId:   testCase.Entity.EntityId,
						DateInsert: testCase.Entity.DateInsert,
						DateUpdate: testCase.Entity.DateUpdate,
						StartTime:  testCase.Entity.StartTime,
						EndTime:    testCase.Entity.EndTime,
						Name:       testCase.Entity.Name,
						Status:     testCase.Entity.Status,
					},
					Recipes: []*PlannerRecipe{
						{
							Entity: &DomainEntity.PlannerRecipe{
								Id:         testCase.Recipes[0].Entity.Id,
								UserId:     testCase.Recipes[0].Entity.UserId,
								EntityId:   testCase.Recipes[0].Entity.EntityId,
								RecipeId:   testCase.Recipes[0].Entity.RecipeId,
								DateInsert: testCase.Recipes[0].Entity.DateInsert,
								DateUpdate: testCase.Recipes[0].Entity.DateUpdate,
								Status:     testCase.Recipes[0].Entity.Status,
							},
							Recipe: &Recipe{
								AltNames: []*DomainEntity.AltName{
									{
										Id:         testCase.Recipes[0].Recipe.AltNames[0].Id,
										UserId:     testCase.Recipes[0].Recipe.AltNames[0].UserId,
										EntityId:   testCase.Recipes[0].Recipe.AltNames[0].EntityId,
										DateInsert: testCase.Recipes[0].Recipe.AltNames[0].DateInsert,
										DateUpdate: testCase.Recipes[0].Recipe.AltNames[0].DateUpdate,
										Name:       testCase.Recipes[0].Recipe.AltNames[0].Name,
										Status:     testCase.Recipes[0].Recipe.AltNames[0].Status,
									},
								},
								Categories: []*RecipeCategory{
									{
										Derive: &Category{
											AltNames: []*DomainEntity.AltName{
												{
													Id:         testCase.Recipes[0].Recipe.Categories[0].Derive.AltNames[0].Id,
													UserId:     testCase.Recipes[0].Recipe.Categories[0].Derive.AltNames[0].UserId,
													EntityId:   testCase.Recipes[0].Recipe.Categories[0].Derive.AltNames[0].EntityId,
													DateInsert: testCase.Recipes[0].Recipe.Categories[0].Derive.AltNames[0].DateInsert,
													DateUpdate: testCase.Recipes[0].Recipe.Categories[0].Derive.AltNames[0].DateUpdate,
													Name:       testCase.Recipes[0].Recipe.Categories[0].Derive.AltNames[0].Name,
													Status:     testCase.Recipes[0].Recipe.Categories[0].Derive.AltNames[0].Status,
												},
											},
											Entity: &DomainEntity.Category{
												Id:         testCase.Recipes[0].Recipe.Categories[0].Derive.Entity.Id,
												UserId:     testCase.Recipes[0].Recipe.Categories[0].Derive.Entity.UserId,
												DateInsert: testCase.Recipes[0].Recipe.Categories[0].Derive.Entity.DateInsert,
												DateUpdate: testCase.Recipes[0].Recipe.Categories[0].Derive.Entity.DateUpdate,
												Name:       testCase.Recipes[0].Recipe.Categories[0].Derive.Entity.Name,
												Status:     testCase.Recipes[0].Recipe.Categories[0].Derive.Entity.Status,
											},
											Pictures: []*Picture{
												{
													AltNames: []*DomainEntity.AltName{
														{
															Id:         testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Id,
															UserId:     testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].UserId,
															EntityId:   testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].EntityId,
															DateInsert: testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].DateInsert,
															DateUpdate: testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].DateUpdate,
															Name:       testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Name,
															Status:     testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Status,
														},
													},
													Entity: &DomainEntity.Picture{
														Id:         testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Id,
														UserId:     testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.UserId,
														EntityId:   testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.EntityId,
														DateInsert: testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.DateInsert,
														DateUpdate: testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.DateUpdate,
														Name:       testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Name,
														URL:        testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.URL,
														Width:      testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Width,
														Height:     testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Height,
														Size:       testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Size,
														Type:       testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Type,
														Status:     testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Status,
													},
												},
											},
										},
										Entity: &DomainEntity.RecipeCategory{
											Id:         testCase.Recipes[0].Recipe.Categories[0].Entity.Id,
											UserId:     testCase.Recipes[0].Recipe.Categories[0].Entity.UserId,
											EntityId:   testCase.Recipes[0].Recipe.Categories[0].Entity.EntityId,
											DeriveId:   testCase.Recipes[0].Recipe.Categories[0].Entity.DeriveId,
											DateInsert: testCase.Recipes[0].Recipe.Categories[0].Entity.DateInsert,
											DateUpdate: testCase.Recipes[0].Recipe.Categories[0].Entity.DateUpdate,
											Status:     testCase.Recipes[0].Recipe.Categories[0].Entity.Status,
										},
									},
								},
								Entity: &DomainEntity.Recipe{
									Id:          testCase.Recipes[0].Recipe.Entity.Id,
									UserId:      testCase.Recipes[0].Recipe.Entity.UserId,
									DateInsert:  testCase.Recipes[0].Recipe.Entity.DateInsert,
									DateUpdate:  testCase.Recipes[0].Recipe.Entity.DateUpdate,
									Name:        testCase.Recipes[0].Recipe.Entity.Name,
									Description: testCase.Recipes[0].Recipe.Entity.Description,
									Notes:       testCase.Recipes[0].Recipe.Entity.Notes,
									Status:      testCase.Recipes[0].Recipe.Entity.Status,
								},
								Ingredients: []*RecipeIngredient{
									{
										AltNames: []*DomainEntity.AltName{
											{
												Id:         testCase.Recipes[0].Recipe.Ingredients[0].AltNames[0].Id,
												UserId:     testCase.Recipes[0].Recipe.Ingredients[0].AltNames[0].UserId,
												EntityId:   testCase.Recipes[0].Recipe.Ingredients[0].AltNames[0].EntityId,
												DateInsert: testCase.Recipes[0].Recipe.Ingredients[0].AltNames[0].DateInsert,
												DateUpdate: testCase.Recipes[0].Recipe.Ingredients[0].AltNames[0].DateUpdate,
												Name:       testCase.Recipes[0].Recipe.Ingredients[0].AltNames[0].Name,
												Status:     testCase.Recipes[0].Recipe.Ingredients[0].AltNames[0].Status,
											},
										},
										Derive: &DomainEntity.Ingredient{
											Id:         testCase.Recipes[0].Recipe.Ingredients[0].Derive.Id,
											UserId:     testCase.Recipes[0].Recipe.Ingredients[0].Derive.UserId,
											DateInsert: testCase.Recipes[0].Recipe.Ingredients[0].Derive.DateInsert,
											DateUpdate: testCase.Recipes[0].Recipe.Ingredients[0].Derive.DateUpdate,
											Name:       testCase.Recipes[0].Recipe.Ingredients[0].Derive.Name,
											Status:     testCase.Recipes[0].Recipe.Ingredients[0].Derive.Status,
										},
										Entity: &DomainEntity.RecipeIngredient{
											Id:         testCase.Recipes[0].Recipe.Ingredients[0].Entity.Id,
											UserId:     testCase.Recipes[0].Recipe.Ingredients[0].Entity.UserId,
											EntityId:   testCase.Recipes[0].Recipe.Ingredients[0].Entity.EntityId,
											DeriveId:   testCase.Recipes[0].Recipe.Ingredients[0].Entity.DeriveId,
											DateInsert: testCase.Recipes[0].Recipe.Ingredients[0].Entity.DateInsert,
											DateUpdate: testCase.Recipes[0].Recipe.Ingredients[0].Entity.DateUpdate,
											Name:       testCase.Recipes[0].Recipe.Ingredients[0].Entity.Name,
											Status:     testCase.Recipes[0].Recipe.Ingredients[0].Entity.Status,
										},
										Measures: []*RecipeMeasure{
											{
												AltNames: []*DomainEntity.AltName{
													{
														Id:         testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].Id,
														UserId:     testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].UserId,
														EntityId:   testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].EntityId,
														DateInsert: testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].DateInsert,
														DateUpdate: testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].DateUpdate,
														Name:       testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].Name,
														Status:     testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].Status,
													},
												},
												Entity: &DomainEntity.RecipeMeasure{
													Id:         testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.Id,
													UserId:     testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.UserId,
													EntityId:   testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.EntityId,
													UnitId:     testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.UnitId,
													DateInsert: testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.DateInsert,
													DateUpdate: testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.DateUpdate,
													Value:      testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.Value,
													Status:     testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.Status,
												},
												Unit: &DomainEntity.Unit{
													Id:         testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.Id,
													DateInsert: testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.DateInsert,
													DateUpdate: testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.DateUpdate,
													Name:       testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.Name,
													Status:     testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.Status,
												},
											},
										},
										Pictures: []*Picture{
											{
												AltNames: []*DomainEntity.AltName{
													{
														Id:         testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].Id,
														UserId:     testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].UserId,
														EntityId:   testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].EntityId,
														DateInsert: testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].DateInsert,
														DateUpdate: testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].DateUpdate,
														Name:       testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].Name,
														Status:     testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].Status,
													},
												},
												Entity: &DomainEntity.Picture{
													Id:         testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Id,
													UserId:     testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.UserId,
													EntityId:   testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.EntityId,
													DateInsert: testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.DateInsert,
													DateUpdate: testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.DateUpdate,
													Name:       testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Name,
													URL:        testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.URL,
													Width:      testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Width,
													Height:     testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Height,
													Size:       testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Size,
													Type:       testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Type,
													Status:     testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Status,
												},
											},
										},
									},
								},
								Processes: []*RecipeProcess{
									{
										AltNames: []*DomainEntity.AltName{
											{
												Id:         testCase.Recipes[0].Recipe.Processes[0].AltNames[0].Id,
												UserId:     testCase.Recipes[0].Recipe.Processes[0].AltNames[0].UserId,
												EntityId:   testCase.Recipes[0].Recipe.Processes[0].AltNames[0].EntityId,
												DateInsert: testCase.Recipes[0].Recipe.Processes[0].AltNames[0].DateInsert,
												DateUpdate: testCase.Recipes[0].Recipe.Processes[0].AltNames[0].DateUpdate,
												Name:       testCase.Recipes[0].Recipe.Processes[0].AltNames[0].Name,
												Status:     testCase.Recipes[0].Recipe.Processes[0].AltNames[0].Status,
											},
										},
										Entity: &DomainEntity.RecipeProcess{
											Id:          testCase.Recipes[0].Recipe.Processes[0].Entity.Id,
											UserId:      testCase.Recipes[0].Recipe.Processes[0].Entity.UserId,
											EntityId:    testCase.Recipes[0].Recipe.Processes[0].Entity.EntityId,
											DateInsert:  testCase.Recipes[0].Recipe.Processes[0].Entity.DateInsert,
											DateUpdate:  testCase.Recipes[0].Recipe.Processes[0].Entity.DateUpdate,
											Name:        testCase.Recipes[0].Recipe.Processes[0].Entity.Name,
											Description: testCase.Recipes[0].Recipe.Processes[0].Entity.Description,
											Notes:       testCase.Recipes[0].Recipe.Processes[0].Entity.Notes,
											Status:      testCase.Recipes[0].Recipe.Processes[0].Entity.Status,
										},
										Pictures: []*Picture{
											{
												AltNames: []*DomainEntity.AltName{
													{
														Id:         testCase.Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].Id,
														UserId:     testCase.Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].UserId,
														EntityId:   testCase.Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].EntityId,
														DateInsert: testCase.Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].DateInsert,
														DateUpdate: testCase.Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].DateUpdate,
														Name:       testCase.Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].Name,
														Status:     testCase.Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].Status,
													},
												},
												Entity: &DomainEntity.Picture{
													Id:         testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Id,
													UserId:     testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.UserId,
													EntityId:   testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.EntityId,
													DateInsert: testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.DateInsert,
													DateUpdate: testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.DateUpdate,
													Name:       testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Name,
													URL:        testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.URL,
													Width:      testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Width,
													Height:     testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Height,
													Size:       testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Size,
													Type:       testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Type,
													Status:     testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Status,
												},
											},
										},
									},
								},
								Pictures: []*Picture{
									{
										AltNames: []*DomainEntity.AltName{
											{
												Id:         testCase.Recipes[0].Recipe.Pictures[0].AltNames[0].Id,
												UserId:     testCase.Recipes[0].Recipe.Pictures[0].AltNames[0].UserId,
												EntityId:   testCase.Recipes[0].Recipe.Pictures[0].AltNames[0].EntityId,
												DateInsert: testCase.Recipes[0].Recipe.Pictures[0].AltNames[0].DateInsert,
												DateUpdate: testCase.Recipes[0].Recipe.Pictures[0].AltNames[0].DateUpdate,
												Name:       testCase.Recipes[0].Recipe.Pictures[0].AltNames[0].Name,
												Status:     testCase.Recipes[0].Recipe.Pictures[0].AltNames[0].Status,
											},
										},
										Entity: &DomainEntity.Picture{
											Id:         testCase.Recipes[0].Recipe.Pictures[0].Entity.Id,
											UserId:     testCase.Recipes[0].Recipe.Pictures[0].Entity.UserId,
											EntityId:   testCase.Recipes[0].Recipe.Pictures[0].Entity.EntityId,
											DateInsert: testCase.Recipes[0].Recipe.Pictures[0].Entity.DateInsert,
											DateUpdate: testCase.Recipes[0].Recipe.Pictures[0].Entity.DateUpdate,
											Name:       testCase.Recipes[0].Recipe.Pictures[0].Entity.Name,
											URL:        testCase.Recipes[0].Recipe.Pictures[0].Entity.URL,
											Width:      testCase.Recipes[0].Recipe.Pictures[0].Entity.Width,
											Height:     testCase.Recipes[0].Recipe.Pictures[0].Entity.Height,
											Size:       testCase.Recipes[0].Recipe.Pictures[0].Entity.Size,
											Type:       testCase.Recipes[0].Recipe.Pictures[0].Entity.Type,
											Status:     testCase.Recipes[0].Recipe.Pictures[0].Entity.Status,
										},
									},
								},
							},
						},
					},
				}

				assert.Equal(t, testCase.Entity.Id, plannerIntervalAggregate.Entity.Id)
				assert.Equal(t, testCase.Entity.UserId, plannerIntervalAggregate.Entity.UserId)
				assert.Equal(t, testCase.Entity.EntityId, plannerIntervalAggregate.Entity.EntityId)
				assert.Equal(t, testCase.Entity.DateInsert, plannerIntervalAggregate.Entity.DateInsert)
				assert.Equal(t, testCase.Entity.DateUpdate, plannerIntervalAggregate.Entity.DateUpdate)
				assert.Equal(t, testCase.Entity.StartTime, plannerIntervalAggregate.Entity.StartTime)
				assert.Equal(t, testCase.Entity.EndTime, plannerIntervalAggregate.Entity.EndTime)
				assert.Equal(t, testCase.Entity.Name, plannerIntervalAggregate.Entity.Name)
				assert.Equal(t, testCase.Entity.Status, plannerIntervalAggregate.Entity.Status)
				assert.Equal(t, testCase.Recipes[0].Entity.Id, plannerIntervalAggregate.Recipes[0].Entity.Id)
				assert.Equal(t, testCase.Recipes[0].Entity.UserId, plannerIntervalAggregate.Recipes[0].Entity.UserId)
				assert.Equal(t, testCase.Recipes[0].Entity.EntityId, plannerIntervalAggregate.Recipes[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipes[0].Entity.RecipeId, plannerIntervalAggregate.Recipes[0].Entity.RecipeId)
				assert.Equal(t, testCase.Recipes[0].Entity.DateInsert, plannerIntervalAggregate.Recipes[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipes[0].Entity.DateUpdate, plannerIntervalAggregate.Recipes[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Entity.Status, plannerIntervalAggregate.Recipes[0].Entity.Status)
				assert.Equal(t, testCase.Recipes[0].Recipe.AltNames[0].Id, plannerIntervalAggregate.Recipes[0].Recipe.AltNames[0].Id)
				assert.Equal(t, testCase.Recipes[0].Recipe.AltNames[0].UserId, plannerIntervalAggregate.Recipes[0].Recipe.AltNames[0].UserId)
				assert.Equal(t, testCase.Recipes[0].Recipe.AltNames[0].EntityId, plannerIntervalAggregate.Recipes[0].Recipe.AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipes[0].Recipe.AltNames[0].DateInsert, plannerIntervalAggregate.Recipes[0].Recipe.AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipes[0].Recipe.AltNames[0].DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.AltNames[0].Name, plannerIntervalAggregate.Recipes[0].Recipe.AltNames[0].Name)
				assert.Equal(t, testCase.Recipes[0].Recipe.AltNames[0].Status, plannerIntervalAggregate.Recipes[0].Recipe.AltNames[0].Status)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.AltNames[0].Id, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.AltNames[0].Id)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.AltNames[0].UserId, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.AltNames[0].UserId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.AltNames[0].EntityId, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.AltNames[0].DateInsert, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.AltNames[0].DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.AltNames[0].Name, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.AltNames[0].Name)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.AltNames[0].Status, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.AltNames[0].Status)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Entity.Id, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Entity.Id)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Entity.UserId, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Entity.UserId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Entity.DateInsert, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Entity.DateInsert)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Entity.DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Entity.DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Entity.Name, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Entity.Name)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Entity.Status, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Entity.Status)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Id, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Id)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].UserId, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].EntityId, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].DateInsert, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Name, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Name)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Status, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Status)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Id, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Id)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.UserId, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.UserId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.EntityId, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.DateInsert, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Name, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Name)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.URL, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.URL)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Width, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Width)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Height, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Height)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Size, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Size)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Type, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Type)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Status, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Derive.Pictures[0].Entity.Status)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Entity.Id, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Entity.Id)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Entity.UserId, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Entity.UserId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Entity.EntityId, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Entity.DeriveId, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Entity.DeriveId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Entity.DateInsert, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Entity.DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.Categories[0].Entity.Status, plannerIntervalAggregate.Recipes[0].Recipe.Categories[0].Entity.Status)
				assert.Equal(t, testCase.Recipes[0].Recipe.Entity.Id, plannerIntervalAggregate.Recipes[0].Recipe.Entity.Id)
				assert.Equal(t, testCase.Recipes[0].Recipe.Entity.UserId, plannerIntervalAggregate.Recipes[0].Recipe.Entity.UserId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Entity.DateInsert, plannerIntervalAggregate.Recipes[0].Recipe.Entity.DateInsert)
				assert.Equal(t, testCase.Recipes[0].Recipe.Entity.DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.Entity.DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.Entity.Name, plannerIntervalAggregate.Recipes[0].Recipe.Entity.Name)
				assert.Equal(t, testCase.Recipes[0].Recipe.Entity.Description, plannerIntervalAggregate.Recipes[0].Recipe.Entity.Description)
				assert.Equal(t, testCase.Recipes[0].Recipe.Entity.Notes, plannerIntervalAggregate.Recipes[0].Recipe.Entity.Notes)
				assert.Equal(t, testCase.Recipes[0].Recipe.Entity.Status, plannerIntervalAggregate.Recipes[0].Recipe.Entity.Status)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].AltNames[0].Id, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].AltNames[0].Id)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].AltNames[0].UserId, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].AltNames[0].EntityId, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].AltNames[0].DateInsert, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].AltNames[0].DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].AltNames[0].Name, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].AltNames[0].Name)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].AltNames[0].Status, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].AltNames[0].Status)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Derive.Id, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Derive.Id)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Derive.UserId, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Derive.UserId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Derive.DateInsert, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Derive.DateInsert)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Derive.DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Derive.DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Derive.Name, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Derive.Name)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Derive.Status, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Derive.Status)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Entity.Id, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Entity.Id)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Entity.UserId, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Entity.UserId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Entity.EntityId, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Entity.DeriveId, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Entity.DeriveId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Entity.DateInsert, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Entity.DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Entity.Name, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Entity.Name)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Entity.Status, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Entity.Status)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].Id, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].Id)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].UserId, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].EntityId, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].DateInsert, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].Name, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].Name)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].Status, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Measures[0].AltNames[0].Status)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.Id, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.Id)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.UserId, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.UserId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.EntityId, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.DateInsert, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.Value, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.Value)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.Status, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Measures[0].Entity.Status)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.Id, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.Id)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.DateInsert, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.DateInsert)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.Name, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.Name)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.Status, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Measures[0].Unit.Status)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].Id, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].Id)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].UserId, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].EntityId, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].DateInsert, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].Name, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].Name)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].Status, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Pictures[0].AltNames[0].Status)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Id, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Id)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.UserId, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.UserId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.EntityId, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.DateInsert, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Name, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Name)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.URL, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.URL)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Width, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Width)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Height, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Height)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Size, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Size)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Type, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Type)
				assert.Equal(t, testCase.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Status, plannerIntervalAggregate.Recipes[0].Recipe.Ingredients[0].Pictures[0].Entity.Status)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].AltNames[0].Id, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].AltNames[0].Id)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].AltNames[0].UserId, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].AltNames[0].EntityId, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].AltNames[0].DateInsert, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].AltNames[0].DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].AltNames[0].Name, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].AltNames[0].Name)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].AltNames[0].Status, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].AltNames[0].Status)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Entity.Id, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Entity.Id)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Entity.UserId, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Entity.UserId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Entity.EntityId, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Entity.DateInsert, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Entity.DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Entity.Name, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Entity.Name)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Entity.Description, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Entity.Description)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Entity.Notes, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Entity.Notes)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Entity.Status, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Entity.Status)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].Id, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].Id)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].UserId, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].EntityId, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].DateInsert, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].Name, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].Name)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].Status, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Pictures[0].AltNames[0].Status)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Id, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Id)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.UserId, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.UserId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.EntityId, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.DateInsert, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Name, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Name)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.URL, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.URL)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Width, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Width)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Height, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Height)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Size, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Size)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Type, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Type)
				assert.Equal(t, testCase.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Status, plannerIntervalAggregate.Recipes[0].Recipe.Processes[0].Pictures[0].Entity.Status)
				assert.Equal(t, testCase.Recipes[0].Recipe.Pictures[0].AltNames[0].Id, plannerIntervalAggregate.Recipes[0].Recipe.Pictures[0].AltNames[0].Id)
				assert.Equal(t, testCase.Recipes[0].Recipe.Pictures[0].AltNames[0].UserId, plannerIntervalAggregate.Recipes[0].Recipe.Pictures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Pictures[0].AltNames[0].EntityId, plannerIntervalAggregate.Recipes[0].Recipe.Pictures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Pictures[0].AltNames[0].DateInsert, plannerIntervalAggregate.Recipes[0].Recipe.Pictures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipes[0].Recipe.Pictures[0].AltNames[0].DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.Pictures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.Pictures[0].AltNames[0].Name, plannerIntervalAggregate.Recipes[0].Recipe.Pictures[0].AltNames[0].Name)
				assert.Equal(t, testCase.Recipes[0].Recipe.Pictures[0].AltNames[0].Status, plannerIntervalAggregate.Recipes[0].Recipe.Pictures[0].AltNames[0].Status)
				assert.Equal(t, testCase.Recipes[0].Recipe.Pictures[0].Entity.Id, plannerIntervalAggregate.Recipes[0].Recipe.Pictures[0].Entity.Id)
				assert.Equal(t, testCase.Recipes[0].Recipe.Pictures[0].Entity.UserId, plannerIntervalAggregate.Recipes[0].Recipe.Pictures[0].Entity.UserId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Pictures[0].Entity.EntityId, plannerIntervalAggregate.Recipes[0].Recipe.Pictures[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipes[0].Recipe.Pictures[0].Entity.DateInsert, plannerIntervalAggregate.Recipes[0].Recipe.Pictures[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipes[0].Recipe.Pictures[0].Entity.DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.Pictures[0].Entity.DateUpdate, plannerIntervalAggregate.Recipes[0].Recipe.Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipes[0].Recipe.Pictures[0].Entity.Name, plannerIntervalAggregate.Recipes[0].Recipe.Pictures[0].Entity.Name)
				assert.Equal(t, testCase.Recipes[0].Recipe.Pictures[0].Entity.URL, plannerIntervalAggregate.Recipes[0].Recipe.Pictures[0].Entity.URL)
				assert.Equal(t, testCase.Recipes[0].Recipe.Pictures[0].Entity.Width, plannerIntervalAggregate.Recipes[0].Recipe.Pictures[0].Entity.Width)
				assert.Equal(t, testCase.Recipes[0].Recipe.Pictures[0].Entity.Height, plannerIntervalAggregate.Recipes[0].Recipe.Pictures[0].Entity.Height)
				assert.Equal(t, testCase.Recipes[0].Recipe.Pictures[0].Entity.Size, plannerIntervalAggregate.Recipes[0].Recipe.Pictures[0].Entity.Size)
				assert.Equal(t, testCase.Recipes[0].Recipe.Pictures[0].Entity.Type, plannerIntervalAggregate.Recipes[0].Recipe.Pictures[0].Entity.Type)
				assert.Equal(t, testCase.Recipes[0].Recipe.Pictures[0].Entity.Status, plannerIntervalAggregate.Recipes[0].Recipe.Pictures[0].Entity.Status)

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(plannerIntervalAggregate)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())

				reflectPlannerRecipeAggregate := reflect.ValueOf(plannerIntervalAggregate)

				for i := 0; i < reflectPlannerRecipeAggregate.NumField(); i++ {
					assert.False(t, reflectPlannerRecipeAggregate.Field(i).IsZero())
				}
			},
		)
	}
}
func TestPlannerRecipe(t *testing.T) {
	tests := []struct {
		name   string
		json   string
		Entity plannerRecipe
		Recipe testRecipeAggregate
	}{
		{
			name: "Test case with active planner recipe properties",
			json: "{\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000003\",\"recipe_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"status\":\"active\"},\"recipe\":{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000005\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"categories\":[{\"derive\":{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000006\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000007\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000007\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Category\",\"status\":\"published\"},\"pictures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000008\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000009\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000009\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000007\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"published\"}}]},\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000010\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"derive_id\":\"00000000-0000-0000-0000-000000000006\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"status\":\"published\"}}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000004\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Recipe\",\"description\":\"Description\",\"notes\":\"Notes\",\"status\":\"published\"},\"ingredients\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000011\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000013\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"derive\":{\"id\":\"00000000-0000-0000-0000-000000000012\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Ingredient\",\"status\":\"published\"},\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000013\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"derive_id\":\"00000000-0000-0000-0000-000000000012\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"RecipeIngredient\",\"status\":\"published\"},\"measures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000014\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000015\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000015\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000013\",\"unit_id\":\"00000000-0000-0000-0000-000000000016\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"value\":42,\"status\":\"published\"},\"unit\":{\"id\":\"00000000-0000-0000-0000-000000000016\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Unit\",\"status\":\"published\"}}],\"pictures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000017\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000018\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000018\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000013\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"published\"}}]}],\"processes\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000019\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000020\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000020\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"RecipeProcess\",\"description\":\"Description\",\"notes\":\"Notes\",\"status\":\"published\"},\"pictures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000021\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000022\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000022\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000020\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"published\"}}]}],\"pictures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000023\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000024\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000024\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"published\"}}]}}\n",
			Entity: plannerRecipe{
				Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000003"),
				RecipeId:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
				DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
				DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
				Status:     kind.PlannerRecipeStatusActive,
			},
			Recipe: struct {
				AltNames   []testAltName
				Categories []struct {
					Derive struct {
						AltNames []testAltName
						Entity   testCategory
						Pictures []testPicture
					}
					Entity testRecipeCategory
				}
				Entity      testRecipe
				Ingredients []struct {
					AltNames []testAltName
					Derive   testIngredient
					Entity   testRecipeIngredient
					Measures []struct {
						AltNames []testAltName
						Entity   testRecipeMeasure
						Unit     testUnit
					}
					Pictures []testPicture
				}
				Processes []struct {
					AltNames []testAltName
					Entity   testRecipeProcess
					Pictures []testPicture
				}
				Pictures []testPicture
			}{
				AltNames: []testAltName{
					{
						Id:         uuid.MustParse("00000000-0000-0000-0000-000000000005"),
						UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
						EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
						DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
						DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
						Name:       "AltName",
						Status:     kind.AltNameStatusPublished,
					},
				},
				Categories: []struct {
					Derive struct {
						AltNames []testAltName
						Entity   testCategory
						Pictures []testPicture
					}
					Entity testRecipeCategory
				}{
					{
						Derive: struct {
							AltNames []testAltName
							Entity   testCategory
							Pictures []testPicture
						}{
							AltNames: []testAltName{
								{
									Id:         uuid.MustParse("00000000-0000-0000-0000-000000000006"),
									UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
									EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000007"),
									DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
									DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
									Name:       "AltName",
									Status:     kind.AltNameStatusPublished,
								},
							},
							Entity: testCategory{
								Id:         uuid.MustParse("00000000-0000-0000-0000-000000000007"),
								UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
								DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).UTC(),
								DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC).UTC(),
								Name:       "Category",
								Status:     kind.CategoryStatusPublished,
							},
							Pictures: []testPicture{
								{
									AltNames: []testAltName{
										{
											Id:         uuid.MustParse("00000000-0000-0000-0000-000000000008"),
											UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
											EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000009"),
											DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
											DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
											Name:       "AltName",
											Status:     kind.AltNameStatusPublished,
										},
									},
									Entity: struct {
										Id         uuid.UUID
										UserId     uuid.UUID
										EntityId   uuid.UUID
										DateInsert time.Time
										DateUpdate time.Time
										Name       string
										URL        string
										Width      int64
										Height     int64
										Size       int64
										Type       string
										Status     kind.PictureStatus
									}{
										Id:         uuid.MustParse("00000000-0000-0000-0000-000000000009"),
										UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
										EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000007"),
										DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
										DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
										Name:       "Picture",
										URL:        "https://google.com/doodle.png",
										Width:      512,
										Height:     512,
										Size:       1024,
										Type:       "image/png",
										Status:     kind.PictureStatusPublished,
									},
								},
							},
						},
						Entity: testRecipeCategory{
							Id:         uuid.MustParse("00000000-0000-0000-0000-000000000010"),
							UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
							EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
							DeriveId:   uuid.MustParse("00000000-0000-0000-0000-000000000006"),
							DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).UTC(),
							DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC).UTC(),
							Status:     kind.RecipeCategoryStatusPublished,
						},
					},
				},
				Entity: testRecipe{
					Id:          uuid.MustParse("00000000-0000-0000-0000-000000000004"),
					UserId:      uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					DateInsert:  time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).UTC(),
					DateUpdate:  time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC).UTC(),
					Name:        "Recipe",
					Description: "Description",
					Notes:       "Notes",
					Status:      kind.RecipeStatusPublished,
				},
				Ingredients: []struct {
					AltNames []testAltName
					Derive   testIngredient
					Entity   testRecipeIngredient
					Measures []struct {
						AltNames []testAltName
						Entity   testRecipeMeasure
						Unit     testUnit
					}
					Pictures []testPicture
				}{
					{
						AltNames: []testAltName{
							{
								Id:         uuid.MustParse("00000000-0000-0000-0000-000000000011"),
								UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
								EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000013"),
								DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
								DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
								Name:       "AltName",
								Status:     kind.AltNameStatusPublished,
							},
						},
						Derive: struct {
							Id         uuid.UUID
							UserId     uuid.UUID
							DateInsert time.Time
							DateUpdate time.Time
							Name       string
							Status     kind.IngredientStatus
						}{
							Id:         uuid.MustParse("00000000-0000-0000-0000-000000000012"),
							UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
							DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).UTC(),
							DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC).UTC(),
							Name:       "Ingredient",
							Status:     kind.IngredientStatusPublished,
						},
						Entity: struct {
							Id         uuid.UUID
							UserId     uuid.UUID
							EntityId   uuid.UUID
							DeriveId   uuid.UUID
							DateInsert time.Time
							DateUpdate time.Time
							Name       string
							Status     kind.RecipeIngredientStatus
						}{
							Id:         uuid.MustParse("00000000-0000-0000-0000-000000000013"),
							UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
							EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
							DeriveId:   uuid.MustParse("00000000-0000-0000-0000-000000000012"),
							DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).UTC(),
							DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC).UTC(),
							Name:       "RecipeIngredient",
							Status:     kind.RecipeIngredientStatusPublished,
						},
						Measures: []struct {
							AltNames []testAltName
							Entity   testRecipeMeasure
							Unit     testUnit
						}{
							{
								AltNames: []testAltName{
									{
										Id:         uuid.MustParse("00000000-0000-0000-0000-000000000014"),
										UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
										EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000015"),
										DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
										DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
										Name:       "AltName",
										Status:     kind.AltNameStatusPublished,
									},
								},
								Entity: testRecipeMeasure{
									Id:         uuid.MustParse("00000000-0000-0000-0000-000000000015"),
									UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
									EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000013"),
									UnitId:     uuid.MustParse("00000000-0000-0000-0000-000000000016"),
									DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).UTC(),
									DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC).UTC(),
									Value:      42,
									Status:     kind.RecipeMeasureStatusPublished,
								},
								Unit: testUnit{
									Id:         uuid.MustParse("00000000-0000-0000-0000-000000000016"),
									DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
									DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
									Name:       "Unit",
									Status:     kind.UnitStatusPublished,
								},
							},
						},
						Pictures: []testPicture{
							{
								AltNames: []testAltName{
									{
										Id:         uuid.MustParse("00000000-0000-0000-0000-000000000017"),
										UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
										EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000018"),
										DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
										DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
										Name:       "AltName",
										Status:     kind.AltNameStatusPublished,
									},
								},
								Entity: struct {
									Id         uuid.UUID
									UserId     uuid.UUID
									EntityId   uuid.UUID
									DateInsert time.Time
									DateUpdate time.Time
									Name       string
									URL        string
									Width      int64
									Height     int64
									Size       int64
									Type       string
									Status     kind.PictureStatus
								}{
									Id:         uuid.MustParse("00000000-0000-0000-0000-000000000018"),
									UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
									EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000013"),
									DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
									DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
									Name:       "Picture",
									URL:        "https://google.com/doodle.png",
									Width:      512,
									Height:     512,
									Size:       1024,
									Type:       "image/png",
									Status:     kind.PictureStatusPublished,
								},
							},
						},
					},
				},
				Processes: []struct {
					AltNames []testAltName
					Entity   testRecipeProcess
					Pictures []testPicture
				}{
					{
						AltNames: []testAltName{
							{
								Id:         uuid.MustParse("00000000-0000-0000-0000-000000000019"),
								UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
								EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000020"),
								DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
								DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
								Name:       "AltName",
								Status:     kind.AltNameStatusPublished,
							},
						},
						Entity: testRecipeProcess{
							Id:          uuid.MustParse("00000000-0000-0000-0000-000000000020"),
							UserId:      uuid.MustParse("00000000-0000-0000-0000-000000000002"),
							EntityId:    uuid.MustParse("00000000-0000-0000-0000-000000000004"),
							DateInsert:  time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).UTC(),
							DateUpdate:  time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC).UTC(),
							Name:        "RecipeProcess",
							Description: "Description",
							Notes:       "Notes",
							Status:      kind.RecipeProcessStatusPublished,
						},
						Pictures: []testPicture{
							{
								AltNames: []testAltName{
									{
										Id:         uuid.MustParse("00000000-0000-0000-0000-000000000021"),
										UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
										EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000022"),
										DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
										DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
										Name:       "AltName",
										Status:     kind.AltNameStatusPublished,
									},
								},
								Entity: struct {
									Id         uuid.UUID
									UserId     uuid.UUID
									EntityId   uuid.UUID
									DateInsert time.Time
									DateUpdate time.Time
									Name       string
									URL        string
									Width      int64
									Height     int64
									Size       int64
									Type       string
									Status     kind.PictureStatus
								}{
									Id:         uuid.MustParse("00000000-0000-0000-0000-000000000022"),
									UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
									EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000020"),
									DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
									DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
									Name:       "Picture",
									URL:        "https://google.com/doodle.png",
									Width:      512,
									Height:     512,
									Size:       1024,
									Type:       "image/png",
									Status:     kind.PictureStatusPublished,
								},
							},
						},
					},
				},
				Pictures: []testPicture{
					{
						AltNames: []testAltName{
							{
								Id:         uuid.MustParse("00000000-0000-0000-0000-000000000023"),
								UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
								EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000024"),
								DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
								DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
								Name:       "AltName",
								Status:     kind.AltNameStatusPublished,
							},
						},
						Entity: struct {
							Id         uuid.UUID
							UserId     uuid.UUID
							EntityId   uuid.UUID
							DateInsert time.Time
							DateUpdate time.Time
							Name       string
							URL        string
							Width      int64
							Height     int64
							Size       int64
							Type       string
							Status     kind.PictureStatus
						}{
							Id:         uuid.MustParse("00000000-0000-0000-0000-000000000024"),
							UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
							EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
							DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
							DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
							Name:       "Picture",
							URL:        "https://google.com/doodle.png",
							Width:      512,
							Height:     512,
							Size:       1024,
							Type:       "image/png",
							Status:     kind.PictureStatusPublished,
						},
					},
				},
			},
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				plannerRecipeAggregate := PlannerRecipe{
					Entity: &DomainEntity.PlannerRecipe{
						Id:         testCase.Entity.Id,
						UserId:     testCase.Entity.UserId,
						EntityId:   testCase.Entity.EntityId,
						RecipeId:   testCase.Entity.RecipeId,
						DateInsert: testCase.Entity.DateInsert,
						DateUpdate: testCase.Entity.DateUpdate,
						Status:     testCase.Entity.Status,
					},
					Recipe: &Recipe{
						AltNames: []*DomainEntity.AltName{
							{
								Id:         testCase.Recipe.AltNames[0].Id,
								UserId:     testCase.Recipe.AltNames[0].UserId,
								EntityId:   testCase.Recipe.AltNames[0].EntityId,
								DateInsert: testCase.Recipe.AltNames[0].DateInsert,
								DateUpdate: testCase.Recipe.AltNames[0].DateUpdate,
								Name:       testCase.Recipe.AltNames[0].Name,
								Status:     testCase.Recipe.AltNames[0].Status,
							},
						},
						Categories: []*RecipeCategory{
							{
								Derive: &Category{
									AltNames: []*DomainEntity.AltName{
										{
											Id:         testCase.Recipe.Categories[0].Derive.AltNames[0].Id,
											UserId:     testCase.Recipe.Categories[0].Derive.AltNames[0].UserId,
											EntityId:   testCase.Recipe.Categories[0].Derive.AltNames[0].EntityId,
											DateInsert: testCase.Recipe.Categories[0].Derive.AltNames[0].DateInsert,
											DateUpdate: testCase.Recipe.Categories[0].Derive.AltNames[0].DateUpdate,
											Name:       testCase.Recipe.Categories[0].Derive.AltNames[0].Name,
											Status:     testCase.Recipe.Categories[0].Derive.AltNames[0].Status,
										},
									},
									Entity: &DomainEntity.Category{
										Id:         testCase.Recipe.Categories[0].Derive.Entity.Id,
										UserId:     testCase.Recipe.Categories[0].Derive.Entity.UserId,
										DateInsert: testCase.Recipe.Categories[0].Derive.Entity.DateInsert,
										DateUpdate: testCase.Recipe.Categories[0].Derive.Entity.DateUpdate,
										Name:       testCase.Recipe.Categories[0].Derive.Entity.Name,
										Status:     testCase.Recipe.Categories[0].Derive.Entity.Status,
									},
									Pictures: []*Picture{
										{
											AltNames: []*DomainEntity.AltName{
												{
													Id:         testCase.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Id,
													UserId:     testCase.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].UserId,
													EntityId:   testCase.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].EntityId,
													DateInsert: testCase.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].DateInsert,
													DateUpdate: testCase.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].DateUpdate,
													Name:       testCase.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Name,
													Status:     testCase.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Status,
												},
											},
											Entity: &DomainEntity.Picture{
												Id:         testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.Id,
												UserId:     testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.UserId,
												EntityId:   testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.EntityId,
												DateInsert: testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.DateInsert,
												DateUpdate: testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.DateUpdate,
												Name:       testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.Name,
												URL:        testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.URL,
												Width:      testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.Width,
												Height:     testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.Height,
												Size:       testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.Size,
												Type:       testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.Type,
												Status:     testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.Status,
											},
										},
									},
								},
								Entity: &DomainEntity.RecipeCategory{
									Id:         testCase.Recipe.Categories[0].Entity.Id,
									UserId:     testCase.Recipe.Categories[0].Entity.UserId,
									EntityId:   testCase.Recipe.Categories[0].Entity.EntityId,
									DeriveId:   testCase.Recipe.Categories[0].Entity.DeriveId,
									DateInsert: testCase.Recipe.Categories[0].Entity.DateInsert,
									DateUpdate: testCase.Recipe.Categories[0].Entity.DateUpdate,
									Status:     testCase.Recipe.Categories[0].Entity.Status,
								},
							},
						},
						Entity: &DomainEntity.Recipe{
							Id:          testCase.Recipe.Entity.Id,
							UserId:      testCase.Recipe.Entity.UserId,
							DateInsert:  testCase.Recipe.Entity.DateInsert,
							DateUpdate:  testCase.Recipe.Entity.DateUpdate,
							Name:        testCase.Recipe.Entity.Name,
							Description: testCase.Recipe.Entity.Description,
							Notes:       testCase.Recipe.Entity.Notes,
							Status:      testCase.Recipe.Entity.Status,
						},
						Ingredients: []*RecipeIngredient{
							{
								AltNames: []*DomainEntity.AltName{
									{
										Id:         testCase.Recipe.Ingredients[0].AltNames[0].Id,
										UserId:     testCase.Recipe.Ingredients[0].AltNames[0].UserId,
										EntityId:   testCase.Recipe.Ingredients[0].AltNames[0].EntityId,
										DateInsert: testCase.Recipe.Ingredients[0].AltNames[0].DateInsert,
										DateUpdate: testCase.Recipe.Ingredients[0].AltNames[0].DateUpdate,
										Name:       testCase.Recipe.Ingredients[0].AltNames[0].Name,
										Status:     testCase.Recipe.Ingredients[0].AltNames[0].Status,
									},
								},
								Derive: &DomainEntity.Ingredient{
									Id:         testCase.Recipe.Ingredients[0].Derive.Id,
									UserId:     testCase.Recipe.Ingredients[0].Derive.UserId,
									DateInsert: testCase.Recipe.Ingredients[0].Derive.DateInsert,
									DateUpdate: testCase.Recipe.Ingredients[0].Derive.DateUpdate,
									Name:       testCase.Recipe.Ingredients[0].Derive.Name,
									Status:     testCase.Recipe.Ingredients[0].Derive.Status,
								},
								Entity: &DomainEntity.RecipeIngredient{
									Id:         testCase.Recipe.Ingredients[0].Entity.Id,
									UserId:     testCase.Recipe.Ingredients[0].Entity.UserId,
									EntityId:   testCase.Recipe.Ingredients[0].Entity.EntityId,
									DeriveId:   testCase.Recipe.Ingredients[0].Entity.DeriveId,
									DateInsert: testCase.Recipe.Ingredients[0].Entity.DateInsert,
									DateUpdate: testCase.Recipe.Ingredients[0].Entity.DateUpdate,
									Name:       testCase.Recipe.Ingredients[0].Entity.Name,
									Status:     testCase.Recipe.Ingredients[0].Entity.Status,
								},
								Measures: []*RecipeMeasure{
									{
										AltNames: []*DomainEntity.AltName{
											{
												Id:         testCase.Recipe.Ingredients[0].Measures[0].AltNames[0].Id,
												UserId:     testCase.Recipe.Ingredients[0].Measures[0].AltNames[0].UserId,
												EntityId:   testCase.Recipe.Ingredients[0].Measures[0].AltNames[0].EntityId,
												DateInsert: testCase.Recipe.Ingredients[0].Measures[0].AltNames[0].DateInsert,
												DateUpdate: testCase.Recipe.Ingredients[0].Measures[0].AltNames[0].DateUpdate,
												Name:       testCase.Recipe.Ingredients[0].Measures[0].AltNames[0].Name,
												Status:     testCase.Recipe.Ingredients[0].Measures[0].AltNames[0].Status,
											},
										},
										Entity: &DomainEntity.RecipeMeasure{
											Id:         testCase.Recipe.Ingredients[0].Measures[0].Entity.Id,
											UserId:     testCase.Recipe.Ingredients[0].Measures[0].Entity.UserId,
											EntityId:   testCase.Recipe.Ingredients[0].Measures[0].Entity.EntityId,
											UnitId:     testCase.Recipe.Ingredients[0].Measures[0].Entity.UnitId,
											DateInsert: testCase.Recipe.Ingredients[0].Measures[0].Entity.DateInsert,
											DateUpdate: testCase.Recipe.Ingredients[0].Measures[0].Entity.DateUpdate,
											Value:      testCase.Recipe.Ingredients[0].Measures[0].Entity.Value,
											Status:     testCase.Recipe.Ingredients[0].Measures[0].Entity.Status,
										},
										Unit: &DomainEntity.Unit{
											Id:         testCase.Recipe.Ingredients[0].Measures[0].Unit.Id,
											DateInsert: testCase.Recipe.Ingredients[0].Measures[0].Unit.DateInsert,
											DateUpdate: testCase.Recipe.Ingredients[0].Measures[0].Unit.DateUpdate,
											Name:       testCase.Recipe.Ingredients[0].Measures[0].Unit.Name,
											Status:     testCase.Recipe.Ingredients[0].Measures[0].Unit.Status,
										},
									},
								},
								Pictures: []*Picture{
									{
										AltNames: []*DomainEntity.AltName{
											{
												Id:         testCase.Recipe.Ingredients[0].Pictures[0].AltNames[0].Id,
												UserId:     testCase.Recipe.Ingredients[0].Pictures[0].AltNames[0].UserId,
												EntityId:   testCase.Recipe.Ingredients[0].Pictures[0].AltNames[0].EntityId,
												DateInsert: testCase.Recipe.Ingredients[0].Pictures[0].AltNames[0].DateInsert,
												DateUpdate: testCase.Recipe.Ingredients[0].Pictures[0].AltNames[0].DateUpdate,
												Name:       testCase.Recipe.Ingredients[0].Pictures[0].AltNames[0].Name,
												Status:     testCase.Recipe.Ingredients[0].Pictures[0].AltNames[0].Status,
											},
										},
										Entity: &DomainEntity.Picture{
											Id:         testCase.Recipe.Ingredients[0].Pictures[0].Entity.Id,
											UserId:     testCase.Recipe.Ingredients[0].Pictures[0].Entity.UserId,
											EntityId:   testCase.Recipe.Ingredients[0].Pictures[0].Entity.EntityId,
											DateInsert: testCase.Recipe.Ingredients[0].Pictures[0].Entity.DateInsert,
											DateUpdate: testCase.Recipe.Ingredients[0].Pictures[0].Entity.DateUpdate,
											Name:       testCase.Recipe.Ingredients[0].Pictures[0].Entity.Name,
											URL:        testCase.Recipe.Ingredients[0].Pictures[0].Entity.URL,
											Width:      testCase.Recipe.Ingredients[0].Pictures[0].Entity.Width,
											Height:     testCase.Recipe.Ingredients[0].Pictures[0].Entity.Height,
											Size:       testCase.Recipe.Ingredients[0].Pictures[0].Entity.Size,
											Type:       testCase.Recipe.Ingredients[0].Pictures[0].Entity.Type,
											Status:     testCase.Recipe.Ingredients[0].Pictures[0].Entity.Status,
										},
									},
								},
							},
						},
						Processes: []*RecipeProcess{
							{
								AltNames: []*DomainEntity.AltName{
									{
										Id:         testCase.Recipe.Processes[0].AltNames[0].Id,
										UserId:     testCase.Recipe.Processes[0].AltNames[0].UserId,
										EntityId:   testCase.Recipe.Processes[0].AltNames[0].EntityId,
										DateInsert: testCase.Recipe.Processes[0].AltNames[0].DateInsert,
										DateUpdate: testCase.Recipe.Processes[0].AltNames[0].DateUpdate,
										Name:       testCase.Recipe.Processes[0].AltNames[0].Name,
										Status:     testCase.Recipe.Processes[0].AltNames[0].Status,
									},
								},
								Entity: &DomainEntity.RecipeProcess{
									Id:          testCase.Recipe.Processes[0].Entity.Id,
									UserId:      testCase.Recipe.Processes[0].Entity.UserId,
									EntityId:    testCase.Recipe.Processes[0].Entity.EntityId,
									DateInsert:  testCase.Recipe.Processes[0].Entity.DateInsert,
									DateUpdate:  testCase.Recipe.Processes[0].Entity.DateUpdate,
									Name:        testCase.Recipe.Processes[0].Entity.Name,
									Description: testCase.Recipe.Processes[0].Entity.Description,
									Notes:       testCase.Recipe.Processes[0].Entity.Notes,
									Status:      testCase.Recipe.Processes[0].Entity.Status,
								},
								Pictures: []*Picture{
									{
										AltNames: []*DomainEntity.AltName{
											{
												Id:         testCase.Recipe.Processes[0].Pictures[0].AltNames[0].Id,
												UserId:     testCase.Recipe.Processes[0].Pictures[0].AltNames[0].UserId,
												EntityId:   testCase.Recipe.Processes[0].Pictures[0].AltNames[0].EntityId,
												DateInsert: testCase.Recipe.Processes[0].Pictures[0].AltNames[0].DateInsert,
												DateUpdate: testCase.Recipe.Processes[0].Pictures[0].AltNames[0].DateUpdate,
												Name:       testCase.Recipe.Processes[0].Pictures[0].AltNames[0].Name,
												Status:     testCase.Recipe.Processes[0].Pictures[0].AltNames[0].Status,
											},
										},
										Entity: &DomainEntity.Picture{
											Id:         testCase.Recipe.Processes[0].Pictures[0].Entity.Id,
											UserId:     testCase.Recipe.Processes[0].Pictures[0].Entity.UserId,
											EntityId:   testCase.Recipe.Processes[0].Pictures[0].Entity.EntityId,
											DateInsert: testCase.Recipe.Processes[0].Pictures[0].Entity.DateInsert,
											DateUpdate: testCase.Recipe.Processes[0].Pictures[0].Entity.DateUpdate,
											Name:       testCase.Recipe.Processes[0].Pictures[0].Entity.Name,
											URL:        testCase.Recipe.Processes[0].Pictures[0].Entity.URL,
											Width:      testCase.Recipe.Processes[0].Pictures[0].Entity.Width,
											Height:     testCase.Recipe.Processes[0].Pictures[0].Entity.Height,
											Size:       testCase.Recipe.Processes[0].Pictures[0].Entity.Size,
											Type:       testCase.Recipe.Processes[0].Pictures[0].Entity.Type,
											Status:     testCase.Recipe.Processes[0].Pictures[0].Entity.Status,
										},
									},
								},
							},
						},
						Pictures: []*Picture{
							{
								AltNames: []*DomainEntity.AltName{
									{
										Id:         testCase.Recipe.Pictures[0].AltNames[0].Id,
										UserId:     testCase.Recipe.Pictures[0].AltNames[0].UserId,
										EntityId:   testCase.Recipe.Pictures[0].AltNames[0].EntityId,
										DateInsert: testCase.Recipe.Pictures[0].AltNames[0].DateInsert,
										DateUpdate: testCase.Recipe.Pictures[0].AltNames[0].DateUpdate,
										Name:       testCase.Recipe.Pictures[0].AltNames[0].Name,
										Status:     testCase.Recipe.Pictures[0].AltNames[0].Status,
									},
								},
								Entity: &DomainEntity.Picture{
									Id:         testCase.Recipe.Pictures[0].Entity.Id,
									UserId:     testCase.Recipe.Pictures[0].Entity.UserId,
									EntityId:   testCase.Recipe.Pictures[0].Entity.EntityId,
									DateInsert: testCase.Recipe.Pictures[0].Entity.DateInsert,
									DateUpdate: testCase.Recipe.Pictures[0].Entity.DateUpdate,
									Name:       testCase.Recipe.Pictures[0].Entity.Name,
									URL:        testCase.Recipe.Pictures[0].Entity.URL,
									Width:      testCase.Recipe.Pictures[0].Entity.Width,
									Height:     testCase.Recipe.Pictures[0].Entity.Height,
									Size:       testCase.Recipe.Pictures[0].Entity.Size,
									Type:       testCase.Recipe.Pictures[0].Entity.Type,
									Status:     testCase.Recipe.Pictures[0].Entity.Status,
								},
							},
						},
					},
				}

				assert.Equal(t, testCase.Entity.Id, plannerRecipeAggregate.Entity.Id)
				assert.Equal(t, testCase.Entity.UserId, plannerRecipeAggregate.Entity.UserId)
				assert.Equal(t, testCase.Entity.EntityId, plannerRecipeAggregate.Entity.EntityId)
				assert.Equal(t, testCase.Entity.RecipeId, plannerRecipeAggregate.Entity.RecipeId)
				assert.Equal(t, testCase.Entity.DateInsert, plannerRecipeAggregate.Entity.DateInsert)
				assert.Equal(t, testCase.Entity.DateUpdate, plannerRecipeAggregate.Entity.DateUpdate)
				assert.Equal(t, testCase.Entity.Status, plannerRecipeAggregate.Entity.Status)
				assert.Equal(t, testCase.Recipe.AltNames[0].Id, plannerRecipeAggregate.Recipe.AltNames[0].Id)
				assert.Equal(t, testCase.Recipe.AltNames[0].UserId, plannerRecipeAggregate.Recipe.AltNames[0].UserId)
				assert.Equal(t, testCase.Recipe.AltNames[0].EntityId, plannerRecipeAggregate.Recipe.AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipe.AltNames[0].DateInsert, plannerRecipeAggregate.Recipe.AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipe.AltNames[0].DateUpdate, plannerRecipeAggregate.Recipe.AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipe.AltNames[0].Name, plannerRecipeAggregate.Recipe.AltNames[0].Name)
				assert.Equal(t, testCase.Recipe.AltNames[0].Status, plannerRecipeAggregate.Recipe.AltNames[0].Status)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.AltNames[0].Id, plannerRecipeAggregate.Recipe.Categories[0].Derive.AltNames[0].Id)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.AltNames[0].UserId, plannerRecipeAggregate.Recipe.Categories[0].Derive.AltNames[0].UserId)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.AltNames[0].EntityId, plannerRecipeAggregate.Recipe.Categories[0].Derive.AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.AltNames[0].DateInsert, plannerRecipeAggregate.Recipe.Categories[0].Derive.AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.AltNames[0].DateUpdate, plannerRecipeAggregate.Recipe.Categories[0].Derive.AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.AltNames[0].Name, plannerRecipeAggregate.Recipe.Categories[0].Derive.AltNames[0].Name)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.AltNames[0].Status, plannerRecipeAggregate.Recipe.Categories[0].Derive.AltNames[0].Status)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Entity.Id, plannerRecipeAggregate.Recipe.Categories[0].Derive.Entity.Id)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Entity.UserId, plannerRecipeAggregate.Recipe.Categories[0].Derive.Entity.UserId)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Entity.DateInsert, plannerRecipeAggregate.Recipe.Categories[0].Derive.Entity.DateInsert)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Entity.DateUpdate, plannerRecipeAggregate.Recipe.Categories[0].Derive.Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Entity.Name, plannerRecipeAggregate.Recipe.Categories[0].Derive.Entity.Name)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Entity.Status, plannerRecipeAggregate.Recipe.Categories[0].Derive.Entity.Status)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Id, plannerRecipeAggregate.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Id)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].UserId, plannerRecipeAggregate.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].EntityId, plannerRecipeAggregate.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].DateInsert, plannerRecipeAggregate.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].DateUpdate, plannerRecipeAggregate.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Name, plannerRecipeAggregate.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Name)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Status, plannerRecipeAggregate.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Status)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.Id, plannerRecipeAggregate.Recipe.Categories[0].Derive.Pictures[0].Entity.Id)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.UserId, plannerRecipeAggregate.Recipe.Categories[0].Derive.Pictures[0].Entity.UserId)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.EntityId, plannerRecipeAggregate.Recipe.Categories[0].Derive.Pictures[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.DateInsert, plannerRecipeAggregate.Recipe.Categories[0].Derive.Pictures[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.DateUpdate, plannerRecipeAggregate.Recipe.Categories[0].Derive.Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.DateUpdate, plannerRecipeAggregate.Recipe.Categories[0].Derive.Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.Name, plannerRecipeAggregate.Recipe.Categories[0].Derive.Pictures[0].Entity.Name)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.URL, plannerRecipeAggregate.Recipe.Categories[0].Derive.Pictures[0].Entity.URL)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.Width, plannerRecipeAggregate.Recipe.Categories[0].Derive.Pictures[0].Entity.Width)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.Height, plannerRecipeAggregate.Recipe.Categories[0].Derive.Pictures[0].Entity.Height)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.Size, plannerRecipeAggregate.Recipe.Categories[0].Derive.Pictures[0].Entity.Size)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.Type, plannerRecipeAggregate.Recipe.Categories[0].Derive.Pictures[0].Entity.Type)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.Status, plannerRecipeAggregate.Recipe.Categories[0].Derive.Pictures[0].Entity.Status)
				assert.Equal(t, testCase.Recipe.Categories[0].Entity.Id, plannerRecipeAggregate.Recipe.Categories[0].Entity.Id)
				assert.Equal(t, testCase.Recipe.Categories[0].Entity.UserId, plannerRecipeAggregate.Recipe.Categories[0].Entity.UserId)
				assert.Equal(t, testCase.Recipe.Categories[0].Entity.EntityId, plannerRecipeAggregate.Recipe.Categories[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipe.Categories[0].Entity.DeriveId, plannerRecipeAggregate.Recipe.Categories[0].Entity.DeriveId)
				assert.Equal(t, testCase.Recipe.Categories[0].Entity.DateInsert, plannerRecipeAggregate.Recipe.Categories[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipe.Categories[0].Entity.DateUpdate, plannerRecipeAggregate.Recipe.Categories[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Categories[0].Entity.Status, plannerRecipeAggregate.Recipe.Categories[0].Entity.Status)
				assert.Equal(t, testCase.Recipe.Entity.Id, plannerRecipeAggregate.Recipe.Entity.Id)
				assert.Equal(t, testCase.Recipe.Entity.UserId, plannerRecipeAggregate.Recipe.Entity.UserId)
				assert.Equal(t, testCase.Recipe.Entity.DateInsert, plannerRecipeAggregate.Recipe.Entity.DateInsert)
				assert.Equal(t, testCase.Recipe.Entity.DateUpdate, plannerRecipeAggregate.Recipe.Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Entity.Name, plannerRecipeAggregate.Recipe.Entity.Name)
				assert.Equal(t, testCase.Recipe.Entity.Description, plannerRecipeAggregate.Recipe.Entity.Description)
				assert.Equal(t, testCase.Recipe.Entity.Notes, plannerRecipeAggregate.Recipe.Entity.Notes)
				assert.Equal(t, testCase.Recipe.Entity.Status, plannerRecipeAggregate.Recipe.Entity.Status)
				assert.Equal(t, testCase.Recipe.Ingredients[0].AltNames[0].Id, plannerRecipeAggregate.Recipe.Ingredients[0].AltNames[0].Id)
				assert.Equal(t, testCase.Recipe.Ingredients[0].AltNames[0].UserId, plannerRecipeAggregate.Recipe.Ingredients[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].AltNames[0].EntityId, plannerRecipeAggregate.Recipe.Ingredients[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].AltNames[0].DateInsert, plannerRecipeAggregate.Recipe.Ingredients[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipe.Ingredients[0].AltNames[0].DateUpdate, plannerRecipeAggregate.Recipe.Ingredients[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipe.Ingredients[0].AltNames[0].Name, plannerRecipeAggregate.Recipe.Ingredients[0].AltNames[0].Name)
				assert.Equal(t, testCase.Recipe.Ingredients[0].AltNames[0].Status, plannerRecipeAggregate.Recipe.Ingredients[0].AltNames[0].Status)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Derive.Id, plannerRecipeAggregate.Recipe.Ingredients[0].Derive.Id)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Derive.UserId, plannerRecipeAggregate.Recipe.Ingredients[0].Derive.UserId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Derive.DateInsert, plannerRecipeAggregate.Recipe.Ingredients[0].Derive.DateInsert)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Derive.DateUpdate, plannerRecipeAggregate.Recipe.Ingredients[0].Derive.DateUpdate)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Derive.Name, plannerRecipeAggregate.Recipe.Ingredients[0].Derive.Name)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Derive.Status, plannerRecipeAggregate.Recipe.Ingredients[0].Derive.Status)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Entity.Id, plannerRecipeAggregate.Recipe.Ingredients[0].Entity.Id)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Entity.UserId, plannerRecipeAggregate.Recipe.Ingredients[0].Entity.UserId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Entity.EntityId, plannerRecipeAggregate.Recipe.Ingredients[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Entity.DeriveId, plannerRecipeAggregate.Recipe.Ingredients[0].Entity.DeriveId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Entity.DateInsert, plannerRecipeAggregate.Recipe.Ingredients[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Entity.DateUpdate, plannerRecipeAggregate.Recipe.Ingredients[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Entity.Name, plannerRecipeAggregate.Recipe.Ingredients[0].Entity.Name)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Entity.Status, plannerRecipeAggregate.Recipe.Ingredients[0].Entity.Status)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].AltNames[0].Id, plannerRecipeAggregate.Recipe.Ingredients[0].Measures[0].AltNames[0].Id)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].AltNames[0].UserId, plannerRecipeAggregate.Recipe.Ingredients[0].Measures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].AltNames[0].EntityId, plannerRecipeAggregate.Recipe.Ingredients[0].Measures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].AltNames[0].DateInsert, plannerRecipeAggregate.Recipe.Ingredients[0].Measures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].AltNames[0].DateUpdate, plannerRecipeAggregate.Recipe.Ingredients[0].Measures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].AltNames[0].Name, plannerRecipeAggregate.Recipe.Ingredients[0].Measures[0].AltNames[0].Name)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].AltNames[0].Status, plannerRecipeAggregate.Recipe.Ingredients[0].Measures[0].AltNames[0].Status)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].Entity.Id, plannerRecipeAggregate.Recipe.Ingredients[0].Measures[0].Entity.Id)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].Entity.UserId, plannerRecipeAggregate.Recipe.Ingredients[0].Measures[0].Entity.UserId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].Entity.EntityId, plannerRecipeAggregate.Recipe.Ingredients[0].Measures[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].Entity.DateInsert, plannerRecipeAggregate.Recipe.Ingredients[0].Measures[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].Entity.DateUpdate, plannerRecipeAggregate.Recipe.Ingredients[0].Measures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].Entity.Value, plannerRecipeAggregate.Recipe.Ingredients[0].Measures[0].Entity.Value)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].Entity.Status, plannerRecipeAggregate.Recipe.Ingredients[0].Measures[0].Entity.Status)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].Unit.Id, plannerRecipeAggregate.Recipe.Ingredients[0].Measures[0].Unit.Id)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].Unit.DateInsert, plannerRecipeAggregate.Recipe.Ingredients[0].Measures[0].Unit.DateInsert)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].Unit.DateUpdate, plannerRecipeAggregate.Recipe.Ingredients[0].Measures[0].Unit.DateUpdate)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].Unit.Name, plannerRecipeAggregate.Recipe.Ingredients[0].Measures[0].Unit.Name)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].Unit.Status, plannerRecipeAggregate.Recipe.Ingredients[0].Measures[0].Unit.Status)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].AltNames[0].Id, plannerRecipeAggregate.Recipe.Ingredients[0].Pictures[0].AltNames[0].Id)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].AltNames[0].UserId, plannerRecipeAggregate.Recipe.Ingredients[0].Pictures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].AltNames[0].EntityId, plannerRecipeAggregate.Recipe.Ingredients[0].Pictures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].AltNames[0].DateInsert, plannerRecipeAggregate.Recipe.Ingredients[0].Pictures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].AltNames[0].DateUpdate, plannerRecipeAggregate.Recipe.Ingredients[0].Pictures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].AltNames[0].Name, plannerRecipeAggregate.Recipe.Ingredients[0].Pictures[0].AltNames[0].Name)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].AltNames[0].Status, plannerRecipeAggregate.Recipe.Ingredients[0].Pictures[0].AltNames[0].Status)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.Id, plannerRecipeAggregate.Recipe.Ingredients[0].Pictures[0].Entity.Id)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.UserId, plannerRecipeAggregate.Recipe.Ingredients[0].Pictures[0].Entity.UserId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.EntityId, plannerRecipeAggregate.Recipe.Ingredients[0].Pictures[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.DateInsert, plannerRecipeAggregate.Recipe.Ingredients[0].Pictures[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.DateUpdate, plannerRecipeAggregate.Recipe.Ingredients[0].Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.DateUpdate, plannerRecipeAggregate.Recipe.Ingredients[0].Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.Name, plannerRecipeAggregate.Recipe.Ingredients[0].Pictures[0].Entity.Name)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.URL, plannerRecipeAggregate.Recipe.Ingredients[0].Pictures[0].Entity.URL)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.Width, plannerRecipeAggregate.Recipe.Ingredients[0].Pictures[0].Entity.Width)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.Height, plannerRecipeAggregate.Recipe.Ingredients[0].Pictures[0].Entity.Height)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.Size, plannerRecipeAggregate.Recipe.Ingredients[0].Pictures[0].Entity.Size)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.Type, plannerRecipeAggregate.Recipe.Ingredients[0].Pictures[0].Entity.Type)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.Status, plannerRecipeAggregate.Recipe.Ingredients[0].Pictures[0].Entity.Status)
				assert.Equal(t, testCase.Recipe.Processes[0].AltNames[0].Id, plannerRecipeAggregate.Recipe.Processes[0].AltNames[0].Id)
				assert.Equal(t, testCase.Recipe.Processes[0].AltNames[0].UserId, plannerRecipeAggregate.Recipe.Processes[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Recipe.Processes[0].AltNames[0].EntityId, plannerRecipeAggregate.Recipe.Processes[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipe.Processes[0].AltNames[0].DateInsert, plannerRecipeAggregate.Recipe.Processes[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipe.Processes[0].AltNames[0].DateUpdate, plannerRecipeAggregate.Recipe.Processes[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipe.Processes[0].AltNames[0].Name, plannerRecipeAggregate.Recipe.Processes[0].AltNames[0].Name)
				assert.Equal(t, testCase.Recipe.Processes[0].AltNames[0].Status, plannerRecipeAggregate.Recipe.Processes[0].AltNames[0].Status)
				assert.Equal(t, testCase.Recipe.Processes[0].Entity.Id, plannerRecipeAggregate.Recipe.Processes[0].Entity.Id)
				assert.Equal(t, testCase.Recipe.Processes[0].Entity.UserId, plannerRecipeAggregate.Recipe.Processes[0].Entity.UserId)
				assert.Equal(t, testCase.Recipe.Processes[0].Entity.EntityId, plannerRecipeAggregate.Recipe.Processes[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipe.Processes[0].Entity.DateInsert, plannerRecipeAggregate.Recipe.Processes[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipe.Processes[0].Entity.DateUpdate, plannerRecipeAggregate.Recipe.Processes[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Processes[0].Entity.Name, plannerRecipeAggregate.Recipe.Processes[0].Entity.Name)
				assert.Equal(t, testCase.Recipe.Processes[0].Entity.Description, plannerRecipeAggregate.Recipe.Processes[0].Entity.Description)
				assert.Equal(t, testCase.Recipe.Processes[0].Entity.Notes, plannerRecipeAggregate.Recipe.Processes[0].Entity.Notes)
				assert.Equal(t, testCase.Recipe.Processes[0].Entity.Status, plannerRecipeAggregate.Recipe.Processes[0].Entity.Status)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].AltNames[0].Id, plannerRecipeAggregate.Recipe.Processes[0].Pictures[0].AltNames[0].Id)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].AltNames[0].UserId, plannerRecipeAggregate.Recipe.Processes[0].Pictures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].AltNames[0].EntityId, plannerRecipeAggregate.Recipe.Processes[0].Pictures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].AltNames[0].DateInsert, plannerRecipeAggregate.Recipe.Processes[0].Pictures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].AltNames[0].DateUpdate, plannerRecipeAggregate.Recipe.Processes[0].Pictures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].AltNames[0].Name, plannerRecipeAggregate.Recipe.Processes[0].Pictures[0].AltNames[0].Name)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].AltNames[0].Status, plannerRecipeAggregate.Recipe.Processes[0].Pictures[0].AltNames[0].Status)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.Id, plannerRecipeAggregate.Recipe.Processes[0].Pictures[0].Entity.Id)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.UserId, plannerRecipeAggregate.Recipe.Processes[0].Pictures[0].Entity.UserId)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.EntityId, plannerRecipeAggregate.Recipe.Processes[0].Pictures[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.DateInsert, plannerRecipeAggregate.Recipe.Processes[0].Pictures[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.DateUpdate, plannerRecipeAggregate.Recipe.Processes[0].Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.DateUpdate, plannerRecipeAggregate.Recipe.Processes[0].Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.Name, plannerRecipeAggregate.Recipe.Processes[0].Pictures[0].Entity.Name)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.URL, plannerRecipeAggregate.Recipe.Processes[0].Pictures[0].Entity.URL)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.Width, plannerRecipeAggregate.Recipe.Processes[0].Pictures[0].Entity.Width)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.Height, plannerRecipeAggregate.Recipe.Processes[0].Pictures[0].Entity.Height)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.Size, plannerRecipeAggregate.Recipe.Processes[0].Pictures[0].Entity.Size)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.Type, plannerRecipeAggregate.Recipe.Processes[0].Pictures[0].Entity.Type)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.Status, plannerRecipeAggregate.Recipe.Processes[0].Pictures[0].Entity.Status)
				assert.Equal(t, testCase.Recipe.Pictures[0].AltNames[0].Id, plannerRecipeAggregate.Recipe.Pictures[0].AltNames[0].Id)
				assert.Equal(t, testCase.Recipe.Pictures[0].AltNames[0].UserId, plannerRecipeAggregate.Recipe.Pictures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Recipe.Pictures[0].AltNames[0].EntityId, plannerRecipeAggregate.Recipe.Pictures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipe.Pictures[0].AltNames[0].DateInsert, plannerRecipeAggregate.Recipe.Pictures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipe.Pictures[0].AltNames[0].DateUpdate, plannerRecipeAggregate.Recipe.Pictures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipe.Pictures[0].AltNames[0].Name, plannerRecipeAggregate.Recipe.Pictures[0].AltNames[0].Name)
				assert.Equal(t, testCase.Recipe.Pictures[0].AltNames[0].Status, plannerRecipeAggregate.Recipe.Pictures[0].AltNames[0].Status)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.Id, plannerRecipeAggregate.Recipe.Pictures[0].Entity.Id)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.UserId, plannerRecipeAggregate.Recipe.Pictures[0].Entity.UserId)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.EntityId, plannerRecipeAggregate.Recipe.Pictures[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.DateInsert, plannerRecipeAggregate.Recipe.Pictures[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.DateUpdate, plannerRecipeAggregate.Recipe.Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.DateUpdate, plannerRecipeAggregate.Recipe.Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.Name, plannerRecipeAggregate.Recipe.Pictures[0].Entity.Name)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.URL, plannerRecipeAggregate.Recipe.Pictures[0].Entity.URL)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.Width, plannerRecipeAggregate.Recipe.Pictures[0].Entity.Width)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.Height, plannerRecipeAggregate.Recipe.Pictures[0].Entity.Height)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.Size, plannerRecipeAggregate.Recipe.Pictures[0].Entity.Size)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.Type, plannerRecipeAggregate.Recipe.Pictures[0].Entity.Type)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.Status, plannerRecipeAggregate.Recipe.Pictures[0].Entity.Status)

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(plannerRecipeAggregate)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())

				reflectPlannerRecipeAggregate := reflect.ValueOf(plannerRecipeAggregate)

				for i := 0; i < reflectPlannerRecipeAggregate.NumField(); i++ {
					assert.False(t, reflectPlannerRecipeAggregate.Field(i).IsZero())
				}
			},
		)
	}
}
