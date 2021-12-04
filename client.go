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
// such as retry count retry wait and max wait time and timeout.
func NewDefaultConfig(apiToken string) *DataGovGrConfig {
	return &DataGovGrConfig{
		apiToken:         apiToken,
		retryCount:       _RETRY_COUNT,
		retryWaitTime:    _RETRY_WAIT_TIME,
		retryMaxWaitTime: _RETRY_MAX_WAIT_TIME,
		timeout:          _TIMEOUT,
	}
}

// GetVaccinationData retrieves the Vaccination data from data.gov.gr.
// The supported query parameters include the start and end dates, as
// well as the area. Shall query parameters be set, the must be given
// in an VaccinationQueryParams struct, as defined in the models.
func (d *DataGovGrClient) GetVaccinationData(
	queryParams *VaccinationQueryParams) (*[]VaccinationData, error) {
	// Handle Query Params.
	vaccinationQueryMap, err := formatQueryParams(queryParams)
	if err != nil {
		return nil, err
	}
	var vaccinationData *[]VaccinationData
	var errorResponse map[string]interface{}
	resp, err := d.client.R().
		SetQueryParams(vaccinationQueryMap).
		SetResult(&vaccinationData).
		SetError(&errorResponse).
		Get(URL_COVID_VACCINATION)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not retrieve vaccination data: %s", err))
	} else if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Request failed with status code [%d]: `%s`",
			resp.StatusCode(), errorResponse["error"].(string)))
	}
	return vaccinationData, nil
}

// GetInternetTraffic retrieves the internet traffic from data.gov.gr.
// The supported query parameters include the start and end dates. The
// query parameters must be given in an InternetTrafficQueryParams struct,
// as defined in the client's models.
func (d *DataGovGrClient) GetInternetTraffic(
	queryParams *InternetTrafficQueryParams) (*[]InternetTraffic, error) {
	// Handle Query Params.
	internetTrafficQueryMap, err := formatQueryParams(queryParams)
	if err != nil {
		return nil, err
	}
	var internetTraffic *[]InternetTraffic
	var errorResponse map[string]interface{}
	resp, err := d.client.R().
		SetQueryParams(internetTrafficQueryMap).
		SetResult(&internetTraffic).
		SetError(&errorResponse).
		Get(URL_INTERNET_TRAFFIC)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not retrieve internet traffic data: %s", err))
	} else if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Request failed with status code [%d]: `%s`",
			resp.StatusCode(), errorResponse["error"].(string)))
	}
	return internetTraffic, nil
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

// GetAccountantNumbers retrieves the accountant numbers from data.gov.gr.
func (d *DataGovGrClient) GetAccountantNumbers() (*[]NumberData, error) {
	return d.GetNumbers(URL_ACCOUNTANT_NUMBERS)
}

// GetEnergyInspectorNumbers retrieves the energy inspector numbers from data.gov.gr.
func (d *DataGovGrClient) GetEnergyInspectorNumbers() (*[]NumberData, error) {
	return d.GetNumbers(URL_ENERGY_INSPECTOR_NUMBERS)
}

// GetLawyerNumbers retrieves the lawyer numbers from data.gov.gr.
func (d *DataGovGrClient) GetLawyerNumbers() (*[]NumberData, error) {
	return d.GetNumbers(URL_LAWYER_NUMBERS)
}

// GetLawFirmNumbers retrieves the law firm numbers from data.gov.gr.
func (d *DataGovGrClient) GetLawFirmNumbers() (*[]NumberData, error) {
	return d.GetNumbers(URL_LAW_FIRM_NUMBERS)
}

// GetPharmacyNumbers retrieves the pharmacy numbers from data.gov.gr.
func (d *DataGovGrClient) GetPharmacyNumbers() (*[]NumberData, error) {
	return d.GetNumbers(URL_PHARMACY_NUMBERS)
}

// GetPharmacistNumbers retrieves the pharmacist numbers from data.gov.gr.
func (d *DataGovGrClient) GetPharmacistNumbers() (*[]NumberData, error) {
	return d.GetNumbers(URL_PHARMACIST_NUMBERS)
}

// GetRealtorNumbers retrieves the realtor numbers from data.gov.gr.
func (d *DataGovGrClient) GetRealtorNumbers() (*[]NumberData, error) {
	return d.GetNumbers(URL_REALTOR_NUMBERS)
}

// GetTouristAgencyNumbers retrieves the tourist agency numbers from data.gov.gr.
func (d *DataGovGrClient) GetTouristAgencyNumbers() (*[]NumberData, error) {
	return d.GetNumbers(URL_TOURIST_AGENCY_NUMBERS)
}
