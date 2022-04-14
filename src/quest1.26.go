package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//Разработать программу, которая проверяет, что все символы в строке
//уникальные (true — если уникальные, false etc). Функция проверки должна быть
//регистронезависимой.
func main() {
	r := bufio.NewReader(os.Stdin)
	var s string
	for {
		fmt.Println("введите строку, можно с пробелами")
		s = readline(r)
		if s != "" {
			break
		}
	}
	runes := []rune(strings.ToUpper(s))
	fmt.Println(isUniq(runes))

}
func isUniq(e []rune) bool {
	for j := 0; j < len(e)-1; j++ {
		for i := j + 1; i < len(e)-1; i++ {
			if e[i] == e[j] {
				return false
			}
		}
	}
	return true
}

func readline(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
