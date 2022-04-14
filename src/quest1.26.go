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
		s = readline(r) // считали, записали
		if s != "" {
			break
		}
	}
	runes := []rune(strings.ToUpper(s)) // поднимаем все руны в строке вверх
	fmt.Println(isUniqMap(runes))       // выводим булевый результат

}

/*
func isUniq(e []rune) bool {
	// Тут мой ум иссяк делать мапы с символами и Я решил просто сделать цикл в цикле
	for j := 0; j < len(e)-1; j++ {
		for i := j + 1; i < len(e)-1; i++ {
			if e[i] == e[j] {
				return false
			}
		}
	}
	return true
}

Когда устал придумывать умные варианты
*/
func isUniqMap(a []rune) bool {
	// создадим карту для хэширования данных
	hash := make(map[rune]struct{})
	for i := 0; i < len(a); i++ {
		if _, ok := hash[a[i]]; ok { //Нашли повтор - выходим
			return false
		} else { // не нашли повтор, записываем значение
			hash[a[i]] = struct{}{}
		}
	}
	return true // если ни разу не вышли по повтору, то возвращаем true
}

// старое доброе чтение из bufio reader'а с удалением символов переноса каретки итд
func readline(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
