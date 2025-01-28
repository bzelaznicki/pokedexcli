package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreaResp, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	if cached, found := c.cache.Get(url); found {
		var locAreaResp LocationAreaResp
		err := json.Unmarshal(cached, &locAreaResp)
		if err != nil {
			return LocationAreaResp{}, err
		}
		return locAreaResp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}

	defer resp.Body.Close()

	cont, err := io.ReadAll(resp.Body)

	if err != nil {
		return LocationAreaResp{}, err
	}

	c.cache.Add(url, cont)

	locationsResp := LocationAreaResp{}

	err = json.Unmarshal(cont, &locationsResp)
	if err != nil {
		return LocationAreaResp{}, err
	}

	return locationsResp, nil

}
