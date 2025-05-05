package commands

import (
	"fmt"

	"github.com/vezzsah/pokedexcli/types"
)

func CommandExplore(c *types.Config) error {
	var err error
	var responseJson types.ExplorePageResponse

	if c.Parameters == "" {
		fmt.Println("Please provide an area to explore. For example: explore pastoria-city-area")
		return nil
	}

	locationToSearch := string(types.PokeapiLocationAreaUrl) + c.Parameters
	body, found := c.CacheMap.Get(locationToSearch)
	if !found {
		body, err = callPokeAPIGet(locationToSearch)
		if err != nil {
			return err
		}

		c.CacheMap.Add(locationToSearch, body)
	}

	err = unmarshalBodyJson(body, &responseJson)
	if err != nil {
		return err
	}

	for _, encounter := range responseJson.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}

/*
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
*/
