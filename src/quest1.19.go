package main

import "fmt"

func main() {
	var v string
	for {
		fmt.Println("введите строку")
		_, err := fmt.Scanln(&v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if v != "" {
			break
		}
	}
	fmt.Println(Reverse(v))
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
