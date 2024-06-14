package pokeapi

import (
	"encoding/json"
	"fmt"
)

func GetMap() error {
	superMap := PokeMap{}
	body, err := getPoket("https://pokeapi.co/api/v2/location/?offset=20&limit=20")

	err = json.Unmarshal(body, &superMap)
	if err != nil {
		fmt.Println(err)
	}

	printCities(superMap)

	return nil
}

func printCities(pokeMap PokeMap) {
	for _, v := range pokeMap.Results {
		fmt.Println(v.Name)
	}
}
