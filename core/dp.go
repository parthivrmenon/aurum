package core

import (
	"time"
)

type Dp struct {
	Provider  string
	Product   string
	Tier      string
	Region    string
	Customer  string
	Timestamp time.Time
	Value     float64
}

func NewDp(provider string, product string, tier string, region string, customer string, timestamp time.Time, value float64) *Dp {
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
