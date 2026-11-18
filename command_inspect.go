package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]

	if captured, ok := cfg.pokedex[name]; ok {
		fmt.Printf("Name: %v\n", captured.Name)
		fmt.Printf("Height: %v\n", captured.Height)
		fmt.Printf("Weight: %v\n", captured.Weight)
		fmt.Printf("Stats: \n")
		for _, s := range captured.Stats {
			fmt.Printf(" -%v: %v\n", s.Stat.Name, s.BaseStat)
		}
		fmt.Printf("Types: \n")
		for _, s := range captured.Types {
			fmt.Printf(" - %v\n", s.Type.Name)
		}
		return nil
	}
	fmt.Println("you have not caught that pokemon")

	return nil
}
