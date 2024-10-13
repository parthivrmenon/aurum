package core

import (
	"time"
)

type Dp struct {
	Provider  string    `json:"provider"`
	Product   string    `json:"product"`
	Tier      string    `json:"tier"`
	Region    string    `json:"region"`
	Customer  string    `json:"customer"`
	Timestamp time.Time `json:"timestamp"`
	Value     float64   `json:"value"`
}

func NewDp(
	provider string,
	product string,
	tier string,
	region string,
	customer string,
	timestamp time.Time,
	value float64,
) *Dp {

	return &Dp{
		Provider:  provider,
		Product:   product,
		Tier:      tier,
		Region:    region,
		Customer:  customer,
		Timestamp: timestamp,
		Value:     value,
	}
}
