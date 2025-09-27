package pokeapi

import (
	"fmt"
	"io"
	"net/http"

	"github.com/simonjwhitlock/bootdev_pokedex/internal/pokecache"
)

func call(url string, cache *pokecache.Cache) ([]byte, error) {
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
	cache.Add(url, body)
	return body, nil
}

func get(url string, cache *pokecache.Cache) ([]byte, error) {
	body, ok := cache.Get(url)

	if !ok {
		body, err := call(url, cache)
		if err != nil {
			return nil, err
		}
		return body, nil
	}

	return body, nil
}
