package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//Разработать программу, которая переворачивает слова в строке.
	//Пример: «snow dog sun — sun dog snow»
	var v string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("введите строку из двух и более слов")
		v = readLine(reader)
		if v != "" {
			break
		}
	}
	varr := strings.Fields(v)
	tmp := make([]string, len(varr))
	fmt.Print("Йода говорит: ")
	for i := len(varr) - 1; i >= 0; i-- {
		fmt.Print(varr[i], " ")
		tmp[len(varr)-i-1] = varr[i]
	}

}
func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
