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

func TestPicture(t *testing.T) {
	tests := []struct {
		name     string
		json     string
		AltNames []struct {
			Id         uuid.UUID
			UserId     uuid.UUID
			EntityId   uuid.UUID
			DateInsert time.Time
			DateUpdate time.Time
			Name       string
			Status     kind.AltNameStatus
		}
		Entity struct {
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
	}{
		{
			name: "Test case with published picture properties",
			json: "{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000004\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000005\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"published\"}}\n",
			AltNames: []struct {
				Id         uuid.UUID
				UserId     uuid.UUID
				EntityId   uuid.UUID
				DateInsert time.Time
				DateUpdate time.Time
				Name       string
				Status     kind.AltNameStatus
			}{
				{
					Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
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
				Id:         uuid.MustParse("00000000-0000-0000-0000-000000000004"),
				UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000005"),
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
		}, {
			name: "Test case with unpublished picture properties",
			json: "{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"unpublished\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000004\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000005\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"unpublished\"}}\n",
			AltNames: []struct {
				Id         uuid.UUID
				UserId     uuid.UUID
				EntityId   uuid.UUID
				DateInsert time.Time
				DateUpdate time.Time
				Name       string
				Status     kind.AltNameStatus
			}{
				{
					Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
					DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
					DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
					Name:       "AltName",
					Status:     kind.AltNameStatusUnPublished,
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
				Id:         uuid.MustParse("00000000-0000-0000-0000-000000000004"),
				UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000005"),
				DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
				DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
				Name:       "Picture",
				URL:        "https://google.com/doodle.png",
				Width:      512,
				Height:     512,
				Size:       1024,
				Type:       "image/png",
				Status:     kind.PictureStatusUnPublished,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				pictureAggregate := Picture{
					AltNames: []*DomainEntity.AltName{
						{
							Id:         testCase.AltNames[0].Id,
							UserId:     testCase.AltNames[0].UserId,
							EntityId:   testCase.AltNames[0].EntityId,
							DateInsert: testCase.AltNames[0].DateInsert,
							DateUpdate: testCase.AltNames[0].DateUpdate,
							Name:       testCase.AltNames[0].Name,
							Status:     testCase.AltNames[0].Status,
						},
					},
					Entity: &DomainEntity.Picture{
						Id:         testCase.Entity.Id,
						UserId:     testCase.Entity.UserId,
						EntityId:   testCase.Entity.EntityId,
						DateInsert: testCase.Entity.DateInsert,
						DateUpdate: testCase.Entity.DateUpdate,
						Name:       testCase.Entity.Name,
						URL:        testCase.Entity.URL,
						Width:      testCase.Entity.Width,
						Height:     testCase.Entity.Height,
						Size:       testCase.Entity.Size,
						Type:       testCase.Entity.Type,
						Status:     testCase.Entity.Status,
					},
				}
				assert.Equal(t, testCase.AltNames[0].Id, pictureAggregate.AltNames[0].Id)
				assert.Equal(t, testCase.AltNames[0].UserId, pictureAggregate.AltNames[0].UserId)
				assert.Equal(t, testCase.AltNames[0].EntityId, pictureAggregate.AltNames[0].EntityId)
				assert.Equal(t, testCase.AltNames[0].DateInsert, pictureAggregate.AltNames[0].DateInsert)
				assert.Equal(t, testCase.AltNames[0].DateUpdate, pictureAggregate.AltNames[0].DateUpdate)
				assert.Equal(t, testCase.AltNames[0].Name, pictureAggregate.AltNames[0].Name)
				assert.Equal(t, testCase.AltNames[0].Status, pictureAggregate.AltNames[0].Status)
				assert.Equal(t, testCase.Entity.Id, pictureAggregate.Entity.Id)
				assert.Equal(t, testCase.Entity.UserId, pictureAggregate.Entity.UserId)
				assert.Equal(t, testCase.Entity.EntityId, pictureAggregate.Entity.EntityId)
				assert.Equal(t, testCase.Entity.DateInsert, pictureAggregate.Entity.DateInsert)
				assert.Equal(t, testCase.Entity.DateUpdate, pictureAggregate.Entity.DateUpdate)
				assert.Equal(t, testCase.Entity.Name, pictureAggregate.Entity.Name)
				assert.Equal(t, testCase.Entity.URL, pictureAggregate.Entity.URL)
				assert.Equal(t, testCase.Entity.Width, pictureAggregate.Entity.Width)
				assert.Equal(t, testCase.Entity.Height, pictureAggregate.Entity.Height)
				assert.Equal(t, testCase.Entity.Size, pictureAggregate.Entity.Size)
				assert.Equal(t, testCase.Entity.Status, pictureAggregate.Entity.Status)

				reflectPicture := reflect.ValueOf(pictureAggregate)
				reflectPictureAltName := reflect.ValueOf(*pictureAggregate.AltNames[0])
				reflectPictureEntity := reflect.ValueOf(*pictureAggregate.Entity)

				for i := 0; i < reflectPicture.NumField(); i++ {
					assert.False(t, reflectPicture.Field(i).IsZero())
				}
				for i := 0; i < reflectPictureAltName.NumField(); i++ {
					assert.False(t, reflectPictureAltName.Field(i).IsZero())
				}
				for i := 0; i < reflectPictureEntity.NumField(); i++ {
					assert.False(t, reflectPictureEntity.Field(i).IsZero())
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(pictureAggregate)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}

func TestCategory(t *testing.T) {
	tests := []struct {
		name     string
		json     string
		AltNames []struct {
			Id         uuid.UUID
			UserId     uuid.UUID
			EntityId   uuid.UUID
			DateInsert time.Time
			DateUpdate time.Time
			Name       string
			Status     kind.AltNameStatus
		}
		Entity struct {
			Id         uuid.UUID
			UserId     uuid.UUID
			DateInsert time.Time
			DateUpdate time.Time
			Name       string
			Status     kind.CategoryStatus
		}
		Pictures []struct {
			AltNames []struct {
				Id         uuid.UUID
				UserId     uuid.UUID
				EntityId   uuid.UUID
				DateInsert time.Time
				DateUpdate time.Time
				Name       string
				Status     kind.AltNameStatus
			}
			Entity struct {
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
	}{
		{
			name: "Test case with published category properties",
			json: "{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000003\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000001\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Category\",\"status\":\"published\"},\"pictures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000005\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000004\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000001\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"published\"}}]}\n",
			AltNames: []struct {
				Id         uuid.UUID
				UserId     uuid.UUID
				EntityId   uuid.UUID
				DateInsert time.Time
				DateUpdate time.Time
				Name       string
				Status     kind.AltNameStatus
			}{
				{
					Id:         uuid.MustParse("00000000-0000-0000-0000-000000000003"),
					UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
					DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
					Name:       "AltName",
					Status:     kind.AltNameStatusPublished,
				},
			},
			Entity: struct {
				Id         uuid.UUID
				UserId     uuid.UUID
				DateInsert time.Time
				DateUpdate time.Time
				Name       string
				Status     kind.CategoryStatus
			}{
				Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
				DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
				Name:       "Category",
				Status:     kind.CategoryStatusPublished,
			},
			Pictures: []struct {
				AltNames []struct {
					Id         uuid.UUID
					UserId     uuid.UUID
					EntityId   uuid.UUID
					DateInsert time.Time
					DateUpdate time.Time
					Name       string
					Status     kind.AltNameStatus
				}
				Entity struct {
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
			}{
				{
					AltNames: []struct {
						Id         uuid.UUID
						UserId     uuid.UUID
						EntityId   uuid.UUID
						DateInsert time.Time
						DateUpdate time.Time
						Name       string
						Status     kind.AltNameStatus
					}{
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
						Id:         uuid.MustParse("00000000-0000-0000-0000-000000000004"),
						UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
						EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000001"),
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
		}, {
			name: "Test case with unpublished category properties",
			json: "{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000003\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000001\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"unpublished\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Category\",\"status\":\"unpublished\"},\"pictures\":[{\"alt_names\":[{\"id\":\"00000000-0000-0000-0000-000000000005\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"unpublished\"}],\"entity\":{\"id\":\"00000000-0000-0000-0000-000000000004\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000001\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"unpublished\"}}]}\n",
			AltNames: []struct {
				Id         uuid.UUID
				UserId     uuid.UUID
				EntityId   uuid.UUID
				DateInsert time.Time
				DateUpdate time.Time
				Name       string
				Status     kind.AltNameStatus
			}{
				{
					Id:         uuid.MustParse("00000000-0000-0000-0000-000000000003"),
					UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
					DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
					Name:       "AltName",
					Status:     kind.AltNameStatusUnPublished,
				},
			},
			Entity: struct {
				Id         uuid.UUID
				UserId     uuid.UUID
				DateInsert time.Time
				DateUpdate time.Time
				Name       string
				Status     kind.CategoryStatus
			}{
				Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
				DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
				Name:       "Category",
				Status:     kind.CategoryStatusUnPublished,
			},
			Pictures: []struct {
				AltNames []struct {
					Id         uuid.UUID
					UserId     uuid.UUID
					EntityId   uuid.UUID
					DateInsert time.Time
					DateUpdate time.Time
					Name       string
					Status     kind.AltNameStatus
				}
				Entity struct {
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
			}{
				{AltNames: []struct {
					Id         uuid.UUID
					UserId     uuid.UUID
					EntityId   uuid.UUID
					DateInsert time.Time
					DateUpdate time.Time
					Name       string
					Status     kind.AltNameStatus
				}{
					{
						Id:         uuid.MustParse("00000000-0000-0000-0000-000000000005"),
						UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
						EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
						DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
						DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
						Name:       "AltName",
						Status:     kind.AltNameStatusUnPublished,
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
						Id:         uuid.MustParse("00000000-0000-0000-0000-000000000004"),
						UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
						EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000001"),
						DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
						DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
						Name:       "Picture",
						URL:        "https://google.com/doodle.png",
						Width:      512,
						Height:     512,
						Size:       1024,
						Type:       "image/png",
						Status:     kind.PictureStatusUnPublished,
					},
				},
			},
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				categoryAggregate := Category{
					AltNames: []*DomainEntity.AltName{
						{
							Id:         testCase.AltNames[0].Id,
							UserId:     testCase.AltNames[0].UserId,
							EntityId:   testCase.AltNames[0].EntityId,
							DateInsert: testCase.AltNames[0].DateInsert,
							DateUpdate: testCase.AltNames[0].DateUpdate,
							Name:       testCase.AltNames[0].Name,
							Status:     testCase.AltNames[0].Status,
						},
					},
					Entity: &DomainEntity.Category{
						Id:         testCase.Entity.Id,
						UserId:     testCase.Entity.UserId,
						DateInsert: testCase.Entity.DateInsert,
						DateUpdate: testCase.Entity.DateUpdate,
						Name:       testCase.Entity.Name,
						Status:     testCase.Entity.Status,
					},
					Pictures: []*Picture{
						{
							AltNames: []*DomainEntity.AltName{
								{
									Id:         testCase.Pictures[0].AltNames[0].Id,
									UserId:     testCase.Pictures[0].AltNames[0].UserId,
									EntityId:   testCase.Pictures[0].AltNames[0].EntityId,
									DateInsert: testCase.Pictures[0].AltNames[0].DateInsert,
									DateUpdate: testCase.Pictures[0].AltNames[0].DateUpdate,
									Name:       testCase.Pictures[0].AltNames[0].Name,
									Status:     testCase.Pictures[0].AltNames[0].Status,
								},
							},
							Entity: &DomainEntity.Picture{
								Id:         testCase.Pictures[0].Entity.Id,
								UserId:     testCase.Pictures[0].Entity.UserId,
								EntityId:   testCase.Pictures[0].Entity.EntityId,
								DateInsert: testCase.Pictures[0].Entity.DateInsert,
								DateUpdate: testCase.Pictures[0].Entity.DateUpdate,
								Name:       testCase.Pictures[0].Entity.Name,
								URL:        testCase.Pictures[0].Entity.URL,
								Width:      testCase.Pictures[0].Entity.Width,
								Height:     testCase.Pictures[0].Entity.Height,
								Size:       testCase.Pictures[0].Entity.Size,
								Type:       testCase.Pictures[0].Entity.Type,
								Status:     testCase.Pictures[0].Entity.Status,
							},
						},
					},
				}
				assert.Equal(t, testCase.AltNames[0].Id, categoryAggregate.AltNames[0].Id)
				assert.Equal(t, testCase.AltNames[0].UserId, categoryAggregate.AltNames[0].UserId)
				assert.Equal(t, testCase.AltNames[0].EntityId, categoryAggregate.AltNames[0].EntityId)
				assert.Equal(t, testCase.AltNames[0].DateInsert, categoryAggregate.AltNames[0].DateInsert)
				assert.Equal(t, testCase.AltNames[0].DateUpdate, categoryAggregate.AltNames[0].DateUpdate)
				assert.Equal(t, testCase.AltNames[0].Name, categoryAggregate.AltNames[0].Name)
				assert.Equal(t, testCase.AltNames[0].Status, categoryAggregate.AltNames[0].Status)
				assert.Equal(t, testCase.Entity.Id, categoryAggregate.Entity.Id)
				assert.Equal(t, testCase.Entity.UserId, categoryAggregate.Entity.UserId)
				assert.Equal(t, testCase.Entity.DateInsert, categoryAggregate.Entity.DateInsert)
				assert.Equal(t, testCase.Entity.DateUpdate, categoryAggregate.Entity.DateUpdate)
				assert.Equal(t, testCase.Entity.Name, categoryAggregate.Entity.Name)
				assert.Equal(t, testCase.Entity.Status, categoryAggregate.Entity.Status)

				assert.Equal(t, testCase.Pictures[0].AltNames[0].Id, categoryAggregate.Pictures[0].AltNames[0].Id)
				assert.Equal(t, testCase.Pictures[0].AltNames[0].UserId, categoryAggregate.Pictures[0].AltNames[0].UserId)
				assert.Equal(t, testCase.Pictures[0].AltNames[0].EntityId, categoryAggregate.Pictures[0].AltNames[0].EntityId)
				assert.Equal(t, testCase.Pictures[0].AltNames[0].DateInsert, categoryAggregate.Pictures[0].AltNames[0].DateInsert)
				assert.Equal(t, testCase.Pictures[0].AltNames[0].DateUpdate, categoryAggregate.Pictures[0].AltNames[0].DateUpdate)
				assert.Equal(t, testCase.Pictures[0].AltNames[0].Name, categoryAggregate.Pictures[0].AltNames[0].Name)
				assert.Equal(t, testCase.Pictures[0].AltNames[0].Status, categoryAggregate.Pictures[0].AltNames[0].Status)
				assert.Equal(t, testCase.Pictures[0].Entity.Id, categoryAggregate.Pictures[0].Entity.Id)
				assert.Equal(t, testCase.Pictures[0].Entity.UserId, categoryAggregate.Pictures[0].Entity.UserId)
				assert.Equal(t, testCase.Pictures[0].Entity.EntityId, categoryAggregate.Pictures[0].Entity.EntityId)
				assert.Equal(t, testCase.Pictures[0].Entity.DateInsert, categoryAggregate.Pictures[0].Entity.DateInsert)
				assert.Equal(t, testCase.Pictures[0].Entity.DateUpdate, categoryAggregate.Pictures[0].Entity.DateUpdate)
				assert.Equal(t, testCase.Pictures[0].Entity.Name, categoryAggregate.Pictures[0].Entity.Name)
				assert.Equal(t, testCase.Pictures[0].Entity.URL, categoryAggregate.Pictures[0].Entity.URL)
				assert.Equal(t, testCase.Pictures[0].Entity.Width, categoryAggregate.Pictures[0].Entity.Width)
				assert.Equal(t, testCase.Pictures[0].Entity.Height, categoryAggregate.Pictures[0].Entity.Height)
				assert.Equal(t, testCase.Pictures[0].Entity.Size, categoryAggregate.Pictures[0].Entity.Size)
				assert.Equal(t, testCase.Pictures[0].Entity.Status, categoryAggregate.Pictures[0].Entity.Status)

				reflectCategory := reflect.ValueOf(categoryAggregate)
				reflectCategoryAltName := reflect.ValueOf(*categoryAggregate.AltNames[0])
				reflectCategoryEntity := reflect.ValueOf(*categoryAggregate.Entity)
				reflectCategoryPicture := reflect.ValueOf(*categoryAggregate.Pictures[0])
				reflectCategoryPictureAltName := reflect.ValueOf(*categoryAggregate.Pictures[0].AltNames[0])
				reflectCategoryPictureEntity := reflect.ValueOf(*categoryAggregate.Pictures[0].Entity)

				for i := 0; i < reflectCategory.NumField(); i++ {
					assert.False(t, reflectCategory.Field(i).IsZero())
				}
				for i := 0; i < reflectCategoryAltName.NumField(); i++ {
					assert.False(t, reflectCategoryAltName.Field(i).IsZero())
				}
				for i := 0; i < reflectCategoryEntity.NumField(); i++ {
					assert.False(t, reflectCategoryEntity.Field(i).IsZero())
				}
				for i := 0; i < reflectCategoryPicture.NumField(); i++ {
					assert.False(t, reflectCategoryPicture.Field(i).IsZero())
				}
				for i := 0; i < reflectCategoryPictureAltName.NumField(); i++ {
					assert.False(t, reflectCategoryPictureAltName.Field(i).IsZero())
				}
				for i := 0; i < reflectCategoryPictureEntity.NumField(); i++ {
					assert.False(t, reflectCategoryPictureEntity.Field(i).IsZero())
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(categoryAggregate)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}
