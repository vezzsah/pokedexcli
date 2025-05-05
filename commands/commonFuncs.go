package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/vezzsah/pokedexcli/types"
)

func unmarshalBodyJson[T any](body []byte, responseJson T) error {
	err := json.Unmarshal(body, &responseJson)
	if err != nil {
		return fmt.Errorf("error while unmarshaling: %s\nBody: %s", err, string(body))
	}
	return nil
}

func callPokeAPIGet(apiURL string) ([]byte, error) {

	res, err := http.Get(apiURL)
	if err != nil {
		return nil, errors.New("error calling API")
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("error while readying body: %s", err)
	}
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %d\n URL = %s\nBody: %s", res.StatusCode, apiURL, body)
	}

	return body, nil
}

func GetAPICallFromCache[T any](c *types.Config, url types.URL, responseJson *T) error {
	var err error
	apiURL := string(url) + c.Parameters
	body, found := c.CacheMap.Get(apiURL)
	if !found {
		body, err = callPokeAPIGet(apiURL)
		if err != nil {
			return err
		}

		c.CacheMap.Add(apiURL, body)
	}

	err = unmarshalBodyJson(body, &responseJson)
	if err != nil {
		return err
	}

	return nil
}
