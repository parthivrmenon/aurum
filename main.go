package main

import (
	"aurum/core"
	"aurum/impl"
	"encoding/json"
	"fmt"
	"time"
)

// Function to print Dp struct as JSON
func printDpAsJSON(dp core.Dp) {
	// Convert the struct to JSON
	jsonData, err := json.Marshal(dp)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	// Print the JSON
	fmt.Println(string(jsonData))
}

func main() {
	// Create a new metric
	metric_one := core.NewDp(
		"aws",
		"capture mobile",
		"dev",
		"us-west-2",
		"multicustomer",
		time.Now(),
		123.45,
	)

	metric_two := core.NewDp(
		"Bar",
		"ProductB",
		"Prod",
		"us-east-1",
		"Customer",
		time.Now(),
		123.45,
	)

	// Print the metric
	//fmt.Printf("Metric: %+v\n", metric)

	var repo core.DpRepository = impl.NewInMemoryDpRepository()
	repo.Insert(metric_one)
	repo.Insert(metric_two)

	results, _ := repo.Retrieve()

	for _, d := range results {
		printDpAsJSON(d)
	}

}
