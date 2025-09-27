package main

import (
	"fmt"

	"github.com/simonjwhitlock/bootdev_pokedex/internal/pokeapi"
)

func commandMap(configuration *config) error {
	newIndex := configuration.mapCurrentIndex + configuration.mapCallCount
	resp, err := pokeapi.MapCall(newIndex, configuration.mapCallCount, configuration.cache)
	if err != nil {
		return err
	}
	configuration.mapCurrentIndex = newIndex
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapBack(configuration *config) error {
	newIndex := configuration.mapCurrentIndex - configuration.mapCallCount
	resp, err := pokeapi.MapCall(newIndex, configuration.mapCallCount, configuration.cache)
	if err != nil {
		return err
	}
	configuration.mapCurrentIndex = newIndex
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandExplore(configuration *config) error {
	fmt.Printf("Exploring %v...\n", configuration.input[1])
	fmt.Println("Found Pokemon:")
	resp, err := pokeapi.ExploreLocation(configuration.input[1], configuration.cache)
	if err != nil {
		return err
	}
	for _, pokemon := range resp.Pokemon_Encounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
