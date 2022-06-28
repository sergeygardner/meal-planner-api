package service

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestMathRandomInt(t *testing.T) {
	tests := []struct {
		Name        string
		Min         int64
		Max         int64
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with MathRandomInt with correct data",
			Min:         100,
			Max:         1000,
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with MathRandomInt with incorrect data",
			Min:         1000,
			Max:         100,
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

				randomInt := MathRandomInt(testCase.Min, testCase.Max)

				if testCase.MustBeFault {
					assert.Greater(t, randomInt, testCase.Max)
					assert.Greater(t, randomInt, testCase.Min)
				} else {
					assert.Greater(t, randomInt, testCase.Min)
					assert.Less(t, randomInt, testCase.Max)
				}
			},
		)
	}
}

func TestMathRandomIntAsString(t *testing.T) {
	tests := []struct {
		Name        string
		Min         int64
		Max         int64
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with MathRandomIntAsString with correct data",
			Min:         100,
			Max:         1000,
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name:        "Test case with MathRandomIntAsString with incorrect data",
			Min:         1000,
			Max:         100,
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

				randomIntAsString := MathRandomIntAsString(testCase.Min, testCase.Max)
				randomInt, errorParseInt := strconv.ParseInt(randomIntAsString, 10, 64)

				assert.Nil(t, errorParseInt)

				if testCase.MustBeFault {
					assert.Greater(t, randomInt, testCase.Max)
					assert.Greater(t, randomInt, testCase.Min)
				} else {
					assert.Greater(t, randomInt, testCase.Min)
					assert.Less(t, randomInt, testCase.Max)
					assert.Equal(t, randomIntAsString, strconv.FormatInt(randomInt, 10))
				}
			},
		)
	}
}
