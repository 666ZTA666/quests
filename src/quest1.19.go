package main

import (
	"fmt"
	"strings"
)

func main() {
	var v string
	for {
		fmt.Println("введите строку")
		_, err := fmt.Scanln(&v) // считали записали
		if err != nil {
			fmt.Println(err)
			continue
		}
		v = strings.TrimSpace(v) // удаляем пробелы, табы и прочую нечисть
		if v != "" {
			break
		}
	}
	// перевернули и вывели
	fmt.Println(Reverse(v))
}

func Reverse(s string) string {
	runes := []rune(s) // делим на руны
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		// проходимся по рунам и меняем их местами
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
