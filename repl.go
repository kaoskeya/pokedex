package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kaoskeya/pokedex/internal/pokeapi"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type config struct {
	commands             map[string]cliCommand
	pokeClient           pokeapi.Client
	nextLocationsURL     string
	previousLocationsURL string
	pokedex              map[string]pokeapi.PokemonResponse
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		split := strings.Fields(strings.ToLower(text))
		if len(split) == 0 {
			continue
		}

		c := split[0]
		args := []string{}
		if len(split) > 1 {
			args = split[1:]
		}

		val, ok := cfg.commands[c]
		if ok == false {
			cfg.commands["help"].callback(cfg)
			continue
		}

		callErr := val.callback(cfg, args...)
		if callErr != nil {
			fmt.Println(callErr)
		}
	}
}
