package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

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
	if c.Option == "previous" {
		pageToLookFor = string(c.Previous)
	} else {
		pageToLookFor = string(c.Next)
	}

	res, err := http.Get(pageToLookFor)
	if err != nil {
		return errors.New("error calling API")
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return fmt.Errorf("error while unmarshaling: %s", err)
	}
	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}

	var responseJson types.LocationsPageResponse
	err = json.Unmarshal(body, &responseJson)
	if err != nil {
		return fmt.Errorf("error while unmarshaling: %s", err)
	}

	if responseJson.Previous == "" && c.Option == "previous" {
		c.Previous = ""
		c.Next = types.PokeapiLocationUrl
	} else {
		c.Previous = responseJson.Previous
		c.Next = responseJson.Next
	}

	for _, location := range responseJson.Results {
		fmt.Println(location.Name)
	}

	return nil
}
