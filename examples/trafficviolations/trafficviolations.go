package main

import (
	"log"

	"github.com/krailis/godatagovgr"
	"github.com/krailis/godatagovgr/queryparams/trafficviolations"
)

func main() {
	// Get API token.
	apiToken := godatagovgr.GetAPIToken()

	// Create client.
	client := godatagovgr.NewClient(godatagovgr.NewDefaultConfig(apiToken))

	// Perform request.
	trafficViolations, err := client.GetTrafficViolations(
		&godatagovgr.TrafficViolationQueryParams{
			Year:      "2018",
			Violation: trafficviolations.DRIVING_LICENSE_SUSPENSION,
		},
	)
	if err != nil {
		log.Panicf("An error occurred while fetching traffic violation data: %s", err)
	}

	// Print traffic violation data.
	for _, trafficViolation := range *trafficViolations {
		log.Printf("In %d, %d traffic violations of type %q were observed.",
			trafficViolation.Year, trafficViolation.Count, trafficViolation.Violation)
	}
}
