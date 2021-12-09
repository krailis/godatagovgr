package godatagovgr

import (
	"log"
	"os"
	"testing"
)

var dataGovGrClient *DataGovGrClient

// TestMain sets up the the tests that follow, by initializing the DataGovGr client.
func TestMain(m *testing.M) {
	// Retrieve API key from the environment.
	apiToken, ok := os.LookupEnv("DATAGOVGR_API_TOKEN")
	if !ok || len(apiToken) == 0 {
		log.Panic("The API token for data.gov.gr has not been properly set.")
	}
	// Initialize client.
	dataGovGrClient = NewClient(NewDefaultConfig(apiToken))

	// Run tests.
	os.Exit(m.Run())
}

// TestGetAcademicProfessors tests the retrieval of the academic professors.
func TestGetAcademicProfessors(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetAcademicProfessors()
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

// TestGetAccountants tests the retrieval of the accountant numbers.
func TestGetAccountants(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetAccountants()
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

// TestGetAuditors tests the retrieval of the auditor numbers.
func TestGetAuditors(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetAuditors()
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

// TestGetCrimes tests the retrieval of the crime data.
func TestGetCrimes(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetCrimes()
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

// TestGetDendists tests the retrieval of the dendist numbers.
func TestGetDendists(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetDendists()
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

// TestGetDoctors tests the retrieval of the doctor numbers.
func TestGetDoctors(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetDoctors()
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

// TestEnergyInspectors tests the retrieval of the energy inspector numbers.
func TestGetEnergyInspectors(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetEnergyInspectors()
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

// TestGetEudoxusApplications tests the retrieval of Eudoxus application data.
func TestGetEudoxusApplications(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetEudoxusApplications()
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

// TestGetFinancialCrimes tests the retrieval of financial crime data.
func TestGetFinancialCrimes(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetFinancialCrimes()
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

// TestHCGIndicents tests the retrieval of the HCG Indicent data.
func TestGetHCGIndidents(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetHCGIncidents()
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

// TestGetInternetTraffic tests the retrieval of internet traffic data.
func TestGetInternetTraffic(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetInternetTraffic(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

// TestGetInternetTrafficByDates tests the retrieval of internet
// traffic data within a specified time frame.
func TestGetInternetTrafficByDates(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetInternetTraffic(
		&InternetTrafficQueryParams{
			DateFrom: "2021-01-01",
			DateTo:   "2021-11-30",
		},
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

// TestLawyerNumbers tests the retrieval of the lawyer numbers.
func TestLawyerNumbers(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetLawyers()
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

// TestGetLawFirmNumbers tests the retrieval of the law firm numbers.
func TestGetLawFirms(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetLawFirms()
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

// TestGetPharmacies tests the retrieval of the pharmacy numbers.
func TestGetPharmacies(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetPharmacies()
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

// TestGetPharmacists tests the retrieval of the pharmacist numbers.
func TestPharmacists(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetPharmacists()
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

// TestGetRealtors tests the retrieval of the realtor numbers.
func TestGetRealtors(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetRealtors()
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

// TestGetTouristAgencies tests the retrieval of the tourist agency numbers.
func TestTouristAgencies(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetTouristAgencies()
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

// TestGetTrafficAccidents tests the retrieval of traffic accident data.
func TestGetTrafficAccidents(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetTrafficAccidents()
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

// TestGetTrafficViolations tests the retrieval of traffic violation data.
func TestGetTrafficViolations(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetTrafficViolations()
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

// TestGetVaccinations tests the retrieval of all the vaccination data.
func TestGetVaccinations(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetVaccinations(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

// TestVaccinationDataByArea tests the retrieval of vaccination data by area.
func TestGetVaccinationsByArea(t *testing.T) {
	// Set area.
	area := "ΡΕΘΥΜΝΟΥ"
	// Perform a request, filtering by area.
	vaccinationDays, err := dataGovGrClient.GetVaccinations(
		&VaccinationQueryParams{Area: area},
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}

	// Ensure that the acquired data include information about the area only.
	for _, vaccinationDay := range *vaccinationDays {
		if vaccinationDay.Area != area {
			t.Errorf("Filtered for area %q, found area %s.", area, vaccinationDay.Area)
		}
	}
}
