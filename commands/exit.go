package commands

import (
	"fmt"
	"os"

	"github.com/vezzsah/pokedexcli/types"
)

func CommandExit(c *types.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
