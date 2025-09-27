package main

import (
	"fmt"
	"math/rand"

	"github.com/simonjwhitlock/bootdev_pokedex/internal/pokeapi"
)

func commandCatch(configuration *config) error {
	fmt.Printf("Throwing a Pokeball at %v...\n", configuration.input[1])
	catchRoll := rand.Intn(configuration.catchRollMax)
	fmt.Printf("rolled a %v\n", catchRoll)
	pokemon, err := pokeapi.PokemonCall(configuration.input[1], configuration.cache)
	if err != nil {
		return err
	}
	fmt.Printf("%v requires %v to catch\n", configuration.input[1], pokemon.BaseExp)
	if catchRoll <= pokemon.BaseExp {
		fmt.Printf("You didnt catch %v :(\n", pokemon.Name)
		return nil
	}
	fmt.Printf("You caught %v!\n", pokemon.Name)
	pokedex[pokemon.Name] = pokemon
	return nil
}

func commandInspect(configuration *config) error {
	pokemon, ok := pokedex[configuration.input[1]]
	if !ok {
		fmt.Printf("%v is not in your pokedex!\n", configuration.input[1])
		return nil
	}
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, pokemonStat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", pokemonStat.Stat.Name, pokemonStat.StatBase)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("  -%v\n", pokemonType.Type.Name)
	}

	return nil
}

func commandPokedex(configuration *config) error {
	if len(pokedex) < 1 {
		fmt.Println("Your pokedex is empty.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range pokedex {
		fmt.Printf(" - %v\n", pokemon.Name)
	}

	return nil
}
