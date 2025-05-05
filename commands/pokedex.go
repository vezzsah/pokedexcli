package commands

import (
	"fmt"

	"github.com/vezzsah/pokedexcli/types"
)

func CommandPokedex(c *types.Config) error {

	if len(c.PokemonCaught) == 0 {
		fmt.Println("You don't have any pokemon caught yet! Use the catch <pokemon-name> command to get new team members")
		return nil
	}

	fmt.Println("Your pokedex is:")
	for k := range c.PokemonCaught {
		fmt.Printf("	-%s\n", k)
	}

	return nil
}
