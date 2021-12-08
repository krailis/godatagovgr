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
	trafficAccidents, err := client.GetTrafficAccidents()
	if err != nil {
		log.Panicf("An error occurred while fetching traffic accident data: %s", err)
	}

	// Print traffic accident data.
	for _, trafficAccident := range *trafficAccidents {
		log.Printf("In %d, %d serious accidents and %d deadly were recorded in %s.",
			trafficAccident.Year, trafficAccident.SeriousAccidents,
			trafficAccident.DeadlyAccidents, trafficAccident.Jurisdiction)
	}
}
