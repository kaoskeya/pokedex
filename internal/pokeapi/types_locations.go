package pokeapi

type LocationAreasResponse struct {
	Count    int
	Next     string
	Previous string
	Results  []NameUrl
}

type NameUrl struct {
	Name string
	Url  string
}
