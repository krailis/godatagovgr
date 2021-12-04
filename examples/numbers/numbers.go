package main

import (
	"fmt"
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
	numbers, err := client.GetAccountantNumbers()
	if err != nil {
		fmt.Printf("An error occurred while fetching accountant numbers: %s", err)
	}

	// Display accountant data.
	for _, accRecord := range *numbers {
		fmt.Printf("In %s of year %d there were %d accountants, with %d new entries and %d exits.\n",
			accRecord.Quarter, accRecord.Year, accRecord.Active, accRecord.Entrants, accRecord.Exits)
	}
}
