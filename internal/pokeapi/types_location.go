package pokeapi

type LocationAreaResponse struct {
	Id                int
	Name              string
	PokemonEncounters []LocationAreaPokemonEncounter `json:"pokemon_encounters"`
}

type LocationAreaPokemonEncounter struct {
	Pokemon NameUrl
}
