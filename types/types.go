package types

import (
	"github.com/vezzsah/pokedexcli/pokecache"
)

type URL string

type Config struct {
	Next       URL
	Previous   URL
	Option     string
	CacheMap   *pokecache.Cache
	Parameters string
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

type LocationsPageResponse struct {
	Count    int             `json:"count"`
	Next     URL             `json:"next"`
	Previous URL             `json:"previous"`
	Results  []NameUrlStruct `json:"results"`
}

type ExplorePageResponse struct {
	EncounterMethodRates []EncounterMethodDetails `json:"encounter_method_rates"`
	GameIndex            int                      `json:"game_index"`
	Id                   int                      `json:"id"`
	Location             NameUrlStruct            `json:"location"`
	Name                 string                   `json:"name"`
	Names                []NameVariation          `json:"names"`
	PokemonEncounters    []PokemonEncounter       `json:"pokemon_encounters"`
}

type NameUrlStruct struct {
	Name string `json:"name"`
	Url  URL    `json:"url"`
}
type EncounterMethodDetails struct {
	EncounterMethod NameUrlStruct `json:"encounter_method"`
	VersionDetails  []RateDetails `json:"version_details"`
}

type RateDetails struct {
	Rate    int           `json:"rate"`
	Version NameUrlStruct `json:"version"`
}

type NameVariation struct {
	Language NameUrlStruct `json:"language"`
	Name     string        `json:"name"`
}

type PokemonEncounter struct {
	Pokemon        NameUrlStruct      `json:"pokemon"`
	VersionDetails []VersionEncounter `json:"version_details"`
}

type VersionEncounter struct {
	EncounterDetails []Encounter   `json:"encounter_details"`
	MaxChance        int           `json:"max_chance"`
	Version          NameUrlStruct `json:"version"`
}

type Encounter struct {
	Chance          int           `json:"chance"`
	ConditionValues []byte        `json:"condition_values"`
	MaxLevel        int           `json:"max_level"`
	Method          NameUrlStruct `json:"method"`
	MinLevel        int           `json:"min_level"`
}

const PokeapiLocationAreaUrl URL = "https://pokeapi.co/api/v2/location-area/"
const PokeapiAreaUrl URL = "https://pokeapi.co/api/v2/location-area/"
