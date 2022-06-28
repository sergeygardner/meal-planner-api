package kind

const (
	UserActive                                               = true
	UserInActive                                             = false
	UserConfirmationActive                                   = true
	UserConfirmationInActive                                 = false
	UserStatusRegister                UserStatus             = "register"
	UserStatusNeedConfirmation        UserStatus             = "need_confirmation"
	UserStatusDisabled                UserStatus             = "disabled"
	UserRoleAdmin                     UserRole               = "admin"
	UserRoleCommon                    UserRole               = "common"
	UserRoleStatusActive              UserRoleStatus         = "active"
	UserRoleStatusInActive            UserRoleStatus         = "inactive"
	UserToRoleStatusActive            UserToRoleStatus       = "active"
	UserToRoleStatusInActive          UserToRoleStatus       = "inactive"
	UserRightRead                     UserRight              = "read"
	UserRightWrite                    UserRight              = "write"
	RecipeStatusPublished             RecipeStatus           = "published"
	RecipeStatusUnPublished           RecipeStatus           = "unpublished"
	RecipeCategoryStatusPublished     RecipeCategoryStatus   = "published"
	RecipeCategoryStatusUnPublished   RecipeCategoryStatus   = "unpublished"
	RecipeProcessStatusPublished      RecipeProcessStatus    = "published"
	RecipeProcessStatusUnPublished    RecipeProcessStatus    = "unpublished"
	RecipeIngredientStatusPublished   RecipeIngredientStatus = "published"
	RecipeIngredientStatusUnPublished RecipeIngredientStatus = "unpublished"
	RecipeMeasureStatusPublished      RecipeMeasureStatus    = "published"
	RecipeMeasureStatusUnPublished    RecipeMeasureStatus    = "unpublished"
	AltNameStatusPublished            AltNameStatus          = "published"
	AltNameStatusUnPublished          AltNameStatus          = "unpublished"
	UnitStatusPublished               UnitStatus             = "published"
	UnitStatusUnPublished             UnitStatus             = "unpublished"
	CategoryStatusPublished           CategoryStatus         = "published"
	CategoryStatusUnPublished         CategoryStatus         = "unpublished"
	IngredientStatusPublished         IngredientStatus       = "published"
	IngredientStatusUnPublished       IngredientStatus       = "unpublished"
	PictureStatusPublished            PictureStatus          = "published"
	PictureStatusUnPublished          PictureStatus          = "unpublished"
	PlannerStatusActive               PlannerStatus          = "active"
	PlannerStatusInActive             PlannerStatus          = "inactive"
	PlannerIntervalStatusActive       PlannerIntervalStatus  = "active"
	PlannerIntervalStatusInActive     PlannerIntervalStatus  = "inactive"
	PlannerRecipeStatusActive         PlannerRecipeStatus    = "active"
	PlannerRecipeStatusInActive       PlannerRecipeStatus    = "inactive"
)

type UserStatus string

func (us UserStatus) String() string {
	switch us {
	case UserStatusRegister:
		return "register"
	case UserStatusNeedConfirmation:
		return "need_confirmation"
	case UserStatusDisabled:
		return "disabled"
	default:
		return "disabled"
	}
}

type UserRole string

func (ur UserRole) String() string {
	switch ur {
	case UserRoleAdmin:
		return "admin"
	case UserRoleCommon:
		return "common"
	default:
		return "common"
	}
}

type UserRoles []UserRole

type UserRoleStatus string

func (urs UserRoleStatus) String() string {
	switch urs {
	case UserRoleStatusActive:
		return "active"
	case UserRoleStatusInActive:
		return "inactive"
	default:
		return "inactive"
	}
}

type UserToRoleStatus string

func (utrs UserToRoleStatus) String() string {
	switch utrs {
	case UserToRoleStatusActive:
		return "active"
	case UserToRoleStatusInActive:
		return "inactive"
	default:
		return "inactive"
	}
}

type UserRight string

func (ur UserRight) String() string {
	switch ur {
	case UserRightWrite:
		return "write"
	case UserRightRead:
		return "read"
	}
	return "unknown"
}

type RecipeStatus string

func (rs RecipeStatus) String() string {
	switch rs {
	case RecipeStatusPublished:
		return "published"
	case RecipeStatusUnPublished:
		return "unpublished"
	default:
		return "unpublished"
	}
}

type RecipeCategoryStatus string

func (rcs RecipeCategoryStatus) String() string {
	switch rcs {
	case RecipeCategoryStatusPublished:
		return "published"
	case RecipeCategoryStatusUnPublished:
		return "unpublished"
	default:
		return "unpublished"
	}
}

type RecipeIngredientStatus string

func (ris RecipeIngredientStatus) String() string {
	switch ris {
	case RecipeIngredientStatusPublished:
		return "published"
	case RecipeIngredientStatusUnPublished:
		return "unpublished"
	default:
		return "unpublished"
	}
}

type RecipeProcessStatus string

func (rps RecipeProcessStatus) String() string {
	switch rps {
	case RecipeProcessStatusPublished:
		return "published"
	case RecipeProcessStatusUnPublished:
		return "unpublished"
	default:
		return "unpublished"
	}
}

type PictureStatus string

func (rps PictureStatus) String() string {
	switch rps {
	case PictureStatusPublished:
		return "published"
	case PictureStatusUnPublished:
		return "unpublished"
	default:
		return "unpublished"
	}
}

type RecipeMeasureStatus string

func (rms RecipeMeasureStatus) String() string {
	switch rms {
	case RecipeMeasureStatusPublished:
		return "published"
	case RecipeMeasureStatusUnPublished:
		return "unpublished"
	default:
		return "unpublished"
	}
}

type AltNameStatus string

func (ans AltNameStatus) String() string {
	switch ans {
	case AltNameStatusPublished:
		return "published"
	case AltNameStatusUnPublished:
		return "unpublished"
	default:
		return "unpublished"
	}
}

type UnitStatus string

func (us UnitStatus) String() string {
	switch us {
	case UnitStatusPublished:
		return "published"
	case UnitStatusUnPublished:
		return "unpublished"
	default:
		return "unpublished"
	}
}

type CategoryStatus string

func (cs CategoryStatus) String() string {
	switch cs {
	case CategoryStatusPublished:
		return "published"
	case CategoryStatusUnPublished:
		return "unpublished"
	default:
		return "unpublished"
	}
}

type IngredientStatus string

func (is IngredientStatus) String() string {
	switch is {
	case IngredientStatusPublished:
		return "published"
	case IngredientStatusUnPublished:
		return "unpublished"
	default:
		return "unpublished"
	}
}

type PlannerStatus string

func (ps PlannerStatus) String() string {
	switch ps {
	case PlannerStatusActive:
		return "active"
	case PlannerStatusInActive:
		return "inactive"
	default:
		return "inactive"
	}
}

type PlannerIntervalStatus string

func (pis PlannerIntervalStatus) String() string {
	switch pis {
	case PlannerIntervalStatusActive:
		return "active"
	case PlannerIntervalStatusInActive:
		return "inactive"
	default:
		return "inactive"
	}
}

type PlannerRecipeStatus string

func (prs PlannerRecipeStatus) String() string {
	switch prs {
	case PlannerRecipeStatusActive:
		return "active"
	case PlannerRecipeStatusInActive:
		return "inactive"
	default:
		return "inactive"
	}
}
