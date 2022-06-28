package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdate(t *testing.T) {
	tests := []struct {
		name   string
		Target *struct {
			Name   string
			Status bool
		}
		Source   any
		Expected *struct {
			Name   string
			Status bool
		}
		MustBePassed bool
	}{
		{
			name: "Test case for Update with success",
			Target: &struct {
				Name   string
				Status bool
			}{
				Name:   "name",
				Status: true,
			},
			Source: &struct {
				Name   string
				Status bool
			}{
				Status: false,
			},
			Expected: &struct {
				Name   string
				Status bool
			}{
				Name:   "name",
				Status: false,
			},
			MustBePassed: true,
		},
		{
			name: "Test case for Update with failure",
			Target: &struct {
				Name   string
				Status bool
			}{
				Name:   "name",
				Status: true,
			},
			Source: &struct{ Id string }{Id: "00000000-0000-0000-0000-000000000001"},
			Expected: &struct {
				Name   string
				Status bool
			}{
				Name:   "name",
				Status: true,
			},
			MustBePassed: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				update, errorUpdate := Update(testCase.Target, testCase.Source)

				if testCase.MustBePassed {
					assert.Equal(t, testCase.Expected, update.Interface())
					assert.Nil(t, errorUpdate)
				} else {
					assert.Nil(t, update)
					assert.NotNil(t, errorUpdate)
				}
			},
		)
	}
}
