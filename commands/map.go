package commands

import (
	"fmt"

	"github.com/vezzsah/pokedexcli/types"
)

func CommandBMap(c *types.Config) error {
	c.Option = "previous"
	CommandMap(c)
	c.Option = ""
	return nil
}

func CommandMap(c *types.Config) error {
	var pageToLookFor string
	var err error
	if c.Option == "previous" {
		pageToLookFor = string(c.Previous)
	} else {
		pageToLookFor = string(c.Next)
	}

	var responseJson types.LocationsPageResponse

	body, found := c.CacheMap.Get(pageToLookFor)
	if !found {
		body, err = callPokeAPIGet(pageToLookFor)
		if err != nil {
			return err
		}

		c.CacheMap.Add(pageToLookFor, body)
	}

	err = unmarshalBodyJson(body, &responseJson)
	if err != nil {
		return err
	}

	if responseJson.Previous == "" && c.Option == "previous" {
		c.Previous = ""
		c.Next = types.PokeapiLocationAreaUrl
	} else {
		c.Previous = responseJson.Previous
		c.Next = responseJson.Next
	}

	for _, location := range responseJson.Results {
		fmt.Println(location.Name)
	}

	return nil
}
