package types

import (
	"github.com/vezzsah/pokedexcli/pokecache"
)

type URL string

type Config struct {
	Next          URL
	Previous      URL
	Option        string
	CacheMap      *pokecache.Cache
	Parameters    string
	PokemonCaught map[string]PokemonPageResponse
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

type NameUrlStruct struct {
	Name string `json:"name"`
	Url  URL    `json:"url"`
}

type LatestLegacyStruct struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
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

type PokemonPageResponse struct {
	Abilities              []AbilityDetails   `json:"abilities"`
	BaseExpirience         int                `json:"base_experience"`
	Cries                  LatestLegacyStruct `json:"cries"`
	Forms                  []NameUrlStruct    `json:"forms"`
	GameIndices            []GameIndexStruct  `json:"game_indices"`
	Height                 int                `json:"height"`
	HeldItems              []HeldItemDetails  `json:"held_items"`
	Id                     int                `json:"id"`
	IsDefault              bool               `json:"is_default"`
	LocationAreaEncounters URL                `json:"location_area_encounters"`
	Moves                  []MoveDetails      `json:"moves"`
	Name                   string             `json:"name"`
	Order                  int                `json:"order"`
	PastAbilities          []PastAbility      `json:"past_abilities"`
	PastTypes              []byte             `json:"past_types"`
	Species                NameUrlStruct      `json:"species"`
	Sprites                SpritesDetails     `json:"sprites"`
	Stats                  []StatDetail       `json:"stats"`
	Types                  []TypeDetail       `json:"types"`
	Weight                 int                `json:"weight"`
}

type TypeDetail struct {
	Slot int           `json:"slot"`
	Type NameUrlStruct `json:"type"`
}

type StatDetail struct {
	BaseStat int           `json:"base_stat"`
	Effort   int           `json:"effort"`
	Stat     NameUrlStruct `json:"stat"`
}

type SpritesDetails struct {
	BackDefault      URL            `json:"back_default"`
	BackFemale       URL            `json:"back_female"`
	BackShiny        URL            `json:"back_shiny"`
	BackShinyFemale  URL            `json:"back_shiny_female"`
	FrontDefault     URL            `json:"front_default"`
	FrontFemale      URL            `json:"front_femare"`
	FrontShiny       URL            `json:"front_shiny"`
	FrontShinyFemale URL            `json:"front_shiny_female"`
	Other            OtherStruct    `json:"other"`
	Versions         VersionsStruct `json:"versions"`
}

type VersionsStruct struct {
	GenerationI struct {
		RedBlue struct {
			BackDefault      string `json:"back_default"`
			BackGray         string `json:"back_gray"`
			BackTransparent  string `json:"back_transparent"`
			FrontDefault     string `json:"front_default"`
			FrontGray        string `json:"front_gray"`
			FrontTransparent string `json:"front_transparent"`
		} `json:"red-blue"`
		Yellow struct {
			BackDefault      string `json:"back_default"`
			BackGray         string `json:"back_gray"`
			BackTransparent  string `json:"back_transparent"`
			FrontDefault     string `json:"front_default"`
			FrontGray        string `json:"front_gray"`
			FrontTransparent string `json:"front_transparent"`
		} `json:"yellow"`
	} `json:"generation-i"`
	GenerationIi struct {
		Crystal struct {
			BackDefault           string `json:"back_default"`
			BackShiny             string `json:"back_shiny"`
			BackShinyTransparent  string `json:"back_shiny_transparent"`
			BackTransparent       string `json:"back_transparent"`
			FrontDefault          string `json:"front_default"`
			FrontShiny            string `json:"front_shiny"`
			FrontShinyTransparent string `json:"front_shiny_transparent"`
			FrontTransparent      string `json:"front_transparent"`
		} `json:"crystal"`
		Gold struct {
			BackDefault      string `json:"back_default"`
			BackShiny        string `json:"back_shiny"`
			FrontDefault     string `json:"front_default"`
			FrontShiny       string `json:"front_shiny"`
			FrontTransparent string `json:"front_transparent"`
		} `json:"gold"`
		Silver struct {
			BackDefault      string `json:"back_default"`
			BackShiny        string `json:"back_shiny"`
			FrontDefault     string `json:"front_default"`
			FrontShiny       string `json:"front_shiny"`
			FrontTransparent string `json:"front_transparent"`
		} `json:"silver"`
	} `json:"generation-ii"`
	GenerationIii struct {
		Emerald struct {
			FrontDefault string `json:"front_default"`
			FrontShiny   string `json:"front_shiny"`
		} `json:"emerald"`
		FireredLeafgreen struct {
			BackDefault  string `json:"back_default"`
			BackShiny    string `json:"back_shiny"`
			FrontDefault string `json:"front_default"`
			FrontShiny   string `json:"front_shiny"`
		} `json:"firered-leafgreen"`
		RubySapphire struct {
			BackDefault  string `json:"back_default"`
			BackShiny    string `json:"back_shiny"`
			FrontDefault string `json:"front_default"`
			FrontShiny   string `json:"front_shiny"`
		} `json:"ruby-sapphire"`
	} `json:"generation-iii"`
	GenerationIv struct {
		DiamondPearl struct {
			BackDefault      string `json:"back_default"`
			BackFemale       string `json:"back_female"`
			BackShiny        string `json:"back_shiny"`
			BackShinyFemale  string `json:"back_shiny_female"`
			FrontDefault     string `json:"front_default"`
			FrontFemale      string `json:"front_female"`
			FrontShiny       string `json:"front_shiny"`
			FrontShinyFemale string `json:"front_shiny_female"`
		} `json:"diamond-pearl"`
		HeartgoldSoulsilver struct {
			BackDefault      string `json:"back_default"`
			BackFemale       string `json:"back_female"`
			BackShiny        string `json:"back_shiny"`
			BackShinyFemale  string `json:"back_shiny_female"`
			FrontDefault     string `json:"front_default"`
			FrontFemale      string `json:"front_female"`
			FrontShiny       string `json:"front_shiny"`
			FrontShinyFemale string `json:"front_shiny_female"`
		} `json:"heartgold-soulsilver"`
		Platinum struct {
			BackDefault      string `json:"back_default"`
			BackFemale       string `json:"back_female"`
			BackShiny        string `json:"back_shiny"`
			BackShinyFemale  string `json:"back_shiny_female"`
			FrontDefault     string `json:"front_default"`
			FrontFemale      string `json:"front_female"`
			FrontShiny       string `json:"front_shiny"`
			FrontShinyFemale string `json:"front_shiny_female"`
		} `json:"platinum"`
	} `json:"generation-iv"`
	GenerationV struct {
		BlackWhite struct {
			Animated struct {
				BackDefault      string `json:"back_default"`
				BackFemale       string `json:"back_female"`
				BackShiny        string `json:"back_shiny"`
				BackShinyFemale  string `json:"back_shiny_female"`
				FrontDefault     string `json:"front_default"`
				FrontFemale      string `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale string `json:"front_shiny_female"`
			} `json:"animated"`
			BackDefault      string `json:"back_default"`
			BackFemale       string `json:"back_female"`
			BackShiny        string `json:"back_shiny"`
			BackShinyFemale  string `json:"back_shiny_female"`
			FrontDefault     string `json:"front_default"`
			FrontFemale      string `json:"front_female"`
			FrontShiny       string `json:"front_shiny"`
			FrontShinyFemale string `json:"front_shiny_female"`
		} `json:"black-white"`
	} `json:"generation-v"`
	GenerationVi struct {
		OmegarubyAlphasapphire struct {
			FrontDefault     string `json:"front_default"`
			FrontFemale      string `json:"front_female"`
			FrontShiny       string `json:"front_shiny"`
			FrontShinyFemale string `json:"front_shiny_female"`
		} `json:"omegaruby-alphasapphire"`
		XY struct {
			FrontDefault     string `json:"front_default"`
			FrontFemale      string `json:"front_female"`
			FrontShiny       string `json:"front_shiny"`
			FrontShinyFemale string `json:"front_shiny_female"`
		} `json:"x-y"`
	} `json:"generation-vi"`
	GenerationVii struct {
		Icons struct {
			FrontDefault string `json:"front_default"`
			FrontFemale  any    `json:"front_female"`
		} `json:"icons"`
		UltraSunUltraMoon struct {
			FrontDefault     string `json:"front_default"`
			FrontFemale      string `json:"front_female"`
			FrontShiny       string `json:"front_shiny"`
			FrontShinyFemale string `json:"front_shiny_female"`
		} `json:"ultra-sun-ultra-moon"`
	} `json:"generation-vii"`
	GenerationViii struct {
		Icons struct {
			FrontDefault string `json:"front_default"`
			FrontFemale  string `json:"front_female"`
		} `json:"icons"`
	} `json:"generation-viii"`
}

type OtherStruct struct {
	DreamWorld struct {
		FrontDefault string `json:"front_default"`
		FrontFemale  any    `json:"front_female"`
	} `json:"dream_world"`
	Home struct {
		FrontDefault     string `json:"front_default"`
		FrontFemale      string `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale string `json:"front_shiny_female"`
	} `json:"home"`
	OfficialArtwork struct {
		FrontDefault string `json:"front_default"`
		FrontShiny   string `json:"front_shiny"`
	} `json:"official-artwork"`
	Showdown struct {
		BackDefault      string `json:"back_default"`
		BackFemale       string `json:"back_female"`
		BackShiny        string `json:"back_shiny"`
		BackShinyFemale  any    `json:"back_shiny_female"`
		FrontDefault     string `json:"front_default"`
		FrontFemale      string `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale string `json:"front_shiny_female"`
	} `json:"showdown"`
}
type PastAbility struct {
	Abilities  []AbilityDetails `json:"abilities"`
	Generation NameUrlStruct    `json:"generation"`
}

type MoveDetails struct {
	Move                NameUrlStruct        `json:"move"`
	VersionGroupDetails []VersionGroupDetail `json:"version_group_details"`
}

type VersionGroupDetail struct {
	LevelLearnedAt int           `json:"level_learned_at"`
	MoveLearMethod NameUrlStruct `json:"move_learn_method"`
	Order          int           `json:"order"`
	VersionGroup   NameUrlStruct `json:"version_group"`
}

type HeldItemDetails struct {
	Item           NameUrlStruct   `json:"item"`
	VersionDetails []RarityDetails `json:"version_details"`
}

type RarityDetails struct {
	Rarity  int           `json:"rarity"`
	Version NameUrlStruct `json:"version"`
}

type GameIndexStruct struct {
	GameIndex int           `json:"game_index"`
	Version   NameUrlStruct `json:"version"`
}

type AbilityDetails struct {
	Ability  NameUrlStruct `json:"ability"`
	IsHidden bool          `json:"is_hidden"`
	Slot     int           `json:"slot"`
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
const PokeapiPokemonUrl URL = "https://pokeapi.co/api/v2/pokemon/"
