package main

import (
	"time"

	"github.com/kaoskeya/pokedex/internal/pokeapi"
)

func main() {

	cfg := config{
		pokeClient:           pokeapi.NewClient(5*time.Second, 5*time.Minute),
		nextLocationsURL:     "",
		previousLocationsURL: "",
		pokedex:              map[string]pokeapi.PokemonResponse{},
		commands: map[string]cliCommand{
			"help": {
				name:        "help",
				description: "Get Help",
				callback:    commandHelp,
			},
			"map": {
				name:        "map",
				description: "Print next 20 location areas",
				callback:    commandMapf,
			},
			"mapb": {
				name:        "mapb",
				description: "Print previous 20 location areas",
				callback:    commandMapb,
			},
			"explore": {
				name:        "explore <location_name>",
				description: "Explore an area by providing its name",
				callback:    commandExplore,
			},
			"inspect": {
				name:        "inspect <location_name>",
				description: "Inspect a captured pokemon",
				callback:    commandInspect,
			},
			"catch": {
				name:        "catch <location_name>",
				description: "Catch a pokemon by providing its name",
				callback:    commandCatch,
			},
			"exit": {
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    commandExit,
			},
		},
	}

	startRepl(&cfg)
}
