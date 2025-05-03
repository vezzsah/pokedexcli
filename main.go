package main

import (
	//"fmt"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/vezzsah/pokedexcli/commands"
	"github.com/vezzsah/pokedexcli/pokecache"
	"github.com/vezzsah/pokedexcli/types"
)

func main() {
	var supportedCommands map[string]types.CliCommand
	commands.InitiateSupportedCommands(&supportedCommands)

	cache := pokecache.NewCache(time.Second * 5)

	scanner := bufio.NewScanner(os.Stdin)
	var userInput []string
	conf := types.Config{
		Next:     types.PokeapiLocationUrl,
		Previous: "",
		Option:   "",
		CacheMap: cache,
	}

	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan() //Here is where we wait for input
		if !ok {
			fmt.Println(scanner.Err())
		}
		userInput = cleanInput(scanner.Text())

		command, ok := supportedCommands[userInput[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.Callback(&conf)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	returnSlice := strings.Fields(text)
	return returnSlice
}
