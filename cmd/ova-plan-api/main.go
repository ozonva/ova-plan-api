package main

import (
	"fmt"
	"github.com/k0kubun/pp"
	"github.com/ozonva/ova-plan-api/internal/utils"
)

func main() {
	fmt.Println("Hello there!")

	demoUtils()
}

func demoUtils() {
	demoReverse()
}

func demoReverse() {
	strToIntMap := map[string]int{"key1": 1, "key2": 2, "key3": 3}
	fmt.Println("Before reversing:")
	pp.Println(strToIntMap)

	reversed, err := utils.ReverseMap(strToIntMap)
	if err != nil {
		return
	}

	fmt.Println("After reverse:")
	pp.Println(reversed)
}
