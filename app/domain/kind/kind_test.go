package kind

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserStatus(t *testing.T) {
	tests := []struct {
		name     string
		status   UserStatus
		expected string
	}{
		{
			name:     "Test case with user status is register",
			status:   UserStatusRegister,
			expected: "register",
		},
		{
			name:     "Test case with user status is need_confirmation",
			status:   UserStatusNeedConfirmation,
			expected: "need_confirmation",
		},
		{
			name:     "Test case with user status is disabled",
			status:   UserStatusDisabled,
			expected: "disabled",
		},
		{
			name:     "Test case with user status is empty",
			status:   "",
			expected: "disabled",
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				assert.Equal(t, testCase.expected, testCase.status.String())
			},
		)
	}
}

func TestUserRole(t *testing.T) {
	tests := []struct {
		name     string
		role     UserRole
		expected string
	}{
		{
			name:     "Test case with user role is admin",
			role:     UserRoleAdmin,
			expected: "admin",
		},
		{
			name:     "Test case with user role is common",
			role:     UserRoleCommon,
			expected: "common",
		},
		{
			name:     "Test case with user role is empty",
			role:     "",
			expected: "common",
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				assert.Equal(t, testCase.expected, testCase.role.String())
			},
		)
	}
}

func TestUserRoles(t *testing.T) {
	tests := []struct {
		name     string
		roles    UserRoles
		expected []string
	}{
		{
			name:     "Test case with user role is admin",
			roles:    UserRoles{UserRoleAdmin},
			expected: []string{"admin"},
		},
		{
			name:     "Test case with user role is common",
			roles:    UserRoles{UserRoleCommon},
			expected: []string{"common"},
		},
		{
			name:     "Test case with user role is empty",
			roles:    UserRoles{""},
			expected: []string{""},
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				var actual []string

				for _, value := range testCase.roles {
					actual = append(actual, string(value))
				}

				assert.Equal(t, testCase.expected, actual)
			},
		)
	}
}

func TestUserRolesAsInterface(t *testing.T) {
	tests := []struct {
		name     string
		roles    any
		expected UserRoles
	}{
		{
			name:     "Test case with user role is admin",
			roles:    []interface{}{"admin"},
			expected: UserRoles{"admin"},
		},
		{
			name:     "Test case with user role is common",
			roles:    []interface{}{"common"},
			expected: UserRoles{"common"},
		},
		{
			name:     "Test case with user role is empty",
			roles:    []interface{}{""},
			expected: UserRoles{""},
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				rolesInterfaces := testCase.roles.([]interface{})
				actual := make(UserRoles, len(rolesInterfaces))

				for i, value := range rolesInterfaces {
					actual[i] = UserRole(value.(string))
				}

				assert.Equal(t, testCase.expected, actual)
			},
		)
	}
}
func TestUserRight(t *testing.T) {
	tests := []struct {
		name     string
		right    UserRight
		expected string
	}{
		{
			name:     "Test case with user right is read",
			right:    UserRightRead,
			expected: "read",
		},
		{
			name:     "Test case with user right is write",
			right:    UserRightWrite,
			expected: "write",
		},
		{
			name:     "Test case with user right is empty",
			right:    "",
			expected: "unknown",
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				assert.Equal(t, testCase.expected, testCase.right.String())
			},
		)
	}
}

func TestRecipeStatus(t *testing.T) {
	tests := []struct {
		name     string
		status   RecipeStatus
		expected string
	}{
		{
			name:     "Test case with recipe status is published",
			status:   RecipeStatusPublished,
			expected: "published",
		},
		{
			name:     "Test case with recipe status is unpublished",
			status:   RecipeStatusUnPublished,
			expected: "unpublished",
		},
		{
			name:     "Test case with recipe status is empty",
			status:   "",
			expected: "unpublished",
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				assert.Equal(t, testCase.expected, testCase.status.String())
			},
		)
	}
}

func TestRecipeCategoryStatus(t *testing.T) {
	tests := []struct {
		name     string
		status   RecipeCategoryStatus
		expected string
	}{
		{
			name:     "Test case with recipe category status is published",
			status:   RecipeCategoryStatusPublished,
			expected: "published",
		},
		{
			name:     "Test case with recipe category status is unpublished",
			status:   RecipeCategoryStatusUnPublished,
			expected: "unpublished",
		},
		{
			name:     "Test case with recipe category status is empty",
			status:   "",
			expected: "unpublished",
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				assert.Equal(t, testCase.expected, testCase.status.String())
			},
		)
	}
}

func TestRecipeIngredientStatus(t *testing.T) {
	tests := []struct {
		name     string
		status   RecipeIngredientStatus
		expected string
	}{
		{
			name:     "Test case with recipe ingredient status is published",
			status:   RecipeIngredientStatusPublished,
			expected: "published",
		},
		{
			name:     "Test case with recipe ingredient status is unpublished",
			status:   RecipeIngredientStatusUnPublished,
			expected: "unpublished",
		},
		{
			name:     "Test case with recipe ingredient status is empty",
			status:   "",
			expected: "unpublished",
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				assert.Equal(t, testCase.expected, testCase.status.String())
			},
		)
	}
}

func TestRecipeProcessStatus(t *testing.T) {
	tests := []struct {
		name     string
		status   RecipeProcessStatus
		expected string
	}{
		{
			name:     "Test case with recipe process status is published",
			status:   RecipeProcessStatusPublished,
			expected: "published",
		},
		{
			name:     "Test case with recipe process status is unpublished",
			status:   RecipeProcessStatusUnPublished,
			expected: "unpublished",
		},
		{
			name:     "Test case with recipe process status is empty",
			status:   "",
			expected: "unpublished",
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				assert.Equal(t, testCase.expected, testCase.status.String())
			},
		)
	}
}

func TestRecipePictureStatus(t *testing.T) {
	tests := []struct {
		name     string
		status   PictureStatus
		expected string
	}{
		{
			name:     "Test case with recipe picture status is published",
			status:   PictureStatusPublished,
			expected: "published",
		},
		{
			name:     "Test case with recipe picture status is unpublished",
			status:   PictureStatusUnPublished,
			expected: "unpublished",
		},
		{
			name:     "Test case with recipe picture status is empty",
			status:   "",
			expected: "unpublished",
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				assert.Equal(t, testCase.expected, testCase.status.String())
			},
		)
	}
}

func TestRecipeMeasureStatus(t *testing.T) {
	tests := []struct {
		name     string
		status   RecipeMeasureStatus
		expected string
	}{
		{
			name:     "Test case with recipe measure status is published",
			status:   RecipeMeasureStatusPublished,
			expected: "published",
		},
		{
			name:     "Test case with recipe measure status is unpublished",
			status:   RecipeMeasureStatusUnPublished,
			expected: "unpublished",
		},
		{
			name:     "Test case with recipe measure status is empty",
			status:   "",
			expected: "unpublished",
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				assert.Equal(t, testCase.expected, testCase.status.String())
			},
		)
	}
}

func TestAltNameStatus(t *testing.T) {
	tests := []struct {
		name     string
		status   AltNameStatus
		expected string
	}{
		{
			name:     "Test case with recipe alt name status is published",
			status:   AltNameStatusPublished,
			expected: "published",
		},
		{
			name:     "Test case with recipe alt name status is unpublished",
			status:   AltNameStatusUnPublished,
			expected: "unpublished",
		},
		{
			name:     "Test case with recipe alt name status is empty",
			status:   "",
			expected: "unpublished",
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				assert.Equal(t, testCase.expected, testCase.status.String())
			},
		)
	}
}

func TestUnitStatus(t *testing.T) {
	tests := []struct {
		name     string
		status   UnitStatus
		expected string
	}{
		{
			name:     "Test case with unit status is published",
			status:   UnitStatusPublished,
			expected: "published",
		},
		{
			name:     "Test case with unit status is unpublished",
			status:   UnitStatusUnPublished,
			expected: "unpublished",
		},
		{
			name:     "Test case with unit status is empty",
			status:   "",
			expected: "unpublished",
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				assert.Equal(t, testCase.expected, testCase.status.String())
			},
		)
	}
}

func TestCategoryStatus(t *testing.T) {
	tests := []struct {
		name     string
		status   CategoryStatus
		expected string
	}{
		{
			name:     "Test case with category status is published",
			status:   CategoryStatusPublished,
			expected: "published",
		},
		{
			name:     "Test case with category status is unpublished",
			status:   CategoryStatusUnPublished,
			expected: "unpublished",
		},
		{
			name:     "Test case with category status is empty",
			status:   "",
			expected: "unpublished",
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				assert.Equal(t, testCase.expected, testCase.status.String())
			},
		)
	}
}

func TestIngredientStatus(t *testing.T) {
	tests := []struct {
		name     string
		status   IngredientStatus
		expected string
	}{
		{
			name:     "Test case with ingredient status is published",
			status:   IngredientStatusPublished,
			expected: "published",
		},
		{
			name:     "Test case with ingredient status is unpublished",
			status:   IngredientStatusUnPublished,
			expected: "unpublished",
		},
		{
			name:     "Test case with ingredient status is empty",
			status:   "",
			expected: "unpublished",
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				assert.Equal(t, testCase.expected, testCase.status.String())
			},
		)
	}
}

func TestPictureStatus(t *testing.T) {
	tests := []struct {
		name     string
		status   PictureStatus
		expected string
	}{
		{
			name:     "Test case with picture status is published",
			status:   PictureStatusPublished,
			expected: "published",
		},
		{
			name:     "Test case with picture status is unpublished",
			status:   PictureStatusUnPublished,
			expected: "unpublished",
		},
		{
			name:     "Test case with picture status is empty",
			status:   "",
			expected: "unpublished",
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				assert.Equal(t, testCase.expected, testCase.status.String())
			},
		)
	}
}

func TestPlannerStatus(t *testing.T) {
	tests := []struct {
		name     string
		status   PlannerStatus
		expected string
	}{
		{
			name:     "Test case with planner status is active",
			status:   PlannerStatusActive,
			expected: "active",
		},
		{
			name:     "Test case with planner status is inactive",
			status:   PlannerStatusInActive,
			expected: "inactive",
		},
		{
			name:     "Test case with planner status is empty",
			status:   "",
			expected: "inactive",
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				assert.Equal(t, testCase.expected, testCase.status.String())
			},
		)
	}
}

func TestPlannerIntervalStatus(t *testing.T) {
	tests := []struct {
		name     string
		status   PlannerIntervalStatus
		expected string
	}{
		{
			name:     "Test case with planner interval status is active",
			status:   PlannerIntervalStatusActive,
			expected: "active",
		},
		{
			name:     "Test case with planner interval status is inactive",
			status:   PlannerIntervalStatusInActive,
			expected: "inactive",
		},
		{
			name:     "Test case with planner interval status is empty",
			status:   "",
			expected: "inactive",
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				assert.Equal(t, testCase.expected, testCase.status.String())
			},
		)
	}
}

func TestPlannerRecipeStatus(t *testing.T) {
	tests := []struct {
		name     string
		status   PlannerRecipeStatus
		expected string
	}{
		{
			name:     "Test case with planner recipe status is active",
			status:   PlannerRecipeStatusActive,
			expected: "active",
		},
		{
			name:     "Test case with planner recipe status is inactive",
			status:   PlannerRecipeStatusInActive,
			expected: "inactive",
		},
		{
			name:     "Test case with planner recipe status is empty",
			status:   "",
			expected: "inactive",
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				assert.Equal(t, testCase.expected, testCase.status.String())
			},
		)
	}
}
