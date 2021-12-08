package main

import (
	"log"
	"os"

	"github.com/krailis/godatagovgr"
)

func main() {
	// Retrieve API token from the environment.
	apiToken, ok := os.LookupEnv("DATAGOVGR_API_TOKEN")
	if !ok || len(apiToken) == 0 {
		log.Panic("The API token for data.gov.gr has not been properly set.")
	}

	// Create client.
	client := godatagovgr.NewClient(godatagovgr.NewDefaultConfig(apiToken))

	// Perform request.
	trafficViolations, err := client.GetTrafficViolations()
	if err != nil {
		log.Panicf("An error occurred while fetching traffic violation data: %s", err)
	}

	// Print traffic violation data.
	for _, trafficViolation := range *trafficViolations {
		log.Printf("In %d, %d traffic violations of type %q were observed.",
			trafficViolation.Year, trafficViolation.Count, trafficViolation.Violation)
	}
}
