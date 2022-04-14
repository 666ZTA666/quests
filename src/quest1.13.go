package main

import "fmt"

func main() {
	var v1, v2 int
	for {
		fmt.Println("введите первое значение")
		_, err := fmt.Scanln(&v1) // считали, записали
		if err != nil {
			fmt.Println(err)
			continue
		}
		if v1 != 0 {
			break
		}
	}
	for {
		fmt.Println("введите второе значение")
		_, err := fmt.Scanln(&v2) // считали, записали
		if err != nil {
			fmt.Println(err)
			continue
		}
		if v2 != 0 {
			break
		}
	}
	// Вся магия ниже.
	v1, v2 = v2, v1
	// Вот и всё.
	//
	//v1 = v1+v2 	Можно использовать такой аналог, но при значениях в районе maxInt могут быть проблемы.
	//v2 = v1-v2	Так же можно использовать пустой интерфейс или дженерики, для работы с разными типами переменных.
	fmt.Print("вывожу значения в порядке ввода\n", v1, " ", v2)
}
