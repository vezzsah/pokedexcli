package commands

import (
	"github.com/vezzsah/pokedexcli/types"
)

func InitiateSupportedCommands(supportedCommands *map[string]types.CliCommand) {
	*supportedCommands = map[string]types.CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Explains other commands and the function of this program",
			Callback:    CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "It displays the Names of 20 location areas in the Pokemon world. Each subsequent call to map should display the next 20 locations, and so on",
			Callback:    CommandMap,
		},
		"bmap": {
			Name:        "bmap",
			Description: "It displays the Names of the prev 20 location areas in the Pokemon world. Each subsequent call to map should display the prev 20 locations, and so on",
			Callback:    CommandBMap,
		},
	}
}
