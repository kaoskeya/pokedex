package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for _, c := range cfg.commands {
		fmt.Printf("%v: %v\n", c.name, c.description)
	}
	return nil
}
