package main

import (
	"fmt"
	"log"

	"github.com/krailis/godatagovgr"
	"github.com/krailis/godatagovgr/queryparams/crime"
)

func main() {
	// Get API token.
	apiToken := godatagovgr.GetAPIToken()

	// Create client.
	client := godatagovgr.NewClient(godatagovgr.NewDefaultConfig(apiToken))

	// Perform request.
	crimeData, err := client.GetCrimes(
		&godatagovgr.CrimeQueryParams{
			Year:  fmt.Sprint(2018),
			Crime: crime.MANSLAUGHTERS,
		},
	)
	if err != nil {
		log.Panicf("An error occurred while fetching crime data: %s", err)
	}

	// Print crime data.
	for _, crime := range *crimeData {
		log.Printf("In %d, %d %s were recorded.", crime.Year, crime.Committed, crime.Crime)
	}
}
