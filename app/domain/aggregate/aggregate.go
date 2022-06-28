package aggregate

import DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"

type Picture struct {
	AltNames []*DomainEntity.AltName `json:"alt_names"`
	Entity   *DomainEntity.Picture   `json:"entity"`
}
type Category struct {
	AltNames []*DomainEntity.AltName `json:"alt_names"`
	Entity   *DomainEntity.Category  `json:"entity"`
	Pictures []*Picture              `json:"pictures"`
}
