package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {
	response, err := cfg.pokeClient.LocationAreas(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = response.Next
	cfg.previousLocationsURL = response.Previous

	for _, p := range response.Results {
		fmt.Println(p.Name)
	}

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousLocationsURL == "" {
		return errors.New("you're on the first page")
	}

	response, err := cfg.pokeClient.LocationAreas(cfg.previousLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = response.Next
	cfg.previousLocationsURL = response.Previous

	for _, p := range response.Results {
		fmt.Println(p.Name)
	}

	return nil
}
