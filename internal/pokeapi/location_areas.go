package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
)

func (c *Client) LocationAreas(pageURL string) (LocationAreasResponse, error) {
	// default url to fetch from
	url := baseURL + "/location-area"
	// if a url is passed, use it
	if pageURL != "" {
		url = pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationAreasResp := LocationAreasResponse{}
		err := json.Unmarshal(val, &locationAreasResp)
		if err == nil {
			// cache hit
			return locationAreasResp, nil
		}
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return LocationAreasResponse{}, errors.New("Could not fetch the areas")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationAreasResp := LocationAreasResponse{}
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	c.cache.Add(url, data)

	return locationAreasResp, nil
}
