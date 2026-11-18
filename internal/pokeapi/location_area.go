package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

func (c *Client) LocationArea(area string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area/" + area

	if val, ok := c.cache.Get(url); ok {
		locationAreaResp := LocationAreaResponse{}
		err := json.Unmarshal(val, &locationAreaResp)
		if err == nil {
			// cache hit
			return locationAreaResp, nil
		}
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return LocationAreaResponse{}, errors.New("Could not fetch the area")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	locationAreaResp := LocationAreaResponse{}
	if err = json.Unmarshal(data, &locationAreaResp); err != nil {
		fmt.Println("err")
		fmt.Println(err)
		return LocationAreaResponse{}, err
	}

	c.cache.Add(url, data)

	return locationAreaResp, nil
}
