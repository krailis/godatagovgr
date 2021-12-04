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
	vaccinationData, err := client.GetVaccinationData(
		&godatagovgr.VaccinationQueryParams{
			DateFrom: "2021-01-01",
			DateTo:   "2021-11-30",
			Area:     "ΡΕΘΥΜΝΟΥ",
		},
	)
	if err != nil {
		fmt.Printf("An error occurred while fetching vaccination data: %s", err)
	}

	// Print vaccination data.
	for _, vaccinationDay := range *vaccinationData {
		fmt.Printf("On %s, %d vaccinations were conducted in regional unit %q.\n",
			vaccinationDay.ReferenceDate, vaccinationDay.DayTotal, vaccinationDay.Area)
	}
}
