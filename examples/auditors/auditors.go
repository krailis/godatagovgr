package main

import (
	"log"

	"github.com/krailis/godatagovgr"
)

func main() {
	// Get API token.
	apiToken := godatagovgr.GetAPIToken()

	// Create client.
	client := godatagovgr.NewClient(godatagovgr.NewDefaultConfig(apiToken))

	// Perform request.
	auditorData, err := client.GetAuditors()
	if err != nil {
		log.Panicf("An error occurred while fetching auditor data: %s", err)
	}

	// Print auditor data.
	for _, auditors := range *auditorData {
		log.Printf("In quarter %s of year %d, active auditors were %d with %d entrants and %d exits.",
			auditors.Quarter, auditors.Year, auditors.Active, auditors.Entrants, auditors.Exits)
	}
}
