package impl

import (
	"aurum/core"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	//
	repo := NewInMemoryDpRepository()

	// Create a sample Dp entity
	dp := &core.Dp{
		Provider:  "AWS",
		Product:   "EC2",
		Tier:      "Premium",
		Region:    "us-west-1",
		Customer:  "CustomerA",
		Timestamp: time.Now(),
		Value:     100.0,
	}

	err := repo.Create(dp)
	if err != nil {
		t.Errorf("Unexpected Error occurred while doing Create() %s", err)

	}

}

func TestGet(t *testing.T) {
	repo := NewInMemoryDpRepository()

	// Create a sample Dp entity
	dp := &core.Dp{
		Provider:  "AWS",
		Product:   "EC2",
		Tier:      "Premium",
		Region:    "us-west-1",
		Customer:  "CustomerA",
		Timestamp: time.Now(),
		Value:     100.0,
	}

	repo.Create(dp)
	_, err := repo.Get()
	if err != nil {
		t.Errorf("Unexpected Error occurred while doing Create() %s", err)

	}

}

func TestSearch(t *testing.T) {

	repo := NewInMemoryDpRepository()
	// Create a sample Dp entity
	dp := &core.Dp{
		Provider:  "AWS",
		Product:   "EC2",
		Tier:      "Premium",
		Region:    "us-west-1",
		Customer:  "CustomerA",
		Timestamp: time.Now(),
		Value:     100.0,
	}

	repo.Create(dp)

	results := repo.Search("AWS",
		"EC2",
		"Premium",
		"us-west-1",
		"CustomerA",
	)

	if len(results) != 1 {
		t.Errorf("Expected 1 result got %d", len(results))
	}

}
