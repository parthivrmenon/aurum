package core

import (
	"testing"
	"time"
)

func TestNewDp(t *testing.T) {
	product := "foo"
	provider := "foo"
	tier := "dev"
	region := "global"
	customer := "multicustomer"
	timestamp := time.Now()
	value := 123.4

	dp := NewDp(provider, product, tier, region, customer, timestamp, value)

	if dp.Product != product {
		t.Errorf("Expected %s but got %s instead.", product, dp.Product)
	}
	if dp.Provider != provider {
		t.Errorf("Expected %s but got %s instead.", provider, dp.Provider)
	}

	if dp.Tier != tier {
		t.Errorf("Expected %s but got %s instead.", tier, dp.Tier)
	}

	if dp.Region != region {
		t.Errorf("Expected %s but got %s instead.", region, dp.Region)
	}

	if dp.Customer != customer {
		t.Errorf("Expected %s but got %s instead.", customer, dp.Customer)
	}

	if dp.Timestamp != timestamp {
		t.Errorf("Expected %s but got %s instead.", timestamp, dp.Timestamp)
	}

	if dp.Value != value {
		t.Errorf("Expected %f but got %f instead.", value, dp.Value)
	}
}
