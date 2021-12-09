package godatagovgr

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

// DataGovGr Client struct definition.
type DataGovGrClient struct {
	config *DataGovGrConfig
	client *resty.Client
}

// DataGovGr Config struct definition.
type DataGovGrConfig struct {
	apiToken         string
	retryCount       int
	retryWaitTime    int
	retryMaxWaitTime int
	timeout          time.Duration
}

const (
	_RETRY_COUNT         = 3
	_RETRY_WAIT_TIME     = 5
	_RETRY_MAX_WAIT_TIME = 5
	_TIMEOUT             = time.Second * 10
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
		retryCount:       _RETRY_COUNT,
		retryWaitTime:    _RETRY_WAIT_TIME,
		retryMaxWaitTime: _RETRY_MAX_WAIT_TIME,
		timeout:          _TIMEOUT,
	}
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
		return errors.New(fmt.Sprintf("Could not retrieve requested data: %s", err))
	} else if resp.StatusCode() != http.StatusOK {
		return errors.New(fmt.Sprintf("Request failed with status code [%d]: `%s`",
			resp.StatusCode(), errorResponse["error"].(string)))
	}
	return nil
}

// GetNumbers performs a request to one of the endpoints that retrieve concern
// data in the form of numbers (entrants, exits etc.) per year.
func (d *DataGovGrClient) GetNumbers(url string) (*[]NumberData, error) {
	var numbers []NumberData
	var errorResponse map[string]interface{}
	resp, err := d.client.R().
		SetResult(&numbers).
		SetError(&errorResponse).
		Get(url)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not retrieve numbers: %s", err))
	} else if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(fmt.Sprintf(
			"Request failed with status code [%d].", resp.StatusCode()))
	}
	return &numbers, nil
}

// GetAcademicProfessors retrieves the Academic Professor data from data.gov.gr.
func (d *DataGovGrClient) GetAcademicProfessors() (*[]AcademicProfessorData, error) {
	// Define empty auditor data.
	academicProfessorData := &[]AcademicProfessorData{}
	// Get vaccination data.
	err := d.GetData(nil, academicProfessorData, URL_ACADEMIC_PROFESSORS)
	if err != nil {
		return nil, err
	}
	return academicProfessorData, nil
}

// GetAccountants retrieves the accountant numbers from data.gov.gr.
func (d *DataGovGrClient) GetAccountants() (*[]NumberData, error) {
	return d.GetNumbers(URL_ACCOUNTANTS)
}

// GetAuditors retrieves the Auditor data from data.gov.gr.
func (d *DataGovGrClient) GetAuditors() (*[]AuditorData, error) {
	// Define empty auditor data.
	auditorData := &[]AuditorData{}
	// Get vaccination data.
	err := d.GetData(nil, auditorData, URL_AUDITORS)
	if err != nil {
		return nil, err
	}
	return auditorData, nil
}

// GetCasinoTickets retrieves the Casino Ticket data from data.gov.gr.
func (d *DataGovGrClient) GetCasinoTickets() (*[]CasinoTicketData, error) {
	// Define empty casino ticket data.
	casinoTicketData := &[]CasinoTicketData{}
	// Get casino ticket data.
	err := d.GetData(nil, casinoTicketData, URL_CASINO_TICKETS)
	if err != nil {
		return nil, err
	}
	return casinoTicketData, nil
}

// GetCrimes retrieves the Crime Statistic data from data.gov.gr.
func (d *DataGovGrClient) GetCrimes() (*[]CrimeStatData, error) {
	// Define empty crime stat data.
	crimeStatData := &[]CrimeStatData{}
	// Get casino ticket data.
	err := d.GetData(nil, crimeStatData, URL_CRIMES)
	if err != nil {
		return nil, err
	}
	return crimeStatData, nil
}

// GetDendists retrieves the dendist numbers from data.gov.gr.
func (d *DataGovGrClient) GetDendists() (*[]NumberData, error) {
	return d.GetNumbers(URL_DENDISTS)
}

// GetDoctors retrieves the doctor numbers from data.gov.gr.
func (d *DataGovGrClient) GetDoctors() (*[]NumberData, error) {
	return d.GetNumbers(URL_DOCTORS)
}

// GetEnergyInspectors retrieves the energy inspector numbers from data.gov.gr.
func (d *DataGovGrClient) GetEnergyInspectors() (*[]NumberData, error) {
	return d.GetNumbers(URL_ENERGY_INSPECTORS)
}

// GetEudoxusApplications retrieves the Eudoxus Application data from data.gov.gr.
func (d *DataGovGrClient) GetEudoxusApplications() (*[]EudoxusApplicationData, error) {
	// Define empty Eudoxus application data.
	eudoxusApplicationData := &[]EudoxusApplicationData{}
	// Get Eudoxus application data.
	err := d.GetData(nil, eudoxusApplicationData, URL_EUDOXUS_APPLICATIONS)
	if err != nil {
		return nil, err
	}
	return eudoxusApplicationData, nil
}

// GetFinancialCrimes retrieves the Financial Crime Statistics from data.gov.gr.
func (d *DataGovGrClient) GetFinancialCrimes() (*[]FinancialCrimeData, error) {
	// Define empty financial crime data.
	financialCrimeData := &[]FinancialCrimeData{}
	// Get Eudoxus application data.
	err := d.GetData(nil, financialCrimeData, URL_FINANCIAL_CRIMES)
	if err != nil {
		return nil, err
	}
	return financialCrimeData, nil
}

// GetInternetTraffic retrieves the Internet Traffic data from data.gov.gr.
// The supported query parameters include the start and end dates. The query
// parameters must be given in an InternetTrafficQueryParams struct, as
// defined in the client's models.
func (d *DataGovGrClient) GetInternetTraffic(
	queryParams *InternetTrafficQueryParams) (*[]InternetTrafficData, error) {
	// Define empty vaccination data.
	internetTrafficData := &[]InternetTrafficData{}
	// Get vaccination data.
	err := d.GetData(queryParams, internetTrafficData, URL_INTERNET_TRAFFIC)
	if err != nil {
		return nil, err
	}
	return internetTrafficData, nil
}

// GetInternships retrieves the Internship data from data.gov.gr.
func (d *DataGovGrClient) GetInternships() (*[]InternshipData, error) {
	// Define empty internship data.
	internshipData := &[]InternshipData{}
	// Get HCG incident data.
	err := d.GetData(nil, internshipData, URL_INTERNSHIPS)
	if err != nil {
		return nil, err
	}
	return internshipData, nil
}

// GetHCGIncidents retrieves the Hellenic Coast Guard Incident data from data.gov.gr.
func (d *DataGovGrClient) GetHCGIncidents() (*[]HellenicCoastGuardIncidentData, error) {
	// Define empty HCG incident data.
	HCGIncidentData := &[]HellenicCoastGuardIncidentData{}
	// Get HCG incident data.
	err := d.GetData(nil, HCGIncidentData, URL_HCG_INCIDENTS)
	if err != nil {
		return nil, err
	}
	return HCGIncidentData, nil
}

// GetLawyers retrieves the lawyer numbers from data.gov.gr.
func (d *DataGovGrClient) GetLawyers() (*[]NumberData, error) {
	return d.GetNumbers(URL_LAWYERS)
}

// GetLawFirms retrieves the law firm numbers from data.gov.gr.
func (d *DataGovGrClient) GetLawFirms() (*[]NumberData, error) {
	return d.GetNumbers(URL_LAW_FIRMS)
}

// GetPharmacies retrieves the pharmacy numbers from data.gov.gr.
func (d *DataGovGrClient) GetPharmacies() (*[]NumberData, error) {
	return d.GetNumbers(URL_PHARMACIES)
}

// GetPharmacists retrieves the pharmacist numbers from data.gov.gr.
func (d *DataGovGrClient) GetPharmacists() (*[]NumberData, error) {
	return d.GetNumbers(URL_PHARMACISTS)
}

// GetRealtors retrieves the realtor numbers from data.gov.gr.
func (d *DataGovGrClient) GetRealtors() (*[]NumberData, error) {
	return d.GetNumbers(URL_REALTORS)
}

// GetStudentSchools retrieves the Student School data from data.gov.gr.
func (d *DataGovGrClient) GetStudentSchools() (*[]StudentSchoolData, error) {
	// Define empty traffic accident data.
	studentSchoolData := &[]StudentSchoolData{}
	// Get traffic accident data.
	err := d.GetData(nil, studentSchoolData, URL_STUDENT_SCHOOLS)
	if err != nil {
		return nil, err
	}
	return studentSchoolData, nil
}

// GetTouristAgencies retrieves the tourist agency numbers from data.gov.gr.
func (d *DataGovGrClient) GetTouristAgencies() (*[]NumberData, error) {
	return d.GetNumbers(URL_TOURIST_AGENCIES)
}

// GetTrafficAccidents retrieves the Traffic Accident data from data.gov.gr.
func (d *DataGovGrClient) GetTrafficAccidents() (*[]TrafficAccidentData, error) {
	// Define empty traffic accident data.
	trafficAccidentData := &[]TrafficAccidentData{}
	// Get traffic accident data.
	err := d.GetData(nil, trafficAccidentData, URL_TRAFFIC_ACCIDENTS)
	if err != nil {
		return nil, err
	}
	return trafficAccidentData, nil
}

// GetTrafficViolations retrieves the Traffic Violation data from data.gov.gr.
func (d *DataGovGrClient) GetTrafficViolations() (*[]TrafficViolationData, error) {
	// Define empty traffic accident data.
	trafficViolationData := &[]TrafficViolationData{}
	// Get traffic accident data.
	err := d.GetData(nil, trafficViolationData, URL_TRAFFIC_VIOLATIONS)
	if err != nil {
		return nil, err
	}
	return trafficViolationData, nil
}

// GetVaccinations retrieves the Vaccination data from data.gov.gr.
// The supported query parameters include the start and end dates, as
// well as the area. Shall query parameters be set, the must be given
// in an VaccinationQueryParams struct, as defined in the models.
func (d *DataGovGrClient) GetVaccinations(
	queryParams *VaccinationQueryParams) (*[]VaccinationData, error) {
	// Define empty vaccination data.
	vaccinationData := &[]VaccinationData{}
	// Get vaccination data.
	err := d.GetData(queryParams, vaccinationData, URL_COVID_VACCINATION)
	if err != nil {
		return nil, err
	}
	return vaccinationData, nil
}
