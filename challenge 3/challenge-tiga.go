package main

import (
	"fmt"
)

func main() {
	text := "selamat malam"
	countMap := make(map[string]int)
	for _, ch := range text {
		fmt.Println(string(ch))
		countMap[string(ch)]++
	}
	fmt.Println(countMap)
}
