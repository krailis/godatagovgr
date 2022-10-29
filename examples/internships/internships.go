package main

import (
	"fmt"
	"log"

	"github.com/krailis/godatagovgr"
	"github.com/krailis/godatagovgr/queryparams/institutions"
)

func main() {
	// Get API token.
	apiToken := godatagovgr.GetAPIToken()

	// Create client.
	client := godatagovgr.NewClient(godatagovgr.NewDefaultConfig(apiToken))

	// Perform request.
	internshipData, err := client.GetInternships(
		&godatagovgr.AcademicQueryParams{
			Year:        fmt.Sprint(2017),
			Institution: institutions.NATIONAL_TECHNICAL_UNIVERSITY_OF_ATHENS,
		},
	)
	if err != nil {
		log.Panicf("An error occurred while fetching internships data: %s", err)
	}

	// Print internship data.
	for _, internships := range *internshipData {
		log.Printf("In year %d, %d internships were conducted in the public sector, "+
			"%d in the private sector and %d in NGOs.",
			internships.Year, internships.PublicSector,
			internships.PrivateSector, internships.NonGovOrganizations)
	}
}
