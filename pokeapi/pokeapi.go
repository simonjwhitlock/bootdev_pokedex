package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func call(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error with PokeAPI call:", err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("PokeAPI call returned status:", res.StatusCode)
	}
	if err != nil {
		return nil, fmt.Errorf("Error reading Body of response:", err)
	}
	return body, nil
}

func MapCall(url string) (MapResponse, error) {
	resp, err := call(url)
	if err != nil {
		return MapResponse{}, err
	}
	var result MapResponse
	err = json.Unmarshal(resp, &result)
	return result, nil
}
