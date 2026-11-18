package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

func (c *Client) Pokemon(name string) (PokemonResponse, error) {
	url := baseURL + "/pokemon/" + name

	if val, ok := c.cache.Get(url); ok {
		pokemonResponse := PokemonResponse{}
		err := json.Unmarshal(val, &pokemonResponse)
		if err == nil {
			// cache hit
			return pokemonResponse, nil
		}
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return PokemonResponse{}, errors.New("Could not find that pokemon")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	pokemonResponse := PokemonResponse{}
	if err = json.Unmarshal(data, &pokemonResponse); err != nil {
		fmt.Println("err")
		fmt.Println(err)
		return PokemonResponse{}, err
	}

	c.cache.Add(url, data)

	return pokemonResponse, nil
}
