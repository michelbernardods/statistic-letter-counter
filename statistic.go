package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	words := os.Args[1:]
	statistics := saveStatistics(words)
	print(statistics)
}

func saveStatistics(words []string) map[string]int {
	statistics := make(map[string]int)
	for _, words := range words {
		initial := strings.ToUpper(string(words[0]))
		counter, found := statistics[initial]

		if found {
			statistics[initial] = counter + 1
		} else {
			statistics[initial] = 1
		}
	}
	return statistics
}

func print(statistics map[string]int) {
	fmt.Println("Count of words started in each letter:")
	for initial, counter := range statistics {
		fmt.Printf("%s = %d\n", initial, counter)
	}
}
