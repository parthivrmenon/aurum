package impl

import (
	"aurum/core"
)

// Simply store DP as a list
type InMemoryDpRepository struct {
	data []core.Dp
}

func NewInMemoryDpRepository() *InMemoryDpRepository {
	return &InMemoryDpRepository{
		data: []core.Dp{},
	}
}

func (r *InMemoryDpRepository) Insert(dp *core.Dp) error {
	// TODO: CHeck if exists
	r.data = append(r.data, *dp)
	return nil

}

func (r *InMemoryDpRepository) Retrieve() ([]core.Dp, error) {
	return r.data, nil
}

func (r *InMemoryDpRepository) Search(provider string, product string, tier string, region string, customer string) []core.Dp {
	var result []core.Dp

	for _, d := range r.data {
		if d.Provider == provider && d.Product == product && d.Tier == tier && d.Region == region && d.Customer == customer {
			result = append(result, d)
		}

	}
	return result
}
