package main

import (
	"fmt"
	"log"

	"github.com/krailis/godatagovgr"
)

func main() {
	// Get API token.
	apiToken := godatagovgr.GetAPIToken()

	// Create client.
	client := godatagovgr.NewClient(godatagovgr.NewDefaultConfig(apiToken))

	// Perform request.
	casinoTickets, err := client.GetCasinoTickets(
		&godatagovgr.YearQueryParams{
			Year: fmt.Sprint(2016),
		},
	)
	if err != nil {
		log.Panicf("An error occurred while fetching casino ticker data: %s", err)
	}

	// Print casino ticket data.
	for _, casinoTicketYear := range *casinoTickets {
		log.Printf("In %d, %d casino tickets were sold.",
			casinoTicketYear.Year, casinoTicketYear.Tickets)
	}
}
