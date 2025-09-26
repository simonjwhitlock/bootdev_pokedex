package pokeapi

import (
	"encoding/json"
	"fmt"
)

type (
	MapResponse struct {
		Count    int
		Next     string
		Previous string
		Results  []Location
	}

	Location struct {
		Name string
		Url  string
	}
)

func MapCall(index, count int) (MapResponse, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=%v", index, count)

	resp, err := call(url)
	if err != nil {
		return MapResponse{}, err
	}
	var result MapResponse
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return MapResponse{}, fmt.Errorf("error unmarshaling json: %v", err)
	}
	return result, nil
}
