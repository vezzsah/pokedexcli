package commands

import (
	"fmt"
	"math/rand"

	"github.com/vezzsah/pokedexcli/types"
)

func CommandCatch(c *types.Config) error {
	if c.Parameters == "" {
		fmt.Println("You need to pick a Pokemon to Catch!")
		return nil
	}

	_, alreadyHave := c.PokemonCaught[c.Parameters]
	if alreadyHave {
		fmt.Printf("You already have a %s\n", c.Parameters)
		return nil
	}

	var responseJson types.PokemonPageResponse
	err := GetAPICallFromCache(c, types.PokeapiPokemonUrl, &responseJson)
	if err != nil {
		return err
	}

	pokemonName := responseJson.Name
	catchRate := 60 - responseJson.BaseExpirience/10
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	if catchRate >= rand.Intn(100) {
		fmt.Printf("%s was caught!\nYou may now inspect it with the inspect command.\n", pokemonName)
		c.PokemonCaught[pokemonName] = responseJson
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
