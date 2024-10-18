package impl

import (
	"aurum/core"
)

// Simply store DP as a list
type InMemoryCostRepository struct {
	data []core.Cost
}

func NewInMemoryCostRepository() *InMemoryCostRepository {
	return &InMemoryCostRepository{
		data: []core.Cost{},
	}
}

func (r *InMemoryCostRepository) Insert(dp *core.Cost) error {
	// TODO: CHeck if exists
	r.data = append(r.data, *dp)
	return nil

}

func (r *InMemoryCostRepository) Retrieve() ([]core.Cost, error) {
	return r.data, nil
}

func (r *InMemoryCostRepository) Search(
	provider string,
	title string,
	product string,
	tier string,
	region string,
	customer string,
) []core.Cost {
	var result []core.Cost

	for _, d := range r.data {
		if d.Provider == provider && d.Resource == title && d.Product == product && d.Tier == tier && d.Region == region && d.Customer == customer {
			result = append(result, d)
		}

	}
	return result
}
