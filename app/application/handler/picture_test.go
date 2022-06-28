package handler

import (
	"github.com/google/uuid"
	DomainAggregate "github.com/sergeygardner/meal-planner-api/domain/aggregate"
	DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	InfrastructureService "github.com/sergeygardner/meal-planner-api/infrastructure/service"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	uuidPictureEntityId = uuid.New()
	testsPictureData    = testsPicture{
		{
			name:     "Test case with correct data",
			id:       nil,
			userId:   &testUserId,
			entityId: &uuidPictureEntityId,
			pictureDTO: &DomainEntity.Picture{
				Name:   "Test case with correct data",
				URL:    "https://www.google.com/images/branding/googlelogo/2x/googlelogo_light_color_272x92dp.png",
				Width:  1024,
				Height: 768,
				Size:   786432,
				Type:   "image/png",
				Status: kind.PictureStatusUnPublished,
			},
			toUpdatingPictureDTO: &DomainEntity.Picture{
				Status: kind.PictureStatusPublished,
			},
		},
	}
)

type testsPicture []struct {
	name                 string
	id                   *uuid.UUID
	userId               *uuid.UUID
	entityId             *uuid.UUID
	pictureDTO           *DomainEntity.Picture
	toUpdatingPictureDTO *DomainEntity.Picture
	picture              *DomainAggregate.Picture
}

func init() {
	InfrastructureService.PrepareCache()
	InfrastructureService.PreparePersistence()
}

func TestPictureCreate(t *testing.T) {
	for index, testCase := range testsPictureData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := PictureCreate(testCase.userId, testCase.entityId, testCase.pictureDTO)

				assert.Nil(t, errorActual)

				testsPictureData[index].picture = actual
				testsPictureData[index].id = &actual.Entity.Id
				testsPictureData[index].entityId = &actual.Entity.EntityId

				assert.NotNil(t, actual.Entity.Id)
				assert.Equal(t, testCase.pictureDTO.UserId, actual.Entity.UserId)
				assert.Equal(t, testCase.pictureDTO.EntityId, actual.Entity.EntityId)
				assert.NotNil(t, actual.Entity.DateInsert)
				assert.NotNil(t, actual.Entity.DateUpdate)
				assert.Equal(t, testCase.pictureDTO.Name, actual.Entity.Name)
				assert.Equal(t, testCase.pictureDTO.URL, actual.Entity.URL)
				assert.Equal(t, testCase.pictureDTO.Width, actual.Entity.Width)
				assert.Equal(t, testCase.pictureDTO.Height, actual.Entity.Height)
				assert.Equal(t, testCase.pictureDTO.Size, actual.Entity.Size)
				assert.Equal(t, testCase.pictureDTO.Type, actual.Entity.Type)
				assert.Equal(t, testCase.pictureDTO.Status, actual.Entity.Status)
			},
		)
	}
}

func TestPicturesInfo(t *testing.T) {
	for _, testCase := range testsPictureData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := PicturesInfo(testCase.userId, testCase.entityId, nil)

				if testCase.picture != nil {
					assert.Nil(t, errorActual)

					for _, actualEntity := range actual {
						assert.Equal(t, testCase.picture.Entity.Id, actualEntity.Entity.Id)
						assert.Equal(t, testCase.picture.Entity.UserId, actualEntity.Entity.UserId)
						assert.Equal(t, testCase.picture.Entity.EntityId, actualEntity.Entity.EntityId)
						assert.Equal(t, testCase.picture.Entity.DateInsert.Format(time.UnixDate), actualEntity.Entity.DateInsert.Format(time.UnixDate))
						assert.Equal(t, testCase.picture.Entity.DateUpdate.Format(time.UnixDate), actualEntity.Entity.DateUpdate.Format(time.UnixDate))
						assert.Equal(t, testCase.picture.Entity.Name, actualEntity.Entity.Name)
						assert.Equal(t, testCase.picture.Entity.URL, actualEntity.Entity.URL)
						assert.Equal(t, testCase.picture.Entity.Width, actualEntity.Entity.Width)
						assert.Equal(t, testCase.picture.Entity.Height, actualEntity.Entity.Height)
						assert.Equal(t, testCase.picture.Entity.Size, actualEntity.Entity.Size)
						assert.Equal(t, testCase.picture.Entity.Type, actualEntity.Entity.Type)
						assert.Equal(t, testCase.picture.Entity.Status, actualEntity.Entity.Status)
					}
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestPictureInfo(t *testing.T) {
	for _, testCase := range testsPictureData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actualEntity, errorActual := PictureInfo(testCase.id, testCase.userId, testCase.entityId, nil)

				if testCase.picture != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.picture.Entity.Id, actualEntity.Entity.Id)
					assert.Equal(t, testCase.picture.Entity.UserId, actualEntity.Entity.UserId)
					assert.Equal(t, testCase.picture.Entity.EntityId, actualEntity.Entity.EntityId)
					assert.Equal(t, testCase.picture.Entity.DateInsert.Format(time.UnixDate), actualEntity.Entity.DateInsert.Format(time.UnixDate))
					assert.Equal(t, testCase.picture.Entity.DateUpdate.Format(time.UnixDate), actualEntity.Entity.DateUpdate.Format(time.UnixDate))
					assert.Equal(t, testCase.picture.Entity.Name, actualEntity.Entity.Name)
					assert.Equal(t, testCase.picture.Entity.URL, actualEntity.Entity.URL)
					assert.Equal(t, testCase.picture.Entity.Width, actualEntity.Entity.Width)
					assert.Equal(t, testCase.picture.Entity.Height, actualEntity.Entity.Height)
					assert.Equal(t, testCase.picture.Entity.Size, actualEntity.Entity.Size)
					assert.Equal(t, testCase.picture.Entity.Type, actualEntity.Entity.Type)
					assert.Equal(t, testCase.picture.Entity.Status, actualEntity.Entity.Status)
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestPictureUpdate(t *testing.T) {
	for index, testCase := range testsPictureData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := PictureUpdate(testCase.id, testCase.userId, testCase.entityId, testCase.toUpdatingPictureDTO)

				if testCase.picture != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.toUpdatingPictureDTO.Status, actual.Entity.Status)
					testsPictureData[index].picture.Entity.Status = actual.Entity.Status
					testsPictureData[index].picture.Entity.DateUpdate = actual.Entity.DateUpdate
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestGetPictureEntity(t *testing.T) {
	TestPictureInfo(t)
}

func TestPictureDelete(t *testing.T) {
	for _, testCase := range testsPictureData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := PictureDelete(testCase.id, testCase.userId, testCase.entityId)

				if testCase.picture != nil {
					assert.Nil(t, errorActual)
					assert.True(t, actual)
				} else {
					assert.NotNil(t, errorActual)
					assert.False(t, actual)
				}
			},
		)
	}
}
