package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	if cached, found := c.cache.Get(url); found {
		var poke Pokemon
		err := json.Unmarshal(cached, &poke)
		if err != nil {
			return Pokemon{}, err
		}
		return poke, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return Pokemon{}, fmt.Errorf("pokemon not found")
	} else if resp.StatusCode != 200 {
		return Pokemon{}, fmt.Errorf("unexpected error: status code %d", resp.StatusCode)
	}

	cont, err := io.ReadAll(resp.Body)

	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, cont)

	poke := Pokemon{}

	err = json.Unmarshal(cont, &poke)
	if err != nil {
		return Pokemon{}, err
	}

	return poke, nil

}
