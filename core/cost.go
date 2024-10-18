package core

import (
	"time"
)

type Cost struct {
	Provider  string    `json:"provider"`
	Resource  string    `json:"resource"`
	Product   string    `json:"product"`
	Tier      string    `json:"tier"`
	Region    string    `json:"region"`
	Customer  string    `json:"customer"`
	Timestamp time.Time `json:"timestamp"`
	Value     float64   `json:"value"`
}

func NewCost(
	provider string,
	resource string,
	product string,
	tier string,
	region string,
	customer string,
	timestamp time.Time,
	value float64,
) *Cost {

	return &Cost{
		Provider:  provider,
		Resource:  resource,
		Product:   product,
		Tier:      tier,
		Region:    region,
		Customer:  customer,
		Timestamp: timestamp,
		Value:     value,
	}
}
