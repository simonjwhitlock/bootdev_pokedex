package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

func call(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error with pokeapi call: %v", err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("pokeapi call returned status: %v", res.StatusCode)
	}
	if err != nil {
		return nil, fmt.Errorf("error reading Body of response: %v", err)
	}
	return body, nil
}
