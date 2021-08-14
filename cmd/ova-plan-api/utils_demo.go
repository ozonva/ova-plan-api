package main

import (
	"bufio"
	"fmt"
	"github.com/k0kubun/pp"
	"github.com/ozonva/ova-plan-api/internal/utils"
	"os"
)

func demoUtils() {
	demoReverse()
	demoBatching()
	demoFilter()
}

func demoReverse() {
	strToIntMap := map[string]int{"key1": 1, "key2": 2, "key3": 3}
	pp.Printf("Before reverse: %v \n", strToIntMap)

	reversed, _ := utils.ReverseMap(strToIntMap)

	pp.Printf("After reverse: %v \n\n", reversed)
}

func demoBatching() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Before batching: %v \n", slice)

	batched, _ := utils.SplitSlice(slice, 2)

	fmt.Printf("Before batching: %v \n\n", batched)
}

func demoFilter() {
	slice := []int{1, 2, 2, 3, 3, 4, 5}
	toDelete := []int{0, 1, 2, 3}
	fmt.Printf("Before filtering: %v \n"+
		"Values to delete: %v \n", slice, toDelete)

	filtered := utils.FilterSlice(slice, toDelete)
	fmt.Printf("After filtering: %v \n\n", filtered)
}

func demoReadFile() {
	err := utils.ReadFiles(printFileContent, "resources/default.config", "resources/custom.config")
	if err != nil {
		print(err.Error())
	}
}

func printFileContent(file *os.File) error {
	fmt.Printf("Start reading %v\n", file.Name())
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Print(err)
	}
	return nil
}
