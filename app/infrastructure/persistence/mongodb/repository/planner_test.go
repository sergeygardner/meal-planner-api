package repository

import (
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence/repository"
	InfrastructureService "github.com/sergeygardner/meal-planner-api/infrastructure/service"
	InfrastructureServiceEntity "github.com/sergeygardner/meal-planner-api/infrastructure/service/entity"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	testPlannerRepository repository.PlannerRepositoryInterface
)

func init() {
	InfrastructureService.PrepareCache()
	InfrastructureService.PreparePersistence()

	testPlannerRepository = &PlannerRepository{Table: "planner", EntityManager: InfrastructureServiceEntity.GetEntityManager()}
}

func TestPlannerRepositoryFindOne(t *testing.T) {
	tests := []struct {
		Name        string
		Criteria    *persistence.Criteria
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with PlannerRepository.FindOne with correct data",
			Criteria:    &persistence.Criteria{},
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with PlannerRepository.FindOne with incorrect data",
			Criteria:    nil,
			MustBePanic: true,
			MustBeFault: true,
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

				one, errorFindOne := testPlannerRepository.FindOne(testCase.Criteria)

				if testCase.MustBeFault {
					assert.NotNil(t, errorFindOne)
					assert.Nil(t, one)
				} else {
					assert.Nil(t, errorFindOne)
					assert.NotNil(t, one)
				}
			},
		)
	}
}

func TestPlannerRepositoryFindAll(t *testing.T) {
	tests := []struct {
		Name        string
		Criteria    *persistence.Criteria
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with PlannerRepository.FindAll with correct data",
			Criteria:    &persistence.Criteria{Limit: 2},
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with PlannerRepository.FindAll with incorrect data",
			Criteria:    nil,
			MustBePanic: true,
			MustBeFault: true,
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

				all, errorFindAll := testPlannerRepository.FindAll(&persistence.Criteria{Limit: 2})

				if testCase.MustBeFault {
					assert.NotNil(t, errorFindAll)
					assert.Nil(t, all)
				} else {
					assert.Nil(t, errorFindAll)
					assert.NotNil(t, all)
				}
			},
		)
	}
}

func TestPlannerRepositoryInsertOne(t *testing.T) {
	tests := []struct {
		Name        string
		Planner     *entity.Planner
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name: "Test case with PlannerRepository.InsertOne with correct data",
			Planner: &entity.Planner{
				UserId:     uuid.New(),
				DateInsert: time.Now().UTC(),
				DateUpdate: time.Now().UTC(),
				StartTime:  time.Now().UTC(),
				EndTime:    time.Now().UTC(),
				Name:       "Planner",
				Status:     kind.PlannerStatusInActive,
			},
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with PlannerRepository.InsertOne with incorrect data",
			Planner:     &entity.Planner{},
			MustBePanic: false,
			MustBeFault: true,
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

				one, errorInsertOne := testPlannerRepository.InsertOne(testCase.Planner)

				if testCase.MustBeFault {
					assert.NotNil(t, errorInsertOne)
					assert.Nil(t, one)
				} else {
					assert.Nil(t, errorInsertOne)
					assert.NotNil(t, one)
					assert.NotEqual(t, uuid.Nil, one.Id)
				}
			},
		)
	}
}

func TestPlannerRepositoryInsertMany(t *testing.T) {
	tests := []struct {
		Name        string
		Planners    []*entity.Planner
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name: "Test case with PlannerRepository.InsertMany with correct data",
			Planners: []*entity.Planner{
				{
					UserId:     uuid.New(),
					DateInsert: time.Now().UTC(),
					DateUpdate: time.Now().UTC(),
					StartTime:  time.Now().UTC(),
					EndTime:    time.Now().UTC(),
					Name:       "Planner1",
					Status:     kind.PlannerStatusInActive,
				}, {
					UserId:     uuid.New(),
					DateInsert: time.Now().UTC(),
					DateUpdate: time.Now().UTC(),
					StartTime:  time.Now().UTC(),
					EndTime:    time.Now().UTC(),
					Name:       "Planner2",
					Status:     kind.PlannerStatusInActive,
				},
			},
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with PlannerRepository.InsertMany with incorrect data",
			Planners:    []*entity.Planner{},
			MustBePanic: false,
			MustBeFault: true,
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

				many, errorInsertMany := testPlannerRepository.InsertMany(testCase.Planners)

				if testCase.MustBeFault {
					assert.NotNil(t, errorInsertMany)
					assert.Nil(t, many)
				} else {
					assert.Nil(t, errorInsertMany)
					assert.NotNil(t, many)

					for _, one := range many {
						assert.NotEqual(t, uuid.Nil, one.Id)
					}
				}
			},
		)
	}
}

func TestPlannerRepositoryUpdateOne(t *testing.T) {
	id := uuid.New()

	tests := []struct {
		Name        string
		Planner     *entity.Planner
		Criteria    *persistence.Criteria
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name: "Test case with PlannerRepository.UpdateOne with correct data",
			Planner: &entity.Planner{
				Id:         id,
				UserId:     uuid.New(),
				DateInsert: time.Now().UTC(),
				DateUpdate: time.Now().UTC(),
				StartTime:  time.Now().UTC(),
				EndTime:    time.Now().UTC(),
				Name:       "Planner",
				Status:     kind.PlannerStatusInActive,
			},
			Criteria: &persistence.Criteria{
				Where: map[string]interface{}{
					"id": &id,
				},
			},
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with PlannerRepository.UpdateOne with incorrect data",
			Planner:     &entity.Planner{},
			Criteria:    &persistence.Criteria{},
			MustBePanic: false,
			MustBeFault: true,
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

				one, errorUpdateOne := testPlannerRepository.UpdateOne(testCase.Criteria, testCase.Planner)

				if testCase.MustBeFault {
					assert.NotNil(t, errorUpdateOne)
					assert.Nil(t, one)
				} else {
					assert.Nil(t, errorUpdateOne)
					assert.NotNil(t, one)
					assert.Equal(t, testCase.Planner.Id, one.Id)
					assert.Equal(t, testCase.Planner.UserId, one.UserId)
					assert.Equal(t, testCase.Planner.DateInsert, one.DateInsert)
					assert.Equal(t, testCase.Planner.DateUpdate, one.DateUpdate)
					assert.Equal(t, testCase.Planner.StartTime, one.StartTime)
					assert.Equal(t, testCase.Planner.EndTime, one.EndTime)
					assert.Equal(t, testCase.Planner.Name, one.Name)
					assert.Equal(t, testCase.Planner.Status, one.Status)
				}
			},
		)
	}
}

func TestPlannerRepositoryUpdateMany(t *testing.T) {
	tests := []struct {
		Name        string
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with PlannerRepository.UpdateMany with correct data",
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with PlannerRepository.UpdateMany with incorrect data",
			MustBePanic: false,
			MustBeFault: true,
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

				if testCase.MustBeFault {

				} else {

				}
			},
		)
	}
}

func TestPlannerRepositoryDeleteOne(t *testing.T) {
	tests := []struct {
		Name        string
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with PlannerRepository.DeleteOne with correct data",
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with PlannerRepository.DeleteOne with incorrect data",
			MustBePanic: false,
			MustBeFault: true,
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

				if testCase.MustBeFault {

				} else {

				}
			},
		)
	}
}

func TestPlannerRepositoryGetCriteria(t *testing.T) {
	tests := []struct {
		Name        string
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with PlannerRepository.GetCriteria with correct data",
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

				criteria := testPlannerRepository.GetCriteria()

				if testCase.MustBeFault {
					assert.Nil(t, criteria)
				} else {
					assert.NotNil(t, criteria)
				}
			},
		)
	}
}

func TestPlannerIntervalRepositoryFindOne(t *testing.T) {
	tests := []struct {
		Name        string
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with  with correct data",
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with  with incorrect data",
			MustBePanic: false,
			MustBeFault: true,
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

				if testCase.MustBeFault {

				} else {

				}
			},
		)
	}
}

func TestPlannerIntervalRepositoryFindAll(t *testing.T) {
	tests := []struct {
		Name        string
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with  with correct data",
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with  with incorrect data",
			MustBePanic: false,
			MustBeFault: true,
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

				if testCase.MustBeFault {

				} else {

				}
			},
		)
	}
}

func TestPlannerIntervalRepositoryInsertOne(t *testing.T) {
	tests := []struct {
		Name        string
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with  with correct data",
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with  with incorrect data",
			MustBePanic: false,
			MustBeFault: true,
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

				if testCase.MustBeFault {

				} else {

				}
			},
		)
	}
}

func TestPlannerIntervalRepositoryInsertMany(t *testing.T) {
	tests := []struct {
		Name        string
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with  with correct data",
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with  with incorrect data",
			MustBePanic: false,
			MustBeFault: true,
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

				if testCase.MustBeFault {

				} else {

				}
			},
		)
	}
}

func TestPlannerIntervalRepositoryUpdateOne(t *testing.T) {
	tests := []struct {
		Name        string
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with  with correct data",
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with  with incorrect data",
			MustBePanic: false,
			MustBeFault: true,
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

				if testCase.MustBeFault {

				} else {

				}
			},
		)
	}
}

func TestPlannerIntervalRepositoryUpdateMany(t *testing.T) {
	tests := []struct {
		Name        string
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with  with correct data",
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with  with incorrect data",
			MustBePanic: false,
			MustBeFault: true,
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

				if testCase.MustBeFault {

				} else {

				}
			},
		)
	}
}

func TestPlannerIntervalRepositoryDeleteOne(t *testing.T) {
	tests := []struct {
		Name        string
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with  with correct data",
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with  with incorrect data",
			MustBePanic: false,
			MustBeFault: true,
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

				if testCase.MustBeFault {

				} else {

				}
			},
		)
	}
}

func TestPlannerIntervalRepositoryGetCriteria(t *testing.T) {
	tests := []struct {
		Name        string
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with  with correct data",
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with  with incorrect data",
			MustBePanic: false,
			MustBeFault: true,
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

				if testCase.MustBeFault {

				} else {

				}
			},
		)
	}
}

func TestPlannerRecipeRepositoryFindOne(t *testing.T) {
	tests := []struct {
		Name        string
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with  with correct data",
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with  with incorrect data",
			MustBePanic: false,
			MustBeFault: true,
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

				if testCase.MustBeFault {

				} else {

				}
			},
		)
	}
}

func TestPlannerRecipeRepositoryFindAll(t *testing.T) {
	tests := []struct {
		Name        string
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with  with correct data",
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with  with incorrect data",
			MustBePanic: false,
			MustBeFault: true,
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

				if testCase.MustBeFault {

				} else {

				}
			},
		)
	}
}

func TestPlannerRecipeRepositoryInsertOne(t *testing.T) {
	tests := []struct {
		Name        string
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with  with correct data",
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with  with incorrect data",
			MustBePanic: false,
			MustBeFault: true,
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

				if testCase.MustBeFault {

				} else {

				}
			},
		)
	}
}

func TestPlannerRecipeRepositoryInsertMany(t *testing.T) {
	tests := []struct {
		Name        string
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with  with correct data",
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with  with incorrect data",
			MustBePanic: false,
			MustBeFault: true,
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

				if testCase.MustBeFault {

				} else {

				}
			},
		)
	}
}

func TestPlannerRecipeRepositoryUpdateOne(t *testing.T) {
	tests := []struct {
		Name        string
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with  with correct data",
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with  with incorrect data",
			MustBePanic: false,
			MustBeFault: true,
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

				if testCase.MustBeFault {

				} else {

				}
			},
		)
	}
}

func TestPlannerRecipeRepositoryUpdateMany(t *testing.T) {
	tests := []struct {
		Name        string
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with  with correct data",
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with  with incorrect data",
			MustBePanic: false,
			MustBeFault: true,
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

				if testCase.MustBeFault {

				} else {

				}
			},
		)
	}
}

func TestPlannerRecipeRepositoryDeleteOne(t *testing.T) {
	tests := []struct {
		Name        string
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with  with correct data",
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with  with incorrect data",
			MustBePanic: false,
			MustBeFault: true,
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

				if testCase.MustBeFault {

				} else {

				}
			},
		)
	}
}

func TestPlannerRecipeRepositoryGetCriteria(t *testing.T) {
	tests := []struct {
		Name        string
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with  with correct data",
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with  with incorrect data",
			MustBePanic: false,
			MustBeFault: true,
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

				if testCase.MustBeFault {

				} else {

				}
			},
		)
	}
}
