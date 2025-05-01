package types

type URL string

type Config struct {
	Next     URL
	Previous URL
	Option   string
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

type LocationsPageResponse struct {
	Count    int               `json:"count"`
	Next     URL               `json:"next"`
	Previous URL               `json:"previous"`
	Results  []LocationNameUrl `json:"results"`
}

type LocationNameUrl struct {
	Name string `json:"name"`
	Url  URL    `json:"url"`
}

const PokeapiLocationUrl URL = "https://pokeapi.co/api/v2/location/"
