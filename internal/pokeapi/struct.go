package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

const defaultPath = "https://pokeapi.co/api/v2/location"

var resultMap PokeMap = PokeMap{}

type PokeMap struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (p *PokeMap) unmarshal(body []byte) {
	err := json.Unmarshal(body, p)
	if err != nil {
		fmt.Println(err)
	}
}

func (p *PokeMap) getNext() string {
	if p.Next == "" {
		return defaultPath
	}

	return p.Next
}

func (p *PokeMap) getPrev() (string, error) {
	if p.Previous == nil {
		fmt.Println("Error: you are on the first page")
		return defaultPath, errors.New("unable to go back")
	}

	return *p.Previous, nil
}
