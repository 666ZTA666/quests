package main

import (
	"fmt"
	"strings"
)

func set(array []string) []string {
	table := make(map[string]bool)
	for _, v := range array {
		table[v] = true
	}
	var result []string
	for key := range table {
		result = append(result, key)
	}
	return result
}

func main() {
	var n int
	for {
		fmt.Println("введите количество строк в срезе")
		_, err := fmt.Scanln(&n)
		if err != nil {
			continue
		}
		if n > 0 {
			break
		}
	}
	array := make([]string, 0)
	for i := 1; i <= n; i++ {
		fmt.Println("введите строку №", i)
		var str string
		_, err := fmt.Scanln(&str)
		if err != nil {
			continue
		}
		str = strings.TrimSpace(str)
		if str == "" {
			i--
			fmt.Println("введите не пустую строку")
		} else {
			array = append(array, str)
		}
	}
	fmt.Printf("Array: %v\n", array)
	fmt.Printf("Set:   %v", set(array))
}
