package password

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func init() {
	salt, errorPasswordSalt := os.LookupEnv("PASSWORD_SALT")

	if !errorPasswordSalt {
		panic("the password salt must be presented in .env files")
	}

	SetPasswordSalt(salt)
}

func TestCheckHashedPassword(t *testing.T) {
	password := "password"
	hashedPassword, _ := CastPassword(password)

	tests := []struct {
		name     string
		password string
		expected bool
	}{
		{
			name:     "Test case with hashed password",
			password: string(hashedPassword),
			expected: true,
		},
		{
			name:     "Test case with plain password",
			password: password,
			expected: false,
		},
		{
			name:     "Test case with empty password",
			password: "",
			expected: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, _ := CheckHashedPassword(&testCase.password)
				assert.Equal(t, testCase.expected, actual)
			},
		)
	}
}
