package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func unmarshalBodyJson[T any](body []byte, responseJson T) error {
	err := json.Unmarshal(body, &responseJson)
	if err != nil {
		return fmt.Errorf("error while unmarshaling: %s\nBody: %s", err, string(body))
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
