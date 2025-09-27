package pokeapi

import (
	"encoding/json"
	"fmt"

	"github.com/simonjwhitlock/bootdev_pokedex/internal/pokecache"
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

type (
	ExporeResponse struct {
		Name               string
		Pokemon_Encounters []PokemonEncounter
	}
	PokemonEncounter struct {
		Pokemon foundPokemon
	}
	foundPokemon struct {
		Name string
		url  string
	}
)

func MapCall(index, count int, cache *pokecache.Cache) (MapResponse, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=%v", index, count)

	resp, err := get(url, cache)
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

func ExploreLocation(location string, cache *pokecache.Cache) (ExporeResponse, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v", location)

	resp, err := get(url, cache)
	if err != nil {
		return ExporeResponse{}, err
	}
	var result ExporeResponse
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return ExporeResponse{}, fmt.Errorf("error unmarshaling json: %v", err)
	}

	return result, nil
}
