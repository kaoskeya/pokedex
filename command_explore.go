package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	response, err := cfg.pokeClient.LocationArea(name)
	if err != nil {
		return err
	}

	for _, p := range response.PokemonEncounters {
		fmt.Println(p.Pokemon.Name)
	}
	return nil
}
