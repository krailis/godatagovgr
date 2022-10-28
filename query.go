package godatagovgr

import (
	"encoding/json"
	"fmt"
)

// AcademicQueryParams struct definition.
type AcademicQueryParams struct {
	Year        string `json:"year,omitempty"`
	Institution string `json:"institution,omitempty"`
}

// CrimeQueryParams struct definition.
type CrimeQueryParams struct {
	Year  string `json:"year,omitempty"`
	Crime string `json:"crime,omitempty"`
}

// DateQueryParams struct definition.
type DateQueryParams struct {
	DateFrom string `json:"date_from,omitempty"`
	DateTo   string `json:"date_to,omitempty"`
}

// ElectricityConsumptionQueryParams struct definition.
type ElectricityConsumptionQueryParams struct {
	DateFrom string `json:"date_from,omitempty"`
	DateTo   string `json:"date_to,omitempty"`
	Area     string `json:"area,omitempty"`
}

// EnergyBalanceQueryParams struct definition.
type EnergyBalanceQueryParams struct {
	DateFrom string `json:"date_from,omitempty"`
	DateTo   string `json:"date_to,omitempty"`
	Fuel     string `json:"fuel,omitempty"`
}

// FinancialCrimeQueryParams struct representation.
type FinancialCrimeQueryParams struct {
	Year  string `json:"year,omitempty"`
	Crime string `json:"crime,omitempty"`
}

// ForestFireQueryParams struct representation.
type ForestFireQueryParams struct {
	DateFrom   string `json:"date_from,omitempty"`
	DateTo     string `json:"date_to,omitempty"`
	Prefecture string `json:"prefecture,omitempty"`
}

// HCGIncidentQueryParams struct representation.
type HCGIncidentQueryParams struct {
	Year     string `json:"year,omitempty"`
	Incident string `json:"incident,omitempty"`
}

// LandRegistryStatusQueryParams struct representation.
type LandRegistryStatusQueryParams struct {
	DateFrom string `json:"date_from,omitempty"`
	DateTo   string `json:"date_to,omitempty"`
	Status   string `json:"status,omitempty"`
}

// RidershipQueryParams struct representation.
type RidershipQueryParams struct {
	DateFrom string `json:"date_from,omitempty"`
	DateTo   string `json:"date_to,omitempty"`
}

// RoadTrafficQueryParams struct representation.
type RoadTrafficQueryParams struct {
	DateFrom string `json:"date_from,omitempty"`
	DateTo   string `json:"date_to,omitempty"`
}

// StudentSchoolQueryParans struct representation.
type StudentSchoolQueryParans struct {
	Year        int    `json:"year,omitempty"`
	SchoolType  string `json:"school_type,omitempty"`
	SchoolClass string `json:"school_class,omitempty"`
}

// SailingTrafficQueryParams struct representation.
type SailingTrafficQueryParams struct {
	DateFrom string `json:"date_from,omitempty"`
	DateTo   string `json:"date_to,omitempty"`
}

// TrafficViolationQueryParams struct representation.
type TrafficViolationQueryParams struct {
	Year      string `json:"year,omitempty"`
	Violation string `json:"violation,omitempty"`
}

// VaccinationQueryParams struct representation.
type VaccinationQueryParams struct {
	DateFrom string `json:"date_from,omitempty"`
	DateTo   string `json:"date_to,omitempty"`
	Area     string `json:"area,omitempty"`
}

// YearQueryParams struct definition.
type YearQueryParams struct {
	Year string `json:"year,omitempty"`
}

// FormatQueryParams formats incoming query parameters according and returns
// a map of query parameters to use with the request.
func FormatQueryParams(queryParams interface{}) (map[string]string, error) {
	if queryParams == nil {
		return map[string]string{}, nil
	}
	jsonBts, err := json.Marshal(queryParams)
	if err != nil {
		return nil, fmt.Errorf("Could not format query params: %s", err)
	}
	var queryMap map[string]string
	err = json.Unmarshal(jsonBts, &queryMap)
	if err != nil {
		return nil, fmt.Errorf("Could not parse query params: %s", err)
	}
	return queryMap, nil
}

// NewDateQueryParams creates and returns a new DateQueryParams object.
func NewDateQueryParams(dateFrom, dateTo string) *DateQueryParams {
	return &DateQueryParams{
		DateFrom: dateFrom,
		DateTo:   dateTo,
	}
}

// NewYearQueryParams creates and returns a new YearQueryParams object.
func NewYearQueryParams(year int) *YearQueryParams {
	return &YearQueryParams{
		Year: fmt.Sprint(year),
	}
}
