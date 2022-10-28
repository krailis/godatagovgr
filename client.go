package godatagovgr

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

// DataGovGrClient struct representation.
type DataGovGrClient struct {
	config *DataGovGrConfig
	client *resty.Client
}

// DataGovGrConfig struct representation.
type DataGovGrConfig struct {
	apiToken         string
	retryCount       int
	retryWaitTime    int
	retryMaxWaitTime int
	timeout          time.Duration
}

const (
	_retryCount       = 3
	_retryWaitTime    = 5
	_retryMaxWaitTime = 15
	_timeout          = time.Second * 20
)

// NewClient creates and returns an new DataGovGr client.
func NewClient(config *DataGovGrConfig) *DataGovGrClient {
	return &DataGovGrClient{
		config: config,
		client: resty.New().
			SetHostURL("https://data.gov.gr/api/v1/query").
			SetHeader("Accept", "application/json").
			SetHeader("Authorization", fmt.Sprintf("Token %s", config.apiToken)).
			SetRetryCount(config.retryCount).
			SetRetryWaitTime(time.Duration(config.retryWaitTime) * time.Second).
			SetRetryMaxWaitTime(time.Duration(config.retryMaxWaitTime) * time.Second).
			SetTimeout(config.timeout).
			SetRetryAfter(nil),
	}
}

// NewDefaultConfig creates and returns a new default DataGovGr config.
// It accepts the API token and uses predefined values for parameters
// such as retry count, retry wait and max wait time, and timeout. The
// default values are a retry count of three (3), a retry wait time and
// retry max wait time of five (5) secs, and a timeout of ten (10) secs.
func NewDefaultConfig(apiToken string) *DataGovGrConfig {
	return &DataGovGrConfig{
		apiToken:         apiToken,
		retryCount:       _retryCount,
		retryWaitTime:    _retryWaitTime,
		retryMaxWaitTime: _retryMaxWaitTime,
		timeout:          _timeout,
	}
}

// GetAPIToken retrieves the API token from the environment.
// It panics if the API is not set properly.
func GetAPIToken() string {
	// Retrieve API token from the environment.
	apiToken, ok := os.LookupEnv("DATAGOVGR_API_TOKEN")
	if !ok || len(apiToken) == 0 {
		log.Panic("The API token for data.gov.gr has not been properly set.")
	}
	return apiToken
}

// GetData retrieves the data for most of the available datasets. It accepts the query
// parameters (if any), a pointer to an array of appropriate structs (according to the
// requested dataset) and the URL. It returns an error.
func (d *DataGovGrClient) GetData(queryParams, data interface{}, url string) error {
	// Handle Query Params.
	queryMap, err := FormatQueryParams(queryParams)
	if err != nil {
		return err
	}
	var errorResponse map[string]interface{}
	resp, err := d.client.R().
		SetQueryParams(queryMap).
		SetResult(data).
		SetError(&errorResponse).
		Get(url)
	if err != nil {
		return fmt.Errorf("Could not retrieve requested data: %s", err)
	} else if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("Request failed with status code [%d]: `%s`",
			resp.StatusCode(), errorResponse["detail"].(string))
	}
	return nil
}

// GetNumbers performs a request to one of the endpoints that retrieve
// data in the form of numbers (entrants, exits etc.) per year. It
// accepts the URL of the dataset.
func (d *DataGovGrClient) GetNumbers(url string) (*[]NumberData, error) {
	var numbers []NumberData
	var errorResponse map[string]interface{}
	resp, err := d.client.R().
		SetResult(&numbers).
		SetError(&errorResponse).
		Get(url)
	if err != nil {
		return nil, fmt.Errorf("Could not retrieve numbers: %s", err)
	} else if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("Request failed with status code [%d]: `%s`",
			resp.StatusCode(), errorResponse["detail"].(string))
	}
	return &numbers, nil
}

// GetAcademicProfessors retrieves the Academic Professor data from data.gov.gr.
//
// Optional query parameters include the year and institution. If query parameters
// are given they should be provided in an AcademicQueryParams struct.
func (d *DataGovGrClient) GetAcademicProfessors(
	queryParams *AcademicQueryParams) (*[]AcademicProfessorData, error) {
	// Define empty auditor data.
	academicProfessorData := &[]AcademicProfessorData{}
	// Get vaccination data.
	err := d.GetData(queryParams, academicProfessorData, "minedu_dep")
	if err != nil {
		return nil, err
	}
	return academicProfessorData, nil
}

// GetAccountants retrieves the accountant numbers from data.gov.gr.
func (d *DataGovGrClient) GetAccountants() (*[]NumberData, error) {
	return d.GetNumbers("oee_accountants")
}

// GetAuditors retrieves the Auditor data from data.gov.gr.
func (d *DataGovGrClient) GetAuditors() (*[]AuditorData, error) {
	// Define empty auditor data.
	auditorData := &[]AuditorData{}
	// Get vaccination data.
	err := d.GetData(nil, auditorData, "elte_auditors")
	if err != nil {
		return nil, err
	}
	return auditorData, nil
}

// GetCadastreNaturaPlots retrieves the Cadastre Natura Plots from data.gov.gr.
//
// Start and end dates are mandatory as query parameters and must be given in
// the defined DateQueryParams struct.
func (d *DataGovGrClient) GetCadastreNaturaPlots(
	queryParams *DateQueryParams) (*[]CadastrePlotData, error) {
	// Define empty cadastre natura plot data.
	cadastreNaturaPlotData := &[]CadastrePlotData{}
	// Get voters per age data.
	err := d.GetData(queryParams, cadastreNaturaPlotData, "cadastre_natura_plot")
	if err != nil {
		return nil, err
	}
	return cadastreNaturaPlotData, nil
}

// GetCadastrePlots retrieves the Cadastre Plots per Municipality from data.gov.gr.
//
// Start and end dates are mandatory as query parameters and must be given in the
// defined DateQueryParams struct.
func (d *DataGovGrClient) GetCadastrePlots(
	queryParams *DateQueryParams) (*[]CadastrePlotData, error) {
	// Define empty cadastre plot data.
	cadastrePlotData := &[]CadastrePlotData{}
	// Get voters per age data.
	err := d.GetData(queryParams, cadastrePlotData, "cadastre_plot")
	if err != nil {
		return nil, err
	}
	return cadastrePlotData, nil
}

// GetCasinoTickets retrieves the Casino Ticket data from data.gov.gr.
//
// The Year is an optional query parameter, which should be given in a
// YearQueryParams struct.
func (d *DataGovGrClient) GetCasinoTickets(
	queryParams *YearQueryParams) (*[]CasinoTicketData, error) {
	// Define empty casino ticket data.
	casinoTicketData := &[]CasinoTicketData{}
	// Get casino ticket data.
	err := d.GetData(queryParams, casinoTicketData, "eeep_casino_tickets")
	if err != nil {
		return nil, err
	}
	return casinoTicketData, nil
}

// GetCrimes retrieves the Crime Statistic data from data.gov.gr.
//
// Optional query parameters include the year and type of crime,
// and should be provided in a CrimeQueryParams struct.
func (d *DataGovGrClient) GetCrimes(
	queryParams *CrimeQueryParams) (*[]CrimeStatData, error) {
	// Define empty crime stat data.
	crimeStatData := &[]CrimeStatData{}
	// Get casino ticket data.
	err := d.GetData(queryParams, crimeStatData, "mcp_crime")
	if err != nil {
		return nil, err
	}
	return crimeStatData, nil
}

// GetDendists retrieves the dendist numbers from data.gov.gr.
func (d *DataGovGrClient) GetDendists() (*[]NumberData, error) {
	return d.GetNumbers("minhealth_dentists")
}

// GetDoctors retrieves the doctor numbers from data.gov.gr.
func (d *DataGovGrClient) GetDoctors() (*[]NumberData, error) {
	return d.GetNumbers("minhealth_doctors")
}

// GetElectricityConsumption retrieves the Electricity Consumption data from data.gov.gr.
//
// The supported query parameters include the start and end dates and area. The query
// parameters must be given in an ElectricityConsumptionQueryParams struct.
func (d *DataGovGrClient) GetElectricityConsumption(
	queryParams *ElectricityConsumptionQueryParams) (*[]ElectricityConsumptionData, error) {
	// Define empty electricity consumption data.
	electricityConsumptionData := &[]ElectricityConsumptionData{}
	// Get electricity consumption data.
	err := d.GetData(queryParams, electricityConsumptionData, "electricity_consumption")
	if err != nil {
		return nil, err
	}
	return electricityConsumptionData, nil
}

// GetEnergyBalance retrieves the Energy Balance data from data.gov.gr.
//
// The supported query parameters include the start and end dates and fuel.
// The query parameters must be given in an EnergyBalanceQueryParams struct.
func (d *DataGovGrClient) GetEnergyBalance(
	queryParams *EnergyBalanceQueryParams) (*[]EnergyBalanceData, error) {
	// Define empty energy balance data.
	energyBalanceData := &[]EnergyBalanceData{}
	// Get energy balance data.
	err := d.GetData(queryParams, energyBalanceData, "admie_dailyenergybalanceanalysis")
	if err != nil {
		return nil, err
	}
	return energyBalanceData, nil
}

// GetEnergyInspectors retrieves the energy inspector numbers from data.gov.gr.
func (d *DataGovGrClient) GetEnergyInspectors() (*[]NumberData, error) {
	return d.GetNumbers("minenv_inspectors")
}

// GetEudoxusApplications retrieves the Eudoxus Application data from data.gov.gr.
//
// Optional query parameters include the year and institution. If query parameters
// are given they should be provided in an AcademicQueryParams struct.
func (d *DataGovGrClient) GetEudoxusApplications(
	queryParams *AcademicQueryParams) (*[]EudoxusApplicationData, error) {
	// Define empty Eudoxus application data.
	eudoxusApplicationData := &[]EudoxusApplicationData{}
	// Get Eudoxus application data.
	err := d.GetData(queryParams, eudoxusApplicationData, "grnet_eudoxus")
	if err != nil {
		return nil, err
	}
	return eudoxusApplicationData, nil
}

// GetFinancialCrimes retrieves the Financial Crime Statistics from data.gov.gr.
//
// The Year is an optional query parameter and should be provided in a
// YearQueryParams struct.
func (d *DataGovGrClient) GetFinancialCrimes(
	queryParams *FinancialCrimeQueryParams) (*[]FinancialCrimeData, error) {
	// Define empty financial crime data.
	financialCrimeData := &[]FinancialCrimeData{}
	// Get Eudoxus application data.
	err := d.GetData(queryParams, financialCrimeData, "mcp_financial_crimes")
	if err != nil {
		return nil, err
	}
	return financialCrimeData, nil
}

// GetFoodInspections retrieves the Food Inspection data from data.gov.gr.
//
// The Year is an optional query parameter, which should be given in a
// YearQueryParams struct.
func (d *DataGovGrClient) GetFoodInspections(
	queryParams *YearQueryParams) (*[]FoodInspectionData, error) {
	// Define empty food inspection data.
	foodInspectionData := &[]FoodInspectionData{}
	// Get Eudoxus application data.
	err := d.GetData(queryParams, foodInspectionData, "efet_inspections")
	if err != nil {
		return nil, err
	}
	return foodInspectionData, nil
}

// GetForestFires retrieves the Forest Fire data from data.gov.gr.
//
// Start and end dates are mandatory as query parameters. An optional
// query parameter is the prefecture. Query parameters must be given
// in a ForestFireQueryParams struct.
func (d *DataGovGrClient) GetForestFires(
	queryParams *ForestFireQueryParams) (*[]ForestFireData, error) {
	// Define empty forest fire data.
	forestFireData := &[]ForestFireData{}
	// Get Forest Fire data.
	err := d.GetData(queryParams, forestFireData, "mcp_forest_fires")
	if err != nil {
		return nil, err
	}
	return forestFireData, nil
}

// GetHCGIncidents retrieves the Hellenic Coast Guard Incident data from data.gov.gr.
//
// The Year and Incident type are optional query parameters, which should be provided
// in a struct of HCGIncidentQueryParams type.
func (d *DataGovGrClient) GetHCGIncidents(
	queryParams *HCGIncidentQueryParams) (*[]HellenicCoastGuardIncidentData, error) {
	// Define empty HCG incident data.
	HCGIncidentData := &[]HellenicCoastGuardIncidentData{}
	// Get HCG incident data.
	err := d.GetData(queryParams, HCGIncidentData, "hcg_incidents")
	if err != nil {
		return nil, err
	}
	return HCGIncidentData, nil
}

// GetInternetTraffic retrieves the Internet Traffic data from data.gov.gr.
//
// Start and end dates are mandatory as query parameters and must be given
// in the defined DateQueryParams struct.
func (d *DataGovGrClient) GetInternetTraffic(
	queryParams *DateQueryParams) (*[]InternetTrafficData, error) {
	// Define empty vaccination data.
	internetTrafficData := &[]InternetTrafficData{}
	// Get vaccination data.
	err := d.GetData(queryParams, internetTrafficData, "internet_traffic")
	if err != nil {
		return nil, err
	}
	return internetTrafficData, nil
}

// GetInternships retrieves the Internship data from data.gov.gr.
//
// Optional query parameters include the year and institution. If
// query parameters are given they should be provided in an
// AcademicQueryParams struct.
func (d *DataGovGrClient) GetInternships(
	queryParams *AcademicQueryParams) (*[]InternshipData, error) {
	// Define empty internship data.
	internshipData := &[]InternshipData{}
	// Get HCG incident data.
	err := d.GetData(queryParams, internshipData, "grnet_atlas")
	if err != nil {
		return nil, err
	}
	return internshipData, nil
}

// GetLandConfiscations retrieves the Confiscation data from data.gov.gr.
//
// Start and end dates are mandatory as query parameters and must be
// given in the defined DateQueryParams struct.
func (d *DataGovGrClient) GetLandConfiscations(
	queryParams *DateQueryParams) (*[]LandConfiscationData, error) {
	// Define empty confiscation data.
	confiscationData := &[]LandConfiscationData{}
	// Get voters per age data.
	err := d.GetData(queryParams, confiscationData, "ktm_confs")
	if err != nil {
		return nil, err
	}
	return confiscationData, nil
}

// GetLandHorizontalArea retrieves the horizontal area data from data.gov.gr.
//
// Start and end dates are mandatory as query parameters and must be given
// in the defined DateQueryParams struct.
func (d *DataGovGrClient) GetLandHorizontalArea(
	queryParams *DateQueryParams) (*[]LandHorizontalAreaData, error) {
	// Define empty horizontal area data.
	horizontalAreaData := &[]LandHorizontalAreaData{}
	// Get horizontal area data.
	err := d.GetData(queryParams, horizontalAreaData, "ktm_harea")
	if err != nil {
		return nil, err
	}
	return horizontalAreaData, nil
}

// GetLandHorizontalPlots retrieves the horizontal plot data from data.gov.gr.
//
// Start and end dates are mandatory as query parameters and must be given in
// the defined DateQueryParams struct.
func (d *DataGovGrClient) GetLandHorizontalPlots(
	queryParams *DateQueryParams) (*[]LandHorizontalPlotData, error) {
	// Define empty horizontal plot data.
	horizontalPlotData := &[]LandHorizontalPlotData{}
	// Get horizontal plot data.
	err := d.GetData(queryParams, horizontalPlotData, "ktm_hplots")
	if err != nil {
		return nil, err
	}
	return horizontalPlotData, nil
}

// GetLandLiens retrieves the Land Liens data from data.gov.gr.
//
// Start and end dates are mandatory as query parameters and must
// be given in the defined DateQueryParams struct.
func (d *DataGovGrClient) GetLandLiens(
	queryParams *DateQueryParams) (*[]LandLienData, error) {
	// Define empty lien data.
	lienData := &[]LandLienData{}
	// Get lien data.
	err := d.GetData(queryParams, lienData, "ktm_liens")
	if err != nil {
		return nil, err
	}
	return lienData, nil
}

// GetLandOwners retrieves the Land Owner data from data.gov.gr.
//
// Start and end dates are mandatory as query parameters and must
// be given in the defined DateQueryParams struct.
func (d *DataGovGrClient) GetLandOwners(
	queryParams *DateQueryParams) (*[]LandOwnerData, error) {
	// Define land owner data.
	landOwnerData := &[]LandOwnerData{}
	// Get land owner data.
	err := d.GetData(queryParams, landOwnerData, "ktm_owners")
	if err != nil {
		return nil, err
	}
	return landOwnerData, nil
}

// GetLandPlots retrieves the Plots per Municipality data from data.gov.gr.
func (d *DataGovGrClient) GetLandPlots() (*[]LandPlotData, error) {
	// Define empty plot data.
	plotData := &[]LandPlotData{}
	// Get voters per age data.
	err := d.GetData(nil, plotData, "ktm_plots")
	if err != nil {
		return nil, err
	}
	return plotData, nil
}

// GetLandRegistryStatus retrieves the Land Registry Status from data.gov.gr.
//
// Start and end dates are mandatory as query parameters and must be given in
// the defined DateQueryParams struct.
func (d *DataGovGrClient) GetLandRegistryStatus(
	queryParams *LandRegistryStatusQueryParams) (*[]LandRegistryStatusData, error) {
	// Define empty cadastre status data.
	cadastreStatusData := &[]LandRegistryStatusData{}
	// Get Cadastre status data.
	err := d.GetData(queryParams, cadastreStatusData, "ktm_status")
	if err != nil {
		return nil, err
	}
	return cadastreStatusData, nil
}

// GetLandTransactions retrieves the Land Transaction data from data.gov.gr.
//
// Start and end dates are mandatory as query parameters and must be given
// in the defined DateQueryParams struct.
func (d *DataGovGrClient) GetLandTransactions(
	queryParams *DateQueryParams) (*[]LandTransactionData, error) {
	// Define empty transaction data.
	transactionData := &[]LandTransactionData{}
	// Get transaction data.
	err := d.GetData(queryParams, transactionData, "ktm_transactions")
	if err != nil {
		return nil, err
	}
	return transactionData, nil
}

// GetLawyers retrieves the lawyer numbers from data.gov.gr.
func (d *DataGovGrClient) GetLawyers() (*[]NumberData, error) {
	return d.GetNumbers("minjust_lawyers")
}

// GetLawFirms retrieves the law firm numbers from data.gov.gr.
func (d *DataGovGrClient) GetLawFirms() (*[]NumberData, error) {
	return d.GetNumbers("minjust_law_firms")
}

// GetPharmacies retrieves the pharmacy numbers from data.gov.gr.
func (d *DataGovGrClient) GetPharmacies() (*[]NumberData, error) {
	return d.GetNumbers("minhealth_pharmacies")
}

// GetPharmacists retrieves the pharmacist numbers from data.gov.gr.
func (d *DataGovGrClient) GetPharmacists() (*[]NumberData, error) {
	return d.GetNumbers("minhealth_pharmacists")
}

// GetPowerSystemLoad retrieves the Power System Load data from data.gov.gr.
//
// Start and end dates are mandatory as query parameters and must be given
// in the defined DateQueryParams struct.
func (d *DataGovGrClient) GetPowerSystemLoad(
	queryParams *DateQueryParams) (*[]PowerSystemLoadData, error) {
	// Define empty power system load data.
	powerSystemLoadData := &[]PowerSystemLoadData{}
	// Get voters per age data.
	err := d.GetData(queryParams, powerSystemLoadData, "admie_realtimescadasystemload")
	if err != nil {
		return nil, err
	}
	return powerSystemLoadData, nil
}

// GetResProduction retrieves the Renewable Energy Sources (RES) Production
// from data.gov.gr.
//
// Start and end dates are mandatory as query parameters and must be given
// in the defined DateQueryParams struct.
func (d *DataGovGrClient) GetResProduction(
	queryParams *DateQueryParams) (*[]ResProductionData, error) {
	// Define empty RES production data.
	resProductionData := &[]ResProductionData{}
	// Get RES production data.
	err := d.GetData(queryParams, resProductionData, "admie_realtimescadares")
	if err != nil {
		return nil, err
	}
	return resProductionData, nil
}

// GetRealtors retrieves the realtor numbers from data.gov.gr.
func (d *DataGovGrClient) GetRealtors() (*[]NumberData, error) {
	return d.GetNumbers("mindev_realtors")
}

// GetRidership retrieves the ridership (OASA) data from data.gov.gr.
//
// The supported query parameters include the start and end dates.
// Query parameters are mandatory and must be given in the defined
// RidershipQueryParams struct.
func (d *DataGovGrClient) GetRidership(
	queryParams *RidershipQueryParams) (*[]RidershipData, error) {
	// Define empty ridership data.
	ridershipData := &[]RidershipData{}
	// Get ridership data.
	err := d.GetData(queryParams, ridershipData, "oasa_ridership")
	if err != nil {
		return nil, err
	}
	return ridershipData, nil
}

// GetRoadTraffic retrieves the road traffic data from data.gov.gr.
//
// The supported query parameters include the start and end dates.
// Query parameters are mandatory and must be given in the defined
// RoadTrafficQueryParams struct.
func (d *DataGovGrClient) GetRoadTraffic(
	queryParams *RoadTrafficQueryParams) (*[]RoadTrafficData, error) {
	// Define empty road traffic data.
	roadTrafficData := &[]RoadTrafficData{}
	// Get road traffic data.
	err := d.GetData(queryParams, roadTrafficData, "road_traffic_attica")
	if err != nil {
		return nil, err
	}
	return roadTrafficData, nil
}

// GetSailingTraffic retrieves the sailing traffic data from data.gov.gr.
//
// The supported query parameters include the start and end dates.
// Query parameters are mandatory and must be given in the defined
// SailingTrafficQueryParams struct.
func (d *DataGovGrClient) GetSailingTraffic(
	queryParams *SailingTrafficQueryParams) (*[]SailingTrafficData, error) {
	// Define empty road traffic data.
	sailingTrafficData := &[]SailingTrafficData{}
	// Get traffic accident data.
	err := d.GetData(queryParams, sailingTrafficData, "sailing_traffic")
	if err != nil {
		return nil, err
	}
	return sailingTrafficData, nil
}

// GetStudentSchools retrieves the Student School data from data.gov.gr.
//
// Supported optional query parameters include the Year, the School Type and
// Class. They should be provided in a StudentSchoolQueryParams struct.
func (d *DataGovGrClient) GetStudentSchools(
	queryParams *StudentSchoolQueryParans) (*[]StudentSchoolData, error) {
	// Define empty traffic accident data.
	studentSchoolData := &[]StudentSchoolData{}
	// Get traffic accident data.
	err := d.GetData(queryParams, studentSchoolData, "minedu_students_school")
	if err != nil {
		return nil, err
	}
	return studentSchoolData, nil
}

// GetTelecomIndicators retrieves the Telecom Indicator data from data.gov.gr.
func (d *DataGovGrClient) GetTelecomIndicators() (*[]TelecomIndicatorData, error) {
	// Define empty telecom indicator data.
	telecomIndicatorData := &[]TelecomIndicatorData{}
	// Get telecom indicator data.
	err := d.GetData(nil, telecomIndicatorData, "eett_telecom_indicators")
	if err != nil {
		return nil, err
	}
	return telecomIndicatorData, nil
}

// GetTouristAgencies retrieves the tourist agency numbers from data.gov.gr.
func (d *DataGovGrClient) GetTouristAgencies() (*[]NumberData, error) {
	return d.GetNumbers("mintour_agencies")
}

// GetTrafficAccidents retrieves the Traffic Accident data from data.gov.gr.
func (d *DataGovGrClient) GetTrafficAccidents() (*[]TrafficAccidentData, error) {
	// Define empty traffic accident data.
	trafficAccidentData := &[]TrafficAccidentData{}
	// Get traffic accident data.
	err := d.GetData(nil, trafficAccidentData, "mcp_traffic_accidents")
	if err != nil {
		return nil, err
	}
	return trafficAccidentData, nil
}

// GetTrafficViolations retrieves the Traffic Violation data from data.gov.gr.
func (d *DataGovGrClient) GetTrafficViolations(
	queryParams *TrafficViolationQueryParams) (*[]TrafficViolationData, error) {
	// Define empty traffic accident data.
	trafficViolationData := &[]TrafficViolationData{}
	// Get traffic accident data.
	err := d.GetData(queryParams, trafficViolationData, "mcp_traffic_violations")
	if err != nil {
		return nil, err
	}
	return trafficViolationData, nil
}

// GetUnemployment retrieves the Unemployment data from data.gov.gr.
//
// Start and end dates are mandatory as query parameters and must be
// given in the defined DateQueryParams struct.
func (d *DataGovGrClient) GetUnemployment(
	queryParams *DateQueryParams) (*[]UnemploymentData, error) {
	// Define empty unemployment data.
	unemploymentData := &[]UnemploymentData{}
	// Get vaccination data.
	err := d.GetData(queryParams, unemploymentData, "oaed_unemployment")
	if err != nil {
		return nil, err
	}
	return unemploymentData, nil
}

// GetUrbanIncidents retrieves the Civil Protection Urban Incident
// data from data.gov.gr.
//
// Start and end dates are mandatory as query parameters and must be
// given in the defined DateQueryParams struct.
func (d *DataGovGrClient) GetUrbanIncidents(
	queryParams *DateQueryParams) (*[]UrbanIncidentData, error) {
	// Define empty CP urban incident data.
	urbanIncidentData := &[]UrbanIncidentData{}
	// Get voters per age data.
	err := d.GetData(queryParams, urbanIncidentData, "mcp_urban_incidents")
	if err != nil {
		return nil, err
	}
	return urbanIncidentData, nil
}

// GetVaccinations retrieves the Vaccination data from data.gov.gr.
//
// The supported query parameters include the start and end dates, as
// well as the area. Shall query parameters be set, the must be given
// in an VaccinationQueryParams struct, as defined in the models.
func (d *DataGovGrClient) GetVaccinations(
	queryParams *VaccinationQueryParams) (*[]VaccinationData, error) {
	// Define empty vaccination data.
	vaccinationData := &[]VaccinationData{}
	// Get vaccination data.
	err := d.GetData(queryParams, vaccinationData, "mdg_emvolio")
	if err != nil {
		return nil, err
	}
	return vaccinationData, nil
}

// GetVotersPerAge retrieves the Voters per Age data from data.gov.gr.
func (d *DataGovGrClient) GetVotersPerAge() (*[]VotersPerAgeData, error) {
	// Define votes per age data.
	votersPerAgeData := &[]VotersPerAgeData{}
	// Get voters per age data.
	err := d.GetData(nil, votersPerAgeData, "minint_election_age")
	if err != nil {
		return nil, err
	}
	return votersPerAgeData, nil
}

// GetVotersPerMunicipality retrieves the Voters per Age data from data.gov.gr.
func (d *DataGovGrClient) GetVotersPerMunicipality() (*[]VotersPerMunicipalityData, error) {
	// Define votes per municipality data.
	votersPerMunicipalityData := &[]VotersPerMunicipalityData{}
	// Get voters per age data.
	err := d.GetData(nil, votersPerMunicipalityData, "minint_election_distribution")
	if err != nil {
		return nil, err
	}
	return votersPerMunicipalityData, nil
}
