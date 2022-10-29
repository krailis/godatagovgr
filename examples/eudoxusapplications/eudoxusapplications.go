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
	eudoxusApplicationData, err := client.GetEudoxusApplications(
		&godatagovgr.AcademicQueryParams{
			Year:        fmt.Sprint(2017),
			Institution: institutions.NATIONAL_TECHNICAL_UNIVERSITY_OF_ATHENS,
		},
	)
	if err != nil {
		log.Panicf("An error occurred while fetching Eudoxus application data: %s", err)
	}

	// Print eudoxus application data.
	for _, eudoxusApplications := range *eudoxusApplicationData {
		log.Printf("In year %d, %d students received their books from %s (%s).",
			eudoxusApplications.Year, eudoxusApplications.StudentWithDeliveries,
			eudoxusApplications.Department, eudoxusApplications.Institution)
	}
}
