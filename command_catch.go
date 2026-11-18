package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	fmt.Printf("Throwing a Pokeball at %v...\n", name)
	response, err := cfg.pokeClient.Pokemon(name)
	if err != nil {
		return err
	}

	// calc random
	capture := rand.Intn(response.BaseExperience * 2)

	// compare and reply
	if capture > response.BaseExperience {
		// captured
		cfg.pokedex[response.Name] = response
		fmt.Printf("%v was caught!\n", name)
		return nil
	}

	// escaped
	fmt.Printf("%v escaped!\n", name)

	return nil
}
