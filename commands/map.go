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
	var err error
	if c.Option == "previous" {
		pageToLookFor = string(c.Previous)
	} else {
		pageToLookFor = string(c.Next)
	}

	var responseJson types.LocationsPageResponse

	body, found := c.CacheMap.Get(pageToLookFor)
	if !found {
		body, err = callPokeAPIGetLocation(pageToLookFor)
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

func callPokeAPIGetLocation(pageToLookFor string) ([]byte, error) {

	res, err := http.Get(pageToLookFor)
	if err != nil {
		return nil, errors.New("error calling API")
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("error while readying body: %s", err)
	}
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %d and\nBody: %s", res.StatusCode, body)
	}

	return body, nil
}

func unmarshalBodyJson(body []byte, responseJson *types.LocationsPageResponse) error {
	err := json.Unmarshal(body, &responseJson)
	if err != nil {
		return fmt.Errorf("error while unmarshaling: %s\nBody: %s", err, string(body))
	}
	return nil
}
