package swapi

import (
	"encoding/json"
	"errors"
	"net/http"
)

const BaseURL = "http://swapi.co/api"

var ErrNotFound = errors.New("404: not found")

// Get makes an API call to the path, it decodes the JSON response
// and stores it into out.
func Get(path string, out interface{}) error {
	url := path
	if path[:4] != "http" {
		url = BaseURL + path
	}

	res, err := http.Get(url)
	defer res.Body.Close()

	if err != nil {
		return err
	}

	if res.StatusCode == 404 {
		return ErrNotFound
	}

	if err = json.NewDecoder(res.Body).Decode(out); err != nil {
		return err
	}

	return nil
}

func GetRoot() (map[string]string, error) {
	var out map[string]string
	return out, Get("/", &out)
}
