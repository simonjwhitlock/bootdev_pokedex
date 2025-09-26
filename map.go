package main

import (
	"fmt"

	"github.com/simonjwhitlock/bootdev_pokedex/internal/pokeapi"
)

func commandMap(configuration *config) error {
	newIndex := configuration.mapCurrentIndex + configuration.mapCallCount
	resp, err := pokeapi.MapCall(newIndex, configuration.mapCallCount)
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
	resp, err := pokeapi.MapCall(newIndex, configuration.mapCallCount)
	if err != nil {
		return err
	}
	configuration.mapCurrentIndex = newIndex
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	return nil
}
