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
}

// TestVaccinationData tests the retrieval of all the vaccination data.
func TestGetVaccinationData(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetVaccinationData(nil)
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

// TestVaccinationDataByArea tests the retrieval of vaccination data by area.
func TestVaccinationDataByArea(t *testing.T) {
	// Set area.
	area := "ΡΕΘΥΜΝΟΥ"
	// Perform a request, filtering by area.
	vaccinationDays, err := dataGovGrClient.GetVaccinationData(
		&VaccinationQueryParams{Area: area},
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}

	// Ensure that the acquired data include information about the area only.
	for _, vaccinationDay := range *vaccinationDays {
		if vaccinationDay.Area != area {
			t.Fatalf("Filtered for area %q, found area %s.", area, vaccinationDay.Area)
		}
	}
}

// TestAccountantNumbers tests the retrieval of the accountant numbers.
func TestAccountantNumbers(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetAccountantNumbers()
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

// TestEnergyInspectorNumbers tests the retrieval of the energy inspector numbers.
func TestEnergyInspectorNumbers(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetEnergyInspectorNumbers()
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

// TestLawyerNumbers tests the retrieval of the lawyer numbers.
func TestLawyerNumbers(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetLawyerNumbers()
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

// TestLawFirmNumbers tests the retrieval of the law firm numbers.
func TestLawFirmNumbers(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetLawFirmNumbers()
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

// TestPharmacyNumbers tests the retrieval of the pharmacy numbers.
func TestPharmacyNumbers(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetPharmacyNumbers()
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

// TestPharmacistNumbers tests the retrieval of the pharmacist numbers.
func TestPharmacistNumbers(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetPharmacistNumbers()
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

// TestRealtorNumbers tests the retrieval of the realtor numbers.
func TestRealtorNumbers(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetRealtorNumbers()
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

// TestTouristAgencyNumbers tests the retrieval of the tourist agency numbers.
func TestTouristAgencyNumbers(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetTouristAgencyNumbers()
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}
