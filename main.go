package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/simonjwhitlock/bootdev_pokedex/internal/pokeapi"
	"github.com/simonjwhitlock/bootdev_pokedex/internal/pokecache"
)

// define the type structure of the command map items
type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

// define sturcture of congfig
type config struct {
	mapCurrentIndex int
	mapCallCount    int
	cache           *pokecache.Cache
	input           []string
	catchRollMax    int
}

// declare and set global configuration
var configuration = config{
	mapCurrentIndex: 0,
	mapCallCount:    20,
	cache:           pokecache.NewCache(5 * time.Minute),
	catchRollMax:    300,
}

// initalise pokedex map
var pokedex = make(map[string]pokeapi.Pokemon)

// declare the command map (main func calls the getCommands func to fill in the global varable)
var commands map[string]cliCommand

func getCommands() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 map locations from pokeapi.co - each call will display the next 20 from the previous call",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the preivous 20 map locations from pokeapi.co - each call will display the previous 20 from the previous call",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "explores the aria and returns a list of found pokemon - takes an area from the map as the input",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "thow pokeball at pokemon to attempt to catch them - takes an pokemon name as the input",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect the stats of a caught pokemon, use pokemon name as argument",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "lists all caught pokemon",
			callback:    commandPokedex,
		},
	}
}

func main() {
	//fill in the command map
	getCommands()

	// create new io scanner for comand line
	scanner := bufio.NewScanner(os.Stdin)
	//print initail cli prompt
	fmt.Print("Pokedex >")
	//capture cli input and parse for commands, execute if first word in input is valid command
	for scanner.Scan() {
		configuration.input = cleanInput(scanner.Text())
		function, ok := commands[configuration.input[0]]
		if ok {
			err := function.callback(&configuration)
			if err != nil {
				fmt.Println("last call returned error:")
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
		//print next cli prompt once executed
		fmt.Print("Pokedex >")
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}

func commandExit(configuration *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(configuration *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Useage:")
	fmt.Println()
	for _, value := range commands {
		fmt.Printf("%v: %v\n", value.name, value.description)
	}
	return nil
}
