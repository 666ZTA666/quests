package main

import (
	"fmt"
	"strings"
)

func set(array []string) []string {
	table := make(map[string]struct{})
	for _, v := range array {
		table[v] = struct{}{}
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
		_, err := fmt.Scanln(&n) // прочитали записали
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
		_, err := fmt.Scanln(&str) // прочитали записали
		if err != nil {
			continue
		}
		str = strings.TrimSpace(str) // удаляет пробелы табы итд, так как лишний пробел может сделать строку уникальной
		if str == "" {
			i--
			fmt.Println("введите не пустую строку") // проверка на пустоту
		} else {
			array = append(array, str)
		}
	}
	fmt.Printf("Array: %v\n", array)    // весь срез строк
	fmt.Printf("Set:   %v", set(array)) // срез уникальных строк
}
