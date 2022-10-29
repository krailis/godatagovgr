package godatagovgr

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/krailis/godatagovgr/queryparams/crime"
	"github.com/krailis/godatagovgr/queryparams/electricityconsumption"
	"github.com/krailis/godatagovgr/queryparams/energybalance"
	"github.com/krailis/godatagovgr/queryparams/financialcrime"
	"github.com/krailis/godatagovgr/queryparams/institutions"
	"github.com/krailis/godatagovgr/queryparams/landregistrystatus"
	"github.com/krailis/godatagovgr/queryparams/prefectures"
	"github.com/krailis/godatagovgr/queryparams/vaccination"
)

var dataGovGrClient *DataGovGrClient

func TestMain(m *testing.M) {
	// Get API token.
	apiToken := GetAPIToken()
	// Initialize client.
	dataGovGrClient = NewClient(
		&DataGovGrConfig{
			apiToken:                     apiToken,
			retryCount:                   3,
			retryWaitTime:                15,
			retryMaxWaitTime:             120,
			retryTooManyRequestsWaitTime: 90,
			timeout:                      time.Duration(20) * time.Second,
		},
	)

	// Run tests.
	os.Exit(m.Run())
}

func TestGetAcademicProfessors(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetAcademicProfessors(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

func TestGetAcademicProfessorsByYear(t *testing.T) {
	// Set year.
	year := 2019
	// Perform a request with query params.
	academicProfessorData, err := dataGovGrClient.GetAcademicProfessors(
		&AcademicQueryParams{
			Year: fmt.Sprint(year),
		},
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
	// Check if there is another year.
	for _, academicProfessor := range *academicProfessorData {
		if academicProfessor.Year != year {
			t.Errorf("Filtered for year %d, found year %d.", year, academicProfessor.Year)
		}
	}
}

func TestGetAcademicProfessorsByInstitution(t *testing.T) {
	// Set Institution.
	institution := institutions.NATIONAL_TECHNICAL_UNIVERSITY_OF_ATHENS
	// Perform a request with query params.
	academicProfessorData, err := dataGovGrClient.GetAcademicProfessors(
		&AcademicQueryParams{
			Institution: institution,
		},
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
	// Check if there is another institution.
	for _, academicProfessor := range *academicProfessorData {
		if academicProfessor.Institution != institution {
			t.Errorf("Filtered for institution %s, found institution %s.",
				institution, academicProfessor.Institution)
		}
	}
}

func TestGetAccountants(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetAccountants()
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

func TestGetAuditors(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetAuditors()
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

func TestGetCasinoTickets(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetCasinoTickets(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

func TestGetCasinoTicketsByYear(t *testing.T) {
	// Set year.
	year := 2016
	// Perform a request with query params.
	casinoTicketData, err := dataGovGrClient.GetCasinoTickets(
		&YearQueryParams{
			Year: fmt.Sprint(year),
		},
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
	// Check if there is another year.
	for _, casinoTickets := range *casinoTicketData {
		if casinoTickets.Year != year {
			t.Errorf("Filtered for year %d, found year %d.", year, casinoTickets.Year)
		}
	}
}

func TestGetCrimes(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetCrimes(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

func TestGetCrimesByYear(t *testing.T) {
	// Set year.
	year := 2018
	// Perform a request with query params.
	crimeData, err := dataGovGrClient.GetCrimes(
		&CrimeQueryParams{
			Year: fmt.Sprint(year),
		},
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
	// Check if there is another year.
	for _, crime := range *crimeData {
		if crime.Year != year {
			t.Errorf("Filtered for year %d, found year %d.", year, crime.Year)
		}
	}
}

func TestGetCrimesByCrime(t *testing.T) {
	// Set Crime.
	crime := crime.MANSLAUGHTERS
	// Perform a request with query params.
	crimeData, err := dataGovGrClient.GetCrimes(
		&CrimeQueryParams{
			Crime: crime,
		},
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
	// Check if there is another crime.
	for _, crimes := range *crimeData {
		if crimes.Crime != crime {
			t.Errorf("Filtered for crime %s, found crime %s.", crime, crimes.Crime)
		}
	}
}

func TestGetDendists(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetDendists()
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

func TestGetDoctors(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetDoctors()
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

func TestGetElectricityConsumption(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetElectricityConsumption(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

func TestGetElectricityConsumptionByDates(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetElectricityConsumption(
		&ElectricityConsumptionQueryParams{
			DateFrom: "2020-08-09",
			DateTo:   "2020-08-20",
		},
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

func TestGetElectricityConsumptionByArea(t *testing.T) {
	// Define area.
	area := electricityconsumption.CRETE

	// Perform a request without query params.
	electricityConsumptionDays, err := dataGovGrClient.GetElectricityConsumption(
		&ElectricityConsumptionQueryParams{
			Area: area,
		},
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}

	// Ensure that the acquired data include information about the area only.
	for _, electricityConsumptionDay := range *electricityConsumptionDays {
		if electricityConsumptionDay.Area != area {
			t.Errorf("Filtered for area %q, found area %s.",
				area, electricityConsumptionDay.Area)
		}
	}
}

func TestGetEnergyBalance(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetEnergyBalance(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

func TestGetEnergyBalanceByDates(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetEnergyBalance(
		&EnergyBalanceQueryParams{
			DateFrom: "2021-12-10",
			DateTo:   "2021-12-11",
		},
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

func TestGetEnergyBalanceByFuel(t *testing.T) {
	// Define fuel.
	fuel := energybalance.LIGNITE

	// Perform a request without query params.
	energyBalanceDays, err := dataGovGrClient.GetEnergyBalance(
		&EnergyBalanceQueryParams{
			Fuel: fuel,
		},
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}

	// Ensure that the acquired data include information about the fuel only.
	for _, energyBalanceDay := range *energyBalanceDays {
		if energyBalanceDay.Fuel != fuel {
			t.Errorf("Filtered for fuel %q, found fuel %s.", fuel, energyBalanceDay.Fuel)
		}
	}
}

func TestGetEnergyInspectors(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetEnergyInspectors()
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

func TestGetEudoxusApplications(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetEudoxusApplications(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

func TestGetFinancialCrimes(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetFinancialCrimes(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

func TestGetFinancialCrimesByYear(t *testing.T) {
	// Set year.
	year := 2018
	// Perform a request with query params.
	financialCrimeData, err := dataGovGrClient.GetFinancialCrimes(
		&FinancialCrimeQueryParams{
			Year: fmt.Sprint(year),
		},
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
	// Check if there is another year.
	for _, financialCrime := range *financialCrimeData {
		if financialCrime.Year != year {
			t.Errorf("Filtered for year %d, found year %d.", year, financialCrime.Year)
		}
	}
}

func TestGetFinancialCrimesByCrime(t *testing.T) {
	// Set crime.
	crime := financialcrime.ALCOHOL_CONTRABAND
	// Perform a request with query params.
	financialCrimeData, err := dataGovGrClient.GetFinancialCrimes(
		&FinancialCrimeQueryParams{
			Crime: crime,
		},
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
	// Check if there is another crime.
	for _, financialCrime := range *financialCrimeData {
		if financialCrime.Crime != crime {
			t.Errorf("Filtered for crime %q, found crime %q.", crime, financialCrime.Crime)
		}
	}
}

func TestGetFoodInspections(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetFoodInspections(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

func TestGetFoodInspectionsByYear(t *testing.T) {
	// Set year.
	year := 2019
	// Perform a request with query params.
	foodInspections, err := dataGovGrClient.GetFoodInspections(NewYearQueryParams(year))
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
	// Ensure only the requested year is present.
	for _, foodIspection := range *foodInspections {
		if foodIspection.Year != year {
			t.Errorf("Filtered for year %d, found year %d.", year, foodIspection.Year)
		}
	}
}

func TestGetForestFires(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetForestFires(
		&ForestFireQueryParams{
			DateFrom: "2018-08-10",
			DateTo:   "2020-08-12",
		},
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

func TestGetForestFiresByPrefecture(t *testing.T) {
	// Set prefecture.
	prefecture := prefectures.ATTIKIS
	// Perform a request with query params.
	forestFires, err := dataGovGrClient.GetForestFires(
		&ForestFireQueryParams{
			DateFrom:   "2018-08-10",
			DateTo:     "2020-08-12",
			Prefecture: prefecture,
		},
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
	// Ensure only the requested prefecture is present.
	for _, forestFires := range *forestFires {
		if forestFires.Prefecture != prefecture {
			t.Errorf("Filtered for prefecture %s, found %s.", prefecture, forestFires.Prefecture)
		}
	}
}

func TestGetHCGIndidents(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetHCGIncidents(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

func TestGetHCGIndidentsByYear(t *testing.T) {
	// Set year.
	year := 2019
	// Perform a request without query params.
	hcgIncidents, err := dataGovGrClient.GetHCGIncidents(
		&HCGIncidentQueryParams{
			Year: fmt.Sprint(year),
		},
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
	// Ensure only the requested year is present.
	for _, hcgIncident := range *hcgIncidents {
		if hcgIncident.Year != year {
			t.Errorf("Filtered for year %d, found year %d.", year, hcgIncident.Year)
		}
	}
}

func TestGetInternetTraffic(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetInternetTraffic(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

func TestGetInternetTrafficByDates(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetInternetTraffic(
		NewDateQueryParams("2021-01-01", "2021-11-30"))
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

func TestGetInternships(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetInternships(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

func TestGetLandLiens(t *testing.T) {
	// Perform a request for given dates.
	_, err := dataGovGrClient.GetLandLiens(
		NewDateQueryParams("2021-12-01", "2021-12-02"))
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

func TestGetLandRegistryStatus(t *testing.T) {
	// Perform a request for given dates.
	_, err := dataGovGrClient.GetLandRegistryStatus(
		&LandRegistryStatusQueryParams{
			DateFrom: "2021-12-01",
			DateTo:   "2021-12-02",
		},
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

func TestGetLandRegistryStatusByStatus(t *testing.T) {
	// Set status.
	status := landregistrystatus.OPERATING
	// Perform a request with params.
	landRegistryStatuses, err := dataGovGrClient.GetLandRegistryStatus(
		&LandRegistryStatusQueryParams{
			DateFrom: "2021-12-01",
			DateTo:   "2021-12-02",
			Status:   status,
		},
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}

	// Ensure that the acquired data include information about the status only.
	for _, landRegistryStatus := range *landRegistryStatuses {
		if landRegistryStatus.Status != status {
			t.Errorf("Filtered for status %q, found status %s.", status, landRegistryStatus.Status)
		}
	}

}

func TestGetLandPlots(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetLandPlots()
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

func TestGetLandTransactions(t *testing.T) {
	// Perform a request with date query params.
	_, err := dataGovGrClient.GetLandTransactions(
		NewDateQueryParams("2021-12-01", "2021-12-01"))
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

func TestGetLawyerNumbers(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetLawyers()
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

func TestGetLawFirms(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetLawFirms()
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

func TestGetPharmacies(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetPharmacies()
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

func TestGetPharmacists(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetPharmacists()
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

func TestGetPowerSystemLoad(t *testing.T) {
	// Perform a request with date query params.
	_, err := dataGovGrClient.GetPowerSystemLoad(
		&DateQueryParams{
			DateFrom: "2021-12-01",
			DateTo:   "2021-12-02",
		},
	)
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

func TestGetResProduction(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetResProduction(
		&DateQueryParams{
			DateFrom: "2021-12-01",
			DateTo:   "2021-12-02",
		},
	)
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

func TestGetRidership(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetRidership(
		&RidershipQueryParams{
			DateFrom: "2021-12-01",
			DateTo:   "2021-12-01",
		},
	)
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

func TestGetRoadTraffic(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetRoadTraffic(
		&RoadTrafficQueryParams{
			DateFrom: "2021-12-01",
			DateTo:   "2021-12-10",
		},
	)
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

func TestGetRealtors(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetRealtors()
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

func TestGetSailingTraffic(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetSailingTraffic(
		&SailingTrafficQueryParams{
			DateFrom: "2021-12-01",
			DateTo:   "2021-12-10",
		},
	)
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

func TestGetStudentSchools(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetStudentSchools(nil)
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

func TestGetTelecomIndicators(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetTelecomIndicators()
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

func TestGetTouristAgencies(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetTouristAgencies()
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

func TestGetTrafficAccidents(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetTrafficAccidents()
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

func TestGetTrafficViolations(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetTrafficViolations(nil)
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

func TestGetUnemploymentData(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetUnemployment(
		NewDateQueryParams("2021-12-09", "2021-12-10"))
	if err != nil {
		t.Errorf("Expected no error, got %v.", err)
	}
}

func TestGetVaccinations(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetVaccinations(
		&VaccinationQueryParams{
			DateFrom: "2020-12-28",
			DateTo:   "2021-12-16",
		},
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

func TestGetVaccinationsByArea(t *testing.T) {
	// Set area.
	area := vaccination.RETHIMNO
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

func TestGetVotersPerAge(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetVotersPerAge()
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}

func TestGetVotersPerMunicipality(t *testing.T) {
	// Perform a request without query params.
	_, err := dataGovGrClient.GetVotersPerMunicipality()
	if err != nil {
		t.Fatalf("Expected no error, got %v.", err)
	}
}
