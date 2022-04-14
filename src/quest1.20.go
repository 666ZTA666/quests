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
	reader := bufio.NewReader(os.Stdin) // считывать строки пришлось посложнее)
	for {
		fmt.Println("введите строку из двух и более слов")
		v = readLine(reader) // считали записали
		if v != "" {
			break
		}
	}
	varr := strings.Fields(v)        // разделили строку на срез строк
	tmp := make([]string, len(varr)) // создали временный массив строк
	fmt.Print("Йода говорит: ")      // просто для веселья)
	for i := len(varr) - 1; i >= 0; i-- {
		// идем по массиву строк с конца в начало и выводим слова
		fmt.Print(varr[i], " ")
		tmp[len(varr)-i-1] = varr[i]
		// так же записываем слова в обратном порядке в массив, если нам нужна обратно строка,
		//а не срез то можно будет их объединить.
	}

}
func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n") //отсекаем всякое переносностроковое заранее
}
