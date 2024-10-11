package impl

import (
	"aurum/core"
	"fmt"
	"os"
	"testing"
	"time"
)

var testDp *core.Dp

func TestMain(m *testing.M) {
	testDp = core.NewDp(
		"aws:ec2",
		"capture-mobile",
		"dev",
		"us-west-2",
		"internal",
		time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		123.4,
	)
	fmt.Printf("Test Datapoint %s", testDp.Provider)
	code := m.Run()
	os.Exit(code)

}

func TestCreate(t *testing.T) {
	repo := NewInMemoryDpRepository()
	err := repo.Insert(testDp)
	if err != nil {
		t.Errorf("Unexpected Error occurred while doing Create() %s", err)

	}

}

func TestRetrieve(t *testing.T) {
	repo := NewInMemoryDpRepository()
	repo.Insert(testDp)
	d, err := repo.Retrieve()
	if err != nil {
		t.Errorf("Unexpected Error occurred while doing Create() %s", err)
	}

	// Check if repo contains exactly 1 element
	n := len(d)
	if n != 1 {
		t.Errorf("Expected length 1 got %d", n)
	}

}

func TestSearch(t *testing.T) {

	repo := NewInMemoryDpRepository()
	repo.Insert(testDp)

	// Test successful search
	results := repo.Search(
		"aws:ec2",
		"capture-mobile",
		"dev",
		"us-west-2",
		"internal",
	)

	if len(results) != 1 {
		t.Errorf("Expected 1 result got %d", len(results))
	}

}
