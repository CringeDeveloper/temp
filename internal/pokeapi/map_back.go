package pokeapi

import (
	"fmt"
)

func MapBack() error {
	prevPage, err := resultMap.getPrev()
	if err != nil {
		return err
	}

	body, err := getPoket(prevPage)
	if err != nil {
		fmt.Println(err)
	}
	resultMap.unmarshal(body)

	printCities(resultMap)

	return nil
}
