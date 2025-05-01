package commands

import (
	"fmt"

	"github.com/vezzsah/pokedexcli/types"
)

func CommandHelp(c *types.Config) error {
	fmt.Println(`Welcome to the Pokedex!
		Usage:
		
		help: Displays a help message
		exit: Exit the Pokedex`)
	return nil
}
