package pokeapi

type PokemonResponse struct {
	Id             int
	Name           string
	BaseExperience int `json:"base_experience"`
	Height         int
	Weight         int
	Stats          []pokemonStat
	Types          []pokemonType
}

type pokemonStat struct {
	BaseStat int `json:"base_stat"`
	Stat     NameUrl
}

type pokemonType struct {
	Type NameUrl
}
