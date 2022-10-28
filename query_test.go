package godatagovgr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatQueryParamsVaccination(t *testing.T) {
	// New assert.
	assert := assert.New(t)

	// Test for empty vaccination query params.
	queryMap, err := FormatQueryParams(VaccinationQueryParams{})
	assert.NoError(err, "Unexpected error.")
	assert.Equal(queryMap, map[string]string{},
		fmt.Sprintf("Expected an empty map, got %v.", queryMap))

	// Test for vaccination query params.
	queryMap, err = FormatQueryParams(VaccinationQueryParams{
		DateFrom: "2021-01-01",
		DateTo:   "2021-09-01",
		Area:     "ΘΕΣΣΑΛΟΝΙΚΗΣ",
	})
	assert.NoError(err, "Unexpected error.")
	assert.Equal(queryMap,
		map[string]string{
			"date_from": "2021-01-01",
			"date_to":   "2021-09-01",
			"area":      "ΘΕΣΣΑΛΟΝΙΚΗΣ",
		},
		fmt.Sprintf("The expected map differs from the actual one: %s", queryMap),
	)

	// Test for half-filled vaccination query params.
	queryMap, err = FormatQueryParams(VaccinationQueryParams{
		Area: "ΘΕΣΣΑΛΟΝΙΚΗΣ",
	})
	assert.NoError(err, "Unexpected error.")
	assert.Equal(queryMap,
		map[string]string{
			"area": "ΘΕΣΣΑΛΟΝΙΚΗΣ",
		},
		fmt.Sprintf("The expected map differs from the actual one: %s", queryMap),
	)
}

func TestFormatQueryParamsDate(t *testing.T) {
	// New assert.
	assert := assert.New(t)

	// Test for empty date query params.
	queryMap, err := FormatQueryParams(DateQueryParams{})
	assert.NoError(err, "Unexpected error.")
	assert.Equal(queryMap, map[string]string{},
		fmt.Sprintf("Expected an empty map, got %v.", queryMap))

	// Test for date query params.
	queryMap, err = FormatQueryParams(DateQueryParams{
		DateFrom: "2021-01-01",
		DateTo:   "2021-09-01",
	})
	assert.NoError(err, "Unexpected error.")
	assert.Equal(queryMap,
		map[string]string{
			"date_from": "2021-01-01",
			"date_to":   "2021-09-01",
		},
		fmt.Sprintf("The expected map differs from the actual one: %s", queryMap),
	)

	// Test for half-filled internet traffic query params.
	queryMap, err = FormatQueryParams(VaccinationQueryParams{
		DateFrom: "2021-01-01",
	})
	assert.NoError(err, "Unexpected error.")
	assert.Equal(queryMap,
		map[string]string{
			"date_from": "2021-01-01",
		},
		fmt.Sprintf("The expected map differs from the actual one: %s", queryMap),
	)
}
