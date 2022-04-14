package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//Удалить i-ый элемент из слайса.
	var x int
	for {
		fmt.Println("Введите количество элементов среза")
		_, err := fmt.Scanln(&x) // считали, записали
		if err != nil {
			fmt.Println(err)
		} else if x > 0 {
			break
		} else {
			fmt.Println("Число должно быть больше нуля!") // проверка на 0
		}
	}
	arr := make([]int, x)
	for i := 0; i < x; i++ {
		arr[i] = rand.Intn(100) + 1 // сгенерировали массив
	}
	fmt.Println(arr) // вывели сгенерированный массив
	var y int
	for {
		fmt.Println("Введите номер(не индекс) элемента который надо удалить")
		_, err := fmt.Scanln(&y) // считали, записали
		if err != nil {
			fmt.Println(err)
		} else if y > 0 && y <= x {
			break
		} else if y <= 0 {
			fmt.Println("Номер != индекс, номера начинаются с 1") // проверка
		} else {
			fmt.Println("Нельзя удалить элемент, которого нет") // проверка
		}
	}
	// так как Я решил сделать 2 метода удаления элемента из среза,
	//мне нужно каждый раз создавать новый срез чтобы данные не менялись.
	fmt.Print("удаляем:[", arr[y-1], "]\n") // выводим удаляемый элемент
	arr1 := make([]int, x)
	copy(arr1, arr)                                                    // копируем срез
	fmt.Print("сохранили порядок\n", removeWithOrder(arr1, y-1), "\n") // удаляем и выводим
	arr2 := make([]int, x)
	copy(arr2, arr)                                                      // копируем срез
	fmt.Print("нарушили порядок\n", removeWithoutOrder(arr2, y-1), "\n") // удаляем и выводим

}
func removeWithOrder(s []int, i int) []int {
	v := append(s[:i], s[i+1:]...) // складываем два среза, исключая i элемент
	return v
}
func removeWithoutOrder(s []int, i int) []int {
	s[i] = s[len(s)-1] // меняем i элемент и последний местами, отрезаем последний
	return s[:len(s)-1]
}
