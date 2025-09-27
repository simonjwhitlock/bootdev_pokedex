package pokeapi

import (
	"encoding/json"
	"fmt"

	"github.com/simonjwhitlock/bootdev_pokedex/internal/pokecache"
)

type (
	Pokemon struct {
		Name    string         `json:"name"`
		BaseExp int            `json:"base_experience"`
		Height  int            `json:"Height"`
		Weight  int            `json:"Weight"`
		Stats   []pokemonStats `json:"stats"`
		Types   []pokemonTypes `json:"types"`
	}
	pokemonStats struct {
		StatBase int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	}
	pokemonTypes struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	}
)

func PokemonCall(pokemon string, cache *pokecache.Cache) (Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", pokemon)

	resp, err := get(url, cache)
	if err != nil {
		return Pokemon{}, err
	}
	var result Pokemon
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error unmarshaling json: %v", err)
	}
	return result, nil
}
