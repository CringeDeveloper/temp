package pokeapi

import (
	"fmt"
)

func GetMap() error {
	body, err := getPoket(resultMap.getNext())
	if err != nil {
		fmt.Println(err)
	}

	resultMap.unmarshal(body)

	printCities(resultMap)

	return nil
}

func printCities(pokeMap PokeMap) {
	for _, v := range pokeMap.Results {
		fmt.Println(v.Name)
	}
}
