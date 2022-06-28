package aggregate

import DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"

type Recipe struct {
	AltNames    []*DomainEntity.AltName `json:"alt_names"`
	Categories  []*RecipeCategory       `json:"categories"`
	Entity      *DomainEntity.Recipe    `json:"entity"`
	Ingredients []*RecipeIngredient     `json:"ingredients"`
	Processes   []*RecipeProcess        `json:"processes"`
	Pictures    []*Picture              `json:"pictures"`
}
type RecipeCategory struct {
	Derive *Category                    `json:"derive"`
	Entity *DomainEntity.RecipeCategory `json:"entity"`
}
type RecipeIngredient struct {
	AltNames []*DomainEntity.AltName        `json:"alt_names"`
	Derive   *DomainEntity.Ingredient       `json:"derive"`
	Entity   *DomainEntity.RecipeIngredient `json:"entity"`
	Measures []*RecipeMeasure               `json:"measures"`
	Pictures []*Picture                     `json:"pictures"`
}
type RecipeMeasure struct {
	AltNames []*DomainEntity.AltName     `json:"alt_names"`
	Entity   *DomainEntity.RecipeMeasure `json:"entity"`
	Unit     *DomainEntity.Unit          `json:"unit"`
}
type RecipeProcess struct {
	AltNames []*DomainEntity.AltName     `json:"alt_names"`
	Entity   *DomainEntity.RecipeProcess `json:"entity"`
	Pictures []*Picture                  `json:"pictures"`
}
