package commands

import (
	"fmt"

	"github.com/vezzsah/pokedexcli/types"
)

func CommandInspect(c *types.Config) error {

	pokemon, ok := c.PokemonCaught[c.Parameters]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	var stats map[string]int = make(map[string]int)

	for _, unknownStat := range pokemon.Stats {
		stats[unknownStat.Stat.Name] = unknownStat.BaseStat
	}

	fmt.Printf(`
		Name: %s
		Height: %d
		Weight: %d
		Stats:
  			-hp: %d
  			-attack: %d
  			-defense: %d
  			-special-attack: %d
  			-special-defense: %d
  			-speed: %d
		Types:
	`, pokemon.Name, pokemon.Height, pokemon.Weight,
		stats["hp"],
		stats["attack"],
		stats["defense"],
		stats["special-attack"],
		stats["special-defense"],
		stats["speed"])

	for _, singleType := range pokemon.Types {
		fmt.Printf("  		- %s\n", singleType.Type.Name)
	}
	return nil
}
