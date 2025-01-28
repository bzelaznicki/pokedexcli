package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(name string) (LocationArea, error) {
	url := baseURL + "/location-area/" + name

	if cached, found := c.cache.Get(url); found {
		var locArea LocationArea
		err := json.Unmarshal(cached, &locArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locArea, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	defer resp.Body.Close()

	cont, err := io.ReadAll(resp.Body)

	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, cont)

	locArea := LocationArea{}

	err = json.Unmarshal(cont, &locArea)
	if err != nil {
		return LocationArea{}, err
	}

	return locArea, nil

}
