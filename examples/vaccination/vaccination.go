package main

import (
	"log"

	"github.com/krailis/godatagovgr"
	"github.com/krailis/godatagovgr/queryparams/vaccination"
)

func main() {
	// Get API token.
	apiToken := godatagovgr.GetAPIToken()

	// Create client.
	client := godatagovgr.NewClient(godatagovgr.NewDefaultConfig(apiToken))

	// Perform request.
	vaccinationData, err := client.GetVaccinations(
		&godatagovgr.VaccinationQueryParams{
			DateFrom: "2021-01-01",
			DateTo:   "2021-11-30",
			Area:     vaccination.RETHIMNO,
		},
	)
	if err != nil {
		log.Panicf("An error occurred while fetching vaccination data: %s", err)
	}

	// Print vaccination data.
	for _, vaccinationDay := range *vaccinationData {
		log.Printf("On %s, %d vaccinations were conducted in regional unit %q.",
			vaccinationDay.ReferenceDate, vaccinationDay.DayTotal, vaccinationDay.Area)
	}
}
