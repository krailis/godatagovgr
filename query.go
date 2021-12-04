package godatagovgr

import (
	"encoding/json"
	"errors"
	"fmt"
)

// formatQuery params formats incoming query parameters according to the given
// type and returns a map of query parameters to use with the request.
func formatQueryParams(queryParams interface{}) (map[string]string, error) {
	if queryParams == nil {
		return map[string]string{}, nil
	}
	jsonBts, err := json.Marshal(queryParams)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not format query params: %s", err))
	}
	var queryMap map[string]string
	err = json.Unmarshal(jsonBts, &queryMap)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not parse query params: %s", err))
	}
	return queryMap, nil
}
