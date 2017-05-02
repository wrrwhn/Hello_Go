package utils

import (
	"fmt"
)

func HelloMap() {
	initMap()
}

func initMap() {

	// init
	initMap := map[string]int{"one": 1, "two": 2}
	// search
	fmt.Println(initMap)
	searchMap(initMap, "two")
	searchMap(initMap, "three")

	// add + delete
	initMap["three"] = 3
	delete(initMap, "two")
	// search
	fmt.Println(initMap)
	searchMap(initMap, "two")
	searchMap(initMap, "three")
}

func searchMap(toSearchMap map[string]int, key string) {

	val, ok := toSearchMap[key]
	if ok {
		fmt.Printf("Map[\"%v\"]= %v\n", key, val)
	} else {
		fmt.Printf("Map[\"%v\"] is not exists\n", key)
	}
}
