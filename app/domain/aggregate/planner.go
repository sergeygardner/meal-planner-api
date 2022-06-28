package aggregate

import "github.com/sergeygardner/meal-planner-api/domain/entity"

type Planner struct {
	Entity    *entity.Planner    `bson:"entity" json:"entity"`
	Intervals []*PlannerInterval `bson:"intervals" json:"intervals"`
}

type PlannerInterval struct {
	Entity  *entity.PlannerInterval `bson:"entity" json:"entity"`
	Recipes []*PlannerRecipe        `bson:"recipes" json:"recipes"`
}

type PlannerRecipe struct {
	Entity *entity.PlannerRecipe `bson:"entity" json:"entity"`
	Recipe *Recipe               `bson:"recipe" json:"recipe"`
}

type PlannerCalculation struct {
	Ingredient *entity.Ingredient
	Unit       *entity.Unit
	Amount     int64
}
