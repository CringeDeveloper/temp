package pokeapi

import (
	"encoding/json"
	"fmt"
)

func MapBack() error {
	superMap := PokeMap{}
	body, err := getPoket("https://pokeapi.co/api/v2/location/?offset=0&limit=20")

	err = json.Unmarshal(body, &superMap)
	if err != nil {
		fmt.Println(err)
	}

	printCities(superMap)

	return nil
}
