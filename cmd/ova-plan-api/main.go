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
	demoBatching()
}

func demoReverse() {
	strToIntMap := map[string]int{"key1": 1, "key2": 2, "key3": 3}
	pp.Printf("Before reverse: %v \n", strToIntMap)

	reversed, _ := utils.ReverseMap(strToIntMap)

	pp.Printf("After reverse: %v \n", reversed)
}

func demoBatching() {
	slice := []int{1, 2, 3, 4, 5}
	pp.Printf("Before batching: %v \n", slice)

	batched, _ := utils.SplitSlice(slice, 2)

	pp.Printf("Before batching: %v \n", batched)
}
