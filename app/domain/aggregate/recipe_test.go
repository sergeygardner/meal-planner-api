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

type testRecipe struct {
	Id          uuid.UUID
	UserId      uuid.UUID
	DateInsert  time.Time
	DateUpdate  time.Time
	Name        string
	Description string
	Notes       string
	Status      kind.RecipeStatus
}
type testRecipeCategory struct {
	Id         uuid.UUID
	UserId     uuid.UUID
	EntityId   uuid.UUID
	DeriveId   uuid.UUID
	DateInsert time.Time
	DateUpdate time.Time
	Status     kind.RecipeCategoryStatus
}
type testRecipeIngredient struct {
	Id         uuid.UUID
	UserId     uuid.UUID
	EntityId   uuid.UUID
	DeriveId   uuid.UUID
	DateInsert time.Time
	DateUpdate time.Time
	Name       string
	Status     kind.RecipeIngredientStatus
}
type testRecipeMeasure struct {
	Id         uuid.UUID
	UserId     uuid.UUID
	EntityId   uuid.UUID
	UnitId     uuid.UUID
	DateInsert time.Time
	DateUpdate time.Time
	Value      int64
	Status     kind.RecipeMeasureStatus
}
type testRecipeProcess struct {
	Id          uuid.UUID
	UserId      uuid.UUID
	EntityId    uuid.UUID
	DateInsert  time.Time
	DateUpdate  time.Time
	Name        string
	Description string
	Notes       string
	Status      kind.RecipeProcessStatus
}
type testPicture struct {
	AltNames []testAltName
	Entity   struct {
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
	}
}
type testAltName struct {
	Id         uuid.UUID
	UserId     uuid.UUID
	EntityId   uuid.UUID
	DateInsert time.Time
	DateUpdate time.Time
	Name       string
	Status     kind.AltNameStatus
}
type testCategory struct {
	Id         uuid.UUID
	UserId     uuid.UUID
	DateInsert time.Time
	DateUpdate time.Time
	Name       string
	Status     kind.CategoryStatus
}
type testIngredient struct {
	Id         uuid.UUID
	UserId     uuid.UUID
	DateInsert time.Time
	DateUpdate time.Time
	Name       string
	Status     kind.IngredientStatus
}
type testUnit struct {
	Id         uuid.UUID
	DateInsert time.Time
	DateUpdate time.Time
	Name       string
	Status     kind.UnitStatus
}
type testRecipeAggregate struct {
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
}

func TestRecipe(t *testing.T) {
	tests := []struct {
		name   string
		json   string
		Recipe testRecipeAggregate
	}{
		{
			name: "Test case with published recipe properties",
			json: "{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000005\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"categories\":[{\"derive\":{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000006\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000007\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000007\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Category\",\"status\":\"published\"},\"pictures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000008\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000009\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000009\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000007\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"published\"}}]},\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000010\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"derive_id\":\"00000000-0000-0000-0000-000000000006\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"status\":\"published\"}}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000004\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Recipe\",\"description\":\"Description\",\"notes\":\"Notes\",\"status\":\"published\"},\"ingredients\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000011\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000013\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"derive\":{\"id\":\"00000000-0000-0000-0000-000000000012\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Ingredient\",\"status\":\"published\"},\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000013\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"derive_id\":\"00000000-0000-0000-0000-000000000012\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"RecipeIngredient\",\"status\":\"published\"},\"measures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000014\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000015\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000015\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000013\",\"unit_id\":\"00000000-0000-0000-0000-000000000016\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"value\":42,\"status\":\"published\"},\"unit\":{\"id\":\"00000000-0000-0000-0000-000000000016\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Unit\",\"status\":\"published\"}}],\"pictures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000017\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000018\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000018\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000013\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"published\"}}]}],\"processes\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000019\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000020\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000020\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"RecipeProcess\",\"description\":\"Description\",\"notes\":\"Notes\",\"status\":\"published\"},\"pictures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000021\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000022\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000022\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000020\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"published\"}}]}],\"pictures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000023\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000024\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000024\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"published\"}}]}\n",
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
				recipeAggregate := Recipe{
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
				}

				assert.Equal(t, testCase.Recipe.AltNames[0].Id, recipeAggregate.AltNames[0].Id)
				assert.Equal(t, testCase.Recipe.AltNames[0].UserId, recipeAggregate.AltNames[0].UserId)
				assert.Equal(t, testCase.Recipe.AltNames[0].EntityId, recipeAggregate.AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipe.AltNames[0].DateInsert, recipeAggregate.AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipe.AltNames[0].DateUpdate, recipeAggregate.AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipe.AltNames[0].Name, recipeAggregate.AltNames[0].Name)
				assert.Equal(t, testCase.Recipe.AltNames[0].Status, recipeAggregate.AltNames[0].Status)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.AltNames[0].Id, recipeAggregate.Categories[0].Derive.AltNames[0].Id)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.AltNames[0].UserId, recipeAggregate.Categories[0].Derive.AltNames[0].UserId)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.AltNames[0].EntityId, recipeAggregate.Categories[0].Derive.AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.AltNames[0].DateInsert, recipeAggregate.Categories[0].Derive.AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.AltNames[0].DateUpdate, recipeAggregate.Categories[0].Derive.AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.AltNames[0].Name, recipeAggregate.Categories[0].Derive.AltNames[0].Name)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.AltNames[0].Status, recipeAggregate.Categories[0].Derive.AltNames[0].Status)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Entity.Id, recipeAggregate.Categories[0].Derive.Entity.Id)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Entity.UserId, recipeAggregate.Categories[0].Derive.Entity.UserId)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Entity.DateInsert, recipeAggregate.Categories[0].Derive.Entity.DateInsert)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Entity.DateUpdate, recipeAggregate.Categories[0].Derive.Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Entity.Name, recipeAggregate.Categories[0].Derive.Entity.Name)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Entity.Status, recipeAggregate.Categories[0].Derive.Entity.Status)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Id, recipeAggregate.Categories[0].Derive.Pictures[0].AltNames[0].Id)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].UserId, recipeAggregate.Categories[0].Derive.Pictures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].EntityId, recipeAggregate.Categories[0].Derive.Pictures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].DateInsert, recipeAggregate.Categories[0].Derive.Pictures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].DateUpdate, recipeAggregate.Categories[0].Derive.Pictures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Name, recipeAggregate.Categories[0].Derive.Pictures[0].AltNames[0].Name)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].AltNames[0].Status, recipeAggregate.Categories[0].Derive.Pictures[0].AltNames[0].Status)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.Id, recipeAggregate.Categories[0].Derive.Pictures[0].Entity.Id)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.UserId, recipeAggregate.Categories[0].Derive.Pictures[0].Entity.UserId)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.EntityId, recipeAggregate.Categories[0].Derive.Pictures[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.DateInsert, recipeAggregate.Categories[0].Derive.Pictures[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.DateUpdate, recipeAggregate.Categories[0].Derive.Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.DateUpdate, recipeAggregate.Categories[0].Derive.Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.Name, recipeAggregate.Categories[0].Derive.Pictures[0].Entity.Name)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.URL, recipeAggregate.Categories[0].Derive.Pictures[0].Entity.URL)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.Width, recipeAggregate.Categories[0].Derive.Pictures[0].Entity.Width)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.Height, recipeAggregate.Categories[0].Derive.Pictures[0].Entity.Height)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.Size, recipeAggregate.Categories[0].Derive.Pictures[0].Entity.Size)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.Type, recipeAggregate.Categories[0].Derive.Pictures[0].Entity.Type)
				assert.Equal(t, testCase.Recipe.Categories[0].Derive.Pictures[0].Entity.Status, recipeAggregate.Categories[0].Derive.Pictures[0].Entity.Status)
				assert.Equal(t, testCase.Recipe.Categories[0].Entity.Id, recipeAggregate.Categories[0].Entity.Id)
				assert.Equal(t, testCase.Recipe.Categories[0].Entity.UserId, recipeAggregate.Categories[0].Entity.UserId)
				assert.Equal(t, testCase.Recipe.Categories[0].Entity.EntityId, recipeAggregate.Categories[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipe.Categories[0].Entity.DeriveId, recipeAggregate.Categories[0].Entity.DeriveId)
				assert.Equal(t, testCase.Recipe.Categories[0].Entity.DateInsert, recipeAggregate.Categories[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipe.Categories[0].Entity.DateUpdate, recipeAggregate.Categories[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Categories[0].Entity.Status, recipeAggregate.Categories[0].Entity.Status)
				assert.Equal(t, testCase.Recipe.Entity.Id, recipeAggregate.Entity.Id)
				assert.Equal(t, testCase.Recipe.Entity.UserId, recipeAggregate.Entity.UserId)
				assert.Equal(t, testCase.Recipe.Entity.DateInsert, recipeAggregate.Entity.DateInsert)
				assert.Equal(t, testCase.Recipe.Entity.DateUpdate, recipeAggregate.Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Entity.Name, recipeAggregate.Entity.Name)
				assert.Equal(t, testCase.Recipe.Entity.Description, recipeAggregate.Entity.Description)
				assert.Equal(t, testCase.Recipe.Entity.Notes, recipeAggregate.Entity.Notes)
				assert.Equal(t, testCase.Recipe.Entity.Status, recipeAggregate.Entity.Status)
				assert.Equal(t, testCase.Recipe.Ingredients[0].AltNames[0].Id, recipeAggregate.Ingredients[0].AltNames[0].Id)
				assert.Equal(t, testCase.Recipe.Ingredients[0].AltNames[0].UserId, recipeAggregate.Ingredients[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].AltNames[0].EntityId, recipeAggregate.Ingredients[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].AltNames[0].DateInsert, recipeAggregate.Ingredients[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipe.Ingredients[0].AltNames[0].DateUpdate, recipeAggregate.Ingredients[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipe.Ingredients[0].AltNames[0].Name, recipeAggregate.Ingredients[0].AltNames[0].Name)
				assert.Equal(t, testCase.Recipe.Ingredients[0].AltNames[0].Status, recipeAggregate.Ingredients[0].AltNames[0].Status)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Derive.Id, recipeAggregate.Ingredients[0].Derive.Id)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Derive.UserId, recipeAggregate.Ingredients[0].Derive.UserId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Derive.DateInsert, recipeAggregate.Ingredients[0].Derive.DateInsert)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Derive.DateUpdate, recipeAggregate.Ingredients[0].Derive.DateUpdate)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Derive.Name, recipeAggregate.Ingredients[0].Derive.Name)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Derive.Status, recipeAggregate.Ingredients[0].Derive.Status)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Entity.Id, recipeAggregate.Ingredients[0].Entity.Id)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Entity.UserId, recipeAggregate.Ingredients[0].Entity.UserId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Entity.EntityId, recipeAggregate.Ingredients[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Entity.DeriveId, recipeAggregate.Ingredients[0].Entity.DeriveId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Entity.DateInsert, recipeAggregate.Ingredients[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Entity.DateUpdate, recipeAggregate.Ingredients[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Entity.Name, recipeAggregate.Ingredients[0].Entity.Name)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Entity.Status, recipeAggregate.Ingredients[0].Entity.Status)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].AltNames[0].Id, recipeAggregate.Ingredients[0].Measures[0].AltNames[0].Id)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].AltNames[0].UserId, recipeAggregate.Ingredients[0].Measures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].AltNames[0].EntityId, recipeAggregate.Ingredients[0].Measures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].AltNames[0].DateInsert, recipeAggregate.Ingredients[0].Measures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].AltNames[0].DateUpdate, recipeAggregate.Ingredients[0].Measures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].AltNames[0].Name, recipeAggregate.Ingredients[0].Measures[0].AltNames[0].Name)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].AltNames[0].Status, recipeAggregate.Ingredients[0].Measures[0].AltNames[0].Status)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].Entity.Id, recipeAggregate.Ingredients[0].Measures[0].Entity.Id)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].Entity.UserId, recipeAggregate.Ingredients[0].Measures[0].Entity.UserId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].Entity.EntityId, recipeAggregate.Ingredients[0].Measures[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].Entity.DateInsert, recipeAggregate.Ingredients[0].Measures[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].Entity.DateUpdate, recipeAggregate.Ingredients[0].Measures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].Entity.Value, recipeAggregate.Ingredients[0].Measures[0].Entity.Value)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].Entity.Status, recipeAggregate.Ingredients[0].Measures[0].Entity.Status)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].Unit.Id, recipeAggregate.Ingredients[0].Measures[0].Unit.Id)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].Unit.DateInsert, recipeAggregate.Ingredients[0].Measures[0].Unit.DateInsert)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].Unit.DateUpdate, recipeAggregate.Ingredients[0].Measures[0].Unit.DateUpdate)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].Unit.Name, recipeAggregate.Ingredients[0].Measures[0].Unit.Name)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Measures[0].Unit.Status, recipeAggregate.Ingredients[0].Measures[0].Unit.Status)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].AltNames[0].Id, recipeAggregate.Ingredients[0].Pictures[0].AltNames[0].Id)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].AltNames[0].UserId, recipeAggregate.Ingredients[0].Pictures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].AltNames[0].EntityId, recipeAggregate.Ingredients[0].Pictures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].AltNames[0].DateInsert, recipeAggregate.Ingredients[0].Pictures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].AltNames[0].DateUpdate, recipeAggregate.Ingredients[0].Pictures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].AltNames[0].Name, recipeAggregate.Ingredients[0].Pictures[0].AltNames[0].Name)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].AltNames[0].Status, recipeAggregate.Ingredients[0].Pictures[0].AltNames[0].Status)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.Id, recipeAggregate.Ingredients[0].Pictures[0].Entity.Id)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.UserId, recipeAggregate.Ingredients[0].Pictures[0].Entity.UserId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.EntityId, recipeAggregate.Ingredients[0].Pictures[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.DateInsert, recipeAggregate.Ingredients[0].Pictures[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.DateUpdate, recipeAggregate.Ingredients[0].Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.DateUpdate, recipeAggregate.Ingredients[0].Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.Name, recipeAggregate.Ingredients[0].Pictures[0].Entity.Name)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.URL, recipeAggregate.Ingredients[0].Pictures[0].Entity.URL)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.Width, recipeAggregate.Ingredients[0].Pictures[0].Entity.Width)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.Height, recipeAggregate.Ingredients[0].Pictures[0].Entity.Height)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.Size, recipeAggregate.Ingredients[0].Pictures[0].Entity.Size)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.Type, recipeAggregate.Ingredients[0].Pictures[0].Entity.Type)
				assert.Equal(t, testCase.Recipe.Ingredients[0].Pictures[0].Entity.Status, recipeAggregate.Ingredients[0].Pictures[0].Entity.Status)
				assert.Equal(t, testCase.Recipe.Processes[0].AltNames[0].Id, recipeAggregate.Processes[0].AltNames[0].Id)
				assert.Equal(t, testCase.Recipe.Processes[0].AltNames[0].UserId, recipeAggregate.Processes[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Recipe.Processes[0].AltNames[0].EntityId, recipeAggregate.Processes[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipe.Processes[0].AltNames[0].DateInsert, recipeAggregate.Processes[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipe.Processes[0].AltNames[0].DateUpdate, recipeAggregate.Processes[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipe.Processes[0].AltNames[0].Name, recipeAggregate.Processes[0].AltNames[0].Name)
				assert.Equal(t, testCase.Recipe.Processes[0].AltNames[0].Status, recipeAggregate.Processes[0].AltNames[0].Status)
				assert.Equal(t, testCase.Recipe.Processes[0].Entity.Id, recipeAggregate.Processes[0].Entity.Id)
				assert.Equal(t, testCase.Recipe.Processes[0].Entity.UserId, recipeAggregate.Processes[0].Entity.UserId)
				assert.Equal(t, testCase.Recipe.Processes[0].Entity.EntityId, recipeAggregate.Processes[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipe.Processes[0].Entity.DateInsert, recipeAggregate.Processes[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipe.Processes[0].Entity.DateUpdate, recipeAggregate.Processes[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Processes[0].Entity.Name, recipeAggregate.Processes[0].Entity.Name)
				assert.Equal(t, testCase.Recipe.Processes[0].Entity.Description, recipeAggregate.Processes[0].Entity.Description)
				assert.Equal(t, testCase.Recipe.Processes[0].Entity.Notes, recipeAggregate.Processes[0].Entity.Notes)
				assert.Equal(t, testCase.Recipe.Processes[0].Entity.Status, recipeAggregate.Processes[0].Entity.Status)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].AltNames[0].Id, recipeAggregate.Processes[0].Pictures[0].AltNames[0].Id)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].AltNames[0].UserId, recipeAggregate.Processes[0].Pictures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].AltNames[0].EntityId, recipeAggregate.Processes[0].Pictures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].AltNames[0].DateInsert, recipeAggregate.Processes[0].Pictures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].AltNames[0].DateUpdate, recipeAggregate.Processes[0].Pictures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].AltNames[0].Name, recipeAggregate.Processes[0].Pictures[0].AltNames[0].Name)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].AltNames[0].Status, recipeAggregate.Processes[0].Pictures[0].AltNames[0].Status)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.Id, recipeAggregate.Processes[0].Pictures[0].Entity.Id)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.UserId, recipeAggregate.Processes[0].Pictures[0].Entity.UserId)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.EntityId, recipeAggregate.Processes[0].Pictures[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.DateInsert, recipeAggregate.Processes[0].Pictures[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.DateUpdate, recipeAggregate.Processes[0].Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.DateUpdate, recipeAggregate.Processes[0].Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.Name, recipeAggregate.Processes[0].Pictures[0].Entity.Name)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.URL, recipeAggregate.Processes[0].Pictures[0].Entity.URL)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.Width, recipeAggregate.Processes[0].Pictures[0].Entity.Width)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.Height, recipeAggregate.Processes[0].Pictures[0].Entity.Height)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.Size, recipeAggregate.Processes[0].Pictures[0].Entity.Size)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.Type, recipeAggregate.Processes[0].Pictures[0].Entity.Type)
				assert.Equal(t, testCase.Recipe.Processes[0].Pictures[0].Entity.Status, recipeAggregate.Processes[0].Pictures[0].Entity.Status)
				assert.Equal(t, testCase.Recipe.Pictures[0].AltNames[0].Id, recipeAggregate.Pictures[0].AltNames[0].Id)
				assert.Equal(t, testCase.Recipe.Pictures[0].AltNames[0].UserId, recipeAggregate.Pictures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Recipe.Pictures[0].AltNames[0].EntityId, recipeAggregate.Pictures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Recipe.Pictures[0].AltNames[0].DateInsert, recipeAggregate.Pictures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Recipe.Pictures[0].AltNames[0].DateUpdate, recipeAggregate.Pictures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Recipe.Pictures[0].AltNames[0].Name, recipeAggregate.Pictures[0].AltNames[0].Name)
				assert.Equal(t, testCase.Recipe.Pictures[0].AltNames[0].Status, recipeAggregate.Pictures[0].AltNames[0].Status)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.Id, recipeAggregate.Pictures[0].Entity.Id)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.UserId, recipeAggregate.Pictures[0].Entity.UserId)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.EntityId, recipeAggregate.Pictures[0].Entity.EntityId)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.DateInsert, recipeAggregate.Pictures[0].Entity.DateInsert)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.DateUpdate, recipeAggregate.Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.DateUpdate, recipeAggregate.Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.Name, recipeAggregate.Pictures[0].Entity.Name)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.URL, recipeAggregate.Pictures[0].Entity.URL)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.Width, recipeAggregate.Pictures[0].Entity.Width)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.Height, recipeAggregate.Pictures[0].Entity.Height)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.Size, recipeAggregate.Pictures[0].Entity.Size)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.Type, recipeAggregate.Pictures[0].Entity.Type)
				assert.Equal(t, testCase.Recipe.Pictures[0].Entity.Status, recipeAggregate.Pictures[0].Entity.Status)

				reflectRecipeAggregate := reflect.ValueOf(recipeAggregate)

				for i := 0; i < reflectRecipeAggregate.NumField(); i++ {
					assert.False(t, reflectRecipeAggregate.Field(i).IsZero())
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(recipeAggregate)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}
func TestRecipeCategory(t *testing.T) {
	tests := []struct {
		name           string
		json           string
		RecipeCategory struct {
			Derive struct {
				AltNames []testAltName
				Entity   testCategory
				Pictures []testPicture
			}
			Entity testRecipeCategory
		}
	}{
		{
			name: "Test case with published recipe category properties",
			json: "{\"derive\":{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000006\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000007\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000007\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Category\",\"status\":\"published\"},\"pictures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000008\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000009\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000009\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000007\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"published\"}}]},\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000010\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"derive_id\":\"00000000-0000-0000-0000-000000000006\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"status\":\"published\"}}\n",
			RecipeCategory: struct {
				Derive struct {
					AltNames []testAltName
					Entity   testCategory
					Pictures []testPicture
				}
				Entity testRecipeCategory
			}{
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
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				recipeCategoryAggregate := RecipeCategory{
					Derive: &Category{
						AltNames: []*DomainEntity.AltName{
							{
								Id:         testCase.RecipeCategory.Derive.AltNames[0].Id,
								UserId:     testCase.RecipeCategory.Derive.AltNames[0].UserId,
								EntityId:   testCase.RecipeCategory.Derive.AltNames[0].EntityId,
								DateInsert: testCase.RecipeCategory.Derive.AltNames[0].DateInsert,
								DateUpdate: testCase.RecipeCategory.Derive.AltNames[0].DateUpdate,
								Name:       testCase.RecipeCategory.Derive.AltNames[0].Name,
								Status:     testCase.RecipeCategory.Derive.AltNames[0].Status,
							},
						},
						Entity: &DomainEntity.Category{
							Id:         testCase.RecipeCategory.Derive.Entity.Id,
							UserId:     testCase.RecipeCategory.Derive.Entity.UserId,
							DateInsert: testCase.RecipeCategory.Derive.Entity.DateInsert,
							DateUpdate: testCase.RecipeCategory.Derive.Entity.DateUpdate,
							Name:       testCase.RecipeCategory.Derive.Entity.Name,
							Status:     testCase.RecipeCategory.Derive.Entity.Status,
						},
						Pictures: []*Picture{
							{
								AltNames: []*DomainEntity.AltName{
									{
										Id:         testCase.RecipeCategory.Derive.Pictures[0].AltNames[0].Id,
										UserId:     testCase.RecipeCategory.Derive.Pictures[0].AltNames[0].UserId,
										EntityId:   testCase.RecipeCategory.Derive.Pictures[0].AltNames[0].EntityId,
										DateInsert: testCase.RecipeCategory.Derive.Pictures[0].AltNames[0].DateInsert,
										DateUpdate: testCase.RecipeCategory.Derive.Pictures[0].AltNames[0].DateUpdate,
										Name:       testCase.RecipeCategory.Derive.Pictures[0].AltNames[0].Name,
										Status:     testCase.RecipeCategory.Derive.Pictures[0].AltNames[0].Status,
									},
								},
								Entity: &DomainEntity.Picture{
									Id:         testCase.RecipeCategory.Derive.Pictures[0].Entity.Id,
									UserId:     testCase.RecipeCategory.Derive.Pictures[0].Entity.UserId,
									EntityId:   testCase.RecipeCategory.Derive.Pictures[0].Entity.EntityId,
									DateInsert: testCase.RecipeCategory.Derive.Pictures[0].Entity.DateInsert,
									DateUpdate: testCase.RecipeCategory.Derive.Pictures[0].Entity.DateUpdate,
									Name:       testCase.RecipeCategory.Derive.Pictures[0].Entity.Name,
									URL:        testCase.RecipeCategory.Derive.Pictures[0].Entity.URL,
									Width:      testCase.RecipeCategory.Derive.Pictures[0].Entity.Width,
									Height:     testCase.RecipeCategory.Derive.Pictures[0].Entity.Height,
									Size:       testCase.RecipeCategory.Derive.Pictures[0].Entity.Size,
									Type:       testCase.RecipeCategory.Derive.Pictures[0].Entity.Type,
									Status:     testCase.RecipeCategory.Derive.Pictures[0].Entity.Status,
								},
							},
						},
					},
					Entity: &DomainEntity.RecipeCategory{
						Id:         testCase.RecipeCategory.Entity.Id,
						UserId:     testCase.RecipeCategory.Entity.UserId,
						EntityId:   testCase.RecipeCategory.Entity.EntityId,
						DeriveId:   testCase.RecipeCategory.Entity.DeriveId,
						DateInsert: testCase.RecipeCategory.Entity.DateInsert,
						DateUpdate: testCase.RecipeCategory.Entity.DateUpdate,
						Status:     testCase.RecipeCategory.Entity.Status,
					},
				}

				assert.Equal(t, testCase.RecipeCategory.Derive.AltNames[0].Id, recipeCategoryAggregate.Derive.AltNames[0].Id)
				assert.Equal(t, testCase.RecipeCategory.Derive.AltNames[0].UserId, recipeCategoryAggregate.Derive.AltNames[0].UserId)
				assert.Equal(t, testCase.RecipeCategory.Derive.AltNames[0].EntityId, recipeCategoryAggregate.Derive.AltNames[0].EntityId)
				assert.Equal(t, testCase.RecipeCategory.Derive.AltNames[0].DateInsert, recipeCategoryAggregate.Derive.AltNames[0].DateInsert)
				assert.Equal(t, testCase.RecipeCategory.Derive.AltNames[0].DateUpdate, recipeCategoryAggregate.Derive.AltNames[0].DateUpdate)
				assert.Equal(t, testCase.RecipeCategory.Derive.AltNames[0].Name, recipeCategoryAggregate.Derive.AltNames[0].Name)
				assert.Equal(t, testCase.RecipeCategory.Derive.AltNames[0].Status, recipeCategoryAggregate.Derive.AltNames[0].Status)
				assert.Equal(t, testCase.RecipeCategory.Derive.Entity.Id, recipeCategoryAggregate.Derive.Entity.Id)
				assert.Equal(t, testCase.RecipeCategory.Derive.Entity.UserId, recipeCategoryAggregate.Derive.Entity.UserId)
				assert.Equal(t, testCase.RecipeCategory.Derive.Entity.DateInsert, recipeCategoryAggregate.Derive.Entity.DateInsert)
				assert.Equal(t, testCase.RecipeCategory.Derive.Entity.DateUpdate, recipeCategoryAggregate.Derive.Entity.DateUpdate)
				assert.Equal(t, testCase.RecipeCategory.Derive.Entity.Name, recipeCategoryAggregate.Derive.Entity.Name)
				assert.Equal(t, testCase.RecipeCategory.Derive.Entity.Status, recipeCategoryAggregate.Derive.Entity.Status)
				assert.Equal(t, testCase.RecipeCategory.Derive.Pictures[0].AltNames[0].Id, recipeCategoryAggregate.Derive.Pictures[0].AltNames[0].Id)
				assert.Equal(t, testCase.RecipeCategory.Derive.Pictures[0].AltNames[0].UserId, recipeCategoryAggregate.Derive.Pictures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.RecipeCategory.Derive.Pictures[0].AltNames[0].EntityId, recipeCategoryAggregate.Derive.Pictures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.RecipeCategory.Derive.Pictures[0].AltNames[0].DateInsert, recipeCategoryAggregate.Derive.Pictures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.RecipeCategory.Derive.Pictures[0].AltNames[0].DateUpdate, recipeCategoryAggregate.Derive.Pictures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.RecipeCategory.Derive.Pictures[0].AltNames[0].Name, recipeCategoryAggregate.Derive.Pictures[0].AltNames[0].Name)
				assert.Equal(t, testCase.RecipeCategory.Derive.Pictures[0].AltNames[0].Status, recipeCategoryAggregate.Derive.Pictures[0].AltNames[0].Status)
				assert.Equal(t, testCase.RecipeCategory.Derive.Pictures[0].Entity.Id, recipeCategoryAggregate.Derive.Pictures[0].Entity.Id)
				assert.Equal(t, testCase.RecipeCategory.Derive.Pictures[0].Entity.UserId, recipeCategoryAggregate.Derive.Pictures[0].Entity.UserId)
				assert.Equal(t, testCase.RecipeCategory.Derive.Pictures[0].Entity.EntityId, recipeCategoryAggregate.Derive.Pictures[0].Entity.EntityId)
				assert.Equal(t, testCase.RecipeCategory.Derive.Pictures[0].Entity.DateInsert, recipeCategoryAggregate.Derive.Pictures[0].Entity.DateInsert)
				assert.Equal(t, testCase.RecipeCategory.Derive.Pictures[0].Entity.DateUpdate, recipeCategoryAggregate.Derive.Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.RecipeCategory.Derive.Pictures[0].Entity.DateUpdate, recipeCategoryAggregate.Derive.Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.RecipeCategory.Derive.Pictures[0].Entity.Name, recipeCategoryAggregate.Derive.Pictures[0].Entity.Name)
				assert.Equal(t, testCase.RecipeCategory.Derive.Pictures[0].Entity.URL, recipeCategoryAggregate.Derive.Pictures[0].Entity.URL)
				assert.Equal(t, testCase.RecipeCategory.Derive.Pictures[0].Entity.Width, recipeCategoryAggregate.Derive.Pictures[0].Entity.Width)
				assert.Equal(t, testCase.RecipeCategory.Derive.Pictures[0].Entity.Height, recipeCategoryAggregate.Derive.Pictures[0].Entity.Height)
				assert.Equal(t, testCase.RecipeCategory.Derive.Pictures[0].Entity.Size, recipeCategoryAggregate.Derive.Pictures[0].Entity.Size)
				assert.Equal(t, testCase.RecipeCategory.Derive.Pictures[0].Entity.Type, recipeCategoryAggregate.Derive.Pictures[0].Entity.Type)
				assert.Equal(t, testCase.RecipeCategory.Derive.Pictures[0].Entity.Status, recipeCategoryAggregate.Derive.Pictures[0].Entity.Status)
				assert.Equal(t, testCase.RecipeCategory.Entity.Id, recipeCategoryAggregate.Entity.Id)
				assert.Equal(t, testCase.RecipeCategory.Entity.UserId, recipeCategoryAggregate.Entity.UserId)
				assert.Equal(t, testCase.RecipeCategory.Entity.EntityId, recipeCategoryAggregate.Entity.EntityId)
				assert.Equal(t, testCase.RecipeCategory.Entity.DeriveId, recipeCategoryAggregate.Entity.DeriveId)
				assert.Equal(t, testCase.RecipeCategory.Entity.DateInsert, recipeCategoryAggregate.Entity.DateInsert)
				assert.Equal(t, testCase.RecipeCategory.Entity.DateUpdate, recipeCategoryAggregate.Entity.DateUpdate)
				assert.Equal(t, testCase.RecipeCategory.Entity.Status, recipeCategoryAggregate.Entity.Status)

				reflectRecipeAggregate := reflect.ValueOf(recipeCategoryAggregate)

				for i := 0; i < reflectRecipeAggregate.NumField(); i++ {
					assert.False(t, reflectRecipeAggregate.Field(i).IsZero())
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(recipeCategoryAggregate)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}
func TestRecipeIngredient(t *testing.T) {
	tests := []struct {
		name             string
		json             string
		RecipeIngredient struct {
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
	}{
		{
			name: "Test case with published recipe ingredient properties",
			json: "{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000011\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000013\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"derive\":{\"id\":\"00000000-0000-0000-0000-000000000012\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Ingredient\",\"status\":\"published\"},\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000013\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"derive_id\":\"00000000-0000-0000-0000-000000000012\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"RecipeIngredient\",\"status\":\"published\"},\"measures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000014\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000015\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000015\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000013\",\"unit_id\":\"00000000-0000-0000-0000-000000000016\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"value\":42,\"status\":\"published\"},\"unit\":{\"id\":\"00000000-0000-0000-0000-000000000016\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Unit\",\"status\":\"published\"}}],\"pictures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000017\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000018\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000018\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000013\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"published\"}}]}\n",
			RecipeIngredient: struct {
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
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				recipeIngredientAggregate := RecipeIngredient{
					AltNames: []*DomainEntity.AltName{
						{
							Id:         testCase.RecipeIngredient.AltNames[0].Id,
							UserId:     testCase.RecipeIngredient.AltNames[0].UserId,
							EntityId:   testCase.RecipeIngredient.AltNames[0].EntityId,
							DateInsert: testCase.RecipeIngredient.AltNames[0].DateInsert,
							DateUpdate: testCase.RecipeIngredient.AltNames[0].DateUpdate,
							Name:       testCase.RecipeIngredient.AltNames[0].Name,
							Status:     testCase.RecipeIngredient.AltNames[0].Status,
						},
					},
					Derive: &DomainEntity.Ingredient{
						Id:         testCase.RecipeIngredient.Derive.Id,
						UserId:     testCase.RecipeIngredient.Derive.UserId,
						DateInsert: testCase.RecipeIngredient.Derive.DateInsert,
						DateUpdate: testCase.RecipeIngredient.Derive.DateUpdate,
						Name:       testCase.RecipeIngredient.Derive.Name,
						Status:     testCase.RecipeIngredient.Derive.Status,
					},
					Entity: &DomainEntity.RecipeIngredient{
						Id:         testCase.RecipeIngredient.Entity.Id,
						UserId:     testCase.RecipeIngredient.Entity.UserId,
						EntityId:   testCase.RecipeIngredient.Entity.EntityId,
						DeriveId:   testCase.RecipeIngredient.Entity.DeriveId,
						DateInsert: testCase.RecipeIngredient.Entity.DateInsert,
						DateUpdate: testCase.RecipeIngredient.Entity.DateUpdate,
						Name:       testCase.RecipeIngredient.Entity.Name,
						Status:     testCase.RecipeIngredient.Entity.Status,
					},
					Measures: []*RecipeMeasure{
						{
							AltNames: []*DomainEntity.AltName{
								{
									Id:         testCase.RecipeIngredient.Measures[0].AltNames[0].Id,
									UserId:     testCase.RecipeIngredient.Measures[0].AltNames[0].UserId,
									EntityId:   testCase.RecipeIngredient.Measures[0].AltNames[0].EntityId,
									DateInsert: testCase.RecipeIngredient.Measures[0].AltNames[0].DateInsert,
									DateUpdate: testCase.RecipeIngredient.Measures[0].AltNames[0].DateUpdate,
									Name:       testCase.RecipeIngredient.Measures[0].AltNames[0].Name,
									Status:     testCase.RecipeIngredient.Measures[0].AltNames[0].Status,
								},
							},
							Entity: &DomainEntity.RecipeMeasure{
								Id:         testCase.RecipeIngredient.Measures[0].Entity.Id,
								UserId:     testCase.RecipeIngredient.Measures[0].Entity.UserId,
								EntityId:   testCase.RecipeIngredient.Measures[0].Entity.EntityId,
								UnitId:     testCase.RecipeIngredient.Measures[0].Entity.UnitId,
								DateInsert: testCase.RecipeIngredient.Measures[0].Entity.DateInsert,
								DateUpdate: testCase.RecipeIngredient.Measures[0].Entity.DateUpdate,
								Value:      testCase.RecipeIngredient.Measures[0].Entity.Value,
								Status:     testCase.RecipeIngredient.Measures[0].Entity.Status,
							},
							Unit: &DomainEntity.Unit{
								Id:         testCase.RecipeIngredient.Measures[0].Unit.Id,
								DateInsert: testCase.RecipeIngredient.Measures[0].Unit.DateInsert,
								DateUpdate: testCase.RecipeIngredient.Measures[0].Unit.DateUpdate,
								Name:       testCase.RecipeIngredient.Measures[0].Unit.Name,
								Status:     testCase.RecipeIngredient.Measures[0].Unit.Status,
							},
						},
					},
					Pictures: []*Picture{
						{
							AltNames: []*DomainEntity.AltName{
								{
									Id:         testCase.RecipeIngredient.Pictures[0].AltNames[0].Id,
									UserId:     testCase.RecipeIngredient.Pictures[0].AltNames[0].UserId,
									EntityId:   testCase.RecipeIngredient.Pictures[0].AltNames[0].EntityId,
									DateInsert: testCase.RecipeIngredient.Pictures[0].AltNames[0].DateInsert,
									DateUpdate: testCase.RecipeIngredient.Pictures[0].AltNames[0].DateUpdate,
									Name:       testCase.RecipeIngredient.Pictures[0].AltNames[0].Name,
									Status:     testCase.RecipeIngredient.Pictures[0].AltNames[0].Status,
								},
							},
							Entity: &DomainEntity.Picture{
								Id:         testCase.RecipeIngredient.Pictures[0].Entity.Id,
								UserId:     testCase.RecipeIngredient.Pictures[0].Entity.UserId,
								EntityId:   testCase.RecipeIngredient.Pictures[0].Entity.EntityId,
								DateInsert: testCase.RecipeIngredient.Pictures[0].Entity.DateInsert,
								DateUpdate: testCase.RecipeIngredient.Pictures[0].Entity.DateUpdate,
								Name:       testCase.RecipeIngredient.Pictures[0].Entity.Name,
								URL:        testCase.RecipeIngredient.Pictures[0].Entity.URL,
								Width:      testCase.RecipeIngredient.Pictures[0].Entity.Width,
								Height:     testCase.RecipeIngredient.Pictures[0].Entity.Height,
								Size:       testCase.RecipeIngredient.Pictures[0].Entity.Size,
								Type:       testCase.RecipeIngredient.Pictures[0].Entity.Type,
								Status:     testCase.RecipeIngredient.Pictures[0].Entity.Status,
							},
						},
					},
				}

				assert.Equal(t, testCase.RecipeIngredient.AltNames[0].Id, recipeIngredientAggregate.AltNames[0].Id)
				assert.Equal(t, testCase.RecipeIngredient.AltNames[0].UserId, recipeIngredientAggregate.AltNames[0].UserId)
				assert.Equal(t, testCase.RecipeIngredient.AltNames[0].EntityId, recipeIngredientAggregate.AltNames[0].EntityId)
				assert.Equal(t, testCase.RecipeIngredient.AltNames[0].DateInsert, recipeIngredientAggregate.AltNames[0].DateInsert)
				assert.Equal(t, testCase.RecipeIngredient.AltNames[0].DateUpdate, recipeIngredientAggregate.AltNames[0].DateUpdate)
				assert.Equal(t, testCase.RecipeIngredient.AltNames[0].Name, recipeIngredientAggregate.AltNames[0].Name)
				assert.Equal(t, testCase.RecipeIngredient.AltNames[0].Status, recipeIngredientAggregate.AltNames[0].Status)
				assert.Equal(t, testCase.RecipeIngredient.Derive.Id, recipeIngredientAggregate.Derive.Id)
				assert.Equal(t, testCase.RecipeIngredient.Derive.UserId, recipeIngredientAggregate.Derive.UserId)
				assert.Equal(t, testCase.RecipeIngredient.Derive.DateInsert, recipeIngredientAggregate.Derive.DateInsert)
				assert.Equal(t, testCase.RecipeIngredient.Derive.DateUpdate, recipeIngredientAggregate.Derive.DateUpdate)
				assert.Equal(t, testCase.RecipeIngredient.Derive.Name, recipeIngredientAggregate.Derive.Name)
				assert.Equal(t, testCase.RecipeIngredient.Derive.Status, recipeIngredientAggregate.Derive.Status)
				assert.Equal(t, testCase.RecipeIngredient.Entity.Id, recipeIngredientAggregate.Entity.Id)
				assert.Equal(t, testCase.RecipeIngredient.Entity.UserId, recipeIngredientAggregate.Entity.UserId)
				assert.Equal(t, testCase.RecipeIngredient.Entity.EntityId, recipeIngredientAggregate.Entity.EntityId)
				assert.Equal(t, testCase.RecipeIngredient.Entity.DeriveId, recipeIngredientAggregate.Entity.DeriveId)
				assert.Equal(t, testCase.RecipeIngredient.Entity.DateInsert, recipeIngredientAggregate.Entity.DateInsert)
				assert.Equal(t, testCase.RecipeIngredient.Entity.DateUpdate, recipeIngredientAggregate.Entity.DateUpdate)
				assert.Equal(t, testCase.RecipeIngredient.Entity.Name, recipeIngredientAggregate.Entity.Name)
				assert.Equal(t, testCase.RecipeIngredient.Entity.Status, recipeIngredientAggregate.Entity.Status)
				assert.Equal(t, testCase.RecipeIngredient.Measures[0].AltNames[0].Id, recipeIngredientAggregate.Measures[0].AltNames[0].Id)
				assert.Equal(t, testCase.RecipeIngredient.Measures[0].AltNames[0].UserId, recipeIngredientAggregate.Measures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.RecipeIngredient.Measures[0].AltNames[0].EntityId, recipeIngredientAggregate.Measures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.RecipeIngredient.Measures[0].AltNames[0].DateInsert, recipeIngredientAggregate.Measures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.RecipeIngredient.Measures[0].AltNames[0].DateUpdate, recipeIngredientAggregate.Measures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.RecipeIngredient.Measures[0].AltNames[0].Name, recipeIngredientAggregate.Measures[0].AltNames[0].Name)
				assert.Equal(t, testCase.RecipeIngredient.Measures[0].AltNames[0].Status, recipeIngredientAggregate.Measures[0].AltNames[0].Status)
				assert.Equal(t, testCase.RecipeIngredient.Measures[0].Entity.Id, recipeIngredientAggregate.Measures[0].Entity.Id)
				assert.Equal(t, testCase.RecipeIngredient.Measures[0].Entity.UserId, recipeIngredientAggregate.Measures[0].Entity.UserId)
				assert.Equal(t, testCase.RecipeIngredient.Measures[0].Entity.EntityId, recipeIngredientAggregate.Measures[0].Entity.EntityId)
				assert.Equal(t, testCase.RecipeIngredient.Measures[0].Entity.DateInsert, recipeIngredientAggregate.Measures[0].Entity.DateInsert)
				assert.Equal(t, testCase.RecipeIngredient.Measures[0].Entity.DateUpdate, recipeIngredientAggregate.Measures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.RecipeIngredient.Measures[0].Entity.Value, recipeIngredientAggregate.Measures[0].Entity.Value)
				assert.Equal(t, testCase.RecipeIngredient.Measures[0].Entity.Status, recipeIngredientAggregate.Measures[0].Entity.Status)
				assert.Equal(t, testCase.RecipeIngredient.Measures[0].Unit.Id, recipeIngredientAggregate.Measures[0].Unit.Id)
				assert.Equal(t, testCase.RecipeIngredient.Measures[0].Unit.DateInsert, recipeIngredientAggregate.Measures[0].Unit.DateInsert)
				assert.Equal(t, testCase.RecipeIngredient.Measures[0].Unit.DateUpdate, recipeIngredientAggregate.Measures[0].Unit.DateUpdate)
				assert.Equal(t, testCase.RecipeIngredient.Measures[0].Unit.Name, recipeIngredientAggregate.Measures[0].Unit.Name)
				assert.Equal(t, testCase.RecipeIngredient.Measures[0].Unit.Status, recipeIngredientAggregate.Measures[0].Unit.Status)
				assert.Equal(t, testCase.RecipeIngredient.Pictures[0].AltNames[0].Id, recipeIngredientAggregate.Pictures[0].AltNames[0].Id)
				assert.Equal(t, testCase.RecipeIngredient.Pictures[0].AltNames[0].UserId, recipeIngredientAggregate.Pictures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.RecipeIngredient.Pictures[0].AltNames[0].EntityId, recipeIngredientAggregate.Pictures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.RecipeIngredient.Pictures[0].AltNames[0].DateInsert, recipeIngredientAggregate.Pictures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.RecipeIngredient.Pictures[0].AltNames[0].DateUpdate, recipeIngredientAggregate.Pictures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.RecipeIngredient.Pictures[0].AltNames[0].Name, recipeIngredientAggregate.Pictures[0].AltNames[0].Name)
				assert.Equal(t, testCase.RecipeIngredient.Pictures[0].AltNames[0].Status, recipeIngredientAggregate.Pictures[0].AltNames[0].Status)
				assert.Equal(t, testCase.RecipeIngredient.Pictures[0].Entity.Id, recipeIngredientAggregate.Pictures[0].Entity.Id)
				assert.Equal(t, testCase.RecipeIngredient.Pictures[0].Entity.UserId, recipeIngredientAggregate.Pictures[0].Entity.UserId)
				assert.Equal(t, testCase.RecipeIngredient.Pictures[0].Entity.EntityId, recipeIngredientAggregate.Pictures[0].Entity.EntityId)
				assert.Equal(t, testCase.RecipeIngredient.Pictures[0].Entity.DateInsert, recipeIngredientAggregate.Pictures[0].Entity.DateInsert)
				assert.Equal(t, testCase.RecipeIngredient.Pictures[0].Entity.DateUpdate, recipeIngredientAggregate.Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.RecipeIngredient.Pictures[0].Entity.DateUpdate, recipeIngredientAggregate.Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.RecipeIngredient.Pictures[0].Entity.Name, recipeIngredientAggregate.Pictures[0].Entity.Name)
				assert.Equal(t, testCase.RecipeIngredient.Pictures[0].Entity.URL, recipeIngredientAggregate.Pictures[0].Entity.URL)
				assert.Equal(t, testCase.RecipeIngredient.Pictures[0].Entity.Width, recipeIngredientAggregate.Pictures[0].Entity.Width)
				assert.Equal(t, testCase.RecipeIngredient.Pictures[0].Entity.Height, recipeIngredientAggregate.Pictures[0].Entity.Height)
				assert.Equal(t, testCase.RecipeIngredient.Pictures[0].Entity.Size, recipeIngredientAggregate.Pictures[0].Entity.Size)
				assert.Equal(t, testCase.RecipeIngredient.Pictures[0].Entity.Type, recipeIngredientAggregate.Pictures[0].Entity.Type)
				assert.Equal(t, testCase.RecipeIngredient.Pictures[0].Entity.Status, recipeIngredientAggregate.Pictures[0].Entity.Status)

				reflectRecipeAggregate := reflect.ValueOf(recipeIngredientAggregate)

				for i := 0; i < reflectRecipeAggregate.NumField(); i++ {
					assert.False(t, reflectRecipeAggregate.Field(i).IsZero())
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(recipeIngredientAggregate)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}
func TestRecipeMeasure(t *testing.T) {
	tests := []struct {
		name          string
		json          string
		RecipeMeasure struct {
			AltNames []testAltName
			Entity   testRecipeMeasure
			Unit     testUnit
		}
	}{
		{
			name: "Test case with published recipe measure properties",
			json: "{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000014\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000015\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000015\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000013\",\"unit_id\":\"00000000-0000-0000-0000-000000000016\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"value\":42,\"status\":\"published\"},\"unit\":{\"id\":\"00000000-0000-0000-0000-000000000016\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Unit\",\"status\":\"published\"}}\n",
			RecipeMeasure: struct {
				AltNames []testAltName
				Entity   testRecipeMeasure
				Unit     testUnit
			}{
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
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				recipeMeasureAggregate := RecipeMeasure{
					AltNames: []*DomainEntity.AltName{
						{
							Id:         testCase.RecipeMeasure.AltNames[0].Id,
							UserId:     testCase.RecipeMeasure.AltNames[0].UserId,
							EntityId:   testCase.RecipeMeasure.AltNames[0].EntityId,
							DateInsert: testCase.RecipeMeasure.AltNames[0].DateInsert,
							DateUpdate: testCase.RecipeMeasure.AltNames[0].DateUpdate,
							Name:       testCase.RecipeMeasure.AltNames[0].Name,
							Status:     testCase.RecipeMeasure.AltNames[0].Status,
						},
					},
					Entity: &DomainEntity.RecipeMeasure{
						Id:         testCase.RecipeMeasure.Entity.Id,
						UserId:     testCase.RecipeMeasure.Entity.UserId,
						EntityId:   testCase.RecipeMeasure.Entity.EntityId,
						UnitId:     testCase.RecipeMeasure.Entity.UnitId,
						DateInsert: testCase.RecipeMeasure.Entity.DateInsert,
						DateUpdate: testCase.RecipeMeasure.Entity.DateUpdate,
						Value:      testCase.RecipeMeasure.Entity.Value,
						Status:     testCase.RecipeMeasure.Entity.Status,
					},
					Unit: &DomainEntity.Unit{
						Id:         testCase.RecipeMeasure.Unit.Id,
						DateInsert: testCase.RecipeMeasure.Unit.DateInsert,
						DateUpdate: testCase.RecipeMeasure.Unit.DateUpdate,
						Name:       testCase.RecipeMeasure.Unit.Name,
						Status:     testCase.RecipeMeasure.Unit.Status,
					},
				}

				assert.Equal(t, testCase.RecipeMeasure.AltNames[0].Id, recipeMeasureAggregate.AltNames[0].Id)
				assert.Equal(t, testCase.RecipeMeasure.AltNames[0].UserId, recipeMeasureAggregate.AltNames[0].UserId)
				assert.Equal(t, testCase.RecipeMeasure.AltNames[0].EntityId, recipeMeasureAggregate.AltNames[0].EntityId)
				assert.Equal(t, testCase.RecipeMeasure.AltNames[0].DateInsert, recipeMeasureAggregate.AltNames[0].DateInsert)
				assert.Equal(t, testCase.RecipeMeasure.AltNames[0].DateUpdate, recipeMeasureAggregate.AltNames[0].DateUpdate)
				assert.Equal(t, testCase.RecipeMeasure.AltNames[0].Name, recipeMeasureAggregate.AltNames[0].Name)
				assert.Equal(t, testCase.RecipeMeasure.AltNames[0].Status, recipeMeasureAggregate.AltNames[0].Status)
				assert.Equal(t, testCase.RecipeMeasure.Entity.Id, recipeMeasureAggregate.Entity.Id)
				assert.Equal(t, testCase.RecipeMeasure.Entity.UserId, recipeMeasureAggregate.Entity.UserId)
				assert.Equal(t, testCase.RecipeMeasure.Entity.EntityId, recipeMeasureAggregate.Entity.EntityId)
				assert.Equal(t, testCase.RecipeMeasure.Entity.DateInsert, recipeMeasureAggregate.Entity.DateInsert)
				assert.Equal(t, testCase.RecipeMeasure.Entity.DateUpdate, recipeMeasureAggregate.Entity.DateUpdate)
				assert.Equal(t, testCase.RecipeMeasure.Entity.Value, recipeMeasureAggregate.Entity.Value)
				assert.Equal(t, testCase.RecipeMeasure.Entity.Status, recipeMeasureAggregate.Entity.Status)
				assert.Equal(t, testCase.RecipeMeasure.Unit.Id, recipeMeasureAggregate.Unit.Id)
				assert.Equal(t, testCase.RecipeMeasure.Unit.DateInsert, recipeMeasureAggregate.Unit.DateInsert)
				assert.Equal(t, testCase.RecipeMeasure.Unit.DateUpdate, recipeMeasureAggregate.Unit.DateUpdate)
				assert.Equal(t, testCase.RecipeMeasure.Unit.Name, recipeMeasureAggregate.Unit.Name)
				assert.Equal(t, testCase.RecipeMeasure.Unit.Status, recipeMeasureAggregate.Unit.Status)

				reflectRecipeAggregate := reflect.ValueOf(recipeMeasureAggregate)

				for i := 0; i < reflectRecipeAggregate.NumField(); i++ {
					assert.False(t, reflectRecipeAggregate.Field(i).IsZero())
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(recipeMeasureAggregate)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}
func TestRecipeProcess(t *testing.T) {
	tests := []struct {
		name          string
		json          string
		RecipeProcess struct {
			AltNames []testAltName
			Entity   testRecipeProcess
			Pictures []testPicture
		}
	}{
		{
			name: "Test case with published recipe process properties",
			json: "{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000019\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000020\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000020\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"RecipeProcess\",\"description\":\"Description\",\"notes\":\"Notes\",\"status\":\"published\"},\"pictures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000021\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000022\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000022\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000020\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"published\"}}]}\n",
			RecipeProcess: struct {
				AltNames []testAltName
				Entity   testRecipeProcess
				Pictures []testPicture
			}{
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
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				recipeProcessAggregate := RecipeProcess{
					AltNames: []*DomainEntity.AltName{
						{
							Id:         testCase.RecipeProcess.AltNames[0].Id,
							UserId:     testCase.RecipeProcess.AltNames[0].UserId,
							EntityId:   testCase.RecipeProcess.AltNames[0].EntityId,
							DateInsert: testCase.RecipeProcess.AltNames[0].DateInsert,
							DateUpdate: testCase.RecipeProcess.AltNames[0].DateUpdate,
							Name:       testCase.RecipeProcess.AltNames[0].Name,
							Status:     testCase.RecipeProcess.AltNames[0].Status,
						},
					},
					Entity: &DomainEntity.RecipeProcess{
						Id:          testCase.RecipeProcess.Entity.Id,
						UserId:      testCase.RecipeProcess.Entity.UserId,
						EntityId:    testCase.RecipeProcess.Entity.EntityId,
						DateInsert:  testCase.RecipeProcess.Entity.DateInsert,
						DateUpdate:  testCase.RecipeProcess.Entity.DateUpdate,
						Name:        testCase.RecipeProcess.Entity.Name,
						Description: testCase.RecipeProcess.Entity.Description,
						Notes:       testCase.RecipeProcess.Entity.Notes,
						Status:      testCase.RecipeProcess.Entity.Status,
					},
					Pictures: []*Picture{
						{
							AltNames: []*DomainEntity.AltName{
								{
									Id:         testCase.RecipeProcess.Pictures[0].AltNames[0].Id,
									UserId:     testCase.RecipeProcess.Pictures[0].AltNames[0].UserId,
									EntityId:   testCase.RecipeProcess.Pictures[0].AltNames[0].EntityId,
									DateInsert: testCase.RecipeProcess.Pictures[0].AltNames[0].DateInsert,
									DateUpdate: testCase.RecipeProcess.Pictures[0].AltNames[0].DateUpdate,
									Name:       testCase.RecipeProcess.Pictures[0].AltNames[0].Name,
									Status:     testCase.RecipeProcess.Pictures[0].AltNames[0].Status,
								},
							},
							Entity: &DomainEntity.Picture{
								Id:         testCase.RecipeProcess.Pictures[0].Entity.Id,
								UserId:     testCase.RecipeProcess.Pictures[0].Entity.UserId,
								EntityId:   testCase.RecipeProcess.Pictures[0].Entity.EntityId,
								DateInsert: testCase.RecipeProcess.Pictures[0].Entity.DateInsert,
								DateUpdate: testCase.RecipeProcess.Pictures[0].Entity.DateUpdate,
								Name:       testCase.RecipeProcess.Pictures[0].Entity.Name,
								URL:        testCase.RecipeProcess.Pictures[0].Entity.URL,
								Width:      testCase.RecipeProcess.Pictures[0].Entity.Width,
								Height:     testCase.RecipeProcess.Pictures[0].Entity.Height,
								Size:       testCase.RecipeProcess.Pictures[0].Entity.Size,
								Type:       testCase.RecipeProcess.Pictures[0].Entity.Type,
								Status:     testCase.RecipeProcess.Pictures[0].Entity.Status,
							},
						},
					},
				}

				assert.Equal(t, testCase.RecipeProcess.AltNames[0].Id, recipeProcessAggregate.AltNames[0].Id)
				assert.Equal(t, testCase.RecipeProcess.AltNames[0].UserId, recipeProcessAggregate.AltNames[0].UserId)
				assert.Equal(t, testCase.RecipeProcess.AltNames[0].EntityId, recipeProcessAggregate.AltNames[0].EntityId)
				assert.Equal(t, testCase.RecipeProcess.AltNames[0].DateInsert, recipeProcessAggregate.AltNames[0].DateInsert)
				assert.Equal(t, testCase.RecipeProcess.AltNames[0].DateUpdate, recipeProcessAggregate.AltNames[0].DateUpdate)
				assert.Equal(t, testCase.RecipeProcess.AltNames[0].Name, recipeProcessAggregate.AltNames[0].Name)
				assert.Equal(t, testCase.RecipeProcess.AltNames[0].Status, recipeProcessAggregate.AltNames[0].Status)
				assert.Equal(t, testCase.RecipeProcess.Entity.Id, recipeProcessAggregate.Entity.Id)
				assert.Equal(t, testCase.RecipeProcess.Entity.UserId, recipeProcessAggregate.Entity.UserId)
				assert.Equal(t, testCase.RecipeProcess.Entity.EntityId, recipeProcessAggregate.Entity.EntityId)
				assert.Equal(t, testCase.RecipeProcess.Entity.DateInsert, recipeProcessAggregate.Entity.DateInsert)
				assert.Equal(t, testCase.RecipeProcess.Entity.DateUpdate, recipeProcessAggregate.Entity.DateUpdate)
				assert.Equal(t, testCase.RecipeProcess.Entity.Name, recipeProcessAggregate.Entity.Name)
				assert.Equal(t, testCase.RecipeProcess.Entity.Description, recipeProcessAggregate.Entity.Description)
				assert.Equal(t, testCase.RecipeProcess.Entity.Notes, recipeProcessAggregate.Entity.Notes)
				assert.Equal(t, testCase.RecipeProcess.Entity.Status, recipeProcessAggregate.Entity.Status)
				assert.Equal(t, testCase.RecipeProcess.Pictures[0].AltNames[0].Id, recipeProcessAggregate.Pictures[0].AltNames[0].Id)
				assert.Equal(t, testCase.RecipeProcess.Pictures[0].AltNames[0].UserId, recipeProcessAggregate.Pictures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.RecipeProcess.Pictures[0].AltNames[0].EntityId, recipeProcessAggregate.Pictures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.RecipeProcess.Pictures[0].AltNames[0].DateInsert, recipeProcessAggregate.Pictures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.RecipeProcess.Pictures[0].AltNames[0].DateUpdate, recipeProcessAggregate.Pictures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.RecipeProcess.Pictures[0].AltNames[0].Name, recipeProcessAggregate.Pictures[0].AltNames[0].Name)
				assert.Equal(t, testCase.RecipeProcess.Pictures[0].AltNames[0].Status, recipeProcessAggregate.Pictures[0].AltNames[0].Status)
				assert.Equal(t, testCase.RecipeProcess.Pictures[0].Entity.Id, recipeProcessAggregate.Pictures[0].Entity.Id)
				assert.Equal(t, testCase.RecipeProcess.Pictures[0].Entity.UserId, recipeProcessAggregate.Pictures[0].Entity.UserId)
				assert.Equal(t, testCase.RecipeProcess.Pictures[0].Entity.EntityId, recipeProcessAggregate.Pictures[0].Entity.EntityId)
				assert.Equal(t, testCase.RecipeProcess.Pictures[0].Entity.DateInsert, recipeProcessAggregate.Pictures[0].Entity.DateInsert)
				assert.Equal(t, testCase.RecipeProcess.Pictures[0].Entity.DateUpdate, recipeProcessAggregate.Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.RecipeProcess.Pictures[0].Entity.DateUpdate, recipeProcessAggregate.Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.RecipeProcess.Pictures[0].Entity.Name, recipeProcessAggregate.Pictures[0].Entity.Name)
				assert.Equal(t, testCase.RecipeProcess.Pictures[0].Entity.URL, recipeProcessAggregate.Pictures[0].Entity.URL)
				assert.Equal(t, testCase.RecipeProcess.Pictures[0].Entity.Width, recipeProcessAggregate.Pictures[0].Entity.Width)
				assert.Equal(t, testCase.RecipeProcess.Pictures[0].Entity.Height, recipeProcessAggregate.Pictures[0].Entity.Height)
				assert.Equal(t, testCase.RecipeProcess.Pictures[0].Entity.Size, recipeProcessAggregate.Pictures[0].Entity.Size)
				assert.Equal(t, testCase.RecipeProcess.Pictures[0].Entity.Type, recipeProcessAggregate.Pictures[0].Entity.Type)
				assert.Equal(t, testCase.RecipeProcess.Pictures[0].Entity.Status, recipeProcessAggregate.Pictures[0].Entity.Status)

				reflectRecipeAggregate := reflect.ValueOf(recipeProcessAggregate)

				for i := 0; i < reflectRecipeAggregate.NumField(); i++ {
					assert.False(t, reflectRecipeAggregate.Field(i).IsZero())
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(recipeProcessAggregate)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}
