package main

import "fmt"

func main() {
	var coutryMap map[string]string
	coutryMap = make(map[string]string)

	coutryMap["China"] = "BeiJing"
	coutryMap["America"] = "NewYork"

	delete(coutryMap, "America")

	mapPrint(coutryMap)
}

func mapPrint(x map[string]string) {
	for key, value := range x {
		fmt.Println(key, value)
	}
}
