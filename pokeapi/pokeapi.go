package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

func Call(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("Error with PokeAPI call:", err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return "", fmt.Errorf("PokeAPI call returned status:", res.StatusCode)
	}
	if err != nil {
		return "", fmt.Errorf("Error reading Body of response:", err)
	}
	return fmt.Sprintf("%s", body), nil
}
