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
		"explore": {
			Name:        "explore",
			Description: "Show the pokemon located in the given area",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Try to catch the requested pokemon",
			Callback:    CommandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect the catched pokemon",
			Callback:    CommandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Inspect your pokemon team",
			Callback:    CommandPokedex,
		},
	}
}
